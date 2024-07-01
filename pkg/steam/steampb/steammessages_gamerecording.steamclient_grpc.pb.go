// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_gamerecording.steamclient.proto

package steampb

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

// GameRecordingClipClient is the client API for GameRecordingClip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameRecordingClipClient interface {
	CreateShareClip(ctx context.Context, in *CGameRecording_CreateShareClip_Request, opts ...grpc.CallOption) (*CGameRecording_CreateShareClip_Response, error)
	DeleteSharedClip(ctx context.Context, in *CGameRecording_DeleteSharedClip_Request, opts ...grpc.CallOption) (*CGameRecording_DeleteSharedClip_Response, error)
	GetSingleSharedClip(ctx context.Context, in *CGameRecording_GetSingleSharedClip_Request, opts ...grpc.CallOption) (*CGameRecording_GetSingleSharedClip_Response, error)
}

type gameRecordingClipClient struct {
	cc grpc.ClientConnInterface
}

func NewGameRecordingClipClient(cc grpc.ClientConnInterface) GameRecordingClipClient {
	return &gameRecordingClipClient{cc}
}

func (c *gameRecordingClipClient) CreateShareClip(ctx context.Context, in *CGameRecording_CreateShareClip_Request, opts ...grpc.CallOption) (*CGameRecording_CreateShareClip_Response, error) {
	out := new(CGameRecording_CreateShareClip_Response)
	err := c.cc.Invoke(ctx, "/GameRecordingClip/CreateShareClip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRecordingClipClient) DeleteSharedClip(ctx context.Context, in *CGameRecording_DeleteSharedClip_Request, opts ...grpc.CallOption) (*CGameRecording_DeleteSharedClip_Response, error) {
	out := new(CGameRecording_DeleteSharedClip_Response)
	err := c.cc.Invoke(ctx, "/GameRecordingClip/DeleteSharedClip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRecordingClipClient) GetSingleSharedClip(ctx context.Context, in *CGameRecording_GetSingleSharedClip_Request, opts ...grpc.CallOption) (*CGameRecording_GetSingleSharedClip_Response, error) {
	out := new(CGameRecording_GetSingleSharedClip_Response)
	err := c.cc.Invoke(ctx, "/GameRecordingClip/GetSingleSharedClip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameRecordingClipServer is the server API for GameRecordingClip service.
// All implementations must embed UnimplementedGameRecordingClipServer
// for forward compatibility
type GameRecordingClipServer interface {
	CreateShareClip(context.Context, *CGameRecording_CreateShareClip_Request) (*CGameRecording_CreateShareClip_Response, error)
	DeleteSharedClip(context.Context, *CGameRecording_DeleteSharedClip_Request) (*CGameRecording_DeleteSharedClip_Response, error)
	GetSingleSharedClip(context.Context, *CGameRecording_GetSingleSharedClip_Request) (*CGameRecording_GetSingleSharedClip_Response, error)
	mustEmbedUnimplementedGameRecordingClipServer()
}

// UnimplementedGameRecordingClipServer must be embedded to have forward compatible implementations.
type UnimplementedGameRecordingClipServer struct {
}

func (UnimplementedGameRecordingClipServer) CreateShareClip(context.Context, *CGameRecording_CreateShareClip_Request) (*CGameRecording_CreateShareClip_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShareClip not implemented")
}
func (UnimplementedGameRecordingClipServer) DeleteSharedClip(context.Context, *CGameRecording_DeleteSharedClip_Request) (*CGameRecording_DeleteSharedClip_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSharedClip not implemented")
}
func (UnimplementedGameRecordingClipServer) GetSingleSharedClip(context.Context, *CGameRecording_GetSingleSharedClip_Request) (*CGameRecording_GetSingleSharedClip_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleSharedClip not implemented")
}
func (UnimplementedGameRecordingClipServer) mustEmbedUnimplementedGameRecordingClipServer() {}

// UnsafeGameRecordingClipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameRecordingClipServer will
// result in compilation errors.
type UnsafeGameRecordingClipServer interface {
	mustEmbedUnimplementedGameRecordingClipServer()
}

func RegisterGameRecordingClipServer(s grpc.ServiceRegistrar, srv GameRecordingClipServer) {
	s.RegisterService(&GameRecordingClip_ServiceDesc, srv)
}

func _GameRecordingClip_CreateShareClip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CGameRecording_CreateShareClip_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRecordingClipServer).CreateShareClip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameRecordingClip/CreateShareClip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRecordingClipServer).CreateShareClip(ctx, req.(*CGameRecording_CreateShareClip_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRecordingClip_DeleteSharedClip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CGameRecording_DeleteSharedClip_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRecordingClipServer).DeleteSharedClip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameRecordingClip/DeleteSharedClip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRecordingClipServer).DeleteSharedClip(ctx, req.(*CGameRecording_DeleteSharedClip_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRecordingClip_GetSingleSharedClip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CGameRecording_GetSingleSharedClip_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRecordingClipServer).GetSingleSharedClip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameRecordingClip/GetSingleSharedClip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRecordingClipServer).GetSingleSharedClip(ctx, req.(*CGameRecording_GetSingleSharedClip_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GameRecordingClip_ServiceDesc is the grpc.ServiceDesc for GameRecordingClip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameRecordingClip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameRecordingClip",
	HandlerType: (*GameRecordingClipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShareClip",
			Handler:    _GameRecordingClip_CreateShareClip_Handler,
		},
		{
			MethodName: "DeleteSharedClip",
			Handler:    _GameRecordingClip_DeleteSharedClip_Handler,
		},
		{
			MethodName: "GetSingleSharedClip",
			Handler:    _GameRecordingClip_GetSingleSharedClip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_gamerecording.steamclient.proto",
}

// VideoClipClient is the client API for VideoClip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClipClient interface {
	BeginGameRecordingSegmentsUpload(ctx context.Context, in *CVideo_BeginGameRecordingSegmentsUpload_Request, opts ...grpc.CallOption) (*CVideo_BeginGameRecordingSegmentsUpload_Response, error)
	CommitGameRecordingSegmentsUpload(ctx context.Context, in *CVideo_CommitGameRecordingSegmentsUpload_Request, opts ...grpc.CallOption) (*CVideo_CommitGameRecordingSegmentsUpload_Response, error)
	GetNextBatchOfSegmentsToUpload(ctx context.Context, in *CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request, opts ...grpc.CallOption) (*CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Response, error)
	CommitSegmentUploads(ctx context.Context, in *CVideo_GameRecordingCommitSegmentUploads_Request, opts ...grpc.CallOption) (*CVideo_GameRecordingCommitSegmentUploads_Response, error)
}

type videoClipClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClipClient(cc grpc.ClientConnInterface) VideoClipClient {
	return &videoClipClient{cc}
}

func (c *videoClipClient) BeginGameRecordingSegmentsUpload(ctx context.Context, in *CVideo_BeginGameRecordingSegmentsUpload_Request, opts ...grpc.CallOption) (*CVideo_BeginGameRecordingSegmentsUpload_Response, error) {
	out := new(CVideo_BeginGameRecordingSegmentsUpload_Response)
	err := c.cc.Invoke(ctx, "/VideoClip/BeginGameRecordingSegmentsUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClipClient) CommitGameRecordingSegmentsUpload(ctx context.Context, in *CVideo_CommitGameRecordingSegmentsUpload_Request, opts ...grpc.CallOption) (*CVideo_CommitGameRecordingSegmentsUpload_Response, error) {
	out := new(CVideo_CommitGameRecordingSegmentsUpload_Response)
	err := c.cc.Invoke(ctx, "/VideoClip/CommitGameRecordingSegmentsUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClipClient) GetNextBatchOfSegmentsToUpload(ctx context.Context, in *CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request, opts ...grpc.CallOption) (*CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Response, error) {
	out := new(CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Response)
	err := c.cc.Invoke(ctx, "/VideoClip/GetNextBatchOfSegmentsToUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClipClient) CommitSegmentUploads(ctx context.Context, in *CVideo_GameRecordingCommitSegmentUploads_Request, opts ...grpc.CallOption) (*CVideo_GameRecordingCommitSegmentUploads_Response, error) {
	out := new(CVideo_GameRecordingCommitSegmentUploads_Response)
	err := c.cc.Invoke(ctx, "/VideoClip/CommitSegmentUploads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoClipServer is the server API for VideoClip service.
// All implementations must embed UnimplementedVideoClipServer
// for forward compatibility
type VideoClipServer interface {
	BeginGameRecordingSegmentsUpload(context.Context, *CVideo_BeginGameRecordingSegmentsUpload_Request) (*CVideo_BeginGameRecordingSegmentsUpload_Response, error)
	CommitGameRecordingSegmentsUpload(context.Context, *CVideo_CommitGameRecordingSegmentsUpload_Request) (*CVideo_CommitGameRecordingSegmentsUpload_Response, error)
	GetNextBatchOfSegmentsToUpload(context.Context, *CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request) (*CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Response, error)
	CommitSegmentUploads(context.Context, *CVideo_GameRecordingCommitSegmentUploads_Request) (*CVideo_GameRecordingCommitSegmentUploads_Response, error)
	mustEmbedUnimplementedVideoClipServer()
}

// UnimplementedVideoClipServer must be embedded to have forward compatible implementations.
type UnimplementedVideoClipServer struct {
}

func (UnimplementedVideoClipServer) BeginGameRecordingSegmentsUpload(context.Context, *CVideo_BeginGameRecordingSegmentsUpload_Request) (*CVideo_BeginGameRecordingSegmentsUpload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginGameRecordingSegmentsUpload not implemented")
}
func (UnimplementedVideoClipServer) CommitGameRecordingSegmentsUpload(context.Context, *CVideo_CommitGameRecordingSegmentsUpload_Request) (*CVideo_CommitGameRecordingSegmentsUpload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitGameRecordingSegmentsUpload not implemented")
}
func (UnimplementedVideoClipServer) GetNextBatchOfSegmentsToUpload(context.Context, *CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request) (*CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextBatchOfSegmentsToUpload not implemented")
}
func (UnimplementedVideoClipServer) CommitSegmentUploads(context.Context, *CVideo_GameRecordingCommitSegmentUploads_Request) (*CVideo_GameRecordingCommitSegmentUploads_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitSegmentUploads not implemented")
}
func (UnimplementedVideoClipServer) mustEmbedUnimplementedVideoClipServer() {}

// UnsafeVideoClipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoClipServer will
// result in compilation errors.
type UnsafeVideoClipServer interface {
	mustEmbedUnimplementedVideoClipServer()
}

func RegisterVideoClipServer(s grpc.ServiceRegistrar, srv VideoClipServer) {
	s.RegisterService(&VideoClip_ServiceDesc, srv)
}

func _VideoClip_BeginGameRecordingSegmentsUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CVideo_BeginGameRecordingSegmentsUpload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoClipServer).BeginGameRecordingSegmentsUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoClip/BeginGameRecordingSegmentsUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoClipServer).BeginGameRecordingSegmentsUpload(ctx, req.(*CVideo_BeginGameRecordingSegmentsUpload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoClip_CommitGameRecordingSegmentsUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CVideo_CommitGameRecordingSegmentsUpload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoClipServer).CommitGameRecordingSegmentsUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoClip/CommitGameRecordingSegmentsUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoClipServer).CommitGameRecordingSegmentsUpload(ctx, req.(*CVideo_CommitGameRecordingSegmentsUpload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoClip_GetNextBatchOfSegmentsToUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoClipServer).GetNextBatchOfSegmentsToUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoClip/GetNextBatchOfSegmentsToUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoClipServer).GetNextBatchOfSegmentsToUpload(ctx, req.(*CVideo_GameRecordingGetNextBatchOfSegmentsToUpload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoClip_CommitSegmentUploads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CVideo_GameRecordingCommitSegmentUploads_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoClipServer).CommitSegmentUploads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoClip/CommitSegmentUploads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoClipServer).CommitSegmentUploads(ctx, req.(*CVideo_GameRecordingCommitSegmentUploads_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoClip_ServiceDesc is the grpc.ServiceDesc for VideoClip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoClip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VideoClip",
	HandlerType: (*VideoClipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeginGameRecordingSegmentsUpload",
			Handler:    _VideoClip_BeginGameRecordingSegmentsUpload_Handler,
		},
		{
			MethodName: "CommitGameRecordingSegmentsUpload",
			Handler:    _VideoClip_CommitGameRecordingSegmentsUpload_Handler,
		},
		{
			MethodName: "GetNextBatchOfSegmentsToUpload",
			Handler:    _VideoClip_GetNextBatchOfSegmentsToUpload_Handler,
		},
		{
			MethodName: "CommitSegmentUploads",
			Handler:    _VideoClip_CommitSegmentUploads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_gamerecording.steamclient.proto",
}
