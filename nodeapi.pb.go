// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeapi.proto

package synerex_nodeapi // import "github.com/synerex/synerex_nodeapi"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeInfo struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	IsServer             bool     `protobuf:"varint,2,opt,name=is_server,json=isServer,proto3" json:"is_server,omitempty"`
	NodePbaseVersion     string   `protobuf:"bytes,3,opt,name=node_pbase_version,json=nodePbaseVersion,proto3" json:"node_pbase_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_17118c58b937a8b6, []int{0}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (dst *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(dst, src)
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

func (m *NodeInfo) GetIsServer() bool {
	if m != nil {
		return m.IsServer
	}
	return false
}

func (m *NodeInfo) GetNodePbaseVersion() string {
	if m != nil {
		return m.NodePbaseVersion
	}
	return ""
}

type NodeID struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64   `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	KeepaliveDuration    int32    `protobuf:"varint,3,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeID) Reset()         { *m = NodeID{} }
func (m *NodeID) String() string { return proto.CompactTextString(m) }
func (*NodeID) ProtoMessage()    {}
func (*NodeID) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_17118c58b937a8b6, []int{1}
}
func (m *NodeID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeID.Unmarshal(m, b)
}
func (m *NodeID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeID.Marshal(b, m, deterministic)
}
func (dst *NodeID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeID.Merge(dst, src)
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

func (m *NodeID) GetKeepaliveDuration() int32 {
	if m != nil {
		return m.KeepaliveDuration
	}
	return 0
}

type NodeUpdate struct {
	NodeId               int32    `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Secret               uint64   `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`
	UpdateCount          int32    `protobuf:"varint,3,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"`
	NodeStatus           int32    `protobuf:"varint,4,opt,name=node_status,json=nodeStatus,proto3" json:"node_status,omitempty"`
	NodeArg              string   `protobuf:"bytes,5,opt,name=node_arg,json=nodeArg,proto3" json:"node_arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUpdate) Reset()         { *m = NodeUpdate{} }
func (m *NodeUpdate) String() string { return proto.CompactTextString(m) }
func (*NodeUpdate) ProtoMessage()    {}
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_17118c58b937a8b6, []int{2}
}
func (m *NodeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeUpdate.Unmarshal(m, b)
}
func (m *NodeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeUpdate.Marshal(b, m, deterministic)
}
func (dst *NodeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUpdate.Merge(dst, src)
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

type Response struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodeapi_17118c58b937a8b6, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
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

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "nodeapi.NodeInfo")
	proto.RegisterType((*NodeID)(nil), "nodeapi.NodeID")
	proto.RegisterType((*NodeUpdate)(nil), "nodeapi.NodeUpdate")
	proto.RegisterType((*Response)(nil), "nodeapi.Response")
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

// NodeServer is the server API for Node service.
type NodeServer interface {
	RegisterNode(context.Context, *NodeInfo) (*NodeID, error)
	QueryNode(context.Context, *NodeID) (*NodeInfo, error)
	KeepAlive(context.Context, *NodeUpdate) (*Response, error)
	UnRegisterNode(context.Context, *NodeID) (*Response, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeapi.proto",
}

func init() { proto.RegisterFile("nodeapi.proto", fileDescriptor_nodeapi_17118c58b937a8b6) }

var fileDescriptor_nodeapi_17118c58b937a8b6 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0xa6, 0xcd, 0xb2, 0x3b, 0x2d, 0xa5, 0x1d, 0x24, 0x58, 0xca, 0x81, 0x62, 0x71, 0xe8,
	0xa1, 0x04, 0xf1, 0x79, 0x2f, 0xe4, 0x52, 0x21, 0x55, 0xe0, 0xaa, 0x1c, 0xb8, 0x58, 0x4e, 0x76,
	0x48, 0xad, 0x10, 0x7b, 0x65, 0x7b, 0x23, 0xfa, 0x5f, 0xf8, 0x6d, 0xfc, 0x16, 0xe4, 0xc9, 0x2e,
	0x10, 0x90, 0x90, 0x7a, 0x4a, 0xde, 0x7b, 0x7e, 0xf3, 0xc6, 0xcf, 0x0b, 0xb7, 0xad, 0xab, 0x49,
	0x37, 0x66, 0xdc, 0x78, 0x17, 0x1d, 0xde, 0xea, 0xa0, 0xf0, 0x50, 0x9c, 0xbb, 0x9a, 0xce, 0xec,
	0x17, 0x87, 0x0f, 0xa1, 0x4c, 0xb4, 0xb2, 0x7a, 0x49, 0x55, 0x76, 0x94, 0x1d, 0x97, 0xb2, 0x48,
	0xc4, 0xb9, 0x5e, 0x52, 0x12, 0x4d, 0x50, 0x81, 0xfc, 0x8a, 0x7c, 0x35, 0x3c, 0xca, 0x8e, 0x0b,
	0x59, 0x98, 0x70, 0xc1, 0x18, 0x4f, 0x00, 0xd9, 0xd9, 0x4c, 0x75, 0x20, 0xb5, 0x22, 0x1f, 0x8c,
	0xb3, 0xd5, 0x16, 0x8f, 0xd8, 0x4f, 0xca, 0x87, 0x24, 0x7c, 0x5a, 0xf3, 0xe2, 0x0a, 0x72, 0xce,
	0x9c, 0xe0, 0x7d, 0xe0, 0x45, 0x94, 0xa9, 0x39, 0x6f, 0x24, 0xf3, 0x04, 0xcf, 0x6a, 0xbc, 0x07,
	0x79, 0xa0, 0x99, 0xa7, 0xc8, 0x51, 0xb9, 0xec, 0x10, 0x3e, 0x05, 0x5c, 0x10, 0x35, 0xfa, 0xab,
	0x59, 0x91, 0xaa, 0x5b, 0xaf, 0x63, 0x1f, 0x34, 0x92, 0x07, 0xbf, 0x94, 0x49, 0x27, 0x88, 0xef,
	0x19, 0x40, 0x8a, 0xba, 0x6c, 0x6a, 0x1d, 0xe9, 0xe6, 0x71, 0x8f, 0x61, 0xb7, 0x65, 0xab, 0x9a,
	0xb9, 0xd6, 0xc6, 0x2e, 0x68, 0x67, 0xcd, 0xbd, 0x4b, 0x14, 0x3e, 0x82, 0x1d, 0x9e, 0x19, 0xa2,
	0x8e, 0x6d, 0xa8, 0xb6, 0xf9, 0x04, 0x24, 0xea, 0x82, 0x19, 0x7c, 0x00, 0x5c, 0xa2, 0xd2, 0x7e,
	0x5e, 0x8d, 0xb8, 0x11, 0x5e, 0xe2, 0xd4, 0xcf, 0xc5, 0x09, 0x14, 0x92, 0x42, 0xe3, 0x6c, 0x20,
	0xdc, 0x83, 0xa1, 0x5b, 0xf0, 0x5a, 0x85, 0x1c, 0xba, 0x05, 0xee, 0xc3, 0x16, 0xf9, 0x75, 0xd3,
	0xa5, 0x4c, 0x7f, 0x5f, 0xfc, 0xc8, 0x60, 0x3b, 0x5d, 0x06, 0x5f, 0xc1, 0xae, 0xa4, 0xb9, 0x09,
	0x91, 0x3c, 0xe3, 0x83, 0x71, 0xff, 0xb8, 0xfd, 0x53, 0x1e, 0xde, 0xd9, 0xa4, 0x26, 0x62, 0x80,
	0xcf, 0xa1, 0xfc, 0xd8, 0x92, 0xbf, 0x66, 0xcb, 0xdf, 0xfa, 0xe1, 0xbf, 0x33, 0xc4, 0x00, 0x5f,
	0x43, 0xf9, 0x9e, 0xa8, 0x39, 0x4d, 0x9d, 0xe2, 0xdd, 0x8d, 0x13, 0xeb, 0x46, 0xff, 0xb0, 0xf5,
	0x17, 0x11, 0x03, 0x7c, 0x03, 0x7b, 0x97, 0x76, 0x63, 0xc3, 0xff, 0xc4, 0xfd, 0xf6, 0xbd, 0x7d,
	0xf2, 0x59, 0xcc, 0x4d, 0xbc, 0x6a, 0xa7, 0xe3, 0x99, 0x5b, 0x3e, 0x0b, 0xd7, 0x96, 0x3c, 0x7d,
	0xeb, 0x7f, 0x55, 0x67, 0x98, 0xe6, 0xfc, 0x05, 0xbf, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x4c, 0x4f, 0xbe, 0xd2, 0x02, 0x00, 0x00,
}
