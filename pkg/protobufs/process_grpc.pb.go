// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: protobufs/process.proto

package protobufs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProcessClient is the client API for Process service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessClient interface {
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error)
	StreamStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Process_StreamStatusClient, error)
	WriteCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StreamLogs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Process_StreamLogsClient, error)
}

type processClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessClient(cc grpc.ClientConnInterface) ProcessClient {
	return &processClient{cc}
}

func (c *processClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Process/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processClient) StreamStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Process_StreamStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Process_ServiceDesc.Streams[0], "/Process/StreamStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &processStreamStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Process_StreamStatusClient interface {
	Recv() (*Status, error)
	grpc.ClientStream
}

type processStreamStatusClient struct {
	grpc.ClientStream
}

func (x *processStreamStatusClient) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processClient) WriteCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Process/WriteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processClient) StreamLogs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Process_StreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Process_ServiceDesc.Streams[1], "/Process/StreamLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &processStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Process_StreamLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type processStreamLogsClient struct {
	grpc.ClientStream
}

func (x *processStreamLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessServer is the server API for Process service.
// All implementations must embed UnimplementedProcessServer
// for forward compatibility
type ProcessServer interface {
	GetStatus(context.Context, *emptypb.Empty) (*Status, error)
	StreamStatus(*emptypb.Empty, Process_StreamStatusServer) error
	WriteCommand(context.Context, *Command) (*emptypb.Empty, error)
	StreamLogs(*emptypb.Empty, Process_StreamLogsServer) error
	mustEmbedUnimplementedProcessServer()
}

// UnimplementedProcessServer must be embedded to have forward compatible implementations.
type UnimplementedProcessServer struct {
}

func (UnimplementedProcessServer) GetStatus(context.Context, *emptypb.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedProcessServer) StreamStatus(*emptypb.Empty, Process_StreamStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStatus not implemented")
}
func (UnimplementedProcessServer) WriteCommand(context.Context, *Command) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteCommand not implemented")
}
func (UnimplementedProcessServer) StreamLogs(*emptypb.Empty, Process_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedProcessServer) mustEmbedUnimplementedProcessServer() {}

// UnsafeProcessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessServer will
// result in compilation errors.
type UnsafeProcessServer interface {
	mustEmbedUnimplementedProcessServer()
}

func RegisterProcessServer(s grpc.ServiceRegistrar, srv ProcessServer) {
	s.RegisterService(&Process_ServiceDesc, srv)
}

func _Process_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Process/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Process_StreamStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessServer).StreamStatus(m, &processStreamStatusServer{stream})
}

type Process_StreamStatusServer interface {
	Send(*Status) error
	grpc.ServerStream
}

type processStreamStatusServer struct {
	grpc.ServerStream
}

func (x *processStreamStatusServer) Send(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func _Process_WriteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).WriteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Process/WriteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).WriteCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Process_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessServer).StreamLogs(m, &processStreamLogsServer{stream})
}

type Process_StreamLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type processStreamLogsServer struct {
	grpc.ServerStream
}

func (x *processStreamLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

// Process_ServiceDesc is the grpc.ServiceDesc for Process service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Process_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Process",
	HandlerType: (*ProcessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Process_GetStatus_Handler,
		},
		{
			MethodName: "WriteCommand",
			Handler:    _Process_WriteCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamStatus",
			Handler:       _Process_StreamStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamLogs",
			Handler:       _Process_StreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobufs/process.proto",
}
