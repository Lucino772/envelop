// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_site_license.steamclient.proto

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

// SiteManagerClientClient is the client API for SiteManagerClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteManagerClientClient interface {
	IncomingClient(ctx context.Context, in *CSiteManagerClient_IncomingClient_Request, opts ...grpc.CallOption) (*CSiteManagerClient_IncomingClient_Response, error)
	ClientSeatCheckoutNotification(ctx context.Context, in *CSiteLicense_ClientSeatCheckout_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	TrackedPaymentsNotification(ctx context.Context, in *CSiteManagerClient_TrackedPayments_Notification, opts ...grpc.CallOption) (*NoResponse, error)
}

type siteManagerClientClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteManagerClientClient(cc grpc.ClientConnInterface) SiteManagerClientClient {
	return &siteManagerClientClient{cc}
}

func (c *siteManagerClientClient) IncomingClient(ctx context.Context, in *CSiteManagerClient_IncomingClient_Request, opts ...grpc.CallOption) (*CSiteManagerClient_IncomingClient_Response, error) {
	out := new(CSiteManagerClient_IncomingClient_Response)
	err := c.cc.Invoke(ctx, "/SiteManagerClient/IncomingClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteManagerClientClient) ClientSeatCheckoutNotification(ctx context.Context, in *CSiteLicense_ClientSeatCheckout_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/SiteManagerClient/ClientSeatCheckoutNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteManagerClientClient) TrackedPaymentsNotification(ctx context.Context, in *CSiteManagerClient_TrackedPayments_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/SiteManagerClient/TrackedPaymentsNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteManagerClientServer is the server API for SiteManagerClient service.
// All implementations must embed UnimplementedSiteManagerClientServer
// for forward compatibility
type SiteManagerClientServer interface {
	IncomingClient(context.Context, *CSiteManagerClient_IncomingClient_Request) (*CSiteManagerClient_IncomingClient_Response, error)
	ClientSeatCheckoutNotification(context.Context, *CSiteLicense_ClientSeatCheckout_Notification) (*NoResponse, error)
	TrackedPaymentsNotification(context.Context, *CSiteManagerClient_TrackedPayments_Notification) (*NoResponse, error)
	mustEmbedUnimplementedSiteManagerClientServer()
}

// UnimplementedSiteManagerClientServer must be embedded to have forward compatible implementations.
type UnimplementedSiteManagerClientServer struct {
}

func (UnimplementedSiteManagerClientServer) IncomingClient(context.Context, *CSiteManagerClient_IncomingClient_Request) (*CSiteManagerClient_IncomingClient_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncomingClient not implemented")
}
func (UnimplementedSiteManagerClientServer) ClientSeatCheckoutNotification(context.Context, *CSiteLicense_ClientSeatCheckout_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientSeatCheckoutNotification not implemented")
}
func (UnimplementedSiteManagerClientServer) TrackedPaymentsNotification(context.Context, *CSiteManagerClient_TrackedPayments_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackedPaymentsNotification not implemented")
}
func (UnimplementedSiteManagerClientServer) mustEmbedUnimplementedSiteManagerClientServer() {}

// UnsafeSiteManagerClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteManagerClientServer will
// result in compilation errors.
type UnsafeSiteManagerClientServer interface {
	mustEmbedUnimplementedSiteManagerClientServer()
}

func RegisterSiteManagerClientServer(s grpc.ServiceRegistrar, srv SiteManagerClientServer) {
	s.RegisterService(&SiteManagerClient_ServiceDesc, srv)
}

func _SiteManagerClient_IncomingClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteManagerClient_IncomingClient_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteManagerClientServer).IncomingClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteManagerClient/IncomingClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteManagerClientServer).IncomingClient(ctx, req.(*CSiteManagerClient_IncomingClient_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteManagerClient_ClientSeatCheckoutNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_ClientSeatCheckout_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteManagerClientServer).ClientSeatCheckoutNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteManagerClient/ClientSeatCheckoutNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteManagerClientServer).ClientSeatCheckoutNotification(ctx, req.(*CSiteLicense_ClientSeatCheckout_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteManagerClient_TrackedPaymentsNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteManagerClient_TrackedPayments_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteManagerClientServer).TrackedPaymentsNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteManagerClient/TrackedPaymentsNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteManagerClientServer).TrackedPaymentsNotification(ctx, req.(*CSiteManagerClient_TrackedPayments_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteManagerClient_ServiceDesc is the grpc.ServiceDesc for SiteManagerClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteManagerClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SiteManagerClient",
	HandlerType: (*SiteManagerClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncomingClient",
			Handler:    _SiteManagerClient_IncomingClient_Handler,
		},
		{
			MethodName: "ClientSeatCheckoutNotification",
			Handler:    _SiteManagerClient_ClientSeatCheckoutNotification_Handler,
		},
		{
			MethodName: "TrackedPaymentsNotification",
			Handler:    _SiteManagerClient_TrackedPaymentsNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_site_license.steamclient.proto",
}

// SiteLicenseClient is the client API for SiteLicense service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteLicenseClient interface {
	InitiateAssociation(ctx context.Context, in *CSiteLicense_InitiateAssociation_Request, opts ...grpc.CallOption) (*CSiteLicense_InitiateAssociation_Response, error)
	LCSAuthenticate(ctx context.Context, in *CSiteLicense_LCSAuthenticate_Request, opts ...grpc.CallOption) (*CSiteLicense_LCSAuthenticate_Response, error)
	LCSAssociateUser(ctx context.Context, in *CSiteLicense_LCSAssociateUser_Request, opts ...grpc.CallOption) (*CSiteLicense_LCSAssociateUser_Response, error)
	ClientSeatCheckout(ctx context.Context, in *CSiteLicense_ClientSeatCheckout_Request, opts ...grpc.CallOption) (*CSiteLicense_ClientSeatCheckout_Response, error)
	ClientGetAvailableSeats(ctx context.Context, in *CSiteLicense_ClientGetAvailableSeats_Request, opts ...grpc.CallOption) (*CSiteLicense_ClientGetAvailableSeats_Response, error)
}

type siteLicenseClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteLicenseClient(cc grpc.ClientConnInterface) SiteLicenseClient {
	return &siteLicenseClient{cc}
}

func (c *siteLicenseClient) InitiateAssociation(ctx context.Context, in *CSiteLicense_InitiateAssociation_Request, opts ...grpc.CallOption) (*CSiteLicense_InitiateAssociation_Response, error) {
	out := new(CSiteLicense_InitiateAssociation_Response)
	err := c.cc.Invoke(ctx, "/SiteLicense/InitiateAssociation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteLicenseClient) LCSAuthenticate(ctx context.Context, in *CSiteLicense_LCSAuthenticate_Request, opts ...grpc.CallOption) (*CSiteLicense_LCSAuthenticate_Response, error) {
	out := new(CSiteLicense_LCSAuthenticate_Response)
	err := c.cc.Invoke(ctx, "/SiteLicense/LCSAuthenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteLicenseClient) LCSAssociateUser(ctx context.Context, in *CSiteLicense_LCSAssociateUser_Request, opts ...grpc.CallOption) (*CSiteLicense_LCSAssociateUser_Response, error) {
	out := new(CSiteLicense_LCSAssociateUser_Response)
	err := c.cc.Invoke(ctx, "/SiteLicense/LCSAssociateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteLicenseClient) ClientSeatCheckout(ctx context.Context, in *CSiteLicense_ClientSeatCheckout_Request, opts ...grpc.CallOption) (*CSiteLicense_ClientSeatCheckout_Response, error) {
	out := new(CSiteLicense_ClientSeatCheckout_Response)
	err := c.cc.Invoke(ctx, "/SiteLicense/ClientSeatCheckout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteLicenseClient) ClientGetAvailableSeats(ctx context.Context, in *CSiteLicense_ClientGetAvailableSeats_Request, opts ...grpc.CallOption) (*CSiteLicense_ClientGetAvailableSeats_Response, error) {
	out := new(CSiteLicense_ClientGetAvailableSeats_Response)
	err := c.cc.Invoke(ctx, "/SiteLicense/ClientGetAvailableSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteLicenseServer is the server API for SiteLicense service.
// All implementations must embed UnimplementedSiteLicenseServer
// for forward compatibility
type SiteLicenseServer interface {
	InitiateAssociation(context.Context, *CSiteLicense_InitiateAssociation_Request) (*CSiteLicense_InitiateAssociation_Response, error)
	LCSAuthenticate(context.Context, *CSiteLicense_LCSAuthenticate_Request) (*CSiteLicense_LCSAuthenticate_Response, error)
	LCSAssociateUser(context.Context, *CSiteLicense_LCSAssociateUser_Request) (*CSiteLicense_LCSAssociateUser_Response, error)
	ClientSeatCheckout(context.Context, *CSiteLicense_ClientSeatCheckout_Request) (*CSiteLicense_ClientSeatCheckout_Response, error)
	ClientGetAvailableSeats(context.Context, *CSiteLicense_ClientGetAvailableSeats_Request) (*CSiteLicense_ClientGetAvailableSeats_Response, error)
	mustEmbedUnimplementedSiteLicenseServer()
}

// UnimplementedSiteLicenseServer must be embedded to have forward compatible implementations.
type UnimplementedSiteLicenseServer struct {
}

func (UnimplementedSiteLicenseServer) InitiateAssociation(context.Context, *CSiteLicense_InitiateAssociation_Request) (*CSiteLicense_InitiateAssociation_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateAssociation not implemented")
}
func (UnimplementedSiteLicenseServer) LCSAuthenticate(context.Context, *CSiteLicense_LCSAuthenticate_Request) (*CSiteLicense_LCSAuthenticate_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LCSAuthenticate not implemented")
}
func (UnimplementedSiteLicenseServer) LCSAssociateUser(context.Context, *CSiteLicense_LCSAssociateUser_Request) (*CSiteLicense_LCSAssociateUser_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LCSAssociateUser not implemented")
}
func (UnimplementedSiteLicenseServer) ClientSeatCheckout(context.Context, *CSiteLicense_ClientSeatCheckout_Request) (*CSiteLicense_ClientSeatCheckout_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientSeatCheckout not implemented")
}
func (UnimplementedSiteLicenseServer) ClientGetAvailableSeats(context.Context, *CSiteLicense_ClientGetAvailableSeats_Request) (*CSiteLicense_ClientGetAvailableSeats_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientGetAvailableSeats not implemented")
}
func (UnimplementedSiteLicenseServer) mustEmbedUnimplementedSiteLicenseServer() {}

// UnsafeSiteLicenseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteLicenseServer will
// result in compilation errors.
type UnsafeSiteLicenseServer interface {
	mustEmbedUnimplementedSiteLicenseServer()
}

func RegisterSiteLicenseServer(s grpc.ServiceRegistrar, srv SiteLicenseServer) {
	s.RegisterService(&SiteLicense_ServiceDesc, srv)
}

func _SiteLicense_InitiateAssociation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_InitiateAssociation_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteLicenseServer).InitiateAssociation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteLicense/InitiateAssociation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteLicenseServer).InitiateAssociation(ctx, req.(*CSiteLicense_InitiateAssociation_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteLicense_LCSAuthenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_LCSAuthenticate_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteLicenseServer).LCSAuthenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteLicense/LCSAuthenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteLicenseServer).LCSAuthenticate(ctx, req.(*CSiteLicense_LCSAuthenticate_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteLicense_LCSAssociateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_LCSAssociateUser_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteLicenseServer).LCSAssociateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteLicense/LCSAssociateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteLicenseServer).LCSAssociateUser(ctx, req.(*CSiteLicense_LCSAssociateUser_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteLicense_ClientSeatCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_ClientSeatCheckout_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteLicenseServer).ClientSeatCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteLicense/ClientSeatCheckout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteLicenseServer).ClientSeatCheckout(ctx, req.(*CSiteLicense_ClientSeatCheckout_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteLicense_ClientGetAvailableSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSiteLicense_ClientGetAvailableSeats_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteLicenseServer).ClientGetAvailableSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SiteLicense/ClientGetAvailableSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteLicenseServer).ClientGetAvailableSeats(ctx, req.(*CSiteLicense_ClientGetAvailableSeats_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteLicense_ServiceDesc is the grpc.ServiceDesc for SiteLicense service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteLicense_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SiteLicense",
	HandlerType: (*SiteLicenseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitiateAssociation",
			Handler:    _SiteLicense_InitiateAssociation_Handler,
		},
		{
			MethodName: "LCSAuthenticate",
			Handler:    _SiteLicense_LCSAuthenticate_Handler,
		},
		{
			MethodName: "LCSAssociateUser",
			Handler:    _SiteLicense_LCSAssociateUser_Handler,
		},
		{
			MethodName: "ClientSeatCheckout",
			Handler:    _SiteLicense_ClientSeatCheckout_Handler,
		},
		{
			MethodName: "ClientGetAvailableSeats",
			Handler:    _SiteLicense_ClientGetAvailableSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_site_license.steamclient.proto",
}
