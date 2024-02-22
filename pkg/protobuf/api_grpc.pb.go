// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api.proto

package protobuf

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

// NetworkSnifferApiClient is the client API for NetworkSnifferApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkSnifferApiClient interface {
	ProcessNetworkData(ctx context.Context, in *NetworkData, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type networkSnifferApiClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkSnifferApiClient(cc grpc.ClientConnInterface) NetworkSnifferApiClient {
	return &networkSnifferApiClient{cc}
}

func (c *networkSnifferApiClient) ProcessNetworkData(ctx context.Context, in *NetworkData, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protobuf.NetworkSnifferApi/ProcessNetworkData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkSnifferApiServer is the server API for NetworkSnifferApi service.
// All implementations should embed UnimplementedNetworkSnifferApiServer
// for forward compatibility
type NetworkSnifferApiServer interface {
	ProcessNetworkData(context.Context, *NetworkData) (*EmptyResponse, error)
}

// UnimplementedNetworkSnifferApiServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkSnifferApiServer struct {
}

func (UnimplementedNetworkSnifferApiServer) ProcessNetworkData(context.Context, *NetworkData) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessNetworkData not implemented")
}

// UnsafeNetworkSnifferApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkSnifferApiServer will
// result in compilation errors.
type UnsafeNetworkSnifferApiServer interface {
	mustEmbedUnimplementedNetworkSnifferApiServer()
}

func RegisterNetworkSnifferApiServer(s grpc.ServiceRegistrar, srv NetworkSnifferApiServer) {
	s.RegisterService(&NetworkSnifferApi_ServiceDesc, srv)
}

func _NetworkSnifferApi_ProcessNetworkData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkSnifferApiServer).ProcessNetworkData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.NetworkSnifferApi/ProcessNetworkData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkSnifferApiServer).ProcessNetworkData(ctx, req.(*NetworkData))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkSnifferApi_ServiceDesc is the grpc.ServiceDesc for NetworkSnifferApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkSnifferApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.NetworkSnifferApi",
	HandlerType: (*NetworkSnifferApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessNetworkData",
			Handler:    _NetworkSnifferApi_ProcessNetworkData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
