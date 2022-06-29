package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "MESSAGE-SERVICE/gen/proto/go/mes/v1"
)

//模拟短信模块服务端代码

type MessageServer struct {
	proto.UnimplementedMessageServer
}

func (m *MessageServer) Message(ctx context.Context, in *proto.MessageRequest) (*proto.MessageResponse, error) {
	//通过传来的request中的phone和验证码查询数据库中是否一致
	//这里写MessageService中的方法
	return &proto.MessageResponse{
		Message: "测试！！",
		Code:    "2",
	}, nil
}

func ServerRun() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer() // 创建gRPC服务器

	// 为 message 服务注册业务实现（将 message 服务绑定到 GRPC 服务容器上）
	proto.RegisterMessageServer(grpcServer, &MessageServer{})

	// 注册反射服务，这个服务是 CLI 使用的，跟服务本身没有关系
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
