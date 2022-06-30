// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bubblebox/note/v1/note.proto

package note

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

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteServiceClient interface {
	//
	// 根据 filter 访问投稿
	// jwt needed in metadata
	GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error)
	//
	// 创建投稿
	// jwt needed in metadata
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	//
	// 更新投稿
	// jwt needed in metadata
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error)
	//
	// 删除投稿
	// jwt needed in metadata
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error)
	//
	// 改变投稿已读状态
	// jwt needed in metadata
	ChangeNoteReadStatus(ctx context.Context, in *ChangeNoteReadStatusRequest, opts ...grpc.CallOption) (*ChangeNoteReadStatusResponse, error)
	//
	// 改变投稿归档状态
	// jwt needed in metadata
	ChangeNoteArchiveStatus(ctx context.Context, in *ChangeNoteArchiveStatusRequest, opts ...grpc.CallOption) (*ChangeNoteArchiveStatusResponse, error)
}

type noteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteServiceClient(cc grpc.ClientConnInterface) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) GetNotes(ctx context.Context, in *GetNotesRequest, opts ...grpc.CallOption) (*GetNotesResponse, error) {
	out := new(GetNotesResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/GetNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...grpc.CallOption) (*UpdateNoteResponse, error) {
	out := new(UpdateNoteResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/UpdateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...grpc.CallOption) (*DeleteNoteResponse, error) {
	out := new(DeleteNoteResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/DeleteNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) ChangeNoteReadStatus(ctx context.Context, in *ChangeNoteReadStatusRequest, opts ...grpc.CallOption) (*ChangeNoteReadStatusResponse, error) {
	out := new(ChangeNoteReadStatusResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/ChangeNoteReadStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) ChangeNoteArchiveStatus(ctx context.Context, in *ChangeNoteArchiveStatusRequest, opts ...grpc.CallOption) (*ChangeNoteArchiveStatusResponse, error) {
	out := new(ChangeNoteArchiveStatusResponse)
	err := c.cc.Invoke(ctx, "/note.v1.NoteService/ChangeNoteArchiveStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServiceServer is the server API for NoteService service.
// All implementations should embed UnimplementedNoteServiceServer
// for forward compatibility
type NoteServiceServer interface {
	//
	// 根据 filter 访问投稿
	// jwt needed in metadata
	GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error)
	//
	// 创建投稿
	// jwt needed in metadata
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	//
	// 更新投稿
	// jwt needed in metadata
	UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error)
	//
	// 删除投稿
	// jwt needed in metadata
	DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error)
	//
	// 改变投稿已读状态
	// jwt needed in metadata
	ChangeNoteReadStatus(context.Context, *ChangeNoteReadStatusRequest) (*ChangeNoteReadStatusResponse, error)
	//
	// 改变投稿归档状态
	// jwt needed in metadata
	ChangeNoteArchiveStatus(context.Context, *ChangeNoteArchiveStatusRequest) (*ChangeNoteArchiveStatusResponse, error)
}

// UnimplementedNoteServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNoteServiceServer struct {
}

func (UnimplementedNoteServiceServer) GetNotes(context.Context, *GetNotesRequest) (*GetNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotes not implemented")
}
func (UnimplementedNoteServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteServiceServer) UpdateNote(context.Context, *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNote not implemented")
}
func (UnimplementedNoteServiceServer) DeleteNote(context.Context, *DeleteNoteRequest) (*DeleteNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNote not implemented")
}
func (UnimplementedNoteServiceServer) ChangeNoteReadStatus(context.Context, *ChangeNoteReadStatusRequest) (*ChangeNoteReadStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNoteReadStatus not implemented")
}
func (UnimplementedNoteServiceServer) ChangeNoteArchiveStatus(context.Context, *ChangeNoteArchiveStatusRequest) (*ChangeNoteArchiveStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNoteArchiveStatus not implemented")
}

// UnsafeNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteServiceServer will
// result in compilation errors.
type UnsafeNoteServiceServer interface {
	mustEmbedUnimplementedNoteServiceServer()
}

func RegisterNoteServiceServer(s grpc.ServiceRegistrar, srv NoteServiceServer) {
	s.RegisterService(&NoteService_ServiceDesc, srv)
}

func _NoteService_GetNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/GetNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNotes(ctx, req.(*GetNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_UpdateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).UpdateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/UpdateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).UpdateNote(ctx, req.(*UpdateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_DeleteNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).DeleteNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/DeleteNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).DeleteNote(ctx, req.(*DeleteNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_ChangeNoteReadStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNoteReadStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).ChangeNoteReadStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/ChangeNoteReadStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).ChangeNoteReadStatus(ctx, req.(*ChangeNoteReadStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_ChangeNoteArchiveStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNoteArchiveStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).ChangeNoteArchiveStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.v1.NoteService/ChangeNoteArchiveStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).ChangeNoteArchiveStatus(ctx, req.(*ChangeNoteArchiveStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteService_ServiceDesc is the grpc.ServiceDesc for NoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "note.v1.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotes",
			Handler:    _NoteService_GetNotes_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _NoteService_CreateNote_Handler,
		},
		{
			MethodName: "UpdateNote",
			Handler:    _NoteService_UpdateNote_Handler,
		},
		{
			MethodName: "DeleteNote",
			Handler:    _NoteService_DeleteNote_Handler,
		},
		{
			MethodName: "ChangeNoteReadStatus",
			Handler:    _NoteService_ChangeNoteReadStatus_Handler,
		},
		{
			MethodName: "ChangeNoteArchiveStatus",
			Handler:    _NoteService_ChangeNoteArchiveStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bubblebox/note/v1/note.proto",
}
