// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: protobufs/events.proto

package protobufs

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*Event_ProcessLogEvent
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protobufs_events_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Event) GetProcessLogEvent() *Event_ProcessLog {
	if x, ok := x.GetEvent().(*Event_ProcessLogEvent); ok {
		return x.ProcessLogEvent
	}
	return nil
}

type isEvent_Event interface {
	isEvent_Event()
}

type Event_ProcessLogEvent struct {
	ProcessLogEvent *Event_ProcessLog `protobuf:"bytes,1,opt,name=process_log_event,json=processLogEvent,proto3,oneof"`
}

func (*Event_ProcessLogEvent) isEvent_Event() {}

type Event_ProcessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Event_ProcessLog) Reset() {
	*x = Event_ProcessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_ProcessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_ProcessLog) ProtoMessage() {}

func (x *Event_ProcessLog) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_ProcessLog.ProtoReflect.Descriptor instead.
func (*Event_ProcessLog) Descriptor() ([]byte, []int) {
	return file_protobufs_events_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_ProcessLog) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

var File_protobufs_events_proto protoreflect.FileDescriptor

var file_protobufs_events_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x3f, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x48,
	0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c,
	0x6f, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x6e, 0x6f,
	0x37, 0x37, 0x32, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protobufs_events_proto_rawDescOnce sync.Once
	file_protobufs_events_proto_rawDescData = file_protobufs_events_proto_rawDesc
)

func file_protobufs_events_proto_rawDescGZIP() []byte {
	file_protobufs_events_proto_rawDescOnce.Do(func() {
		file_protobufs_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_events_proto_rawDescData)
	})
	return file_protobufs_events_proto_rawDescData
}

var file_protobufs_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobufs_events_proto_goTypes = []interface{}{
	(*Event)(nil),            // 0: Event
	(*Event_ProcessLog)(nil), // 1: Event.ProcessLog
}
var file_protobufs_events_proto_depIdxs = []int32{
	1, // 0: Event.process_log_event:type_name -> Event.ProcessLog
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobufs_events_proto_init() }
func file_protobufs_events_proto_init() {
	if File_protobufs_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufs_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_protobufs_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_ProcessLog); i {
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
	file_protobufs_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_ProcessLogEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobufs_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufs_events_proto_goTypes,
		DependencyIndexes: file_protobufs_events_proto_depIdxs,
		MessageInfos:      file_protobufs_events_proto_msgTypes,
	}.Build()
	File_protobufs_events_proto = out.File
	file_protobufs_events_proto_rawDesc = nil
	file_protobufs_events_proto_goTypes = nil
	file_protobufs_events_proto_depIdxs = nil
}
