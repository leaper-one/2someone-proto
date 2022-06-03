// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: idl/proto/bubble-box/rank.proto

package rank

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

// RankServiceClient is the client API for RankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankServiceClient interface {
	// 获取排序
	GetRank(ctx context.Context, in *GetRankRequest, opts ...grpc.CallOption) (RankService_GetRankClient, error)
	// 因其有 Note 而标记他
	HasNote(ctx context.Context, in *HasNoteRequest, opts ...grpc.CallOption) (*HasNoteResponse, error)
	// 爬虫提交抓取的排名数据
	HandInRankInfo(ctx context.Context, opts ...grpc.CallOption) (RankService_HandInRankInfoClient, error)
}

type rankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRankServiceClient(cc grpc.ClientConnInterface) RankServiceClient {
	return &rankServiceClient{cc}
}

func (c *rankServiceClient) GetRank(ctx context.Context, in *GetRankRequest, opts ...grpc.CallOption) (RankService_GetRankClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankService_ServiceDesc.Streams[0], "/rank.RankService/GetRank", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankServiceGetRankClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RankService_GetRankClient interface {
	Recv() (*GetRankResponse, error)
	grpc.ClientStream
}

type rankServiceGetRankClient struct {
	grpc.ClientStream
}

func (x *rankServiceGetRankClient) Recv() (*GetRankResponse, error) {
	m := new(GetRankResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankServiceClient) HasNote(ctx context.Context, in *HasNoteRequest, opts ...grpc.CallOption) (*HasNoteResponse, error) {
	out := new(HasNoteResponse)
	err := c.cc.Invoke(ctx, "/rank.RankService/HasNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankServiceClient) HandInRankInfo(ctx context.Context, opts ...grpc.CallOption) (RankService_HandInRankInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankService_ServiceDesc.Streams[1], "/rank.RankService/HandInRankInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankServiceHandInRankInfoClient{stream}
	return x, nil
}

type RankService_HandInRankInfoClient interface {
	Send(*HandInRankInfoRequest) error
	CloseAndRecv() (*HandInRankInfoResponse, error)
	grpc.ClientStream
}

type rankServiceHandInRankInfoClient struct {
	grpc.ClientStream
}

func (x *rankServiceHandInRankInfoClient) Send(m *HandInRankInfoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rankServiceHandInRankInfoClient) CloseAndRecv() (*HandInRankInfoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HandInRankInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RankServiceServer is the server API for RankService service.
// All implementations must embed UnimplementedRankServiceServer
// for forward compatibility
type RankServiceServer interface {
	// 获取排序
	GetRank(*GetRankRequest, RankService_GetRankServer) error
	// 因其有 Note 而标记他
	HasNote(context.Context, *HasNoteRequest) (*HasNoteResponse, error)
	// 爬虫提交抓取的排名数据
	HandInRankInfo(RankService_HandInRankInfoServer) error
	mustEmbedUnimplementedRankServiceServer()
}

// UnimplementedRankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRankServiceServer struct {
}

func (UnimplementedRankServiceServer) GetRank(*GetRankRequest, RankService_GetRankServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRank not implemented")
}
func (UnimplementedRankServiceServer) HasNote(context.Context, *HasNoteRequest) (*HasNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasNote not implemented")
}
func (UnimplementedRankServiceServer) HandInRankInfo(RankService_HandInRankInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method HandInRankInfo not implemented")
}
func (UnimplementedRankServiceServer) mustEmbedUnimplementedRankServiceServer() {}

// UnsafeRankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServiceServer will
// result in compilation errors.
type UnsafeRankServiceServer interface {
	mustEmbedUnimplementedRankServiceServer()
}

func RegisterRankServiceServer(s grpc.ServiceRegistrar, srv RankServiceServer) {
	s.RegisterService(&RankService_ServiceDesc, srv)
}

func _RankService_GetRank_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRankRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServiceServer).GetRank(m, &rankServiceGetRankServer{stream})
}

type RankService_GetRankServer interface {
	Send(*GetRankResponse) error
	grpc.ServerStream
}

type rankServiceGetRankServer struct {
	grpc.ServerStream
}

func (x *rankServiceGetRankServer) Send(m *GetRankResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RankService_HasNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).HasNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rank.RankService/HasNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).HasNote(ctx, req.(*HasNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankService_HandInRankInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RankServiceServer).HandInRankInfo(&rankServiceHandInRankInfoServer{stream})
}

type RankService_HandInRankInfoServer interface {
	SendAndClose(*HandInRankInfoResponse) error
	Recv() (*HandInRankInfoRequest, error)
	grpc.ServerStream
}

type rankServiceHandInRankInfoServer struct {
	grpc.ServerStream
}

func (x *rankServiceHandInRankInfoServer) SendAndClose(m *HandInRankInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rankServiceHandInRankInfoServer) Recv() (*HandInRankInfoRequest, error) {
	m := new(HandInRankInfoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RankService_ServiceDesc is the grpc.ServiceDesc for RankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rank.RankService",
	HandlerType: (*RankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasNote",
			Handler:    _RankService_HasNote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRank",
			Handler:       _RankService_GetRank_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HandInRankInfo",
			Handler:       _RankService_HandInRankInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "idl/proto/bubble-box/rank.proto",
}
