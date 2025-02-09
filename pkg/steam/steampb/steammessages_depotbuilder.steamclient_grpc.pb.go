// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_depotbuilder.steamclient.proto

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

// ContentBuilderClient is the client API for ContentBuilder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentBuilderClient interface {
	InitDepotBuild(ctx context.Context, in *CContentBuilder_InitDepotBuild_Request, opts ...grpc.CallOption) (*CContentBuilder_InitDepotBuild_Response, error)
	StartDepotUpload(ctx context.Context, in *CContentBuilder_StartDepotUpload_Request, opts ...grpc.CallOption) (*CContentBuilder_StartDepotUpload_Response, error)
	GetMissingDepotChunks(ctx context.Context, in *CContentBuilder_GetMissingDepotChunks_Request, opts ...grpc.CallOption) (*CContentBuilder_GetMissingDepotChunks_Response, error)
	FinishDepotUpload(ctx context.Context, in *CContentBuilder_FinishDepotUpload_Request, opts ...grpc.CallOption) (*CContentBuilder_FinishDepotUpload_Response, error)
	CommitAppBuild(ctx context.Context, in *CContentBuilder_CommitAppBuild_Request, opts ...grpc.CallOption) (*CContentBuilder_CommitAppBuild_Response, error)
	SignInstallScript(ctx context.Context, in *CContentBuilder_SignInstallScript_Request, opts ...grpc.CallOption) (*CContentBuilder_SignInstallScript_Response, error)
}

type contentBuilderClient struct {
	cc grpc.ClientConnInterface
}

func NewContentBuilderClient(cc grpc.ClientConnInterface) ContentBuilderClient {
	return &contentBuilderClient{cc}
}

func (c *contentBuilderClient) InitDepotBuild(ctx context.Context, in *CContentBuilder_InitDepotBuild_Request, opts ...grpc.CallOption) (*CContentBuilder_InitDepotBuild_Response, error) {
	out := new(CContentBuilder_InitDepotBuild_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/InitDepotBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentBuilderClient) StartDepotUpload(ctx context.Context, in *CContentBuilder_StartDepotUpload_Request, opts ...grpc.CallOption) (*CContentBuilder_StartDepotUpload_Response, error) {
	out := new(CContentBuilder_StartDepotUpload_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/StartDepotUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentBuilderClient) GetMissingDepotChunks(ctx context.Context, in *CContentBuilder_GetMissingDepotChunks_Request, opts ...grpc.CallOption) (*CContentBuilder_GetMissingDepotChunks_Response, error) {
	out := new(CContentBuilder_GetMissingDepotChunks_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/GetMissingDepotChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentBuilderClient) FinishDepotUpload(ctx context.Context, in *CContentBuilder_FinishDepotUpload_Request, opts ...grpc.CallOption) (*CContentBuilder_FinishDepotUpload_Response, error) {
	out := new(CContentBuilder_FinishDepotUpload_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/FinishDepotUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentBuilderClient) CommitAppBuild(ctx context.Context, in *CContentBuilder_CommitAppBuild_Request, opts ...grpc.CallOption) (*CContentBuilder_CommitAppBuild_Response, error) {
	out := new(CContentBuilder_CommitAppBuild_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/CommitAppBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentBuilderClient) SignInstallScript(ctx context.Context, in *CContentBuilder_SignInstallScript_Request, opts ...grpc.CallOption) (*CContentBuilder_SignInstallScript_Response, error) {
	out := new(CContentBuilder_SignInstallScript_Response)
	err := c.cc.Invoke(ctx, "/ContentBuilder/SignInstallScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentBuilderServer is the server API for ContentBuilder service.
// All implementations must embed UnimplementedContentBuilderServer
// for forward compatibility
type ContentBuilderServer interface {
	InitDepotBuild(context.Context, *CContentBuilder_InitDepotBuild_Request) (*CContentBuilder_InitDepotBuild_Response, error)
	StartDepotUpload(context.Context, *CContentBuilder_StartDepotUpload_Request) (*CContentBuilder_StartDepotUpload_Response, error)
	GetMissingDepotChunks(context.Context, *CContentBuilder_GetMissingDepotChunks_Request) (*CContentBuilder_GetMissingDepotChunks_Response, error)
	FinishDepotUpload(context.Context, *CContentBuilder_FinishDepotUpload_Request) (*CContentBuilder_FinishDepotUpload_Response, error)
	CommitAppBuild(context.Context, *CContentBuilder_CommitAppBuild_Request) (*CContentBuilder_CommitAppBuild_Response, error)
	SignInstallScript(context.Context, *CContentBuilder_SignInstallScript_Request) (*CContentBuilder_SignInstallScript_Response, error)
	mustEmbedUnimplementedContentBuilderServer()
}

// UnimplementedContentBuilderServer must be embedded to have forward compatible implementations.
type UnimplementedContentBuilderServer struct {
}

func (UnimplementedContentBuilderServer) InitDepotBuild(context.Context, *CContentBuilder_InitDepotBuild_Request) (*CContentBuilder_InitDepotBuild_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDepotBuild not implemented")
}
func (UnimplementedContentBuilderServer) StartDepotUpload(context.Context, *CContentBuilder_StartDepotUpload_Request) (*CContentBuilder_StartDepotUpload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDepotUpload not implemented")
}
func (UnimplementedContentBuilderServer) GetMissingDepotChunks(context.Context, *CContentBuilder_GetMissingDepotChunks_Request) (*CContentBuilder_GetMissingDepotChunks_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissingDepotChunks not implemented")
}
func (UnimplementedContentBuilderServer) FinishDepotUpload(context.Context, *CContentBuilder_FinishDepotUpload_Request) (*CContentBuilder_FinishDepotUpload_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishDepotUpload not implemented")
}
func (UnimplementedContentBuilderServer) CommitAppBuild(context.Context, *CContentBuilder_CommitAppBuild_Request) (*CContentBuilder_CommitAppBuild_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitAppBuild not implemented")
}
func (UnimplementedContentBuilderServer) SignInstallScript(context.Context, *CContentBuilder_SignInstallScript_Request) (*CContentBuilder_SignInstallScript_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInstallScript not implemented")
}
func (UnimplementedContentBuilderServer) mustEmbedUnimplementedContentBuilderServer() {}

// UnsafeContentBuilderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentBuilderServer will
// result in compilation errors.
type UnsafeContentBuilderServer interface {
	mustEmbedUnimplementedContentBuilderServer()
}

func RegisterContentBuilderServer(s grpc.ServiceRegistrar, srv ContentBuilderServer) {
	s.RegisterService(&ContentBuilder_ServiceDesc, srv)
}

func _ContentBuilder_InitDepotBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_InitDepotBuild_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).InitDepotBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/InitDepotBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).InitDepotBuild(ctx, req.(*CContentBuilder_InitDepotBuild_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentBuilder_StartDepotUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_StartDepotUpload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).StartDepotUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/StartDepotUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).StartDepotUpload(ctx, req.(*CContentBuilder_StartDepotUpload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentBuilder_GetMissingDepotChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_GetMissingDepotChunks_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).GetMissingDepotChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/GetMissingDepotChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).GetMissingDepotChunks(ctx, req.(*CContentBuilder_GetMissingDepotChunks_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentBuilder_FinishDepotUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_FinishDepotUpload_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).FinishDepotUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/FinishDepotUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).FinishDepotUpload(ctx, req.(*CContentBuilder_FinishDepotUpload_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentBuilder_CommitAppBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_CommitAppBuild_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).CommitAppBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/CommitAppBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).CommitAppBuild(ctx, req.(*CContentBuilder_CommitAppBuild_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentBuilder_SignInstallScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CContentBuilder_SignInstallScript_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentBuilderServer).SignInstallScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentBuilder/SignInstallScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentBuilderServer).SignInstallScript(ctx, req.(*CContentBuilder_SignInstallScript_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentBuilder_ServiceDesc is the grpc.ServiceDesc for ContentBuilder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentBuilder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ContentBuilder",
	HandlerType: (*ContentBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitDepotBuild",
			Handler:    _ContentBuilder_InitDepotBuild_Handler,
		},
		{
			MethodName: "StartDepotUpload",
			Handler:    _ContentBuilder_StartDepotUpload_Handler,
		},
		{
			MethodName: "GetMissingDepotChunks",
			Handler:    _ContentBuilder_GetMissingDepotChunks_Handler,
		},
		{
			MethodName: "FinishDepotUpload",
			Handler:    _ContentBuilder_FinishDepotUpload_Handler,
		},
		{
			MethodName: "CommitAppBuild",
			Handler:    _ContentBuilder_CommitAppBuild_Handler,
		},
		{
			MethodName: "SignInstallScript",
			Handler:    _ContentBuilder_SignInstallScript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_depotbuilder.steamclient.proto",
}
