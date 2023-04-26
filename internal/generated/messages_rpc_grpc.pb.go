// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: protobuf/messages_rpc.proto

package generated

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	SwitchMessageType(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*empty.Empty, error)
	PutInProducer(ctx context.Context, in *WebSocketMessage, opts ...grpc.CallOption) (*empty.Empty, error)
	PullFromConsumer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Bytes, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) SwitchMessageType(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Messages/SwitchMessageType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) PutInProducer(ctx context.Context, in *WebSocketMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Messages/PutInProducer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) PullFromConsumer(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Bytes, error) {
	out := new(Bytes)
	err := c.cc.Invoke(ctx, "/protobuf.Messages/PullFromConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServer is the server API for Messages service.
// All implementations should embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	SwitchMessageType(context.Context, *Bytes) (*empty.Empty, error)
	PutInProducer(context.Context, *WebSocketMessage) (*empty.Empty, error)
	PullFromConsumer(context.Context, *empty.Empty) (*Bytes, error)
}

// UnimplementedMessagesServer should be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) SwitchMessageType(context.Context, *Bytes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchMessageType not implemented")
}
func (UnimplementedMessagesServer) PutInProducer(context.Context, *WebSocketMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutInProducer not implemented")
}
func (UnimplementedMessagesServer) PullFromConsumer(context.Context, *empty.Empty) (*Bytes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullFromConsumer not implemented")
}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_SwitchMessageType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).SwitchMessageType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Messages/SwitchMessageType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).SwitchMessageType(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_PutInProducer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebSocketMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).PutInProducer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Messages/PutInProducer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).PutInProducer(ctx, req.(*WebSocketMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_PullFromConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).PullFromConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Messages/PullFromConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).PullFromConsumer(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwitchMessageType",
			Handler:    _Messages_SwitchMessageType_Handler,
		},
		{
			MethodName: "PutInProducer",
			Handler:    _Messages_PutInProducer_Handler,
		},
		{
			MethodName: "PullFromConsumer",
			Handler:    _Messages_PullFromConsumer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/messages_rpc.proto",
}
