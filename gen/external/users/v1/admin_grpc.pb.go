// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: external/users/v1/admin.proto

package usersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UsersAdminService_SearchUsers_FullMethodName         = "/usersservice.v1.UsersAdminService/SearchUsers"
	UsersAdminService_GetUserByIdentifier_FullMethodName = "/usersservice.v1.UsersAdminService/GetUserByIdentifier"
	UsersAdminService_UpdateUserRole_FullMethodName      = "/usersservice.v1.UsersAdminService/UpdateUserRole"
	UsersAdminService_BanUser_FullMethodName             = "/usersservice.v1.UsersAdminService/BanUser"
	UsersAdminService_UnbanUser_FullMethodName           = "/usersservice.v1.UsersAdminService/UnbanUser"
)

// UsersAdminServiceClient is the client API for UsersAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *
// Users Admin Service has methods for manage user's profiles via Admin access
type UsersAdminServiceClient interface {
	//*
	// Search users profile with pagination and filters
	// Auth: Admin
	SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error)
	//*
	// Get User profile with Admin view, possible use one of profile identifier.
	// It can be id, email or username.
	// Auth: Admin
	GetUserByIdentifier(ctx context.Context, in *GetUserByIdentifierRequest, opts ...grpc.CallOption) (*UserAdmin, error)
	//*
	// Set new user profile role
	// Auth: Super
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//*
	// Method for ban user, it's happens by set delete_at a date. After ban, impossible login.
	// Auth: Super
	BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//*
	// Method for unban user profile, it's happens by set delete_at a null.
	// Auth: Super
	UnbanUser(ctx context.Context, in *UnbanUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type usersAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersAdminServiceClient(cc grpc.ClientConnInterface) UsersAdminServiceClient {
	return &usersAdminServiceClient{cc}
}

func (c *usersAdminServiceClient) SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUsersResponse)
	err := c.cc.Invoke(ctx, UsersAdminService_SearchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAdminServiceClient) GetUserByIdentifier(ctx context.Context, in *GetUserByIdentifierRequest, opts ...grpc.CallOption) (*UserAdmin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserAdmin)
	err := c.cc.Invoke(ctx, UsersAdminService_GetUserByIdentifier_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAdminServiceClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UsersAdminService_UpdateUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAdminServiceClient) BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UsersAdminService_BanUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAdminServiceClient) UnbanUser(ctx context.Context, in *UnbanUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UsersAdminService_UnbanUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersAdminServiceServer is the server API for UsersAdminService service.
// All implementations should embed UnimplementedUsersAdminServiceServer
// for forward compatibility.
//
// *
// Users Admin Service has methods for manage user's profiles via Admin access
type UsersAdminServiceServer interface {
	//*
	// Search users profile with pagination and filters
	// Auth: Admin
	SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error)
	//*
	// Get User profile with Admin view, possible use one of profile identifier.
	// It can be id, email or username.
	// Auth: Admin
	GetUserByIdentifier(context.Context, *GetUserByIdentifierRequest) (*UserAdmin, error)
	//*
	// Set new user profile role
	// Auth: Super
	UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error)
	//*
	// Method for ban user, it's happens by set delete_at a date. After ban, impossible login.
	// Auth: Super
	BanUser(context.Context, *BanUserRequest) (*emptypb.Empty, error)
	//*
	// Method for unban user profile, it's happens by set delete_at a null.
	// Auth: Super
	UnbanUser(context.Context, *UnbanUserRequest) (*emptypb.Empty, error)
}

// UnimplementedUsersAdminServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersAdminServiceServer struct{}

func (UnimplementedUsersAdminServiceServer) SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedUsersAdminServiceServer) GetUserByIdentifier(context.Context, *GetUserByIdentifierRequest) (*UserAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByIdentifier not implemented")
}
func (UnimplementedUsersAdminServiceServer) UpdateUserRole(context.Context, *UpdateUserRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedUsersAdminServiceServer) BanUser(context.Context, *BanUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanUser not implemented")
}
func (UnimplementedUsersAdminServiceServer) UnbanUser(context.Context, *UnbanUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanUser not implemented")
}
func (UnimplementedUsersAdminServiceServer) testEmbeddedByValue() {}

// UnsafeUsersAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersAdminServiceServer will
// result in compilation errors.
type UnsafeUsersAdminServiceServer interface {
	mustEmbedUnimplementedUsersAdminServiceServer()
}

func RegisterUsersAdminServiceServer(s grpc.ServiceRegistrar, srv UsersAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedUsersAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UsersAdminService_ServiceDesc, srv)
}

func _UsersAdminService_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAdminServiceServer).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAdminService_SearchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAdminServiceServer).SearchUsers(ctx, req.(*SearchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAdminService_GetUserByIdentifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdentifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAdminServiceServer).GetUserByIdentifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAdminService_GetUserByIdentifier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAdminServiceServer).GetUserByIdentifier(ctx, req.(*GetUserByIdentifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAdminService_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAdminServiceServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAdminService_UpdateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAdminServiceServer).UpdateUserRole(ctx, req.(*UpdateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAdminService_BanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAdminServiceServer).BanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAdminService_BanUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAdminServiceServer).BanUser(ctx, req.(*BanUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAdminService_UnbanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbanUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAdminServiceServer).UnbanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAdminService_UnbanUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAdminServiceServer).UnbanUser(ctx, req.(*UnbanUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersAdminService_ServiceDesc is the grpc.ServiceDesc for UsersAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usersservice.v1.UsersAdminService",
	HandlerType: (*UsersAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUsers",
			Handler:    _UsersAdminService_SearchUsers_Handler,
		},
		{
			MethodName: "GetUserByIdentifier",
			Handler:    _UsersAdminService_GetUserByIdentifier_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _UsersAdminService_UpdateUserRole_Handler,
		},
		{
			MethodName: "BanUser",
			Handler:    _UsersAdminService_BanUser_Handler,
		},
		{
			MethodName: "UnbanUser",
			Handler:    _UsersAdminService_UnbanUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/users/v1/admin.proto",
}
