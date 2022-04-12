// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/bootstrap.proto

package bootstrap

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

// BootstrapClient is the client API for Bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BootstrapClient interface {
	Hello(ctx context.Context, in *HelloQuery, opts ...grpc.CallOption) (*HelloDocument, error)
	Login(ctx context.Context, in *LoginCommand, opts ...grpc.CallOption) (*LoginDocument, error)
	Session(ctx context.Context, in *SessionCommand, opts ...grpc.CallOption) (*SessionDocument, error)
}

type bootstrapClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstrapClient(cc grpc.ClientConnInterface) BootstrapClient {
	return &bootstrapClient{cc}
}

func (c *bootstrapClient) Hello(ctx context.Context, in *HelloQuery, opts ...grpc.CallOption) (*HelloDocument, error) {
	out := new(HelloDocument)
	err := c.cc.Invoke(ctx, "/Bootstrap/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootstrapClient) Login(ctx context.Context, in *LoginCommand, opts ...grpc.CallOption) (*LoginDocument, error) {
	out := new(LoginDocument)
	err := c.cc.Invoke(ctx, "/Bootstrap/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootstrapClient) Session(ctx context.Context, in *SessionCommand, opts ...grpc.CallOption) (*SessionDocument, error) {
	out := new(SessionDocument)
	err := c.cc.Invoke(ctx, "/Bootstrap/session", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapServer is the server API for Bootstrap service.
// All implementations must embed UnimplementedBootstrapServer
// for forward compatibility
type BootstrapServer interface {
	Hello(context.Context, *HelloQuery) (*HelloDocument, error)
	Login(context.Context, *LoginCommand) (*LoginDocument, error)
	Session(context.Context, *SessionCommand) (*SessionDocument, error)
	mustEmbedUnimplementedBootstrapServer()
}

// UnimplementedBootstrapServer must be embedded to have forward compatible implementations.
type UnimplementedBootstrapServer struct {
}

func (UnimplementedBootstrapServer) Hello(context.Context, *HelloQuery) (*HelloDocument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedBootstrapServer) Login(context.Context, *LoginCommand) (*LoginDocument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBootstrapServer) Session(context.Context, *SessionCommand) (*SessionDocument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Session not implemented")
}
func (UnimplementedBootstrapServer) mustEmbedUnimplementedBootstrapServer() {}

// UnsafeBootstrapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BootstrapServer will
// result in compilation errors.
type UnsafeBootstrapServer interface {
	mustEmbedUnimplementedBootstrapServer()
}

func RegisterBootstrapServer(s grpc.ServiceRegistrar, srv BootstrapServer) {
	s.RegisterService(&Bootstrap_ServiceDesc, srv)
}

func _Bootstrap_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bootstrap/hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).Hello(ctx, req.(*HelloQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bootstrap_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bootstrap/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).Login(ctx, req.(*LoginCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bootstrap_Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bootstrap/session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).Session(ctx, req.(*SessionCommand))
	}
	return interceptor(ctx, in, info, handler)
}

// Bootstrap_ServiceDesc is the grpc.ServiceDesc for Bootstrap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bootstrap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Bootstrap",
	HandlerType: (*BootstrapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hello",
			Handler:    _Bootstrap_Hello_Handler,
		},
		{
			MethodName: "login",
			Handler:    _Bootstrap_Login_Handler,
		},
		{
			MethodName: "session",
			Handler:    _Bootstrap_Session_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bootstrap.proto",
}
