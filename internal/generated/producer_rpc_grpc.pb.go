// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: protobuf/producer_rpc.proto

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

// ProducerClient is the client API for Producer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProducerClient interface {
	ProduceMessage(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*empty.Empty, error)
}

type producerClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerClient(cc grpc.ClientConnInterface) ProducerClient {
	return &producerClient{cc}
}

func (c *producerClient) ProduceMessage(ctx context.Context, in *Bytes, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Producer/ProduceMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServer is the server API for Producer service.
// All implementations should embed UnimplementedProducerServer
// for forward compatibility
type ProducerServer interface {
	ProduceMessage(context.Context, *Bytes) (*empty.Empty, error)
}

// UnimplementedProducerServer should be embedded to have forward compatible implementations.
type UnimplementedProducerServer struct {
}

func (UnimplementedProducerServer) ProduceMessage(context.Context, *Bytes) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProduceMessage not implemented")
}

// UnsafeProducerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProducerServer will
// result in compilation errors.
type UnsafeProducerServer interface {
	mustEmbedUnimplementedProducerServer()
}

func RegisterProducerServer(s grpc.ServiceRegistrar, srv ProducerServer) {
	s.RegisterService(&Producer_ServiceDesc, srv)
}

func _Producer_ProduceMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServer).ProduceMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Producer/ProduceMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServer).ProduceMessage(ctx, req.(*Bytes))
	}
	return interceptor(ctx, in, info, handler)
}

// Producer_ServiceDesc is the grpc.ServiceDesc for Producer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Producer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Producer",
	HandlerType: (*ProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProduceMessage",
			Handler:    _Producer_ProduceMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/producer_rpc.proto",
}
