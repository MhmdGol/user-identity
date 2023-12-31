// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: identity/userapi/v1/userapi.proto

package userapiv1

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

const (
	UserService_CreateUser_FullMethodName       = "/identity.userapi.v1.UserService/CreateUser"
	UserService_UserByID_FullMethodName         = "/identity.userapi.v1.UserService/UserByID"
	UserService_ChangePermission_FullMethodName = "/identity.userapi.v1.UserService/ChangePermission"
	UserService_ChangeStatus_FullMethodName     = "/identity.userapi.v1.UserService/ChangeStatus"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UserByID(ctx context.Context, in *UserByIDRequest, opts ...grpc.CallOption) (*UserByIDResponse, error)
	ChangePermission(ctx context.Context, in *ChangePermissionRequest, opts ...grpc.CallOption) (*ChangePermissionResponse, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserByID(ctx context.Context, in *UserByIDRequest, opts ...grpc.CallOption) (*UserByIDResponse, error) {
	out := new(UserByIDResponse)
	err := c.cc.Invoke(ctx, UserService_UserByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePermission(ctx context.Context, in *ChangePermissionRequest, opts ...grpc.CallOption) (*ChangePermissionResponse, error) {
	out := new(ChangePermissionResponse)
	err := c.cc.Invoke(ctx, UserService_ChangePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusResponse, error) {
	out := new(ChangeStatusResponse)
	err := c.cc.Invoke(ctx, UserService_ChangeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UserByID(context.Context, *UserByIDRequest) (*UserByIDResponse, error)
	ChangePermission(context.Context, *ChangePermissionRequest) (*ChangePermissionResponse, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UserByID(context.Context, *UserByIDRequest) (*UserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserByID not implemented")
}
func (UnimplementedUserServiceServer) ChangePermission(context.Context, *ChangePermissionRequest) (*ChangePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePermission not implemented")
}
func (UnimplementedUserServiceServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserByID(ctx, req.(*UserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePermission(ctx, req.(*ChangePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.userapi.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UserByID",
			Handler:    _UserService_UserByID_Handler,
		},
		{
			MethodName: "ChangePermission",
			Handler:    _UserService_ChangePermission_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _UserService_ChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/userapi/v1/userapi.proto",
}
