package mobile

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golang/protobuf/proto"
	ipld "github.com/ipfs/go-ipld-format"
	ipfspath "github.com/ipfs/go-path"
	mh "github.com/multiformats/go-multihash"
	"github.com/textileio/go-textile/core"
	"github.com/textileio/go-textile/ipfs"
	"github.com/textileio/go-textile/mill"
	"github.com/textileio/go-textile/pb"
	"github.com/textileio/go-textile/schema"
)

var fileConfigOpt fileConfigOption

type fileConfigSettings struct {
	Use       string
	Data      []byte
	Path      string
	Plaintext bool
}

type fileConfigOption func(*fileConfigSettings)

func (fileConfigOption) Use(val string) fileConfigOption {
	return func(settings *fileConfigSettings) {
		settings.Use = val
	}
}

func (fileConfigOption) Data(val []byte) fileConfigOption {
	return func(settings *fileConfigSettings) {
		settings.Data = val
	}
}

func (fileConfigOption) Path(val string) fileConfigOption {
	return func(settings *fileConfigSettings) {
		settings.Path = val
	}
}

func (fileConfigOption) Plaintext(val bool) fileConfigOption {
	return func(settings *fileConfigSettings) {
		settings.Plaintext = val
	}
}

func fileConfigOptions(opts ...fileConfigOption) *fileConfigSettings {
	options := &fileConfigSettings{}

	for _, opt := range opts {
		opt(options)
	}
	return options
}

func (m *Mobile) AddData(data []byte, threadId string, caption string, cb Callback) {
	go func() {
		hash, err := m.addData(data, threadId, caption)
		if err != nil {
			cb.Call(nil, err)
			return
		}

		cb.Call(m.blockView(hash))
	}()
}

func (m *Mobile) AddFiles(paths []byte, threadId string, caption string, cb Callback) {
	go func() {
		pths := new(pb.Strings)
		err := proto.Unmarshal(paths, pths)
		if err != nil {
			cb.Call(nil, err)
			return
		}

		hash, err := m.addFiles(pths.Values, threadId, caption)
		if err != nil {
			cb.Call(nil, err)
			return
		}

		cb.Call(m.blockView(hash))
	}()
}

func (m *Mobile) ShareFiles(target string, threadId string, caption string, cb Callback) {
	go func() {
		hash, err := m.shareFiles(target, threadId, caption)
		if err != nil {
			cb.Call(nil, err)
			return
		}

		cb.Call(m.blockView(hash))
	}()
}

// Files calls core Files
func (m *Mobile) Files(threadId string, offset string, limit int) ([]byte, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	files, err := m.node.Files(offset, limit, threadId)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(files)
}

// FileData returns a data url of a raw file under a path
func (m *Mobile) FileData(hash string) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	reader, file, err := m.node.FileData(hash)
	if err != nil {
		if err == core.ErrFileNotFound || err == ipld.ErrNotFound {
			return "", nil
		}
		return "", err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	prefix := "data:" + file.Media + ";base64,"
	return prefix + base64.StdEncoding.EncodeToString(data), nil
}

type img struct {
	hash  string
	width int
}

// ImageFileDataForMinWidth returns a data url of an image at or above requested size,
// or the next best option.
// Note: Now that consumers are in control of image sizes via schemas,
// handling this here doesn't feel right. We can eventually push this up to RN, Obj-C, Java.
// Note: pth is <target>/<index>, e.g., "Qm.../0"
func (m *Mobile) ImageFileDataForMinWidth(pth string, minWidth int) (string, error) {
	if !m.node.Started() {
		return "", core.ErrStopped
	}

	node, err := ipfs.NodeAtPath(m.node.Ipfs(), pth)
	if err != nil {
		if err == ipld.ErrNotFound {
			return "", nil
		}
		return "", err
	}

	var imgs []img
	for _, link := range node.Links() {
		nd, err := ipfs.NodeAtLink(m.node.Ipfs(), link)
		if err != nil {
			if err == ipld.ErrNotFound {
				return "", nil
			}
			return "", err
		}

		dlink := schema.LinkByName(nd.Links(), core.ValidContentLinkNames)
		if dlink == nil {
			continue
		}

		file, err := m.node.FileIndex(dlink.Cid.Hash().B58String())
		if err != nil {
			if err == core.ErrFileNotFound {
				return "", nil
			}
			return "", err
		}

		if file.Mill == "/image/resize" {
			width := file.Meta.Fields["width"]
			if width != nil {
				imgs = append(imgs, img{
					hash:  file.Hash,
					width: int(width.GetNumberValue()),
				})
			}
		}
	}

	if len(imgs) == 0 {
		return "", nil
	}

	sort.SliceStable(imgs, func(i, j int) bool {
		return imgs[i].width < imgs[j].width
	})

	var hash string
	for _, img := range imgs {
		if img.width >= minWidth {
			hash = img.hash
			break
		}
	}
	if hash == "" {
		hash = imgs[len(imgs)-1].hash
	}

	return m.FileData(hash)
}

func (m *Mobile) addData(data []byte, threadId string, caption string) (mh.Multihash, error) {
	dir, err := m.buildDirectory(data, "", threadId)
	if err != nil {
		return nil, err
	}

	return m.writeFiles(&pb.DirectoryList{Items: []*pb.Directory{dir}}, threadId, caption)
}

func (m *Mobile) addFiles(paths []string, threadId string, caption string) (mh.Multihash, error) {
	dirs := &pb.DirectoryList{Items: make([]*pb.Directory, 0)}
	for _, pth := range paths {
		dir, err := m.buildDirectory(nil, pth, threadId)
		if err != nil {
			return nil, err
		}
		dirs.Items = append(dirs.Items, dir)
	}

	return m.writeFiles(dirs, threadId, caption)
}

func (m *Mobile) shareFiles(target string, threadId string, caption string) (mh.Multihash, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return nil, core.ErrThreadNotFound
	}

	node, err := ipfs.NodeAtPath(m.node.Ipfs(), target)
	if err != nil {
		return nil, err
	}

	keys, err := m.node.TargetNodeKeys(node)
	if err != nil {
		return nil, err
	}

	return thrd.AddFiles(node, caption, keys.Files)
}

func (m *Mobile) buildDirectory(data []byte, path string, threadId string) (*pb.Directory, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return nil, core.ErrThreadNotFound
	}

	if thrd.Schema == nil {
		return nil, core.ErrThreadSchemaRequired
	}

	var use string
	if data == nil {
		if ref, err := ipfspath.ParsePath(path); err == nil {
			parts := strings.Split(ref.String(), "/")
			use = parts[len(parts)-1]
		}
	}

	dir := &pb.Directory{
		Files: make(map[string]*pb.FileIndex),
	}

	mil, err := getMill(thrd.Schema.Mill, thrd.Schema.Opts)
	if err != nil {
		return nil, err
	}
	if mil != nil {
		conf, err := m.getFileConfig(mil,
			fileConfigOpt.Data(data),
			fileConfigOpt.Path(path),
			fileConfigOpt.Use(use),
			fileConfigOpt.Plaintext(thrd.Schema.Plaintext),
		)
		if err != nil {
			return nil, err
		}

		added, err := m.node.AddFileIndex(mil, *conf)
		if err != nil {
			return nil, err
		}
		dir.Files[schema.SingleFileTag] = added

	} else if len(thrd.Schema.Links) > 0 {

		// determine order
		steps, err := schema.Steps(thrd.Schema.Links)
		if err != nil {
			return nil, err
		}

		// send each link
		for _, step := range steps {
			mil, err := getMill(step.Link.Mill, step.Link.Opts)
			if err != nil {
				return nil, err
			}
			var conf *core.AddFileConfig

			if step.Link.Use == schema.FileTag {
				conf, err = m.getFileConfig(mil,
					fileConfigOpt.Data(data),
					fileConfigOpt.Path(path),
					fileConfigOpt.Use(use),
					fileConfigOpt.Plaintext(step.Link.Plaintext),
				)
				if err != nil {
					return nil, err
				}

			} else {
				if dir.Files[step.Link.Use] == nil {
					return nil, fmt.Errorf(step.Link.Use + " not found")
				}

				conf, err = m.getFileConfig(mil,
					fileConfigOpt.Data(data),
					fileConfigOpt.Path(path),
					fileConfigOpt.Use(dir.Files[step.Link.Use].Hash),
					fileConfigOpt.Plaintext(step.Link.Plaintext),
				)
				if err != nil {
					return nil, err
				}
			}

			added, err := m.node.AddFileIndex(mil, *conf)
			if err != nil {
				return nil, err
			}
			dir.Files[step.Name] = added
		}
	} else {
		return nil, schema.ErrEmptySchema
	}

	return dir, nil
}

func (m *Mobile) getFileConfig(mil mill.Mill, opts ...fileConfigOption) (*core.AddFileConfig, error) {
	var reader io.ReadSeeker
	conf := &core.AddFileConfig{}
	settings := fileConfigOptions(opts...)

	if settings.Use == "" {
		if settings.Data != nil {
			reader = bytes.NewReader(settings.Data)
		} else {
			f, err := os.Open(settings.Path)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			reader = f

			_, file := filepath.Split(f.Name())
			conf.Name = file
		}
	} else {
		var file *pb.FileIndex
		var err error
		reader, file, err = m.node.FileData(settings.Use)
		if err != nil {
			return nil, err
		}

		conf.Name = file.Name
		conf.Use = file.Checksum
	}

	var err error
	if mil.ID() == "/json" {
		conf.Media = "application/json"
	} else {
		conf.Media, err = m.node.GetMedia(reader, mil)
		if err != nil {
			return nil, err
		}
	}
	_, _ = reader.Seek(0, 0)

	input, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	conf.Input = input
	conf.Plaintext = settings.Plaintext

	return conf, nil
}

func getMill(id string, opts map[string]string) (mill.Mill, error) {
	switch id {
	case "/blob":
		return &mill.Blob{}, nil
	case "/image/resize":
		width := opts["width"]
		if width == "" {
			return nil, fmt.Errorf("missing width")
		}
		quality := opts["quality"]
		if quality == "" {
			quality = "75"
		}
		return &mill.ImageResize{
			Opts: mill.ImageResizeOpts{
				Width:   width,
				Quality: quality,
			},
		}, nil
	case "/image/exif":
		return &mill.ImageExif{}, nil
	case "/json":
		return &mill.Json{}, nil
	default:
		return nil, nil
	}
}

func (m *Mobile) writeFiles(dirs *pb.DirectoryList, threadId string, caption string) (mh.Multihash, error) {
	if !m.node.Started() {
		return nil, core.ErrStopped
	}

	if len(dirs.Items) == 0 || len(dirs.Items[0].Files) == 0 {
		return nil, fmt.Errorf("no files found")
	}

	thrd := m.node.Thread(threadId)
	if thrd == nil {
		return nil, core.ErrThreadNotFound
	}

	var node ipld.Node
	var keys *pb.Keys

	var err error
	file := dirs.Items[0].Files[schema.SingleFileTag]
	if file != nil {
		node, keys, err = m.node.AddNodeFromFiles([]*pb.FileIndex{file})
	} else {
		node, keys, err = m.node.AddNodeFromDirs(dirs)
	}
	if err != nil {
		return nil, err
	}

	if node == nil {
		return nil, fmt.Errorf("no files found")
	}

	return thrd.AddFiles(node, caption, keys.Files)
}
