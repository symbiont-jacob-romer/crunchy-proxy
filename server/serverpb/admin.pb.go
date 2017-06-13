// Code generated by protoc-gen-go.
// source: admin.proto
// DO NOT EDIT!

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	admin.proto

It has these top-level messages:
	NodeRequest
	NodeResponse
	PoolRequest
	PoolResponse
	HealthRequest
	HealthResponse
	StatisticsRequest
	StatisticsResponse
	ShutdownRequest
	ShutdownResponse
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// NodeRequest requests a list of nodes.
type NodeRequest struct {
}

func (m *NodeRequest) Reset()                    { *m = NodeRequest{} }
func (m *NodeRequest) String() string            { return proto.CompactTextString(m) }
func (*NodeRequest) ProtoMessage()               {}
func (*NodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// NodeResponse contains a list of nodes.
type NodeResponse struct {
	Nodes map[string]string `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NodeResponse) Reset()                    { *m = NodeResponse{} }
func (m *NodeResponse) String() string            { return proto.CompactTextString(m) }
func (*NodeResponse) ProtoMessage()               {}
func (*NodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NodeResponse) GetNodes() map[string]string {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// PoolRequest requests a list of pools.
type PoolRequest struct {
}

func (m *PoolRequest) Reset()                    { *m = PoolRequest{} }
func (m *PoolRequest) String() string            { return proto.CompactTextString(m) }
func (*PoolRequest) ProtoMessage()               {}
func (*PoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// PoolResponse contains a list of pools.
type PoolResponse struct {
	Pools []string `protobuf:"bytes,1,rep,name=pools" json:"pools,omitempty"`
}

func (m *PoolResponse) Reset()                    { *m = PoolResponse{} }
func (m *PoolResponse) String() string            { return proto.CompactTextString(m) }
func (*PoolResponse) ProtoMessage()               {}
func (*PoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PoolResponse) GetPools() []string {
	if m != nil {
		return m.Pools
	}
	return nil
}

type HealthRequest struct {
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type HealthResponse struct {
	Health map[string]bool `protobuf:"bytes,1,rep,name=health" json:"health,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HealthResponse) GetHealth() map[string]bool {
	if m != nil {
		return m.Health
	}
	return nil
}

type StatisticsRequest struct {
}

func (m *StatisticsRequest) Reset()                    { *m = StatisticsRequest{} }
func (m *StatisticsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatisticsRequest) ProtoMessage()               {}
func (*StatisticsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StatisticsResponse struct {
	Stats []string `protobuf:"bytes,1,rep,name=stats" json:"stats,omitempty"`
}

func (m *StatisticsResponse) Reset()                    { *m = StatisticsResponse{} }
func (m *StatisticsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatisticsResponse) ProtoMessage()               {}
func (*StatisticsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StatisticsResponse) GetStats() []string {
	if m != nil {
		return m.Stats
	}
	return nil
}

// ShutdownRequest requests the server to shutdown.
type ShutdownRequest struct {
}

func (m *ShutdownRequest) Reset()                    { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string            { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()               {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// ShutdownResponse contains the the state of the proxy.
type ShutdownResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *ShutdownResponse) Reset()                    { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string            { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()               {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ShutdownResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*NodeRequest)(nil), "crunchyproxy.server.serverpb.NodeRequest")
	proto.RegisterType((*NodeResponse)(nil), "crunchyproxy.server.serverpb.NodeResponse")
	proto.RegisterType((*PoolRequest)(nil), "crunchyproxy.server.serverpb.PoolRequest")
	proto.RegisterType((*PoolResponse)(nil), "crunchyproxy.server.serverpb.PoolResponse")
	proto.RegisterType((*HealthRequest)(nil), "crunchyproxy.server.serverpb.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "crunchyproxy.server.serverpb.HealthResponse")
	proto.RegisterType((*StatisticsRequest)(nil), "crunchyproxy.server.serverpb.StatisticsRequest")
	proto.RegisterType((*StatisticsResponse)(nil), "crunchyproxy.server.serverpb.StatisticsResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "crunchyproxy.server.serverpb.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "crunchyproxy.server.serverpb.ShutdownResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Admin service

type AdminClient interface {
	Nodes(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	Pools(ctx context.Context, in *PoolRequest, opts ...grpc.CallOption) (*PoolResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (Admin_ShutdownClient, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Nodes(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error) {
	out := new(NodeResponse)
	err := grpc.Invoke(ctx, "/crunchyproxy.server.serverpb.Admin/Nodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Pools(ctx context.Context, in *PoolRequest, opts ...grpc.CallOption) (*PoolResponse, error) {
	out := new(PoolResponse)
	err := grpc.Invoke(ctx, "/crunchyproxy.server.serverpb.Admin/Pools", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/crunchyproxy.server.serverpb.Admin/Health", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	out := new(StatisticsResponse)
	err := grpc.Invoke(ctx, "/crunchyproxy.server.serverpb.Admin/Statistics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (Admin_ShutdownClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Admin_serviceDesc.Streams[0], c.cc, "/crunchyproxy.server.serverpb.Admin/Shutdown", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminShutdownClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Admin_ShutdownClient interface {
	Recv() (*ShutdownResponse, error)
	grpc.ClientStream
}

type adminShutdownClient struct {
	grpc.ClientStream
}

func (x *adminShutdownClient) Recv() (*ShutdownResponse, error) {
	m := new(ShutdownResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Admin service

type AdminServer interface {
	Nodes(context.Context, *NodeRequest) (*NodeResponse, error)
	Pools(context.Context, *PoolRequest) (*PoolResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	Statistics(context.Context, *StatisticsRequest) (*StatisticsResponse, error)
	Shutdown(*ShutdownRequest, Admin_ShutdownServer) error
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crunchyproxy.server.serverpb.Admin/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Nodes(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Pools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Pools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crunchyproxy.server.serverpb.Admin/Pools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Pools(ctx, req.(*PoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crunchyproxy.server.serverpb.Admin/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Statistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Statistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crunchyproxy.server.serverpb.Admin/Statistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Statistics(ctx, req.(*StatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Shutdown_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShutdownRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServer).Shutdown(m, &adminShutdownServer{stream})
}

type Admin_ShutdownServer interface {
	Send(*ShutdownResponse) error
	grpc.ServerStream
}

type adminShutdownServer struct {
	grpc.ServerStream
}

func (x *adminShutdownServer) Send(m *ShutdownResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crunchyproxy.server.serverpb.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nodes",
			Handler:    _Admin_Nodes_Handler,
		},
		{
			MethodName: "Pools",
			Handler:    _Admin_Pools_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Admin_Health_Handler,
		},
		{
			MethodName: "Statistics",
			Handler:    _Admin_Statistics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Shutdown",
			Handler:       _Admin_Shutdown_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin.proto",
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0xb6, 0x64, 0xdd, 0x7e, 0xdb, 0xed, 0xa6, 0x63, 0x2b, 0x21, 0xf6, 0x50, 0x82, 0x87,
	0xba, 0xd6, 0xa4, 0x54, 0x84, 0xb5, 0x37, 0x05, 0x41, 0x10, 0xa4, 0xa4, 0x37, 0x2f, 0x92, 0x26,
	0x43, 0x13, 0x8c, 0x33, 0x31, 0x33, 0xa9, 0x06, 0x6f, 0x1e, 0x8a, 0x77, 0xf1, 0xe8, 0x5f, 0xe5,
	0xbf, 0xe0, 0x1f, 0x22, 0x99, 0x1f, 0x4d, 0xca, 0xe2, 0x26, 0x3d, 0xed, 0xbe, 0x8f, 0x79, 0xdf,
	0x7b, 0x6f, 0xbf, 0xc7, 0xc2, 0x34, 0x4a, 0x3e, 0x65, 0xd4, 0x2f, 0x4a, 0x26, 0x18, 0xde, 0x8f,
	0xcb, 0x8a, 0xc6, 0x69, 0x5d, 0x94, 0xec, 0x6b, 0xed, 0x73, 0x52, 0x5e, 0x91, 0x52, 0x7f, 0x14,
	0x17, 0xee, 0xfe, 0x25, 0x63, 0x97, 0x39, 0x09, 0xa2, 0x22, 0x0b, 0x22, 0x4a, 0x99, 0x88, 0x44,
	0xc6, 0x28, 0x57, 0x5c, 0x6f, 0x06, 0xd3, 0x77, 0x2c, 0x21, 0x21, 0xf9, 0x5c, 0x11, 0x2e, 0xbc,
	0x5f, 0x08, 0xb6, 0x14, 0xe6, 0x05, 0xa3, 0x9c, 0xe0, 0xb7, 0x60, 0x51, 0x96, 0x10, 0xee, 0xa0,
	0x83, 0x8d, 0xc3, 0xe9, 0xc9, 0x73, 0x7f, 0x9d, 0x96, 0xdf, 0xa5, 0x4a, 0xc0, 0x5f, 0x53, 0x51,
	0xd6, 0xa1, 0xda, 0xe1, 0x2e, 0x01, 0xda, 0x21, 0xb6, 0x61, 0xe3, 0x23, 0xa9, 0x1d, 0x74, 0x80,
	0x0e, 0x37, 0xc3, 0xe6, 0x2b, 0xde, 0x05, 0xeb, 0x2a, 0xca, 0x2b, 0xe2, 0x8c, 0xe4, 0x4c, 0x81,
	0xd3, 0xd1, 0x12, 0x35, 0x36, 0xcf, 0x18, 0xcb, 0x8d, 0xcd, 0x47, 0xb0, 0xa5, 0xa0, 0x76, 0xb9,
	0x0b, 0x56, 0xc1, 0x58, 0xae, 0x5c, 0x6e, 0x86, 0x0a, 0x78, 0x73, 0x98, 0xbd, 0x21, 0x51, 0x2e,
	0x52, 0x43, 0xfb, 0x8d, 0x60, 0xdb, 0x4c, 0x34, 0xf3, 0x0c, 0xc6, 0xa9, 0x9c, 0xe8, 0x80, 0xcb,
	0xf5, 0x01, 0x6f, 0xb3, 0x35, 0x54, 0x19, 0xf5, 0x1e, 0xf7, 0x05, 0x4c, 0x3b, 0xe3, 0xbe, 0x94,
	0x93, 0x6e, 0xca, 0xfb, 0xb0, 0x73, 0xde, 0x9c, 0x87, 0x8b, 0x2c, 0xe6, 0xc6, 0xf4, 0x02, 0x70,
	0x77, 0xd8, 0x26, 0xe6, 0x22, 0x12, 0x37, 0x89, 0x25, 0xf0, 0x76, 0x60, 0x7e, 0x9e, 0x56, 0x22,
	0x61, 0x5f, 0xa8, 0xa1, 0x1f, 0x81, 0xdd, 0x8e, 0x34, 0xd9, 0x81, 0x7b, 0xbc, 0x8a, 0x63, 0xc2,
	0xb9, 0xf4, 0x35, 0x09, 0x0d, 0x3c, 0xb9, 0xb6, 0xc0, 0x7a, 0xd9, 0x54, 0x0b, 0x57, 0x60, 0xc9,
	0x5b, 0xe1, 0xc7, 0x43, 0x4e, 0x2e, 0xb5, 0xdc, 0xc5, 0xf0, 0x76, 0x78, 0x7b, 0xdf, 0xff, 0xfc,
	0xfd, 0x39, 0x9a, 0xe3, 0x59, 0xf0, 0x41, 0x76, 0x39, 0x90, 0x15, 0x69, 0x64, 0x9b, 0xcb, 0xf6,
	0xca, 0x76, 0xda, 0xd0, 0x27, 0xdb, 0x6d, 0xca, 0xaa, 0xac, 0xac, 0x0a, 0xfe, 0x06, 0x63, 0x75,
	0x34, 0xfc, 0x64, 0x58, 0x01, 0x94, 0xf2, 0xd1, 0x5d, 0xda, 0xe2, 0x3d, 0x90, 0xda, 0x36, 0xde,
	0x36, 0xda, 0xaa, 0x31, 0xf8, 0x1a, 0x01, 0xb4, 0x27, 0xc6, 0xc1, 0xfa, 0xa5, 0x2b, 0x0d, 0x71,
	0x8f, 0x87, 0x13, 0xfe, 0xf7, 0x2b, 0xc8, 0xfa, 0xe0, 0x1f, 0x08, 0x26, 0xa6, 0x2c, 0xf8, 0x69,
	0xcf, 0xd6, 0xdb, 0x3d, 0x73, 0xfd, 0xa1, 0xcf, 0xb5, 0x85, 0x87, 0xd2, 0xc2, 0x9e, 0x67, 0xdf,
	0x58, 0xd0, 0x2f, 0x4e, 0xd1, 0xe2, 0x18, 0xbd, 0x82, 0xf7, 0x13, 0xc3, 0xbd, 0x18, 0xcb, 0xbf,
	0xaa, 0x67, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x52, 0xce, 0xb4, 0xf5, 0x04, 0x00, 0x00,
}
