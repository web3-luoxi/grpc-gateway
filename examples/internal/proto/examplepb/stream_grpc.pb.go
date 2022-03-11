// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package examplepb

import (
	context "context"
	sub "github.com/web3-luoxi/grpc-gateway/v2/examples/internal/proto/sub"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error)
	Download(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StreamService_DownloadClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/grpc.gateway.examples.internal.proto.examplepb.StreamService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkCreateClient{stream}
	return x, nil
}

type StreamService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type streamServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkCreateClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/grpc.gateway.examples.internal.proto.examplepb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/grpc.gateway.examples.internal.proto.examplepb.StreamService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkEchoClient{stream}
	return x, nil
}

type StreamService_BulkEchoClient interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ClientStream
}

type streamServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkEchoClient) Send(m *sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkEchoClient) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Download(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (StreamService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[3], "/grpc.gateway.examples.internal.proto.examplepb.StreamService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_DownloadClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type streamServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *streamServiceDownloadClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations should embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	BulkCreate(StreamService_BulkCreateServer) error
	List(*emptypb.Empty, StreamService_ListServer) error
	BulkEcho(StreamService_BulkEchoServer) error
	Download(*emptypb.Empty, StreamService_DownloadServer) error
}

// UnimplementedStreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) BulkCreate(StreamService_BulkCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkCreate not implemented")
}
func (UnimplementedStreamServiceServer) List(*emptypb.Empty, StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStreamServiceServer) BulkEcho(StreamService_BulkEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkEcho not implemented")
}
func (UnimplementedStreamServiceServer) Download(*emptypb.Empty, StreamService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkCreate(&streamServiceBulkCreateServer{stream})
}

type StreamService_BulkCreateServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type streamServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkCreateServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkEcho(&streamServiceBulkEchoServer{stream})
}

type StreamService_BulkEchoServer interface {
	Send(*sub.StringMessage) error
	Recv() (*sub.StringMessage, error)
	grpc.ServerStream
}

type streamServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkEchoServer) Send(m *sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkEchoServer) Recv() (*sub.StringMessage, error) {
	m := new(sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).Download(m, &streamServiceDownloadServer{stream})
}

type StreamService_DownloadServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type streamServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *streamServiceDownloadServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _StreamService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _StreamService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _StreamService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/stream.proto",
}
