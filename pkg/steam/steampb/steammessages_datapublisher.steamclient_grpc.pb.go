// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_datapublisher.steamclient.proto

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

// DataPublisherClient is the client API for DataPublisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataPublisherClient interface {
	ClientContentCorruptionReport(ctx context.Context, in *CDataPublisher_ClientContentCorruptionReport_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	ClientUpdateAppJobReport(ctx context.Context, in *CDataPublisher_ClientUpdateAppJob_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	GetVRDeviceInfo(ctx context.Context, in *CDataPublisher_GetVRDeviceInfo_Request, opts ...grpc.CallOption) (*CDataPublisher_GetVRDeviceInfo_Response, error)
	SetVRDeviceInfoAggregationReference(ctx context.Context, in *CDataPublisher_SetVRDeviceInfoAggregationReference_Request, opts ...grpc.CallOption) (*CDataPublisher_SetVRDeviceInfoAggregationReference_Response, error)
	AddVRDeviceInfo(ctx context.Context, in *CDataPublisher_AddVRDeviceInfo_Request, opts ...grpc.CallOption) (*CDataPublisher_AddVRDeviceInfo_Response, error)
}

type dataPublisherClient struct {
	cc grpc.ClientConnInterface
}

func NewDataPublisherClient(cc grpc.ClientConnInterface) DataPublisherClient {
	return &dataPublisherClient{cc}
}

func (c *dataPublisherClient) ClientContentCorruptionReport(ctx context.Context, in *CDataPublisher_ClientContentCorruptionReport_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/DataPublisher/ClientContentCorruptionReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPublisherClient) ClientUpdateAppJobReport(ctx context.Context, in *CDataPublisher_ClientUpdateAppJob_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/DataPublisher/ClientUpdateAppJobReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPublisherClient) GetVRDeviceInfo(ctx context.Context, in *CDataPublisher_GetVRDeviceInfo_Request, opts ...grpc.CallOption) (*CDataPublisher_GetVRDeviceInfo_Response, error) {
	out := new(CDataPublisher_GetVRDeviceInfo_Response)
	err := c.cc.Invoke(ctx, "/DataPublisher/GetVRDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPublisherClient) SetVRDeviceInfoAggregationReference(ctx context.Context, in *CDataPublisher_SetVRDeviceInfoAggregationReference_Request, opts ...grpc.CallOption) (*CDataPublisher_SetVRDeviceInfoAggregationReference_Response, error) {
	out := new(CDataPublisher_SetVRDeviceInfoAggregationReference_Response)
	err := c.cc.Invoke(ctx, "/DataPublisher/SetVRDeviceInfoAggregationReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPublisherClient) AddVRDeviceInfo(ctx context.Context, in *CDataPublisher_AddVRDeviceInfo_Request, opts ...grpc.CallOption) (*CDataPublisher_AddVRDeviceInfo_Response, error) {
	out := new(CDataPublisher_AddVRDeviceInfo_Response)
	err := c.cc.Invoke(ctx, "/DataPublisher/AddVRDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataPublisherServer is the server API for DataPublisher service.
// All implementations must embed UnimplementedDataPublisherServer
// for forward compatibility
type DataPublisherServer interface {
	ClientContentCorruptionReport(context.Context, *CDataPublisher_ClientContentCorruptionReport_Notification) (*NoResponse, error)
	ClientUpdateAppJobReport(context.Context, *CDataPublisher_ClientUpdateAppJob_Notification) (*NoResponse, error)
	GetVRDeviceInfo(context.Context, *CDataPublisher_GetVRDeviceInfo_Request) (*CDataPublisher_GetVRDeviceInfo_Response, error)
	SetVRDeviceInfoAggregationReference(context.Context, *CDataPublisher_SetVRDeviceInfoAggregationReference_Request) (*CDataPublisher_SetVRDeviceInfoAggregationReference_Response, error)
	AddVRDeviceInfo(context.Context, *CDataPublisher_AddVRDeviceInfo_Request) (*CDataPublisher_AddVRDeviceInfo_Response, error)
	mustEmbedUnimplementedDataPublisherServer()
}

// UnimplementedDataPublisherServer must be embedded to have forward compatible implementations.
type UnimplementedDataPublisherServer struct {
}

func (UnimplementedDataPublisherServer) ClientContentCorruptionReport(context.Context, *CDataPublisher_ClientContentCorruptionReport_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientContentCorruptionReport not implemented")
}
func (UnimplementedDataPublisherServer) ClientUpdateAppJobReport(context.Context, *CDataPublisher_ClientUpdateAppJob_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUpdateAppJobReport not implemented")
}
func (UnimplementedDataPublisherServer) GetVRDeviceInfo(context.Context, *CDataPublisher_GetVRDeviceInfo_Request) (*CDataPublisher_GetVRDeviceInfo_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVRDeviceInfo not implemented")
}
func (UnimplementedDataPublisherServer) SetVRDeviceInfoAggregationReference(context.Context, *CDataPublisher_SetVRDeviceInfoAggregationReference_Request) (*CDataPublisher_SetVRDeviceInfoAggregationReference_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVRDeviceInfoAggregationReference not implemented")
}
func (UnimplementedDataPublisherServer) AddVRDeviceInfo(context.Context, *CDataPublisher_AddVRDeviceInfo_Request) (*CDataPublisher_AddVRDeviceInfo_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVRDeviceInfo not implemented")
}
func (UnimplementedDataPublisherServer) mustEmbedUnimplementedDataPublisherServer() {}

// UnsafeDataPublisherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataPublisherServer will
// result in compilation errors.
type UnsafeDataPublisherServer interface {
	mustEmbedUnimplementedDataPublisherServer()
}

func RegisterDataPublisherServer(s grpc.ServiceRegistrar, srv DataPublisherServer) {
	s.RegisterService(&DataPublisher_ServiceDesc, srv)
}

func _DataPublisher_ClientContentCorruptionReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDataPublisher_ClientContentCorruptionReport_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPublisherServer).ClientContentCorruptionReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataPublisher/ClientContentCorruptionReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPublisherServer).ClientContentCorruptionReport(ctx, req.(*CDataPublisher_ClientContentCorruptionReport_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPublisher_ClientUpdateAppJobReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDataPublisher_ClientUpdateAppJob_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPublisherServer).ClientUpdateAppJobReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataPublisher/ClientUpdateAppJobReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPublisherServer).ClientUpdateAppJobReport(ctx, req.(*CDataPublisher_ClientUpdateAppJob_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPublisher_GetVRDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDataPublisher_GetVRDeviceInfo_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPublisherServer).GetVRDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataPublisher/GetVRDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPublisherServer).GetVRDeviceInfo(ctx, req.(*CDataPublisher_GetVRDeviceInfo_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPublisher_SetVRDeviceInfoAggregationReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDataPublisher_SetVRDeviceInfoAggregationReference_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPublisherServer).SetVRDeviceInfoAggregationReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataPublisher/SetVRDeviceInfoAggregationReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPublisherServer).SetVRDeviceInfoAggregationReference(ctx, req.(*CDataPublisher_SetVRDeviceInfoAggregationReference_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPublisher_AddVRDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDataPublisher_AddVRDeviceInfo_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPublisherServer).AddVRDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataPublisher/AddVRDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPublisherServer).AddVRDeviceInfo(ctx, req.(*CDataPublisher_AddVRDeviceInfo_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// DataPublisher_ServiceDesc is the grpc.ServiceDesc for DataPublisher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataPublisher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataPublisher",
	HandlerType: (*DataPublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientContentCorruptionReport",
			Handler:    _DataPublisher_ClientContentCorruptionReport_Handler,
		},
		{
			MethodName: "ClientUpdateAppJobReport",
			Handler:    _DataPublisher_ClientUpdateAppJobReport_Handler,
		},
		{
			MethodName: "GetVRDeviceInfo",
			Handler:    _DataPublisher_GetVRDeviceInfo_Handler,
		},
		{
			MethodName: "SetVRDeviceInfoAggregationReference",
			Handler:    _DataPublisher_SetVRDeviceInfoAggregationReference_Handler,
		},
		{
			MethodName: "AddVRDeviceInfo",
			Handler:    _DataPublisher_AddVRDeviceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_datapublisher.steamclient.proto",
}

// ValveHWSurveyClient is the client API for ValveHWSurvey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValveHWSurveyClient interface {
	GetSurveySchedule(ctx context.Context, in *CValveHWSurvey_GetSurveySchedule_Request, opts ...grpc.CallOption) (*CValveHWSurvey_GetSurveySchedule_Response, error)
}

type valveHWSurveyClient struct {
	cc grpc.ClientConnInterface
}

func NewValveHWSurveyClient(cc grpc.ClientConnInterface) ValveHWSurveyClient {
	return &valveHWSurveyClient{cc}
}

func (c *valveHWSurveyClient) GetSurveySchedule(ctx context.Context, in *CValveHWSurvey_GetSurveySchedule_Request, opts ...grpc.CallOption) (*CValveHWSurvey_GetSurveySchedule_Response, error) {
	out := new(CValveHWSurvey_GetSurveySchedule_Response)
	err := c.cc.Invoke(ctx, "/ValveHWSurvey/GetSurveySchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValveHWSurveyServer is the server API for ValveHWSurvey service.
// All implementations must embed UnimplementedValveHWSurveyServer
// for forward compatibility
type ValveHWSurveyServer interface {
	GetSurveySchedule(context.Context, *CValveHWSurvey_GetSurveySchedule_Request) (*CValveHWSurvey_GetSurveySchedule_Response, error)
	mustEmbedUnimplementedValveHWSurveyServer()
}

// UnimplementedValveHWSurveyServer must be embedded to have forward compatible implementations.
type UnimplementedValveHWSurveyServer struct {
}

func (UnimplementedValveHWSurveyServer) GetSurveySchedule(context.Context, *CValveHWSurvey_GetSurveySchedule_Request) (*CValveHWSurvey_GetSurveySchedule_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSurveySchedule not implemented")
}
func (UnimplementedValveHWSurveyServer) mustEmbedUnimplementedValveHWSurveyServer() {}

// UnsafeValveHWSurveyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValveHWSurveyServer will
// result in compilation errors.
type UnsafeValveHWSurveyServer interface {
	mustEmbedUnimplementedValveHWSurveyServer()
}

func RegisterValveHWSurveyServer(s grpc.ServiceRegistrar, srv ValveHWSurveyServer) {
	s.RegisterService(&ValveHWSurvey_ServiceDesc, srv)
}

func _ValveHWSurvey_GetSurveySchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CValveHWSurvey_GetSurveySchedule_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValveHWSurveyServer).GetSurveySchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ValveHWSurvey/GetSurveySchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValveHWSurveyServer).GetSurveySchedule(ctx, req.(*CValveHWSurvey_GetSurveySchedule_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ValveHWSurvey_ServiceDesc is the grpc.ServiceDesc for ValveHWSurvey service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValveHWSurvey_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ValveHWSurvey",
	HandlerType: (*ValveHWSurveyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSurveySchedule",
			Handler:    _ValveHWSurvey_GetSurveySchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_datapublisher.steamclient.proto",
}
