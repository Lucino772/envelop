// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_parties.steamclient.proto

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

// PartiesClient is the client API for Parties service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartiesClient interface {
	JoinParty(ctx context.Context, in *CParties_JoinParty_Request, opts ...grpc.CallOption) (*CParties_JoinParty_Response, error)
	CreateBeacon(ctx context.Context, in *CParties_CreateBeacon_Request, opts ...grpc.CallOption) (*CParties_CreateBeacon_Response, error)
	OnReservationCompleted(ctx context.Context, in *CParties_OnReservationCompleted_Request, opts ...grpc.CallOption) (*CParties_OnReservationCompleted_Response, error)
	CancelReservation(ctx context.Context, in *CParties_CancelReservation_Request, opts ...grpc.CallOption) (*CParties_CancelReservation_Response, error)
	ChangeNumOpenSlots(ctx context.Context, in *CParties_ChangeNumOpenSlots_Request, opts ...grpc.CallOption) (*CParties_ChangeNumOpenSlots_Response, error)
	DestroyBeacon(ctx context.Context, in *CParties_DestroyBeacon_Request, opts ...grpc.CallOption) (*CParties_DestroyBeacon_Response, error)
}

type partiesClient struct {
	cc grpc.ClientConnInterface
}

func NewPartiesClient(cc grpc.ClientConnInterface) PartiesClient {
	return &partiesClient{cc}
}

func (c *partiesClient) JoinParty(ctx context.Context, in *CParties_JoinParty_Request, opts ...grpc.CallOption) (*CParties_JoinParty_Response, error) {
	out := new(CParties_JoinParty_Response)
	err := c.cc.Invoke(ctx, "/Parties/JoinParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesClient) CreateBeacon(ctx context.Context, in *CParties_CreateBeacon_Request, opts ...grpc.CallOption) (*CParties_CreateBeacon_Response, error) {
	out := new(CParties_CreateBeacon_Response)
	err := c.cc.Invoke(ctx, "/Parties/CreateBeacon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesClient) OnReservationCompleted(ctx context.Context, in *CParties_OnReservationCompleted_Request, opts ...grpc.CallOption) (*CParties_OnReservationCompleted_Response, error) {
	out := new(CParties_OnReservationCompleted_Response)
	err := c.cc.Invoke(ctx, "/Parties/OnReservationCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesClient) CancelReservation(ctx context.Context, in *CParties_CancelReservation_Request, opts ...grpc.CallOption) (*CParties_CancelReservation_Response, error) {
	out := new(CParties_CancelReservation_Response)
	err := c.cc.Invoke(ctx, "/Parties/CancelReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesClient) ChangeNumOpenSlots(ctx context.Context, in *CParties_ChangeNumOpenSlots_Request, opts ...grpc.CallOption) (*CParties_ChangeNumOpenSlots_Response, error) {
	out := new(CParties_ChangeNumOpenSlots_Response)
	err := c.cc.Invoke(ctx, "/Parties/ChangeNumOpenSlots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesClient) DestroyBeacon(ctx context.Context, in *CParties_DestroyBeacon_Request, opts ...grpc.CallOption) (*CParties_DestroyBeacon_Response, error) {
	out := new(CParties_DestroyBeacon_Response)
	err := c.cc.Invoke(ctx, "/Parties/DestroyBeacon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartiesServer is the server API for Parties service.
// All implementations must embed UnimplementedPartiesServer
// for forward compatibility
type PartiesServer interface {
	JoinParty(context.Context, *CParties_JoinParty_Request) (*CParties_JoinParty_Response, error)
	CreateBeacon(context.Context, *CParties_CreateBeacon_Request) (*CParties_CreateBeacon_Response, error)
	OnReservationCompleted(context.Context, *CParties_OnReservationCompleted_Request) (*CParties_OnReservationCompleted_Response, error)
	CancelReservation(context.Context, *CParties_CancelReservation_Request) (*CParties_CancelReservation_Response, error)
	ChangeNumOpenSlots(context.Context, *CParties_ChangeNumOpenSlots_Request) (*CParties_ChangeNumOpenSlots_Response, error)
	DestroyBeacon(context.Context, *CParties_DestroyBeacon_Request) (*CParties_DestroyBeacon_Response, error)
	mustEmbedUnimplementedPartiesServer()
}

// UnimplementedPartiesServer must be embedded to have forward compatible implementations.
type UnimplementedPartiesServer struct {
}

func (UnimplementedPartiesServer) JoinParty(context.Context, *CParties_JoinParty_Request) (*CParties_JoinParty_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinParty not implemented")
}
func (UnimplementedPartiesServer) CreateBeacon(context.Context, *CParties_CreateBeacon_Request) (*CParties_CreateBeacon_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBeacon not implemented")
}
func (UnimplementedPartiesServer) OnReservationCompleted(context.Context, *CParties_OnReservationCompleted_Request) (*CParties_OnReservationCompleted_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnReservationCompleted not implemented")
}
func (UnimplementedPartiesServer) CancelReservation(context.Context, *CParties_CancelReservation_Request) (*CParties_CancelReservation_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedPartiesServer) ChangeNumOpenSlots(context.Context, *CParties_ChangeNumOpenSlots_Request) (*CParties_ChangeNumOpenSlots_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeNumOpenSlots not implemented")
}
func (UnimplementedPartiesServer) DestroyBeacon(context.Context, *CParties_DestroyBeacon_Request) (*CParties_DestroyBeacon_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyBeacon not implemented")
}
func (UnimplementedPartiesServer) mustEmbedUnimplementedPartiesServer() {}

// UnsafePartiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartiesServer will
// result in compilation errors.
type UnsafePartiesServer interface {
	mustEmbedUnimplementedPartiesServer()
}

func RegisterPartiesServer(s grpc.ServiceRegistrar, srv PartiesServer) {
	s.RegisterService(&Parties_ServiceDesc, srv)
}

func _Parties_JoinParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_JoinParty_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).JoinParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/JoinParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).JoinParty(ctx, req.(*CParties_JoinParty_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parties_CreateBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_CreateBeacon_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).CreateBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/CreateBeacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).CreateBeacon(ctx, req.(*CParties_CreateBeacon_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parties_OnReservationCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_OnReservationCompleted_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).OnReservationCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/OnReservationCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).OnReservationCompleted(ctx, req.(*CParties_OnReservationCompleted_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parties_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_CancelReservation_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/CancelReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).CancelReservation(ctx, req.(*CParties_CancelReservation_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parties_ChangeNumOpenSlots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_ChangeNumOpenSlots_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).ChangeNumOpenSlots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/ChangeNumOpenSlots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).ChangeNumOpenSlots(ctx, req.(*CParties_ChangeNumOpenSlots_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parties_DestroyBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CParties_DestroyBeacon_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServer).DestroyBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Parties/DestroyBeacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServer).DestroyBeacon(ctx, req.(*CParties_DestroyBeacon_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Parties_ServiceDesc is the grpc.ServiceDesc for Parties service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parties_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Parties",
	HandlerType: (*PartiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinParty",
			Handler:    _Parties_JoinParty_Handler,
		},
		{
			MethodName: "CreateBeacon",
			Handler:    _Parties_CreateBeacon_Handler,
		},
		{
			MethodName: "OnReservationCompleted",
			Handler:    _Parties_OnReservationCompleted_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _Parties_CancelReservation_Handler,
		},
		{
			MethodName: "ChangeNumOpenSlots",
			Handler:    _Parties_ChangeNumOpenSlots_Handler,
		},
		{
			MethodName: "DestroyBeacon",
			Handler:    _Parties_DestroyBeacon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_parties.steamclient.proto",
}
