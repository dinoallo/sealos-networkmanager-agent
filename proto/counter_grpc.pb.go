// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/counter.proto

package sealos_networkmanager_agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CountingServiceClient is the client API for CountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountingServiceClient interface {
	CreateCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*Counter, error)
	DumpCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (CountingService_DumpCounterClient, error)
	RemoveCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*Empty, error)
}

type countingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountingServiceClient(cc grpc.ClientConnInterface) CountingServiceClient {
	return &countingServiceClient{cc}
}

func (c *countingServiceClient) CreateCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*Counter, error) {
	out := new(Counter)
	err := c.cc.Invoke(ctx, "/proto.CountingService/CreateCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countingServiceClient) DumpCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (CountingService_DumpCounterClient, error) {
	stream, err := c.cc.NewStream(ctx, &CountingService_ServiceDesc.Streams[0], "/proto.CountingService/DumpCounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &countingServiceDumpCounterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CountingService_DumpCounterClient interface {
	Recv() (*CounterDump, error)
	grpc.ClientStream
}

type countingServiceDumpCounterClient struct {
	grpc.ClientStream
}

func (x *countingServiceDumpCounterClient) Recv() (*CounterDump, error) {
	m := new(CounterDump)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *countingServiceClient) RemoveCounter(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.CountingService/RemoveCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountingServiceServer is the server API for CountingService service.
// All implementations must embed UnimplementedCountingServiceServer
// for forward compatibility
type CountingServiceServer interface {
	CreateCounter(context.Context, *Counter) (*Counter, error)
	DumpCounter(*Counter, CountingService_DumpCounterServer) error
	RemoveCounter(context.Context, *Counter) (*Empty, error)
	mustEmbedUnimplementedCountingServiceServer()
}

// UnimplementedCountingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountingServiceServer struct {
}

func (UnimplementedCountingServiceServer) CreateCounter(context.Context, *Counter) (*Counter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCounter not implemented")
}
func (UnimplementedCountingServiceServer) DumpCounter(*Counter, CountingService_DumpCounterServer) error {
	return status.Errorf(codes.Unimplemented, "method DumpCounter not implemented")
}
func (UnimplementedCountingServiceServer) RemoveCounter(context.Context, *Counter) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCounter not implemented")
}
func (UnimplementedCountingServiceServer) mustEmbedUnimplementedCountingServiceServer() {}

// UnsafeCountingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountingServiceServer will
// result in compilation errors.
type UnsafeCountingServiceServer interface {
	mustEmbedUnimplementedCountingServiceServer()
}

func RegisterCountingServiceServer(s grpc.ServiceRegistrar, srv CountingServiceServer) {
	s.RegisterService(&CountingService_ServiceDesc, srv)
}

func _CountingService_CreateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Counter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).CreateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/CreateCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).CreateCounter(ctx, req.(*Counter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountingService_DumpCounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Counter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CountingServiceServer).DumpCounter(m, &countingServiceDumpCounterServer{stream})
}

type CountingService_DumpCounterServer interface {
	Send(*CounterDump) error
	grpc.ServerStream
}

type countingServiceDumpCounterServer struct {
	grpc.ServerStream
}

func (x *countingServiceDumpCounterServer) Send(m *CounterDump) error {
	return x.ServerStream.SendMsg(m)
}

func _CountingService_RemoveCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Counter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).RemoveCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/RemoveCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).RemoveCounter(ctx, req.(*Counter))
	}
	return interceptor(ctx, in, info, handler)
}

// CountingService_ServiceDesc is the grpc.ServiceDesc for CountingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CountingService",
	HandlerType: (*CountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCounter",
			Handler:    _CountingService_CreateCounter_Handler,
		},
		{
			MethodName: "RemoveCounter",
			Handler:    _CountingService_RemoveCounter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DumpCounter",
			Handler:       _CountingService_DumpCounter_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/counter.proto",
}