// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: getrates_req_grpc.proto

package protomodel

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetRatesServiceWithGrpc_GetRatesGrpc_FullMethodName = "/protomodel.GetRatesServiceWithGrpc/GetRatesGrpc"
)

// GetRatesServiceWithGrpcClient is the client API for GetRatesServiceWithGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetRatesServiceWithGrpcClient interface {
	GetRatesGrpc(ctx context.Context, in *RatesRequestGrpc, opts ...grpc.CallOption) (*RatesResponseGrpc, error)
}

type getRatesServiceWithGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGetRatesServiceWithGrpcClient(cc grpc.ClientConnInterface) GetRatesServiceWithGrpcClient {
	return &getRatesServiceWithGrpcClient{cc}
}

func (c *getRatesServiceWithGrpcClient) GetRatesGrpc(ctx context.Context, in *RatesRequestGrpc, opts ...grpc.CallOption) (*RatesResponseGrpc, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RatesResponseGrpc)
	err := c.cc.Invoke(ctx, GetRatesServiceWithGrpc_GetRatesGrpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetRatesServiceWithGrpcServer is the server API for GetRatesServiceWithGrpc service.
// All implementations must embed UnimplementedGetRatesServiceWithGrpcServer
// for forward compatibility.
type GetRatesServiceWithGrpcServer interface {
	GetRatesGrpc(context.Context, *RatesRequestGrpc) (*RatesResponseGrpc, error)
	mustEmbedUnimplementedGetRatesServiceWithGrpcServer()
}

// UnimplementedGetRatesServiceWithGrpcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetRatesServiceWithGrpcServer struct{}

func (UnimplementedGetRatesServiceWithGrpcServer) GetRatesGrpc(context.Context, *RatesRequestGrpc) (*RatesResponseGrpc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatesGrpc not implemented")
}
func (UnimplementedGetRatesServiceWithGrpcServer) mustEmbedUnimplementedGetRatesServiceWithGrpcServer() {
}
func (UnimplementedGetRatesServiceWithGrpcServer) testEmbeddedByValue() {}

// UnsafeGetRatesServiceWithGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetRatesServiceWithGrpcServer will
// result in compilation errors.
type UnsafeGetRatesServiceWithGrpcServer interface {
	mustEmbedUnimplementedGetRatesServiceWithGrpcServer()
}

func RegisterGetRatesServiceWithGrpcServer(s grpc.ServiceRegistrar, srv GetRatesServiceWithGrpcServer) {
	// If the following call pancis, it indicates UnimplementedGetRatesServiceWithGrpcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetRatesServiceWithGrpc_ServiceDesc, srv)
}

func _GetRatesServiceWithGrpc_GetRatesGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatesRequestGrpc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetRatesServiceWithGrpcServer).GetRatesGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetRatesServiceWithGrpc_GetRatesGrpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetRatesServiceWithGrpcServer).GetRatesGrpc(ctx, req.(*RatesRequestGrpc))
	}
	return interceptor(ctx, in, info, handler)
}

// GetRatesServiceWithGrpc_ServiceDesc is the grpc.ServiceDesc for GetRatesServiceWithGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetRatesServiceWithGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protomodel.GetRatesServiceWithGrpc",
	HandlerType: (*GetRatesServiceWithGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRatesGrpc",
			Handler:    _GetRatesServiceWithGrpc_GetRatesGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getrates_req_grpc.proto",
}
