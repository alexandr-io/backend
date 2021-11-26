// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcuser

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// User is used to retrieve a user by it's ID or email.
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
	// Login is used to retrieve a user by it's login(email or username) and password.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	// Register is used to create a user.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	// UpdatePassword change a user's password in DB.
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UserReply, error)
	// UpdatePasswordLogged changed a logged user's password in DB.
	UpdatePasswordLogged(ctx context.Context, in *UpdatePasswordLoggedRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.User/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.User/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdatePasswordLogged(ctx context.Context, in *UpdatePasswordLoggedRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.User/UpdatePasswordLogged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// User is used to retrieve a user by it's ID or email.
	User(context.Context, *UserRequest) (*UserReply, error)
	// Login is used to retrieve a user by it's login(email or username) and password.
	Login(context.Context, *LoginRequest) (*UserReply, error)
	// Register is used to create a user.
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	// UpdatePassword change a user's password in DB.
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UserReply, error)
	// UpdatePasswordLogged changed a logged user's password in DB.
	UpdatePasswordLogged(context.Context, *UpdatePasswordLoggedRequest) (*UserReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) User(context.Context, *UserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserServer) UpdatePasswordLogged(context.Context, *UpdatePasswordLoggedRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePasswordLogged not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdatePasswordLogged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordLoggedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdatePasswordLogged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdatePasswordLogged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdatePasswordLogged(ctx, req.(*UpdatePasswordLoggedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _User_User_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _User_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdatePasswordLogged",
			Handler:    _User_UpdatePasswordLogged_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/user.proto",
}
