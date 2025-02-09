// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: steammessages_friendmessages.steamclient.proto

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

// FriendMessagesClient is the client API for FriendMessages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendMessagesClient interface {
	GetRecentMessages(ctx context.Context, in *CFriendMessages_GetRecentMessages_Request, opts ...grpc.CallOption) (*CFriendMessages_GetRecentMessages_Response, error)
	GetActiveMessageSessions(ctx context.Context, in *CFriendsMessages_GetActiveMessageSessions_Request, opts ...grpc.CallOption) (*CFriendsMessages_GetActiveMessageSessions_Response, error)
	SendMessage(ctx context.Context, in *CFriendMessages_SendMessage_Request, opts ...grpc.CallOption) (*CFriendMessages_SendMessage_Response, error)
	AckMessage(ctx context.Context, in *CFriendMessages_AckMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	IsInFriendsUIBeta(ctx context.Context, in *CFriendMessages_IsInFriendsUIBeta_Request, opts ...grpc.CallOption) (*CFriendMessages_IsInFriendsUIBeta_Response, error)
	UpdateMessageReaction(ctx context.Context, in *CFriendMessages_UpdateMessageReaction_Request, opts ...grpc.CallOption) (*CFriendMessages_UpdateMessageReaction_Response, error)
}

type friendMessagesClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendMessagesClient(cc grpc.ClientConnInterface) FriendMessagesClient {
	return &friendMessagesClient{cc}
}

func (c *friendMessagesClient) GetRecentMessages(ctx context.Context, in *CFriendMessages_GetRecentMessages_Request, opts ...grpc.CallOption) (*CFriendMessages_GetRecentMessages_Response, error) {
	out := new(CFriendMessages_GetRecentMessages_Response)
	err := c.cc.Invoke(ctx, "/FriendMessages/GetRecentMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClient) GetActiveMessageSessions(ctx context.Context, in *CFriendsMessages_GetActiveMessageSessions_Request, opts ...grpc.CallOption) (*CFriendsMessages_GetActiveMessageSessions_Response, error) {
	out := new(CFriendsMessages_GetActiveMessageSessions_Response)
	err := c.cc.Invoke(ctx, "/FriendMessages/GetActiveMessageSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClient) SendMessage(ctx context.Context, in *CFriendMessages_SendMessage_Request, opts ...grpc.CallOption) (*CFriendMessages_SendMessage_Response, error) {
	out := new(CFriendMessages_SendMessage_Response)
	err := c.cc.Invoke(ctx, "/FriendMessages/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClient) AckMessage(ctx context.Context, in *CFriendMessages_AckMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/FriendMessages/AckMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClient) IsInFriendsUIBeta(ctx context.Context, in *CFriendMessages_IsInFriendsUIBeta_Request, opts ...grpc.CallOption) (*CFriendMessages_IsInFriendsUIBeta_Response, error) {
	out := new(CFriendMessages_IsInFriendsUIBeta_Response)
	err := c.cc.Invoke(ctx, "/FriendMessages/IsInFriendsUIBeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClient) UpdateMessageReaction(ctx context.Context, in *CFriendMessages_UpdateMessageReaction_Request, opts ...grpc.CallOption) (*CFriendMessages_UpdateMessageReaction_Response, error) {
	out := new(CFriendMessages_UpdateMessageReaction_Response)
	err := c.cc.Invoke(ctx, "/FriendMessages/UpdateMessageReaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendMessagesServer is the server API for FriendMessages service.
// All implementations must embed UnimplementedFriendMessagesServer
// for forward compatibility
type FriendMessagesServer interface {
	GetRecentMessages(context.Context, *CFriendMessages_GetRecentMessages_Request) (*CFriendMessages_GetRecentMessages_Response, error)
	GetActiveMessageSessions(context.Context, *CFriendsMessages_GetActiveMessageSessions_Request) (*CFriendsMessages_GetActiveMessageSessions_Response, error)
	SendMessage(context.Context, *CFriendMessages_SendMessage_Request) (*CFriendMessages_SendMessage_Response, error)
	AckMessage(context.Context, *CFriendMessages_AckMessage_Notification) (*NoResponse, error)
	IsInFriendsUIBeta(context.Context, *CFriendMessages_IsInFriendsUIBeta_Request) (*CFriendMessages_IsInFriendsUIBeta_Response, error)
	UpdateMessageReaction(context.Context, *CFriendMessages_UpdateMessageReaction_Request) (*CFriendMessages_UpdateMessageReaction_Response, error)
	mustEmbedUnimplementedFriendMessagesServer()
}

// UnimplementedFriendMessagesServer must be embedded to have forward compatible implementations.
type UnimplementedFriendMessagesServer struct {
}

func (UnimplementedFriendMessagesServer) GetRecentMessages(context.Context, *CFriendMessages_GetRecentMessages_Request) (*CFriendMessages_GetRecentMessages_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentMessages not implemented")
}
func (UnimplementedFriendMessagesServer) GetActiveMessageSessions(context.Context, *CFriendsMessages_GetActiveMessageSessions_Request) (*CFriendsMessages_GetActiveMessageSessions_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveMessageSessions not implemented")
}
func (UnimplementedFriendMessagesServer) SendMessage(context.Context, *CFriendMessages_SendMessage_Request) (*CFriendMessages_SendMessage_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedFriendMessagesServer) AckMessage(context.Context, *CFriendMessages_AckMessage_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckMessage not implemented")
}
func (UnimplementedFriendMessagesServer) IsInFriendsUIBeta(context.Context, *CFriendMessages_IsInFriendsUIBeta_Request) (*CFriendMessages_IsInFriendsUIBeta_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInFriendsUIBeta not implemented")
}
func (UnimplementedFriendMessagesServer) UpdateMessageReaction(context.Context, *CFriendMessages_UpdateMessageReaction_Request) (*CFriendMessages_UpdateMessageReaction_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageReaction not implemented")
}
func (UnimplementedFriendMessagesServer) mustEmbedUnimplementedFriendMessagesServer() {}

// UnsafeFriendMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendMessagesServer will
// result in compilation errors.
type UnsafeFriendMessagesServer interface {
	mustEmbedUnimplementedFriendMessagesServer()
}

func RegisterFriendMessagesServer(s grpc.ServiceRegistrar, srv FriendMessagesServer) {
	s.RegisterService(&FriendMessages_ServiceDesc, srv)
}

func _FriendMessages_GetRecentMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_GetRecentMessages_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).GetRecentMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/GetRecentMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).GetRecentMessages(ctx, req.(*CFriendMessages_GetRecentMessages_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessages_GetActiveMessageSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendsMessages_GetActiveMessageSessions_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).GetActiveMessageSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/GetActiveMessageSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).GetActiveMessageSessions(ctx, req.(*CFriendsMessages_GetActiveMessageSessions_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessages_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_SendMessage_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).SendMessage(ctx, req.(*CFriendMessages_SendMessage_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessages_AckMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_AckMessage_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).AckMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/AckMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).AckMessage(ctx, req.(*CFriendMessages_AckMessage_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessages_IsInFriendsUIBeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_IsInFriendsUIBeta_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).IsInFriendsUIBeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/IsInFriendsUIBeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).IsInFriendsUIBeta(ctx, req.(*CFriendMessages_IsInFriendsUIBeta_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessages_UpdateMessageReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_UpdateMessageReaction_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesServer).UpdateMessageReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessages/UpdateMessageReaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesServer).UpdateMessageReaction(ctx, req.(*CFriendMessages_UpdateMessageReaction_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendMessages_ServiceDesc is the grpc.ServiceDesc for FriendMessages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendMessages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FriendMessages",
	HandlerType: (*FriendMessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecentMessages",
			Handler:    _FriendMessages_GetRecentMessages_Handler,
		},
		{
			MethodName: "GetActiveMessageSessions",
			Handler:    _FriendMessages_GetActiveMessageSessions_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _FriendMessages_SendMessage_Handler,
		},
		{
			MethodName: "AckMessage",
			Handler:    _FriendMessages_AckMessage_Handler,
		},
		{
			MethodName: "IsInFriendsUIBeta",
			Handler:    _FriendMessages_IsInFriendsUIBeta_Handler,
		},
		{
			MethodName: "UpdateMessageReaction",
			Handler:    _FriendMessages_UpdateMessageReaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_friendmessages.steamclient.proto",
}

// FriendMessagesClientClient is the client API for FriendMessagesClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendMessagesClientClient interface {
	IncomingMessage(ctx context.Context, in *CFriendMessages_IncomingMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	NotifyAckMessageEcho(ctx context.Context, in *CFriendMessages_AckMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error)
	MessageReaction(ctx context.Context, in *CFriendMessages_MessageReaction_Notification, opts ...grpc.CallOption) (*NoResponse, error)
}

type friendMessagesClientClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendMessagesClientClient(cc grpc.ClientConnInterface) FriendMessagesClientClient {
	return &friendMessagesClientClient{cc}
}

func (c *friendMessagesClientClient) IncomingMessage(ctx context.Context, in *CFriendMessages_IncomingMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/FriendMessagesClient/IncomingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClientClient) NotifyAckMessageEcho(ctx context.Context, in *CFriendMessages_AckMessage_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/FriendMessagesClient/NotifyAckMessageEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendMessagesClientClient) MessageReaction(ctx context.Context, in *CFriendMessages_MessageReaction_Notification, opts ...grpc.CallOption) (*NoResponse, error) {
	out := new(NoResponse)
	err := c.cc.Invoke(ctx, "/FriendMessagesClient/MessageReaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendMessagesClientServer is the server API for FriendMessagesClient service.
// All implementations must embed UnimplementedFriendMessagesClientServer
// for forward compatibility
type FriendMessagesClientServer interface {
	IncomingMessage(context.Context, *CFriendMessages_IncomingMessage_Notification) (*NoResponse, error)
	NotifyAckMessageEcho(context.Context, *CFriendMessages_AckMessage_Notification) (*NoResponse, error)
	MessageReaction(context.Context, *CFriendMessages_MessageReaction_Notification) (*NoResponse, error)
	mustEmbedUnimplementedFriendMessagesClientServer()
}

// UnimplementedFriendMessagesClientServer must be embedded to have forward compatible implementations.
type UnimplementedFriendMessagesClientServer struct {
}

func (UnimplementedFriendMessagesClientServer) IncomingMessage(context.Context, *CFriendMessages_IncomingMessage_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncomingMessage not implemented")
}
func (UnimplementedFriendMessagesClientServer) NotifyAckMessageEcho(context.Context, *CFriendMessages_AckMessage_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAckMessageEcho not implemented")
}
func (UnimplementedFriendMessagesClientServer) MessageReaction(context.Context, *CFriendMessages_MessageReaction_Notification) (*NoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageReaction not implemented")
}
func (UnimplementedFriendMessagesClientServer) mustEmbedUnimplementedFriendMessagesClientServer() {}

// UnsafeFriendMessagesClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendMessagesClientServer will
// result in compilation errors.
type UnsafeFriendMessagesClientServer interface {
	mustEmbedUnimplementedFriendMessagesClientServer()
}

func RegisterFriendMessagesClientServer(s grpc.ServiceRegistrar, srv FriendMessagesClientServer) {
	s.RegisterService(&FriendMessagesClient_ServiceDesc, srv)
}

func _FriendMessagesClient_IncomingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_IncomingMessage_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesClientServer).IncomingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessagesClient/IncomingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesClientServer).IncomingMessage(ctx, req.(*CFriendMessages_IncomingMessage_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessagesClient_NotifyAckMessageEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_AckMessage_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesClientServer).NotifyAckMessageEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessagesClient/NotifyAckMessageEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesClientServer).NotifyAckMessageEcho(ctx, req.(*CFriendMessages_AckMessage_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendMessagesClient_MessageReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CFriendMessages_MessageReaction_Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendMessagesClientServer).MessageReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendMessagesClient/MessageReaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendMessagesClientServer).MessageReaction(ctx, req.(*CFriendMessages_MessageReaction_Notification))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendMessagesClient_ServiceDesc is the grpc.ServiceDesc for FriendMessagesClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendMessagesClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FriendMessagesClient",
	HandlerType: (*FriendMessagesClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncomingMessage",
			Handler:    _FriendMessagesClient_IncomingMessage_Handler,
		},
		{
			MethodName: "NotifyAckMessageEcho",
			Handler:    _FriendMessagesClient_NotifyAckMessageEcho_Handler,
		},
		{
			MethodName: "MessageReaction",
			Handler:    _FriendMessagesClient_MessageReaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "steammessages_friendmessages.steamclient.proto",
}
