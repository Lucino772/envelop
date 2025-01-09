// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: steammessages_workshop.steamclient.proto

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

type CWorkshop_GetEULAStatus_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
}

func (x *CWorkshop_GetEULAStatus_Request) Reset() {
	*x = CWorkshop_GetEULAStatus_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_workshop_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CWorkshop_GetEULAStatus_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CWorkshop_GetEULAStatus_Request) ProtoMessage() {}

func (x *CWorkshop_GetEULAStatus_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_workshop_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CWorkshop_GetEULAStatus_Request.ProtoReflect.Descriptor instead.
func (*CWorkshop_GetEULAStatus_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_workshop_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CWorkshop_GetEULAStatus_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type CWorkshop_GetEULAStatus_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         *uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	TimestampAction *uint32 `protobuf:"varint,2,opt,name=timestamp_action,json=timestampAction" json:"timestamp_action,omitempty"`
	Accepted        *bool   `protobuf:"varint,3,opt,name=accepted" json:"accepted,omitempty"`
	NeedsAction     *bool   `protobuf:"varint,4,opt,name=needs_action,json=needsAction" json:"needs_action,omitempty"`
}

func (x *CWorkshop_GetEULAStatus_Response) Reset() {
	*x = CWorkshop_GetEULAStatus_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_workshop_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CWorkshop_GetEULAStatus_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CWorkshop_GetEULAStatus_Response) ProtoMessage() {}

func (x *CWorkshop_GetEULAStatus_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_workshop_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CWorkshop_GetEULAStatus_Response.ProtoReflect.Descriptor instead.
func (*CWorkshop_GetEULAStatus_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_workshop_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CWorkshop_GetEULAStatus_Response) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *CWorkshop_GetEULAStatus_Response) GetTimestampAction() uint32 {
	if x != nil && x.TimestampAction != nil {
		return *x.TimestampAction
	}
	return 0
}

func (x *CWorkshop_GetEULAStatus_Response) GetAccepted() bool {
	if x != nil && x.Accepted != nil {
		return *x.Accepted
	}
	return false
}

func (x *CWorkshop_GetEULAStatus_Response) GetNeedsAction() bool {
	if x != nil && x.NeedsAction != nil {
		return *x.NeedsAction
	}
	return false
}

var File_steammessages_workshop_steamclient_proto protoreflect.FileDescriptor

var file_steammessages_workshop_steamclient_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x37, 0x0a, 0x1f, 0x43, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x47, 0x65, 0x74, 0x45, 0x55, 0x4c, 0x41, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x20,
	0x43, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x47, 0x65, 0x74, 0x45, 0x55, 0x4c,
	0x41, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x60, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70,
	0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x55, 0x4c, 0x41, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x2e, 0x43, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x47, 0x65,
	0x74, 0x45, 0x55, 0x4c, 0x41, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x43, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x47, 0x65, 0x74, 0x45, 0x55, 0x4c, 0x41, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x6e, 0x6f, 0x37, 0x37, 0x32, 0x2f, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x80, 0x01, 0x01,
}

var (
	file_steammessages_workshop_steamclient_proto_rawDescOnce sync.Once
	file_steammessages_workshop_steamclient_proto_rawDescData = file_steammessages_workshop_steamclient_proto_rawDesc
)

func file_steammessages_workshop_steamclient_proto_rawDescGZIP() []byte {
	file_steammessages_workshop_steamclient_proto_rawDescOnce.Do(func() {
		file_steammessages_workshop_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_workshop_steamclient_proto_rawDescData)
	})
	return file_steammessages_workshop_steamclient_proto_rawDescData
}

var file_steammessages_workshop_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_steammessages_workshop_steamclient_proto_goTypes = []interface{}{
	(*CWorkshop_GetEULAStatus_Request)(nil),  // 0: CWorkshop_GetEULAStatus_Request
	(*CWorkshop_GetEULAStatus_Response)(nil), // 1: CWorkshop_GetEULAStatus_Response
}
var file_steammessages_workshop_steamclient_proto_depIdxs = []int32{
	0, // 0: Workshop.GetEULAStatus:input_type -> CWorkshop_GetEULAStatus_Request
	1, // 1: Workshop.GetEULAStatus:output_type -> CWorkshop_GetEULAStatus_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steammessages_workshop_steamclient_proto_init() }
func file_steammessages_workshop_steamclient_proto_init() {
	if File_steammessages_workshop_steamclient_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	file_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_workshop_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CWorkshop_GetEULAStatus_Request); i {
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
		file_steammessages_workshop_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CWorkshop_GetEULAStatus_Response); i {
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
			RawDescriptor: file_steammessages_workshop_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_workshop_steamclient_proto_goTypes,
		DependencyIndexes: file_steammessages_workshop_steamclient_proto_depIdxs,
		MessageInfos:      file_steammessages_workshop_steamclient_proto_msgTypes,
	}.Build()
	File_steammessages_workshop_steamclient_proto = out.File
	file_steammessages_workshop_steamclient_proto_rawDesc = nil
	file_steammessages_workshop_steamclient_proto_goTypes = nil
	file_steammessages_workshop_steamclient_proto_depIdxs = nil
}
