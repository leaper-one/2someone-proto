// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: 2someone/user/v1/user.proto

package user

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 通过手机号注册，需要验证码
	SignUpByPhone(ctx context.Context, in *SignUpByPhoneRequest, opts ...grpc.CallOption) (*SignUpByPhoneResponse, error)
	// 通过手机号登录
	SignInByPhone(ctx context.Context, in *SignInByPhoneRequest, opts ...grpc.CallOption) (*SignInByPhoneResponse, error)
	// Get current user infomation by metadata with auth token
	// jwt needed in metadata
	GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error)
	// jwt needed in metadata
	SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*SetInfoResponse, error)
	// 根据 buid 获取 user_id
	GetUserIDByBuid(ctx context.Context, in *GetUserIDByBuidRequest, opts ...grpc.CallOption) (*GetUserIDByBuidResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SignUpByPhone(ctx context.Context, in *SignUpByPhoneRequest, opts ...grpc.CallOption) (*SignUpByPhoneResponse, error) {
	out := new(SignUpByPhoneResponse)
	err := c.cc.Invoke(ctx, "/UserService/SignUpByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignInByPhone(ctx context.Context, in *SignInByPhoneRequest, opts ...grpc.CallOption) (*SignInByPhoneResponse, error) {
	out := new(SignInByPhoneResponse)
	err := c.cc.Invoke(ctx, "/UserService/SignInByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*GetMeResponse, error) {
	out := new(GetMeResponse)
	err := c.cc.Invoke(ctx, "/UserService/GetMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetInfo(ctx context.Context, in *SetInfoRequest, opts ...grpc.CallOption) (*SetInfoResponse, error) {
	out := new(SetInfoResponse)
	err := c.cc.Invoke(ctx, "/UserService/SetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserIDByBuid(ctx context.Context, in *GetUserIDByBuidRequest, opts ...grpc.CallOption) (*GetUserIDByBuidResponse, error) {
	out := new(GetUserIDByBuidResponse)
	err := c.cc.Invoke(ctx, "/UserService/GetUserIDByBuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 通过手机号注册，需要验证码
	SignUpByPhone(context.Context, *SignUpByPhoneRequest) (*SignUpByPhoneResponse, error)
	// 通过手机号登录
	SignInByPhone(context.Context, *SignInByPhoneRequest) (*SignInByPhoneResponse, error)
	// Get current user infomation by metadata with auth token
	// jwt needed in metadata
	GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error)
	// jwt needed in metadata
	SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error)
	// 根据 buid 获取 user_id
	GetUserIDByBuid(context.Context, *GetUserIDByBuidRequest) (*GetUserIDByBuidResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) SignUpByPhone(context.Context, *SignUpByPhoneRequest) (*SignUpByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpByPhone not implemented")
}
func (UnimplementedUserServiceServer) SignInByPhone(context.Context, *SignInByPhoneRequest) (*SignInByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInByPhone not implemented")
}
func (UnimplementedUserServiceServer) GetMe(context.Context, *GetMeRequest) (*GetMeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserServiceServer) SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserIDByBuid(context.Context, *GetUserIDByBuidRequest) (*GetUserIDByBuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIDByBuid not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_SignUpByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUpByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/SignUpByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUpByPhone(ctx, req.(*SignUpByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignInByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignInByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/SignInByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignInByPhone(ctx, req.(*SignInByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMe(ctx, req.(*GetMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/SetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetInfo(ctx, req.(*SetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserIDByBuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIDByBuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserIDByBuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetUserIDByBuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserIDByBuid(ctx, req.(*GetUserIDByBuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUpByPhone",
			Handler:    _UserService_SignUpByPhone_Handler,
		},
		{
			MethodName: "SignInByPhone",
			Handler:    _UserService_SignInByPhone_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _UserService_GetMe_Handler,
		},
		{
			MethodName: "SetInfo",
			Handler:    _UserService_SetInfo_Handler,
		},
		{
			MethodName: "GetUserIDByBuid",
			Handler:    _UserService_GetUserIDByBuid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "2someone/user/v1/user.proto",
}
