// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeapi.proto

package synerex_nodeapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeType int32

const (
	NodeType_PROVIDER NodeType = 0
	NodeType_SERVER   NodeType = 1
	NodeType_GATEWAY  NodeType = 2
)

var NodeType_name = map[int32]string{
	0: "PROVIDER",
	1: "SERVER",
	2: "GATEWAY",
}

var NodeType_value = map[string]int32{
	"PROVIDER": 0,
	"SERVER":   1,
	"GATEWAY":  2,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}

func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{0}
}

type KeepAliveCommand int32

const (
	KeepAliveCommand_NONE          KeepAliveCommand = 0
	KeepAliveCommand_RECONNECT     KeepAliveCommand = 1
	KeepAliveCommand_SERVER_CHANGE KeepAliveCommand = 2
)

var KeepAliveCommand_name = map[int32]string{
	0: "NONE",
	1: "RECONNECT",
	2: "SERVER_CHANGE",
}

var KeepAliveCommand_value = map[string]int32{
	"NONE":          0,
	"RECONNECT":     1,
	"SERVER_CHANGE": 2,
}

func (x KeepAliveCommand) String() string {
	return proto.EnumName(KeepAliveCommand_name, int32(x))
}

func (KeepAliveCommand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{1}
}

// information for synerex servers and providers, gateways (nodes)
type NodeInfo struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NodeType             NodeType `protobuf:"varint,2,opt,name=node_type,json=nodeType,proto3,enum=nodeapi.NodeType" json:"node_type,omitempty"`
	ServerInfo           string   `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	NodePbaseVersion     string   `protobuf:"bytes,4,opt,name=node_pbase_version,json=nodePbaseVersion,proto3" json:"node_pbase_version,omitempty"`
	WithNodeId           int32    `protobuf:"varint,5,opt,name=with_node_id,json=withNodeId,proto3" json:"with_node_id,omitempty"`
	ClusterId            int32    `protobuf:"varint,6,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	AreaId               string   `protobuf:"bytes,7,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	ChannelTypes         []uint32 `protobuf:"varint,8,rep,packed,name=channelTypes,proto3" json:"channelTypes,omitempty"`
	GwInfo               string   `protobuf:"bytes,9,opt,name=gw_info,json=gwInfo,proto3" json:"gw_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{0}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeInfo) GetNodeType() NodeType {
	if m != nil {
		return m.NodeType
	}
	return NodeType_PROVIDER
}

func (m *NodeInfo) GetServerInfo() string {
	if m != nil {
		return m.ServerInfo
	}
	return ""
}

func (m *NodeInfo) GetNodePbaseVersion() string {
	if m != nil {
		return m.NodePbaseVersion
	}
	return ""
}

func (m *NodeInfo) GetWithNodeId() int32 {
	if m != nil {
		return m.WithNodeId
	}
	return 0
}

func (m *NodeInfo) GetClusterId() int32 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *NodeInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *NodeInfo) GetChannelTypes() []uint32 {
	if m != nil {
		return m.ChannelTypes
	}
	return nil
}

func (m *NodeInfo) GetGwInfo() string {
	if m != nil {
		return m.GwInfo
	}
	return ""
}

type NodeID struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64   `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	ServerInfo           string   `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	KeepaliveDuration    int32    `protobuf:"varint,4,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{1}
}

func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (m *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(m, src)
}
func (m *NodeID) XXX_Size() int {
	return xxx_messageInfo_NodeID.Size(m)
}
func (m *NodeID) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeID.DiscardUnknown(m)
}

var xxx_messageInfo_NodeID proto.InternalMessageInfo

func (m *NodeID) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeID) GetSecret() uint64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

func (m *NodeID) GetServerInfo() string {
	if m != nil {
		return m.ServerInfo
	}
	return ""
}

func (m *NodeID) GetKeepaliveDuration() int32 {
	if m != nil {
		return m.KeepaliveDuration
	}
	return 0
}

type ServerStatus struct {
	Cpu                  float64  `protobuf:"fixed64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               float64  `protobuf:"fixed64,2,opt,name=memory,proto3" json:"memory,omitempty"`
	MsgCount             uint64   `protobuf:"varint,3,opt,name=msg_count,json=msgCount,proto3" json:"msg_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStatus) Reset()         { *m = ServerStatus{} }
func (m *ServerStatus) String() string { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()    {}
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{2}
}

func (m *ServerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatus.Unmarshal(m, b)
}
func (m *ServerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatus.Marshal(b, m, deterministic)
}
func (m *ServerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatus.Merge(m, src)
}
func (m *ServerStatus) XXX_Size() int {
	return xxx_messageInfo_ServerStatus.Size(m)
}
func (m *ServerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatus proto.InternalMessageInfo

func (m *ServerStatus) GetCpu() float64 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *ServerStatus) GetMemory() float64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ServerStatus) GetMsgCount() uint64 {
	if m != nil {
		return m.MsgCount
	}
	return 0
}

type NodeUpdate struct {
	NodeId               int32         `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64        `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	UpdateCount          int32         `protobuf:"varint,3,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"`
	NodeStatus           int32         `protobuf:"varint,4,opt,name=node_status,json=nodeStatus,proto3" json:"node_status,omitempty"`
	NodeArg              string        `protobuf:"bytes,5,opt,name=node_arg,json=nodeArg,proto3" json:"node_arg,omitempty"`
	Status               *ServerStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeUpdate) Reset()         { *m = NodeUpdate{} }
func (m *NodeUpdate) String() string { return proto.CompactTextString(m) }
func (*NodeUpdate) ProtoMessage()    {}
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{3}
}

func (m *NodeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUpdate.Unmarshal(m, b)
}
func (m *NodeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUpdate.Marshal(b, m, deterministic)
}
func (m *NodeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUpdate.Merge(m, src)
}
func (m *NodeUpdate) XXX_Size() int {
	return xxx_messageInfo_NodeUpdate.Size(m)
}
func (m *NodeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUpdate proto.InternalMessageInfo

func (m *NodeUpdate) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *NodeUpdate) GetSecret() uint64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

func (m *NodeUpdate) GetUpdateCount() int32 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func (m *NodeUpdate) GetNodeStatus() int32 {
	if m != nil {
		return m.NodeStatus
	}
	return 0
}

func (m *NodeUpdate) GetNodeArg() string {
	if m != nil {
		return m.NodeArg
	}
	return ""
}

func (m *NodeUpdate) GetStatus() *ServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type Response struct {
	Ok                   bool             `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Command              KeepAliveCommand `protobuf:"varint,2,opt,name=command,proto3,enum=nodeapi.KeepAliveCommand" json:"command,omitempty"`
	Err                  string           `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetCommand() KeepAliveCommand {
	if m != nil {
		return m.Command
	}
	return KeepAliveCommand_NONE
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Connection struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{5}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Connection) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type ConnectionInfo struct {
	Connections          []*Connection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConnectionInfo) Reset()         { *m = ConnectionInfo{} }
func (m *ConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*ConnectionInfo) ProtoMessage()    {}
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad9a996d16f0fd5c, []int{6}
}

func (m *ConnectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionInfo.Unmarshal(m, b)
}
func (m *ConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionInfo.Marshal(b, m, deterministic)
}
func (m *ConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionInfo.Merge(m, src)
}
func (m *ConnectionInfo) XXX_Size() int {
	return xxx_messageInfo_ConnectionInfo.Size(m)
}
func (m *ConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionInfo proto.InternalMessageInfo

func (m *ConnectionInfo) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

func init() {
	proto.RegisterEnum("nodeapi.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("nodeapi.KeepAliveCommand", KeepAliveCommand_name, KeepAliveCommand_value)
	proto.RegisterType((*NodeInfo)(nil), "nodeapi.NodeInfo")
	proto.RegisterType((*NodeID)(nil), "nodeapi.NodeID")
	proto.RegisterType((*ServerStatus)(nil), "nodeapi.ServerStatus")
	proto.RegisterType((*NodeUpdate)(nil), "nodeapi.NodeUpdate")
	proto.RegisterType((*Response)(nil), "nodeapi.Response")
	proto.RegisterType((*Connection)(nil), "nodeapi.Connection")
	proto.RegisterType((*ConnectionInfo)(nil), "nodeapi.ConnectionInfo")
}

func init() { proto.RegisterFile("nodeapi.proto", fileDescriptor_ad9a996d16f0fd5c) }

var fileDescriptor_ad9a996d16f0fd5c = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x8f, 0xd3, 0xd6, 0xb1, 0x27, 0x7f, 0x70, 0xf7, 0xc4, 0x9d, 0x2f, 0x27, 0x44, 0xb0, 0xf8,
	0x10, 0x9d, 0xb8, 0x54, 0x97, 0xe3, 0x10, 0x9f, 0x10, 0x21, 0xb1, 0x4a, 0x84, 0x94, 0x1e, 0xdb,
	0x3f, 0x08, 0xbe, 0x58, 0x8e, 0x3d, 0x75, 0xac, 0xc6, 0xbb, 0xd6, 0xda, 0x4e, 0xc9, 0x23, 0xf0,
	0x50, 0x3c, 0x01, 0x0f, 0xc1, 0xab, 0xa0, 0x5d, 0xdb, 0x69, 0x52, 0x2a, 0xd0, 0x7d, 0xb2, 0xe7,
	0x37, 0xf3, 0x9b, 0xf9, 0xcd, 0xcc, 0xee, 0x42, 0x97, 0xf1, 0x10, 0xfd, 0x34, 0x1e, 0xa5, 0x82,
	0xe7, 0x9c, 0xb4, 0x2a, 0xb3, 0xff, 0x2a, 0xe2, 0x3c, 0x5a, 0xe3, 0x99, 0x82, 0x97, 0xc5, 0xed,
	0x19, 0x26, 0x69, 0xbe, 0x2d, 0xa3, 0x9c, 0x3f, 0x9b, 0x60, 0x2c, 0x78, 0x88, 0x73, 0x76, 0xcb,
	0xc9, 0x2b, 0x30, 0x25, 0xc9, 0x63, 0x7e, 0x82, 0xb6, 0x36, 0xd0, 0x86, 0x26, 0x35, 0x24, 0xb0,
	0xf0, 0x13, 0x24, 0xa3, 0xca, 0x99, 0x6f, 0x53, 0xb4, 0x9b, 0x03, 0x6d, 0xd8, 0x1b, 0x9f, 0x8e,
	0xea, 0x92, 0x32, 0xc5, 0xd5, 0x36, 0xc5, 0x32, 0x5e, 0xfe, 0x91, 0xcf, 0xa1, 0x9d, 0xa1, 0xd8,
	0xa0, 0xf0, 0x62, 0x76, 0xcb, 0xed, 0x23, 0x95, 0x0e, 0x4a, 0x48, 0x55, 0xfb, 0x0a, 0x88, 0x4a,
	0x98, 0x2e, 0xfd, 0x0c, 0xbd, 0x0d, 0x8a, 0x2c, 0xe6, 0xcc, 0x3e, 0x56, 0x71, 0x96, 0xf4, 0x7c,
	0x90, 0x8e, 0x9b, 0x12, 0x27, 0x03, 0xe8, 0xdc, 0xc7, 0xf9, 0xca, 0x53, 0x94, 0x38, 0xb4, 0x4f,
	0x06, 0xda, 0xf0, 0x84, 0x82, 0xc4, 0x94, 0xfe, 0x90, 0x7c, 0x06, 0x10, 0xac, 0x8b, 0x2c, 0x97,
	0x15, 0x43, 0x5b, 0x57, 0x7e, 0xb3, 0x42, 0xe6, 0x21, 0x79, 0x01, 0x2d, 0x5f, 0xa0, 0x2f, 0x7d,
	0x2d, 0x55, 0x43, 0x97, 0xe6, 0x3c, 0x24, 0x0e, 0x74, 0x82, 0x95, 0xcf, 0x18, 0xae, 0xa5, 0xee,
	0xcc, 0x36, 0x06, 0x47, 0xc3, 0x2e, 0x3d, 0xc0, 0x24, 0x39, 0xba, 0x2f, 0x1b, 0x31, 0x4b, 0x72,
	0x74, 0x2f, 0x9b, 0x70, 0xfe, 0xd0, 0x40, 0x57, 0xf5, 0x67, 0x32, 0xa6, 0x16, 0xa7, 0xa9, 0xe2,
	0x3a, 0x2b, 0x85, 0x3d, 0x07, 0x3d, 0xc3, 0x40, 0x60, 0xae, 0xc6, 0xa6, 0xd3, 0xca, 0xfa, 0xff,
	0x09, 0xbd, 0x01, 0x72, 0x87, 0x98, 0xfa, 0xeb, 0x78, 0x83, 0x5e, 0x58, 0x08, 0x3f, 0xaf, 0x27,
	0x74, 0x42, 0x4f, 0x77, 0x9e, 0x59, 0xe5, 0x70, 0xae, 0xa1, 0x73, 0xa9, 0xc8, 0x97, 0xb9, 0x9f,
	0x17, 0x19, 0xb1, 0xe0, 0x28, 0x48, 0x0b, 0x25, 0x46, 0xa3, 0xf2, 0x57, 0x2a, 0x49, 0x30, 0xe1,
	0x62, 0xab, 0x94, 0x68, 0xb4, 0xb2, 0xe4, 0xe2, 0x93, 0x2c, 0xf2, 0x02, 0x5e, 0xb0, 0x5c, 0xe9,
	0x38, 0xa6, 0x46, 0x92, 0x45, 0x53, 0x69, 0x3b, 0x7f, 0x69, 0x00, 0xb2, 0xc5, 0xeb, 0x34, 0xf4,
	0x73, 0xfc, 0xf8, 0x36, 0xbf, 0x80, 0x4e, 0xa1, 0xa8, 0x7b, 0xf9, 0x4f, 0x68, 0xbb, 0xc4, 0x54,
	0x09, 0x39, 0x09, 0x95, 0x33, 0x53, 0xc2, 0xab, 0x0e, 0x41, 0x42, 0x55, 0x2b, 0x2f, 0x41, 0x1d,
	0x2c, 0xcf, 0x17, 0x91, 0xda, 0xbc, 0x49, 0x95, 0x88, 0x89, 0x88, 0xc8, 0x1b, 0xd0, 0x2b, 0x9a,
	0x5c, 0x79, 0x7b, 0xfc, 0xe9, 0xee, 0x50, 0xee, 0x0f, 0x83, 0x56, 0x41, 0x8e, 0x0f, 0x06, 0xc5,
	0x2c, 0xe5, 0x2c, 0x43, 0xd2, 0x83, 0x26, 0xbf, 0x53, 0x5d, 0x18, 0xb4, 0xc9, 0xef, 0xc8, 0x3b,
	0x68, 0x05, 0x3c, 0x49, 0x7c, 0x16, 0x56, 0x07, 0xfc, 0xe5, 0x2e, 0xd7, 0x4f, 0x88, 0xe9, 0x44,
	0x4e, 0x7b, 0x5a, 0x06, 0xd0, 0x3a, 0x52, 0x4e, 0x19, 0x85, 0xa8, 0xb6, 0x27, 0x7f, 0x9d, 0xef,
	0x01, 0xa6, 0x9c, 0x31, 0x0c, 0xe4, 0x56, 0x48, 0x1f, 0x8c, 0x54, 0xf0, 0x4d, 0x1c, 0xa2, 0xa8,
	0xef, 0x54, 0x6d, 0x97, 0x23, 0x93, 0x22, 0x55, 0x3d, 0x93, 0x56, 0x96, 0x73, 0x0e, 0xbd, 0x87,
	0x0c, 0xea, 0x28, 0xbc, 0x87, 0x76, 0xb0, 0x43, 0x32, 0x5b, 0x1b, 0x1c, 0x0d, 0xdb, 0xe3, 0x67,
	0x3b, 0x79, 0x0f, 0xd1, 0x74, 0x3f, 0xee, 0xf5, 0xdb, 0xf2, 0x76, 0xab, 0x0b, 0xd9, 0x01, 0xe3,
	0x03, 0xbd, 0xb8, 0x99, 0xcf, 0x5c, 0x6a, 0x35, 0x08, 0x80, 0x7e, 0xe9, 0xd2, 0x1b, 0x97, 0x5a,
	0x1a, 0x69, 0x43, 0xeb, 0x7c, 0x72, 0xe5, 0xfe, 0x32, 0xf9, 0xd5, 0x6a, 0xbe, 0xfe, 0x0e, 0xac,
	0xc7, 0xcd, 0x12, 0x03, 0x8e, 0x17, 0x17, 0x0b, 0xd7, 0x6a, 0x90, 0x2e, 0x98, 0xd4, 0x9d, 0x5e,
	0x2c, 0x16, 0xee, 0xf4, 0xca, 0xd2, 0xc8, 0x29, 0x74, 0xcb, 0x2c, 0xde, 0xf4, 0xc7, 0xc9, 0xe2,
	0xdc, 0xb5, 0x9a, 0xe3, 0xbf, 0x9b, 0x70, 0x2c, 0x6b, 0x92, 0xaf, 0xa1, 0x43, 0x31, 0x8a, 0xe5,
	0xf5, 0x53, 0xf6, 0xe1, 0x6b, 0x21, 0xbb, 0xea, 0x7f, 0x72, 0x08, 0xcd, 0x9c, 0x06, 0x79, 0x0b,
	0xe6, 0xcf, 0x05, 0x8a, 0xad, 0xa2, 0x3c, 0xf6, 0xf7, 0xff, 0x9d, 0xc3, 0x69, 0x90, 0xf7, 0x60,
	0xee, 0x14, 0x93, 0x67, 0x07, 0x11, 0xe5, 0x99, 0xdd, 0xa3, 0xd5, 0xbb, 0x77, 0x1a, 0xe4, 0x1b,
	0xe8, 0x5d, 0xb3, 0x03, 0x85, 0xff, 0x51, 0x6e, 0x8f, 0xe7, 0x82, 0xa5, 0x14, 0x3e, 0xcc, 0x3c,
	0x23, 0xcf, 0x47, 0xe5, 0x23, 0x3b, 0xaa, 0x1f, 0xd9, 0x91, 0x2b, 0x1f, 0xd9, 0xfe, 0x8b, 0x27,
	0x36, 0x54, 0xa9, 0xfe, 0x16, 0x3a, 0xd3, 0x95, 0xcf, 0x22, 0x2c, 0x8f, 0x29, 0x79, 0x6a, 0x99,
	0x4f, 0x0a, 0xf8, 0xe1, 0xcb, 0xdf, 0x9c, 0x28, 0xce, 0x57, 0xc5, 0x72, 0x14, 0xf0, 0xe4, 0x2c,
	0xdb, 0x32, 0x14, 0xf8, 0x7b, 0xfd, 0xf5, 0x2a, 0xc2, 0x52, 0x57, 0x52, 0xde, 0xfd, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xed, 0x41, 0x5b, 0xb4, 0x17, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error)
	QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error)
	KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error)
	UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error)
	// get connection info (synerex server and provider)
	QueryConnections(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConnectionInfo, error)
	// change relations between provider and synerex server
	ChangeServer(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Response, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error) {
	out := new(NodeID)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/QueryNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/UnRegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) QueryConnections(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConnectionInfo, error) {
	out := new(ConnectionInfo)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/QueryConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ChangeServer(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/ChangeServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	RegisterNode(context.Context, *NodeInfo) (*NodeID, error)
	QueryNode(context.Context, *NodeID) (*NodeInfo, error)
	KeepAlive(context.Context, *NodeUpdate) (*Response, error)
	UnRegisterNode(context.Context, *NodeID) (*Response, error)
	// get connection info (synerex server and provider)
	QueryConnections(context.Context, *empty.Empty) (*ConnectionInfo, error)
	// change relations between provider and synerex server
	ChangeServer(context.Context, *Connection) (*Response, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) RegisterNode(ctx context.Context, req *NodeInfo) (*NodeID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedNodeServer) QueryNode(ctx context.Context, req *NodeID) (*NodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNode not implemented")
}
func (*UnimplementedNodeServer) KeepAlive(ctx context.Context, req *NodeUpdate) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedNodeServer) UnRegisterNode(ctx context.Context, req *NodeID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterNode not implemented")
}
func (*UnimplementedNodeServer) QueryConnections(ctx context.Context, req *empty.Empty) (*ConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryConnections not implemented")
}
func (*UnimplementedNodeServer) ChangeServer(ctx context.Context, req *Connection) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeServer not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).RegisterNode(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_QueryNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).QueryNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/QueryNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).QueryNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).KeepAlive(ctx, req.(*NodeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_UnRegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UnRegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/UnRegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UnRegisterNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_QueryConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).QueryConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/QueryConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).QueryConnections(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ChangeServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ChangeServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/ChangeServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ChangeServer(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeapi.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Node_RegisterNode_Handler,
		},
		{
			MethodName: "QueryNode",
			Handler:    _Node_QueryNode_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Node_KeepAlive_Handler,
		},
		{
			MethodName: "UnRegisterNode",
			Handler:    _Node_UnRegisterNode_Handler,
		},
		{
			MethodName: "QueryConnections",
			Handler:    _Node_QueryConnections_Handler,
		},
		{
			MethodName: "ChangeServer",
			Handler:    _Node_ChangeServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeapi.proto",
}
