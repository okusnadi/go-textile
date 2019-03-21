// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryOptions_FilterType int32

const (
	QueryOptions_NO_FILTER  QueryOptions_FilterType = 0
	QueryOptions_HIDE_OLDER QueryOptions_FilterType = 1
)

var QueryOptions_FilterType_name = map[int32]string{
	0: "NO_FILTER",
	1: "HIDE_OLDER",
}
var QueryOptions_FilterType_value = map[string]int32{
	"NO_FILTER":  0,
	"HIDE_OLDER": 1,
}

func (x QueryOptions_FilterType) String() string {
	return proto.EnumName(QueryOptions_FilterType_name, int32(x))
}
func (QueryOptions_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{0, 0}
}

type Query_Type int32

const (
	Query_THREAD_SNAPSHOTS Query_Type = 0
	Query_CONTACTS         Query_Type = 1
)

var Query_Type_name = map[int32]string{
	0: "THREAD_SNAPSHOTS",
	1: "CONTACTS",
}
var Query_Type_value = map[string]int32{
	"THREAD_SNAPSHOTS": 0,
	"CONTACTS":         1,
}

func (x Query_Type) String() string {
	return proto.EnumName(Query_Type_name, int32(x))
}
func (Query_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{1, 0}
}

type PubSubQuery_ResponseType int32

const (
	PubSubQuery_P2P    PubSubQuery_ResponseType = 0
	PubSubQuery_PUBSUB PubSubQuery_ResponseType = 1
)

var PubSubQuery_ResponseType_name = map[int32]string{
	0: "P2P",
	1: "PUBSUB",
}
var PubSubQuery_ResponseType_value = map[string]int32{
	"P2P":    0,
	"PUBSUB": 1,
}

func (x PubSubQuery_ResponseType) String() string {
	return proto.EnumName(PubSubQuery_ResponseType_name, int32(x))
}
func (PubSubQuery_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{2, 0}
}

type QueryOptions struct {
	Local                bool                    `protobuf:"varint,1,opt,name=local,proto3" json:"local,omitempty"`
	Limit                int32                   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Wait                 int32                   `protobuf:"varint,3,opt,name=wait,proto3" json:"wait,omitempty"`
	Filter               QueryOptions_FilterType `protobuf:"varint,4,opt,name=filter,proto3,enum=QueryOptions_FilterType" json:"filter,omitempty"`
	Exclude              []string                `protobuf:"bytes,5,rep,name=exclude,proto3" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QueryOptions) Reset()         { *m = QueryOptions{} }
func (m *QueryOptions) String() string { return proto.CompactTextString(m) }
func (*QueryOptions) ProtoMessage()    {}
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{0}
}
func (m *QueryOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOptions.Unmarshal(m, b)
}
func (m *QueryOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOptions.Marshal(b, m, deterministic)
}
func (dst *QueryOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOptions.Merge(dst, src)
}
func (m *QueryOptions) XXX_Size() int {
	return xxx_messageInfo_QueryOptions.Size(m)
}
func (m *QueryOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOptions proto.InternalMessageInfo

func (m *QueryOptions) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *QueryOptions) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryOptions) GetWait() int32 {
	if m != nil {
		return m.Wait
	}
	return 0
}

func (m *QueryOptions) GetFilter() QueryOptions_FilterType {
	if m != nil {
		return m.Filter
	}
	return QueryOptions_NO_FILTER
}

func (m *QueryOptions) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type Query struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string        `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Type                 Query_Type    `protobuf:"varint,3,opt,name=type,proto3,enum=Query_Type" json:"type,omitempty"`
	Options              *QueryOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	Payload              *any.Any      `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Query) GetType() Query_Type {
	if m != nil {
		return m.Type
	}
	return Query_THREAD_SNAPSHOTS
}

func (m *Query) GetOptions() *QueryOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Query) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PubSubQuery struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Query_Type               `protobuf:"varint,2,opt,name=type,proto3,enum=Query_Type" json:"type,omitempty"`
	Payload              *any.Any                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	ResponseType         PubSubQuery_ResponseType `protobuf:"varint,4,opt,name=responseType,proto3,enum=PubSubQuery_ResponseType" json:"responseType,omitempty"`
	Exclude              []string                 `protobuf:"bytes,5,rep,name=exclude,proto3" json:"exclude,omitempty"`
	Topic                string                   `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PubSubQuery) Reset()         { *m = PubSubQuery{} }
func (m *PubSubQuery) String() string { return proto.CompactTextString(m) }
func (*PubSubQuery) ProtoMessage()    {}
func (*PubSubQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{2}
}
func (m *PubSubQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubQuery.Unmarshal(m, b)
}
func (m *PubSubQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubQuery.Marshal(b, m, deterministic)
}
func (dst *PubSubQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubQuery.Merge(dst, src)
}
func (m *PubSubQuery) XXX_Size() int {
	return xxx_messageInfo_PubSubQuery.Size(m)
}
func (m *PubSubQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubQuery proto.InternalMessageInfo

func (m *PubSubQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PubSubQuery) GetType() Query_Type {
	if m != nil {
		return m.Type
	}
	return Query_THREAD_SNAPSHOTS
}

func (m *PubSubQuery) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PubSubQuery) GetResponseType() PubSubQuery_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return PubSubQuery_P2P
}

func (m *PubSubQuery) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

func (m *PubSubQuery) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type QueryResult struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Local                bool                 `protobuf:"varint,3,opt,name=local,proto3" json:"local,omitempty"`
	Value                *any.Any             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryResult) Reset()         { *m = QueryResult{} }
func (m *QueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()    {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{3}
}
func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResult.Unmarshal(m, b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
}
func (dst *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(dst, src)
}
func (m *QueryResult) XXX_Size() int {
	return xxx_messageInfo_QueryResult.Size(m)
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryResult) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *QueryResult) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *QueryResult) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type QueryResults struct {
	Type                 Query_Type     `protobuf:"varint,1,opt,name=type,proto3,enum=Query_Type" json:"type,omitempty"`
	Items                []*QueryResult `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueryResults) Reset()         { *m = QueryResults{} }
func (m *QueryResults) String() string { return proto.CompactTextString(m) }
func (*QueryResults) ProtoMessage()    {}
func (*QueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{4}
}
func (m *QueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResults.Unmarshal(m, b)
}
func (m *QueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResults.Marshal(b, m, deterministic)
}
func (dst *QueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResults.Merge(dst, src)
}
func (m *QueryResults) XXX_Size() int {
	return xxx_messageInfo_QueryResults.Size(m)
}
func (m *QueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResults proto.InternalMessageInfo

func (m *QueryResults) GetType() Query_Type {
	if m != nil {
		return m.Type
	}
	return Query_THREAD_SNAPSHOTS
}

func (m *QueryResults) GetItems() []*QueryResult {
	if m != nil {
		return m.Items
	}
	return nil
}

type PubSubQueryResults struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Results              *QueryResults `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PubSubQueryResults) Reset()         { *m = PubSubQueryResults{} }
func (m *PubSubQueryResults) String() string { return proto.CompactTextString(m) }
func (*PubSubQueryResults) ProtoMessage()    {}
func (*PubSubQueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{5}
}
func (m *PubSubQueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubQueryResults.Unmarshal(m, b)
}
func (m *PubSubQueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubQueryResults.Marshal(b, m, deterministic)
}
func (dst *PubSubQueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubQueryResults.Merge(dst, src)
}
func (m *PubSubQueryResults) XXX_Size() int {
	return xxx_messageInfo_PubSubQueryResults.Size(m)
}
func (m *PubSubQueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubQueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubQueryResults proto.InternalMessageInfo

func (m *PubSubQueryResults) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PubSubQueryResults) GetResults() *QueryResults {
	if m != nil {
		return m.Results
	}
	return nil
}

type ContactQuery struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactQuery) Reset()         { *m = ContactQuery{} }
func (m *ContactQuery) String() string { return proto.CompactTextString(m) }
func (*ContactQuery) ProtoMessage()    {}
func (*ContactQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{6}
}
func (m *ContactQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactQuery.Unmarshal(m, b)
}
func (m *ContactQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactQuery.Marshal(b, m, deterministic)
}
func (dst *ContactQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactQuery.Merge(dst, src)
}
func (m *ContactQuery) XXX_Size() int {
	return xxx_messageInfo_ContactQuery.Size(m)
}
func (m *ContactQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContactQuery proto.InternalMessageInfo

func (m *ContactQuery) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactQuery) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ThreadSnapshotQuery struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadSnapshotQuery) Reset()         { *m = ThreadSnapshotQuery{} }
func (m *ThreadSnapshotQuery) String() string { return proto.CompactTextString(m) }
func (*ThreadSnapshotQuery) ProtoMessage()    {}
func (*ThreadSnapshotQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_2523d4da09040998, []int{7}
}
func (m *ThreadSnapshotQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadSnapshotQuery.Unmarshal(m, b)
}
func (m *ThreadSnapshotQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadSnapshotQuery.Marshal(b, m, deterministic)
}
func (dst *ThreadSnapshotQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadSnapshotQuery.Merge(dst, src)
}
func (m *ThreadSnapshotQuery) XXX_Size() int {
	return xxx_messageInfo_ThreadSnapshotQuery.Size(m)
}
func (m *ThreadSnapshotQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadSnapshotQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadSnapshotQuery proto.InternalMessageInfo

func (m *ThreadSnapshotQuery) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryOptions)(nil), "QueryOptions")
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterType((*PubSubQuery)(nil), "PubSubQuery")
	proto.RegisterType((*QueryResult)(nil), "QueryResult")
	proto.RegisterType((*QueryResults)(nil), "QueryResults")
	proto.RegisterType((*PubSubQueryResults)(nil), "PubSubQueryResults")
	proto.RegisterType((*ContactQuery)(nil), "ContactQuery")
	proto.RegisterType((*ThreadSnapshotQuery)(nil), "ThreadSnapshotQuery")
	proto.RegisterEnum("QueryOptions_FilterType", QueryOptions_FilterType_name, QueryOptions_FilterType_value)
	proto.RegisterEnum("Query_Type", Query_Type_name, Query_Type_value)
	proto.RegisterEnum("PubSubQuery_ResponseType", PubSubQuery_ResponseType_name, PubSubQuery_ResponseType_value)
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_query_2523d4da09040998) }

var fileDescriptor_query_2523d4da09040998 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x4e, 0x9c, 0x90, 0xeb, 0x24, 0xb2, 0x06, 0x16, 0x86, 0x0d, 0x91, 0xbf, 0x05, 0x11,
	0x9f, 0x34, 0x54, 0xe9, 0xba, 0x8b, 0x40, 0x82, 0x40, 0xa2, 0x24, 0x1d, 0x9b, 0x4d, 0x37, 0x68,
	0x1c, 0x0f, 0x30, 0xaa, 0xe3, 0x71, 0xed, 0x71, 0x4b, 0xb6, 0x7d, 0x81, 0x3e, 0x56, 0x1f, 0xa0,
	0x2f, 0x54, 0x65, 0xc6, 0x06, 0xf3, 0x13, 0xb5, 0xbb, 0x5c, 0x9f, 0x93, 0xe3, 0x73, 0xce, 0xbd,
	0x06, 0xfb, 0x6b, 0xc1, 0xb2, 0x15, 0x4e, 0x33, 0x21, 0xc5, 0xfe, 0xde, 0x9d, 0x10, 0x77, 0x31,
	0x3b, 0x56, 0x53, 0x58, 0xdc, 0x1e, 0xd3, 0xa4, 0x82, 0x0e, 0x5e, 0x42, 0x92, 0x2f, 0x59, 0x2e,
	0xe9, 0x32, 0xd5, 0x04, 0xef, 0x97, 0x01, 0xdd, 0x4f, 0x6b, 0xad, 0x59, 0x2a, 0xb9, 0x48, 0x72,
	0xb4, 0x0b, 0x56, 0x2c, 0x16, 0x34, 0x76, 0x8d, 0x81, 0x31, 0xdc, 0x26, 0x7a, 0x50, 0x4f, 0xf9,
	0x92, 0x4b, 0xd7, 0x1c, 0x18, 0x43, 0x8b, 0xe8, 0x01, 0x21, 0x68, 0x7e, 0xa7, 0x5c, 0xba, 0x0d,
	0xf5, 0x50, 0xfd, 0x46, 0xef, 0xa0, 0x75, 0xcb, 0x63, 0xc9, 0x32, 0xb7, 0x39, 0x30, 0x86, 0xfd,
	0x91, 0x8b, 0xeb, 0xf2, 0xf8, 0x4c, 0x61, 0xc1, 0x2a, 0x65, 0xa4, 0xe4, 0x21, 0x17, 0xda, 0xec,
	0x61, 0x11, 0x17, 0x11, 0x73, 0xad, 0x41, 0x63, 0xd8, 0x21, 0xd5, 0xe8, 0xfd, 0x0f, 0xf0, 0xc4,
	0x47, 0x3d, 0xe8, 0x5c, 0xcd, 0x6e, 0xce, 0x2e, 0x2e, 0x83, 0x29, 0x71, 0xb6, 0x50, 0x1f, 0xe0,
	0xfc, 0x62, 0x32, 0xbd, 0x99, 0x5d, 0x4e, 0xa6, 0xc4, 0x31, 0xbc, 0xdf, 0x06, 0x58, 0xea, 0x55,
	0xa8, 0x0f, 0x26, 0x8f, 0x94, 0xff, 0x0e, 0x31, 0x79, 0xb4, 0x36, 0x2f, 0xc5, 0x17, 0x96, 0x28,
	0xf3, 0x1d, 0xa2, 0x07, 0x74, 0x00, 0x4d, 0xb9, 0x4a, 0x99, 0x32, 0xdf, 0x1f, 0xd9, 0xda, 0x26,
	0x56, 0xce, 0x14, 0x80, 0x0e, 0xa1, 0x2d, 0xb4, 0x6b, 0x15, 0xc5, 0x1e, 0xf5, 0x9e, 0x45, 0x21,
	0x15, 0x8a, 0x30, 0xb4, 0x53, 0xba, 0x8a, 0x05, 0x8d, 0x5c, 0x4b, 0x11, 0x77, 0xb1, 0xae, 0x1d,
	0x57, 0xb5, 0xe3, 0x71, 0xb2, 0x22, 0x15, 0xc9, 0x3b, 0x82, 0xa6, 0x0a, 0xb4, 0x0b, 0x4e, 0x70,
	0x4e, 0xa6, 0xe3, 0xc9, 0x8d, 0x7f, 0x35, 0x9e, 0xfb, 0xe7, 0xb3, 0xc0, 0x77, 0xb6, 0x50, 0x17,
	0xb6, 0x4f, 0x67, 0x57, 0xc1, 0xf8, 0x34, 0xf0, 0x1d, 0xc3, 0xfb, 0x61, 0x82, 0x3d, 0x2f, 0x42,
	0xbf, 0x08, 0xdf, 0xce, 0x56, 0xa5, 0x30, 0x37, 0xa5, 0xa8, 0x99, 0x6b, 0xfc, 0x83, 0x39, 0xf4,
	0x01, 0xba, 0x19, 0xcb, 0x53, 0x91, 0xe4, 0x6c, 0xad, 0x52, 0x6e, 0x71, 0x0f, 0xd7, 0x4c, 0x60,
	0x52, 0x23, 0x90, 0x67, 0xf4, 0xcd, 0xcb, 0xd4, 0x5b, 0x48, 0xf9, 0xc2, 0x6d, 0x55, 0x5b, 0x48,
	0xf9, 0xc2, 0xfb, 0x0f, 0xba, 0x75, 0x35, 0xd4, 0x86, 0xc6, 0x7c, 0x34, 0x77, 0xb6, 0x10, 0x40,
	0x6b, 0x7e, 0x7d, 0xe2, 0x5f, 0x9f, 0x38, 0x86, 0xf7, 0xd3, 0x00, 0x5b, 0xbd, 0x99, 0xb0, 0xbc,
	0x88, 0xe5, 0xab, 0x12, 0x30, 0x34, 0x23, 0x2a, 0x75, 0x09, 0xf6, 0x68, 0xff, 0x55, 0xc0, 0xa0,
	0x3a, 0x7a, 0xa2, 0x78, 0x4f, 0x37, 0xde, 0xa8, 0xdf, 0xf8, 0x11, 0x58, 0xdf, 0x68, 0x5c, 0xb0,
	0x72, 0xdb, 0x6f, 0xf7, 0xa4, 0x29, 0x9e, 0x5f, 0x7e, 0x35, 0xda, 0x50, 0xfe, 0xb8, 0x06, 0x63,
	0xd3, 0x1a, 0x3c, 0xb0, 0xb8, 0x64, 0xcb, 0xdc, 0x35, 0x07, 0x8d, 0xa1, 0x3d, 0xea, 0xe2, 0xda,
	0xdf, 0x89, 0x86, 0xbc, 0x8f, 0x80, 0x6a, 0x2d, 0x57, 0xd2, 0x2f, 0xc3, 0x1e, 0x42, 0x3b, 0xd3,
	0x50, 0x99, 0xb7, 0x57, 0xd7, 0xca, 0x49, 0x85, 0x7a, 0x13, 0xe8, 0x9e, 0x8a, 0x44, 0xd2, 0x85,
	0xd4, 0xa7, 0xe3, 0x42, 0x9b, 0x46, 0x51, 0xc6, 0xf2, 0xbc, 0x54, 0xab, 0x46, 0xb4, 0x0f, 0xdb,
	0x45, 0xce, 0xb2, 0x84, 0x2e, 0x59, 0xf9, 0x8d, 0x3c, 0xce, 0xde, 0x31, 0xec, 0x04, 0xf7, 0x19,
	0xa3, 0x91, 0x9f, 0xd0, 0x34, 0xbf, 0x17, 0x7f, 0x13, 0x3b, 0xd9, 0x81, 0x1e, 0x17, 0x58, 0xb2,
	0x07, 0xc9, 0xd7, 0xe5, 0x85, 0x9f, 0xcd, 0x34, 0x0c, 0x5b, 0xaa, 0xc4, 0xf7, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x3f, 0x48, 0x41, 0xb8, 0x04, 0x00, 0x00,
}
