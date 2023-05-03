// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: protobuf/user_rpc.proto

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	DeleteUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckExistUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	AddUserContact(ctx context.Context, in *AddUserContactArguments, opts ...grpc.CallOption) (*Contacts, error)
	GetUserContacts(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Contacts, error)
	PutUserById(ctx context.Context, in *PutUserArguments, opts ...grpc.CallOption) (*User, error)
	GetAllUsersExceptCurrentUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Contacts, error)
	GetSearchUsers(ctx context.Context, in *String, opts ...grpc.CallOption) (*Contacts, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) DeleteUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Users/DeleteUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CheckExistUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.Users/CheckExistUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserById(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protobuf.Users/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddUserContact(ctx context.Context, in *AddUserContactArguments, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/protobuf.Users/AddUserContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserContacts(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/protobuf.Users/GetUserContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) PutUserById(ctx context.Context, in *PutUserArguments, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protobuf.Users/PutUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAllUsersExceptCurrentUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/protobuf.Users/GetAllUsersExceptCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetSearchUsers(ctx context.Context, in *String, opts ...grpc.CallOption) (*Contacts, error) {
	out := new(Contacts)
	err := c.cc.Invoke(ctx, "/protobuf.Users/GetSearchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations should embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	DeleteUserById(context.Context, *UserID) (*empty.Empty, error)
	CheckExistUserById(context.Context, *UserID) (*empty.Empty, error)
	GetUserById(context.Context, *UserID) (*User, error)
	AddUserContact(context.Context, *AddUserContactArguments) (*Contacts, error)
	GetUserContacts(context.Context, *UserID) (*Contacts, error)
	PutUserById(context.Context, *PutUserArguments) (*User, error)
	GetAllUsersExceptCurrentUser(context.Context, *UserID) (*Contacts, error)
	GetSearchUsers(context.Context, *String) (*Contacts, error)
}

// UnimplementedUsersServer should be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) DeleteUserById(context.Context, *UserID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUsersServer) CheckExistUserById(context.Context, *UserID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistUserById not implemented")
}
func (UnimplementedUsersServer) GetUserById(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersServer) AddUserContact(context.Context, *AddUserContactArguments) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserContact not implemented")
}
func (UnimplementedUsersServer) GetUserContacts(context.Context, *UserID) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserContacts not implemented")
}
func (UnimplementedUsersServer) PutUserById(context.Context, *PutUserArguments) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutUserById not implemented")
}
func (UnimplementedUsersServer) GetAllUsersExceptCurrentUser(context.Context, *UserID) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsersExceptCurrentUser not implemented")
}
func (UnimplementedUsersServer) GetSearchUsers(context.Context, *String) (*Contacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchUsers not implemented")
}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/DeleteUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUserById(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CheckExistUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CheckExistUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/CheckExistUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CheckExistUserById(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserById(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddUserContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserContactArguments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddUserContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/AddUserContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddUserContact(ctx, req.(*AddUserContactArguments))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/GetUserContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserContacts(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_PutUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutUserArguments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).PutUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/PutUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).PutUserById(ctx, req.(*PutUserArguments))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAllUsersExceptCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAllUsersExceptCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/GetAllUsersExceptCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAllUsersExceptCurrentUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetSearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetSearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Users/GetSearchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetSearchUsers(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserById",
			Handler:    _Users_DeleteUserById_Handler,
		},
		{
			MethodName: "CheckExistUserById",
			Handler:    _Users_CheckExistUserById_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _Users_GetUserById_Handler,
		},
		{
			MethodName: "AddUserContact",
			Handler:    _Users_AddUserContact_Handler,
		},
		{
			MethodName: "GetUserContacts",
			Handler:    _Users_GetUserContacts_Handler,
		},
		{
			MethodName: "PutUserById",
			Handler:    _Users_PutUserById_Handler,
		},
		{
			MethodName: "GetAllUsersExceptCurrentUser",
			Handler:    _Users_GetAllUsersExceptCurrentUser_Handler,
		},
		{
			MethodName: "GetSearchUsers",
			Handler:    _Users_GetSearchUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/user_rpc.proto",
}
