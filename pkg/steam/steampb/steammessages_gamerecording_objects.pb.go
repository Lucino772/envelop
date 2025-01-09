// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: steammessages_gamerecording_objects.proto

package steampb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CGameRecording_AudioSessionsChanged_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sessions []*CGameRecording_AudioSessionsChanged_Notification_Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
}

func (x *CGameRecording_AudioSessionsChanged_Notification) Reset() {
	*x = CGameRecording_AudioSessionsChanged_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamerecording_objects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameRecording_AudioSessionsChanged_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameRecording_AudioSessionsChanged_Notification) ProtoMessage() {}

func (x *CGameRecording_AudioSessionsChanged_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamerecording_objects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameRecording_AudioSessionsChanged_Notification.ProtoReflect.Descriptor instead.
func (*CGameRecording_AudioSessionsChanged_Notification) Descriptor() ([]byte, []int) {
	return file_steammessages_gamerecording_objects_proto_rawDescGZIP(), []int{0}
}

func (x *CGameRecording_AudioSessionsChanged_Notification) GetSessions() []*CGameRecording_AudioSessionsChanged_Notification_Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type CGameRecording_AudioSessionsChanged_Notification_Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IsSystem   *bool    `protobuf:"varint,3,opt,name=is_system,json=isSystem" json:"is_system,omitempty"`
	IsMuted    *bool    `protobuf:"varint,4,opt,name=is_muted,json=isMuted" json:"is_muted,omitempty"`
	IsActive   *bool    `protobuf:"varint,5,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
	IsCaptured *bool    `protobuf:"varint,6,opt,name=is_captured,json=isCaptured" json:"is_captured,omitempty"`
	RecentPeak *float32 `protobuf:"fixed32,7,opt,name=recent_peak,json=recentPeak" json:"recent_peak,omitempty"`
	IsGame     *bool    `protobuf:"varint,8,opt,name=is_game,json=isGame" json:"is_game,omitempty"`
	IsSteam    *bool    `protobuf:"varint,9,opt,name=is_steam,json=isSteam" json:"is_steam,omitempty"`
	IsSaved    *bool    `protobuf:"varint,10,opt,name=is_saved,json=isSaved" json:"is_saved,omitempty"`
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) Reset() {
	*x = CGameRecording_AudioSessionsChanged_Notification_Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamerecording_objects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameRecording_AudioSessionsChanged_Notification_Session) ProtoMessage() {}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamerecording_objects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameRecording_AudioSessionsChanged_Notification_Session.ProtoReflect.Descriptor instead.
func (*CGameRecording_AudioSessionsChanged_Notification_Session) Descriptor() ([]byte, []int) {
	return file_steammessages_gamerecording_objects_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsSystem() bool {
	if x != nil && x.IsSystem != nil {
		return *x.IsSystem
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsMuted() bool {
	if x != nil && x.IsMuted != nil {
		return *x.IsMuted
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsCaptured() bool {
	if x != nil && x.IsCaptured != nil {
		return *x.IsCaptured
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetRecentPeak() float32 {
	if x != nil && x.RecentPeak != nil {
		return *x.RecentPeak
	}
	return 0
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsGame() bool {
	if x != nil && x.IsGame != nil {
		return *x.IsGame
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsSteam() bool {
	if x != nil && x.IsSteam != nil {
		return *x.IsSteam
	}
	return false
}

func (x *CGameRecording_AudioSessionsChanged_Notification_Session) GetIsSaved() bool {
	if x != nil && x.IsSaved != nil {
		return *x.IsSaved
	}
	return false
}

var File_steammessages_gamerecording_objects_proto protoreflect.FileDescriptor

var file_steammessages_gamerecording_objects_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x03, 0x0a, 0x30, 0x43, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x93, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x61, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x61, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x47, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x53, 0x61, 0x76, 0x65, 0x64, 0x42, 0x35, 0x48, 0x01, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x6e,
	0x6f, 0x37, 0x37, 0x32, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x80, 0x01,
	0x00,
}

var (
	file_steammessages_gamerecording_objects_proto_rawDescOnce sync.Once
	file_steammessages_gamerecording_objects_proto_rawDescData = file_steammessages_gamerecording_objects_proto_rawDesc
)

func file_steammessages_gamerecording_objects_proto_rawDescGZIP() []byte {
	file_steammessages_gamerecording_objects_proto_rawDescOnce.Do(func() {
		file_steammessages_gamerecording_objects_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_gamerecording_objects_proto_rawDescData)
	})
	return file_steammessages_gamerecording_objects_proto_rawDescData
}

var file_steammessages_gamerecording_objects_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_steammessages_gamerecording_objects_proto_goTypes = []interface{}{
	(*CGameRecording_AudioSessionsChanged_Notification)(nil),         // 0: CGameRecording_AudioSessionsChanged_Notification
	(*CGameRecording_AudioSessionsChanged_Notification_Session)(nil), // 1: CGameRecording_AudioSessionsChanged_Notification.Session
}
var file_steammessages_gamerecording_objects_proto_depIdxs = []int32{
	1, // 0: CGameRecording_AudioSessionsChanged_Notification.sessions:type_name -> CGameRecording_AudioSessionsChanged_Notification.Session
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_steammessages_gamerecording_objects_proto_init() }
func file_steammessages_gamerecording_objects_proto_init() {
	if File_steammessages_gamerecording_objects_proto != nil {
		return
	}
	file_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_gamerecording_objects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CGameRecording_AudioSessionsChanged_Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steammessages_gamerecording_objects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CGameRecording_AudioSessionsChanged_Notification_Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steammessages_gamerecording_objects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_gamerecording_objects_proto_goTypes,
		DependencyIndexes: file_steammessages_gamerecording_objects_proto_depIdxs,
		MessageInfos:      file_steammessages_gamerecording_objects_proto_msgTypes,
	}.Build()
	File_steammessages_gamerecording_objects_proto = out.File
	file_steammessages_gamerecording_objects_proto_rawDesc = nil
	file_steammessages_gamerecording_objects_proto_goTypes = nil
	file_steammessages_gamerecording_objects_proto_depIdxs = nil
}
