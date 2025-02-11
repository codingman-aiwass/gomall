// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: auth/rpc/auth.proto

package auth

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
	AuthService_DeliverTokenByRPC_FullMethodName     = "/auth.AuthService/DeliverTokenByRPC"
	AuthService_VerifyTokenByRPC_FullMethodName      = "/auth.AuthService/VerifyTokenByRPC"
	AuthService_RefreshTokenByRPC_FullMethodName     = "/auth.AuthService/RefreshTokenByRPC"
	AuthService_ExpireTokenByRPC_FullMethodName      = "/auth.AuthService/ExpireTokenByRPC"
	AuthService_VerifyPermissionByRPC_FullMethodName = "/auth.AuthService/VerifyPermissionByRPC"
	AuthService_VerifyPathInWhiteList_FullMethodName = "/auth.AuthService/VerifyPathInWhiteList"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error)
	VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error)
	RefreshTokenByRPC(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
	ExpireTokenByRPC(ctx context.Context, in *ExpireTokenReq, opts ...grpc.CallOption) (*ExpireTokenResp, error)
	VerifyPermissionByRPC(ctx context.Context, in *VerifyPermissionReq, opts ...grpc.CallOption) (*VerifyPermissionResp, error)
	VerifyPathInWhiteList(ctx context.Context, in *VerifyPathInWhiteListReq, opts ...grpc.CallOption) (*VerifyPathInWhiteListResp, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResp)
	err := c.cc.Invoke(ctx, AuthService_DeliverTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyResp)
	err := c.cc.Invoke(ctx, AuthService_VerifyTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshTokenByRPC(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResp)
	err := c.cc.Invoke(ctx, AuthService_RefreshTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ExpireTokenByRPC(ctx context.Context, in *ExpireTokenReq, opts ...grpc.CallOption) (*ExpireTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpireTokenResp)
	err := c.cc.Invoke(ctx, AuthService_ExpireTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyPermissionByRPC(ctx context.Context, in *VerifyPermissionReq, opts ...grpc.CallOption) (*VerifyPermissionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPermissionResp)
	err := c.cc.Invoke(ctx, AuthService_VerifyPermissionByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyPathInWhiteList(ctx context.Context, in *VerifyPathInWhiteListReq, opts ...grpc.CallOption) (*VerifyPathInWhiteListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyPathInWhiteListResp)
	err := c.cc.Invoke(ctx, AuthService_VerifyPathInWhiteList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error)
	VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error)
	RefreshTokenByRPC(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error)
	ExpireTokenByRPC(context.Context, *ExpireTokenReq) (*ExpireTokenResp, error)
	VerifyPermissionByRPC(context.Context, *VerifyPermissionReq) (*VerifyPermissionResp, error)
	VerifyPathInWhiteList(context.Context, *VerifyPathInWhiteListReq) (*VerifyPathInWhiteListResp, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) RefreshTokenByRPC(context.Context, *RefreshTokenReq) (*RefreshTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) ExpireTokenByRPC(context.Context, *ExpireTokenReq) (*ExpireTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) VerifyPermissionByRPC(context.Context, *VerifyPermissionReq) (*VerifyPermissionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPermissionByRPC not implemented")
}
func (UnimplementedAuthServiceServer) VerifyPathInWhiteList(context.Context, *VerifyPathInWhiteListReq) (*VerifyPathInWhiteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPathInWhiteList not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_DeliverTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeliverTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, req.(*DeliverTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshTokenByRPC(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ExpireTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpireTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ExpireTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ExpireTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ExpireTokenByRPC(ctx, req.(*ExpireTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyPermissionByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPermissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyPermissionByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyPermissionByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyPermissionByRPC(ctx, req.(*VerifyPermissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyPathInWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPathInWhiteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyPathInWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyPathInWhiteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyPathInWhiteList(ctx, req.(*VerifyPathInWhiteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverTokenByRPC",
			Handler:    _AuthService_DeliverTokenByRPC_Handler,
		},
		{
			MethodName: "VerifyTokenByRPC",
			Handler:    _AuthService_VerifyTokenByRPC_Handler,
		},
		{
			MethodName: "RefreshTokenByRPC",
			Handler:    _AuthService_RefreshTokenByRPC_Handler,
		},
		{
			MethodName: "ExpireTokenByRPC",
			Handler:    _AuthService_ExpireTokenByRPC_Handler,
		},
		{
			MethodName: "VerifyPermissionByRPC",
			Handler:    _AuthService_VerifyPermissionByRPC_Handler,
		},
		{
			MethodName: "VerifyPathInWhiteList",
			Handler:    _AuthService_VerifyPathInWhiteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/rpc/auth.proto",
}
