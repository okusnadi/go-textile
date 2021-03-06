package util

import (
	"io"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("tex-util")

func UnmarshalString(body io.ReadCloser) (string, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	return TrimQuotes(string(data)), nil
}

func SplitString(in string, sep string) []string {
	list := make([]string, 0)
	for _, s := range strings.Split(in, sep) {
		t := strings.TrimSpace(s)
		if t != "" {
			list = append(list, t)
		}
	}
	return list
}

func ListContainsString(list []string, i string) bool {
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return false
}

func ProtoTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func ProtoNanos(ts *timestamp.Timestamp) int64 {
	if ts == nil {
		ts = ptypes.TimestampNow()
	}
	return int64(ts.Nanos) + ts.Seconds*1e9
}

func ProtoTs(nsec int64) *timestamp.Timestamp {
	n := nsec / 1e9
	sec := n
	nsec -= n * 1e9
	if nsec < 0 {
		nsec += 1e9
		sec--
	}

	return &timestamp.Timestamp{
		Seconds: sec,
		Nanos:   int32(nsec),
	}
}

func ProtoTsIsNewer(ts1 *timestamp.Timestamp, ts2 *timestamp.Timestamp) bool {
	return ProtoNanos(ts1) > ProtoNanos(ts2)
}

func TrimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

func LogMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Infof("Alloc = %v MiB", bToMb(m.Alloc))
	log.Infof("TotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	log.Infof("Sys = %v MiB", bToMb(m.Sys))
	log.Infof("NumGC = %v", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
