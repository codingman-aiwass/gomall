// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: mq/rpc/mq.proto

package mq

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
	Mq_SendDelayMessage_FullMethodName = "/mq.Mq/SendDelayMessage"
	Mq_SendMessage_FullMethodName      = "/mq.Mq/SendMessage"
)

// MqClient is the client API for Mq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 消息队列服务
type MqClient interface {
	// 发送延时消息
	SendDelayMessage(ctx context.Context, in *SendDelayMessageReq, opts ...grpc.CallOption) (*SendDelayMessageResp, error)
	// 发送普通消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type mqClient struct {
	cc grpc.ClientConnInterface
}

func NewMqClient(cc grpc.ClientConnInterface) MqClient {
	return &mqClient{cc}
}

func (c *mqClient) SendDelayMessage(ctx context.Context, in *SendDelayMessageReq, opts ...grpc.CallOption) (*SendDelayMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendDelayMessageResp)
	err := c.cc.Invoke(ctx, Mq_SendDelayMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, Mq_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqServer is the server API for Mq service.
// All implementations must embed UnimplementedMqServer
// for forward compatibility.
//
// 消息队列服务
type MqServer interface {
	// 发送延时消息
	SendDelayMessage(context.Context, *SendDelayMessageReq) (*SendDelayMessageResp, error)
	// 发送普通消息
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
	mustEmbedUnimplementedMqServer()
}

// UnimplementedMqServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMqServer struct{}

func (UnimplementedMqServer) SendDelayMessage(context.Context, *SendDelayMessageReq) (*SendDelayMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDelayMessage not implemented")
}
func (UnimplementedMqServer) SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMqServer) mustEmbedUnimplementedMqServer() {}
func (UnimplementedMqServer) testEmbeddedByValue()            {}

// UnsafeMqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MqServer will
// result in compilation errors.
type UnsafeMqServer interface {
	mustEmbedUnimplementedMqServer()
}

func RegisterMqServer(s grpc.ServiceRegistrar, srv MqServer) {
	// If the following call pancis, it indicates UnimplementedMqServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mq_ServiceDesc, srv)
}

func _Mq_SendDelayMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDelayMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServer).SendDelayMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mq_SendDelayMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServer).SendDelayMessage(ctx, req.(*SendDelayMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mq_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mq_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Mq_ServiceDesc is the grpc.ServiceDesc for Mq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mq.Mq",
	HandlerType: (*MqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDelayMessage",
			Handler:    _Mq_SendDelayMessage_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Mq_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mq/rpc/mq.proto",
}
