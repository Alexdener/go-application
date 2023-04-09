// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/server1.proto

package server1

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

// Api Endpoints for Server1 service

func NewServer1Endpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Server1 service

type Server1Service interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Server1_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Server1_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Server1_BidiStreamService, error)
}

type server1Service struct {
	c    client.Client
	name string
}

func NewServer1Service(name string, c client.Client) Server1Service {
	return &server1Service{
		c:    c,
		name: name,
	}
}

func (c *server1Service) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Server1.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *server1Service) ClientStream(ctx context.Context, opts ...client.CallOption) (Server1_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Server1.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &server1ServiceClientStream{stream}, nil
}

type Server1_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type server1ServiceClientStream struct {
	stream client.Stream
}

func (x *server1ServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *server1ServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *server1ServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1ServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1ServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1ServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *server1Service) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Server1_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Server1.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &server1ServiceServerStream{stream}, nil
}

type Server1_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type server1ServiceServerStream struct {
	stream client.Stream
}

func (x *server1ServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *server1ServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *server1ServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1ServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1ServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1ServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *server1Service) BidiStream(ctx context.Context, opts ...client.CallOption) (Server1_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Server1.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &server1ServiceBidiStream{stream}, nil
}

type Server1_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type server1ServiceBidiStream struct {
	stream client.Stream
}

func (x *server1ServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *server1ServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *server1ServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1ServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1ServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1ServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *server1ServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Server1 service

type Server1Handler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Server1_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Server1_ServerStreamStream) error
	BidiStream(context.Context, Server1_BidiStreamStream) error
}

func RegisterServer1Handler(s server.Server, hdlr Server1Handler, opts ...server.HandlerOption) error {
	type server1 interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Server1 struct {
		server1
	}
	h := &server1Handler{hdlr}
	return s.Handle(s.NewHandler(&Server1{h}, opts...))
}

type server1Handler struct {
	Server1Handler
}

func (h *server1Handler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.Server1Handler.Call(ctx, in, out)
}

func (h *server1Handler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.Server1Handler.ClientStream(ctx, &server1ClientStreamStream{stream})
}

type Server1_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type server1ClientStreamStream struct {
	stream server.Stream
}

func (x *server1ClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *server1ClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1ClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1ClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1ClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *server1Handler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.Server1Handler.ServerStream(ctx, m, &server1ServerStreamStream{stream})
}

type Server1_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type server1ServerStreamStream struct {
	stream server.Stream
}

func (x *server1ServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *server1ServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1ServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1ServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1ServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *server1Handler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.Server1Handler.BidiStream(ctx, &server1BidiStreamStream{stream})
}

type Server1_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type server1BidiStreamStream struct {
	stream server.Stream
}

func (x *server1BidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *server1BidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *server1BidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *server1BidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *server1BidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *server1BidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
