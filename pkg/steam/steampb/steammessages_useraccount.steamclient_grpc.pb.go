// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_useraccount.steamclient.proto

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

// UserAccountClient is the client API for UserAccount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAccountClient interface {
	GetAvailableValveDiscountPromotions(ctx context.Context, in *CUserAccount_GetAvailableValveDiscountPromotions_Request, opts ...grpc.CallOption) (*CUserAccount_GetAvailableValveDiscountPromotions_Response, error)
	GetClientWalletDetails(ctx context.Context, in *CUserAccount_GetClientWalletDetails_Request, opts ...grpc.CallOption) (*CUserAccount_GetWalletDetails_Response, error)
	GetAccountLinkStatus(ctx context.Context, in *CUserAccount_GetAccountLinkStatus_Request, opts ...grpc.CallOption) (*CUserAccount_GetAccountLinkStatus_Response, error)
	CancelLicenseForApp(ctx context.Context, in *CUserAccount_CancelLicenseForApp_Request, opts ...grpc.CallOption) (*CUserAccount_CancelLicenseForApp_Response, error)
	GetUserCountry(ctx context.Context, in *CUserAccount_GetUserCountry_Request, opts ...grpc.CallOption) (*CUserAccount_GetUserCountry_Response, error)
	CreateFriendInviteToken(ctx context.Context, in *CUserAccount_CreateFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_CreateFriendInviteToken_Response, error)
	GetFriendInviteTokens(ctx context.Context, in *CUserAccount_GetFriendInviteTokens_Request, opts ...grpc.CallOption) (*CUserAccount_GetFriendInviteTokens_Response, error)
	ViewFriendInviteToken(ctx context.Context, in *CUserAccount_ViewFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_ViewFriendInviteToken_Response, error)
	RedeemFriendInviteToken(ctx context.Context, in *CUserAccount_RedeemFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_RedeemFriendInviteToken_Response, error)
	RevokeFriendInviteToken(ctx context.Context, in *CUserAccount_RevokeFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_RevokeFriendInviteToken_Response, error)
	RegisterCompatTool(ctx context.Context, in *CUserAccount_RegisterCompatTool_Request, opts ...grpc.CallOption) (*CUserAccount_RegisterCompatTool_Response, error)
}

type userAccountClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountClient(cc grpc.ClientConnInterface) UserAccountClient {
	return &userAccountClient{cc}
}

func (c *userAccountClient) GetAvailableValveDiscountPromotions(ctx context.Context, in *CUserAccount_GetAvailableValveDiscountPromotions_Request, opts ...grpc.CallOption) (*CUserAccount_GetAvailableValveDiscountPromotions_Response, error) {
	out := new(CUserAccount_GetAvailableValveDiscountPromotions_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/GetAvailableValveDiscountPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) GetClientWalletDetails(ctx context.Context, in *CUserAccount_GetClientWalletDetails_Request, opts ...grpc.CallOption) (*CUserAccount_GetWalletDetails_Response, error) {
	out := new(CUserAccount_GetWalletDetails_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/GetClientWalletDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) GetAccountLinkStatus(ctx context.Context, in *CUserAccount_GetAccountLinkStatus_Request, opts ...grpc.CallOption) (*CUserAccount_GetAccountLinkStatus_Response, error) {
	out := new(CUserAccount_GetAccountLinkStatus_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/GetAccountLinkStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) CancelLicenseForApp(ctx context.Context, in *CUserAccount_CancelLicenseForApp_Request, opts ...grpc.CallOption) (*CUserAccount_CancelLicenseForApp_Response, error) {
	out := new(CUserAccount_CancelLicenseForApp_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/CancelLicenseForApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) GetUserCountry(ctx context.Context, in *CUserAccount_GetUserCountry_Request, opts ...grpc.CallOption) (*CUserAccount_GetUserCountry_Response, error) {
	out := new(CUserAccount_GetUserCountry_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/GetUserCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) CreateFriendInviteToken(ctx context.Context, in *CUserAccount_CreateFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_CreateFriendInviteToken_Response, error) {
	out := new(CUserAccount_CreateFriendInviteToken_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/CreateFriendInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) GetFriendInviteTokens(ctx context.Context, in *CUserAccount_GetFriendInviteTokens_Request, opts ...grpc.CallOption) (*CUserAccount_GetFriendInviteTokens_Response, error) {
	out := new(CUserAccount_GetFriendInviteTokens_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/GetFriendInviteTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) ViewFriendInviteToken(ctx context.Context, in *CUserAccount_ViewFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_ViewFriendInviteToken_Response, error) {
	out := new(CUserAccount_ViewFriendInviteToken_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/ViewFriendInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) RedeemFriendInviteToken(ctx context.Context, in *CUserAccount_RedeemFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_RedeemFriendInviteToken_Response, error) {
	out := new(CUserAccount_RedeemFriendInviteToken_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/RedeemFriendInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) RevokeFriendInviteToken(ctx context.Context, in *CUserAccount_RevokeFriendInviteToken_Request, opts ...grpc.CallOption) (*CUserAccount_RevokeFriendInviteToken_Response, error) {
	out := new(CUserAccount_RevokeFriendInviteToken_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/RevokeFriendInviteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountClient) RegisterCompatTool(ctx context.Context, in *CUserAccount_RegisterCompatTool_Request, opts ...grpc.CallOption) (*CUserAccount_RegisterCompatTool_Response, error) {
	out := new(CUserAccount_RegisterCompatTool_Response)
	err := c.cc.Invoke(ctx, "/UserAccount/RegisterCompatTool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountServer is the server API for UserAccount service.
// All implementations must embed UnimplementedUserAccountServer
// for forward compatibility
type UserAccountServer interface {
	GetAvailableValveDiscountPromotions(context.Context, *CUserAccount_GetAvailableValveDiscountPromotions_Request) (*CUserAccount_GetAvailableValveDiscountPromotions_Response, error)
	GetClientWalletDetails(context.Context, *CUserAccount_GetClientWalletDetails_Request) (*CUserAccount_GetWalletDetails_Response, error)
	GetAccountLinkStatus(context.Context, *CUserAccount_GetAccountLinkStatus_Request) (*CUserAccount_GetAccountLinkStatus_Response, error)
	CancelLicenseForApp(context.Context, *CUserAccount_CancelLicenseForApp_Request) (*CUserAccount_CancelLicenseForApp_Response, error)
	GetUserCountry(context.Context, *CUserAccount_GetUserCountry_Request) (*CUserAccount_GetUserCountry_Response, error)
	CreateFriendInviteToken(context.Context, *CUserAccount_CreateFriendInviteToken_Request) (*CUserAccount_CreateFriendInviteToken_Response, error)
	GetFriendInviteTokens(context.Context, *CUserAccount_GetFriendInviteTokens_Request) (*CUserAccount_GetFriendInviteTokens_Response, error)
	ViewFriendInviteToken(context.Context, *CUserAccount_ViewFriendInviteToken_Request) (*CUserAccount_ViewFriendInviteToken_Response, error)
	RedeemFriendInviteToken(context.Context, *CUserAccount_RedeemFriendInviteToken_Request) (*CUserAccount_RedeemFriendInviteToken_Response, error)
	RevokeFriendInviteToken(context.Context, *CUserAccount_RevokeFriendInviteToken_Request) (*CUserAccount_RevokeFriendInviteToken_Response, error)
	RegisterCompatTool(context.Context, *CUserAccount_RegisterCompatTool_Request) (*CUserAccount_RegisterCompatTool_Response, error)
	mustEmbedUnimplementedUserAccountServer()
}

// UnimplementedUserAccountServer must be embedded to have forward compatible implementations.
type UnimplementedUserAccountServer struct {
}

func (UnimplementedUserAccountServer) GetAvailableValveDiscountPromotions(context.Context, *CUserAccount_GetAvailableValveDiscountPromotions_Request) (*CUserAccount_GetAvailableValveDiscountPromotions_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableValveDiscountPromotions not implemented")
}
func (UnimplementedUserAccountServer) GetClientWalletDetails(context.Context, *CUserAccount_GetClientWalletDetails_Request) (*CUserAccount_GetWalletDetails_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientWalletDetails not implemented")
}
func (UnimplementedUserAccountServer) GetAccountLinkStatus(context.Context, *CUserAccount_GetAccountLinkStatus_Request) (*CUserAccount_GetAccountLinkStatus_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountLinkStatus not implemented")
}
func (UnimplementedUserAccountServer) CancelLicenseForApp(context.Context, *CUserAccount_CancelLicenseForApp_Request) (*CUserAccount_CancelLicenseForApp_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLicenseForApp not implemented")
}
func (UnimplementedUserAccountServer) GetUserCountry(context.Context, *CUserAccount_GetUserCountry_Request) (*CUserAccount_GetUserCountry_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCountry not implemented")
}
func (UnimplementedUserAccountServer) CreateFriendInviteToken(context.Context, *CUserAccount_CreateFriendInviteToken_Request) (*CUserAccount_CreateFriendInviteToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFriendInviteToken not implemented")
}
func (UnimplementedUserAccountServer) GetFriendInviteTokens(context.Context, *CUserAccount_GetFriendInviteTokens_Request) (*CUserAccount_GetFriendInviteTokens_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendInviteTokens not implemented")
}
func (UnimplementedUserAccountServer) ViewFriendInviteToken(context.Context, *CUserAccount_ViewFriendInviteToken_Request) (*CUserAccount_ViewFriendInviteToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewFriendInviteToken not implemented")
}
func (UnimplementedUserAccountServer) RedeemFriendInviteToken(context.Context, *CUserAccount_RedeemFriendInviteToken_Request) (*CUserAccount_RedeemFriendInviteToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemFriendInviteToken not implemented")
}
func (UnimplementedUserAccountServer) RevokeFriendInviteToken(context.Context, *CUserAccount_RevokeFriendInviteToken_Request) (*CUserAccount_RevokeFriendInviteToken_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeFriendInviteToken not implemented")
}
func (UnimplementedUserAccountServer) RegisterCompatTool(context.Context, *CUserAccount_RegisterCompatTool_Request) (*CUserAccount_RegisterCompatTool_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCompatTool not implemented")
}
func (UnimplementedUserAccountServer) mustEmbedUnimplementedUserAccountServer() {}

// UnsafeUserAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAccountServer will
// result in compilation errors.
type UnsafeUserAccountServer interface {
	mustEmbedUnimplementedUserAccountServer()
}

func RegisterUserAccountServer(s grpc.ServiceRegistrar, srv UserAccountServer) {
	s.RegisterService(&UserAccount_ServiceDesc, srv)
}

func _UserAccount_GetAvailableValveDiscountPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_GetAvailableValveDiscountPromotions_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).GetAvailableValveDiscountPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/GetAvailableValveDiscountPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).GetAvailableValveDiscountPromotions(ctx, req.(*CUserAccount_GetAvailableValveDiscountPromotions_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_GetClientWalletDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_GetClientWalletDetails_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).GetClientWalletDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/GetClientWalletDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).GetClientWalletDetails(ctx, req.(*CUserAccount_GetClientWalletDetails_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_GetAccountLinkStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_GetAccountLinkStatus_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).GetAccountLinkStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/GetAccountLinkStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).GetAccountLinkStatus(ctx, req.(*CUserAccount_GetAccountLinkStatus_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_CancelLicenseForApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_CancelLicenseForApp_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).CancelLicenseForApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/CancelLicenseForApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).CancelLicenseForApp(ctx, req.(*CUserAccount_CancelLicenseForApp_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_GetUserCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_GetUserCountry_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).GetUserCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/GetUserCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).GetUserCountry(ctx, req.(*CUserAccount_GetUserCountry_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_CreateFriendInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_CreateFriendInviteToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).CreateFriendInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/CreateFriendInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).CreateFriendInviteToken(ctx, req.(*CUserAccount_CreateFriendInviteToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_GetFriendInviteTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_GetFriendInviteTokens_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).GetFriendInviteTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/GetFriendInviteTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).GetFriendInviteTokens(ctx, req.(*CUserAccount_GetFriendInviteTokens_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_ViewFriendInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_ViewFriendInviteToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).ViewFriendInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/ViewFriendInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).ViewFriendInviteToken(ctx, req.(*CUserAccount_ViewFriendInviteToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_RedeemFriendInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_RedeemFriendInviteToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).RedeemFriendInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/RedeemFriendInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).RedeemFriendInviteToken(ctx, req.(*CUserAccount_RedeemFriendInviteToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_RevokeFriendInviteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_RevokeFriendInviteToken_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).RevokeFriendInviteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/RevokeFriendInviteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).RevokeFriendInviteToken(ctx, req.(*CUserAccount_RevokeFriendInviteToken_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccount_RegisterCompatTool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CUserAccount_RegisterCompatTool_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServer).RegisterCompatTool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserAccount/RegisterCompatTool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServer).RegisterCompatTool(ctx, req.(*CUserAccount_RegisterCompatTool_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAccount_ServiceDesc is the grpc.ServiceDesc for UserAccount service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAccount_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserAccount",
	HandlerType: (*UserAccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailableValveDiscountPromotions",
			Handler:    _UserAccount_GetAvailableValveDiscountPromotions_Handler,
		},
		{
			MethodName: "GetClientWalletDetails",
			Handler:    _UserAccount_GetClientWalletDetails_Handler,
		},
		{
			MethodName: "GetAccountLinkStatus",
			Handler:    _UserAccount_GetAccountLinkStatus_Handler,
		},
		{
			MethodName: "CancelLicenseForApp",
			Handler:    _UserAccount_CancelLicenseForApp_Handler,
		},
		{
			MethodName: "GetUserCountry",
			Handler:    _UserAccount_GetUserCountry_Handler,
		},
		{
			MethodName: "CreateFriendInviteToken",
			Handler:    _UserAccount_CreateFriendInviteToken_Handler,
		},
		{
			MethodName: "GetFriendInviteTokens",
			Handler:    _UserAccount_GetFriendInviteTokens_Handler,
		},
		{
			MethodName: "ViewFriendInviteToken",
			Handler:    _UserAccount_ViewFriendInviteToken_Handler,
		},
		{
			MethodName: "RedeemFriendInviteToken",
			Handler:    _UserAccount_RedeemFriendInviteToken_Handler,
		},
		{
			MethodName: "RevokeFriendInviteToken",
			Handler:    _UserAccount_RevokeFriendInviteToken_Handler,
		},
		{
			MethodName: "RegisterCompatTool",
			Handler:    _UserAccount_RegisterCompatTool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_useraccount.steamclient.proto",
}

// AccountLinkingClient is the client API for AccountLinking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountLinkingClient interface {
	GetLinkedAccountInfo(ctx context.Context, in *CAccountLinking_GetLinkedAccountInfo_Request, opts ...grpc.CallOption) (*CAccountLinking_GetLinkedAccountInfo_Response, error)
}

type accountLinkingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountLinkingClient(cc grpc.ClientConnInterface) AccountLinkingClient {
	return &accountLinkingClient{cc}
}

func (c *accountLinkingClient) GetLinkedAccountInfo(ctx context.Context, in *CAccountLinking_GetLinkedAccountInfo_Request, opts ...grpc.CallOption) (*CAccountLinking_GetLinkedAccountInfo_Response, error) {
	out := new(CAccountLinking_GetLinkedAccountInfo_Response)
	err := c.cc.Invoke(ctx, "/AccountLinking/GetLinkedAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountLinkingServer is the server API for AccountLinking service.
// All implementations must embed UnimplementedAccountLinkingServer
// for forward compatibility
type AccountLinkingServer interface {
	GetLinkedAccountInfo(context.Context, *CAccountLinking_GetLinkedAccountInfo_Request) (*CAccountLinking_GetLinkedAccountInfo_Response, error)
	mustEmbedUnimplementedAccountLinkingServer()
}

// UnimplementedAccountLinkingServer must be embedded to have forward compatible implementations.
type UnimplementedAccountLinkingServer struct {
}

func (UnimplementedAccountLinkingServer) GetLinkedAccountInfo(context.Context, *CAccountLinking_GetLinkedAccountInfo_Request) (*CAccountLinking_GetLinkedAccountInfo_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedAccountInfo not implemented")
}
func (UnimplementedAccountLinkingServer) mustEmbedUnimplementedAccountLinkingServer() {}

// UnsafeAccountLinkingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountLinkingServer will
// result in compilation errors.
type UnsafeAccountLinkingServer interface {
	mustEmbedUnimplementedAccountLinkingServer()
}

func RegisterAccountLinkingServer(s grpc.ServiceRegistrar, srv AccountLinkingServer) {
	s.RegisterService(&AccountLinking_ServiceDesc, srv)
}

func _AccountLinking_GetLinkedAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CAccountLinking_GetLinkedAccountInfo_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountLinkingServer).GetLinkedAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountLinking/GetLinkedAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountLinkingServer).GetLinkedAccountInfo(ctx, req.(*CAccountLinking_GetLinkedAccountInfo_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountLinking_ServiceDesc is the grpc.ServiceDesc for AccountLinking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountLinking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccountLinking",
	HandlerType: (*AccountLinkingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLinkedAccountInfo",
			Handler:    _AccountLinking_GetLinkedAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_useraccount.steamclient.proto",
}

// EmbeddedClientClient is the client API for EmbeddedClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbeddedClientClient interface {
	AuthorizeCurrentDevice(ctx context.Context, in *CEmbeddedClient_AuthorizeCurrentDevice_Request, opts ...grpc.CallOption) (*CEmbeddedClient_AuthorizeDevice_Response, error)
}

type embeddedClientClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbeddedClientClient(cc grpc.ClientConnInterface) EmbeddedClientClient {
	return &embeddedClientClient{cc}
}

func (c *embeddedClientClient) AuthorizeCurrentDevice(ctx context.Context, in *CEmbeddedClient_AuthorizeCurrentDevice_Request, opts ...grpc.CallOption) (*CEmbeddedClient_AuthorizeDevice_Response, error) {
	out := new(CEmbeddedClient_AuthorizeDevice_Response)
	err := c.cc.Invoke(ctx, "/EmbeddedClient/AuthorizeCurrentDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbeddedClientServer is the server API for EmbeddedClient service.
// All implementations must embed UnimplementedEmbeddedClientServer
// for forward compatibility
type EmbeddedClientServer interface {
	AuthorizeCurrentDevice(context.Context, *CEmbeddedClient_AuthorizeCurrentDevice_Request) (*CEmbeddedClient_AuthorizeDevice_Response, error)
	mustEmbedUnimplementedEmbeddedClientServer()
}

// UnimplementedEmbeddedClientServer must be embedded to have forward compatible implementations.
type UnimplementedEmbeddedClientServer struct {
}

func (UnimplementedEmbeddedClientServer) AuthorizeCurrentDevice(context.Context, *CEmbeddedClient_AuthorizeCurrentDevice_Request) (*CEmbeddedClient_AuthorizeDevice_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeCurrentDevice not implemented")
}
func (UnimplementedEmbeddedClientServer) mustEmbedUnimplementedEmbeddedClientServer() {}

// UnsafeEmbeddedClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbeddedClientServer will
// result in compilation errors.
type UnsafeEmbeddedClientServer interface {
	mustEmbedUnimplementedEmbeddedClientServer()
}

func RegisterEmbeddedClientServer(s grpc.ServiceRegistrar, srv EmbeddedClientServer) {
	s.RegisterService(&EmbeddedClient_ServiceDesc, srv)
}

func _EmbeddedClient_AuthorizeCurrentDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CEmbeddedClient_AuthorizeCurrentDevice_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddedClientServer).AuthorizeCurrentDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmbeddedClient/AuthorizeCurrentDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddedClientServer).AuthorizeCurrentDevice(ctx, req.(*CEmbeddedClient_AuthorizeCurrentDevice_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// EmbeddedClient_ServiceDesc is the grpc.ServiceDesc for EmbeddedClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmbeddedClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EmbeddedClient",
	HandlerType: (*EmbeddedClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeCurrentDevice",
			Handler:    _EmbeddedClient_AuthorizeCurrentDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_useraccount.steamclient.proto",
}
