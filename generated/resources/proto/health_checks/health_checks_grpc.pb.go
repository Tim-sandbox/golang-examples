// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package health_checks

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthChecksClient is the client API for HealthChecks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthChecksClient interface {
	IsAlive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IsAliveOut, error)
}

type healthChecksClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthChecksClient(cc grpc.ClientConnInterface) HealthChecksClient {
	return &healthChecksClient{cc}
}

func (c *healthChecksClient) IsAlive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IsAliveOut, error) {
	out := new(IsAliveOut)
	err := c.cc.Invoke(ctx, "/health_checks.HealthChecks/IsAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthChecksServer is the server API for HealthChecks service.
// All implementations must embed UnimplementedHealthChecksServer
// for forward compatibility
type HealthChecksServer interface {
	IsAlive(context.Context, *emptypb.Empty) (*IsAliveOut, error)
	mustEmbedUnimplementedHealthChecksServer()
}

// UnimplementedHealthChecksServer must be embedded to have forward compatible implementations.
type UnimplementedHealthChecksServer struct {
}

func (UnimplementedHealthChecksServer) IsAlive(context.Context, *emptypb.Empty) (*IsAliveOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAlive not implemented")
}
func (UnimplementedHealthChecksServer) mustEmbedUnimplementedHealthChecksServer() {}

// UnsafeHealthChecksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthChecksServer will
// result in compilation errors.
type UnsafeHealthChecksServer interface {
	mustEmbedUnimplementedHealthChecksServer()
}

func RegisterHealthChecksServer(s grpc.ServiceRegistrar, srv HealthChecksServer) {
	s.RegisterService(&HealthChecks_ServiceDesc, srv)
}

func _HealthChecks_IsAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthChecksServer).IsAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_checks.HealthChecks/IsAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthChecksServer).IsAlive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthChecks_ServiceDesc is the grpc.ServiceDesc for HealthChecks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthChecks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health_checks.HealthChecks",
	HandlerType: (*HealthChecksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAlive",
			Handler:    _HealthChecks_IsAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources/proto/health_checks/health_checks.proto",
}
