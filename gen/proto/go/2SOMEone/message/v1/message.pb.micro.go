// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: 2someone/message/v1/message.proto

package message

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Message service

func NewMessageEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Message service

type MessageService interface {
	// 向手机号发送验证码
	SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...client.CallOption) (*SentMessageCodeResponse, error)
	// 校验验证码
	CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...client.CallOption) (*CheckMessageCodeResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...client.CallOption) (*SentMessageCodeResponse, error) {
	req := c.c.NewRequest(c.name, "Message.SentMessageCode", in)
	out := new(SentMessageCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...client.CallOption) (*CheckMessageCodeResponse, error) {
	req := c.c.NewRequest(c.name, "Message.CheckMessageCode", in)
	out := new(CheckMessageCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	// 向手机号发送验证码
	SentMessageCode(context.Context, *SentMessageCodeRequest, *SentMessageCodeResponse) error
	// 校验验证码
	CheckMessageCode(context.Context, *CheckMessageCodeRequest, *CheckMessageCodeResponse) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, out *SentMessageCodeResponse) error
		CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, out *CheckMessageCodeResponse) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, out *SentMessageCodeResponse) error {
	return h.MessageHandler.SentMessageCode(ctx, in, out)
}

func (h *messageHandler) CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, out *CheckMessageCodeResponse) error {
	return h.MessageHandler.CheckMessageCode(ctx, in, out)
}
