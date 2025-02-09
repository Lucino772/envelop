// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: players.proto

package protobufs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_players_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_players_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PlayerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumPlayers uint32    `protobuf:"varint,1,opt,name=num_players,json=numPlayers,proto3" json:"num_players,omitempty"`
	MaxPlayers uint32    `protobuf:"varint,2,opt,name=max_players,json=maxPlayers,proto3" json:"max_players,omitempty"`
	Players    []*Player `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *PlayerList) Reset() {
	*x = PlayerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerList) ProtoMessage() {}

func (x *PlayerList) ProtoReflect() protoreflect.Message {
	mi := &file_players_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerList.ProtoReflect.Descriptor instead.
func (*PlayerList) Descriptor() ([]byte, []int) {
	return file_players_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerList) GetNumPlayers() uint32 {
	if x != nil {
		return x.NumPlayers
	}
	return 0
}

func (x *PlayerList) GetMaxPlayers() uint32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

func (x *PlayerList) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_players_proto protoreflect.FileDescriptor

var file_players_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x32, 0x85,
	0x01, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x13, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x42, 0x83, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x42, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x6e, 0x6f, 0x37, 0x37, 0x32, 0x2f, 0x65, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x45, 0x6e, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0xca, 0x02, 0x07, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0xe2, 0x02, 0x13, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x07, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_players_proto_rawDescOnce sync.Once
	file_players_proto_rawDescData = file_players_proto_rawDesc
)

func file_players_proto_rawDescGZIP() []byte {
	file_players_proto_rawDescOnce.Do(func() {
		file_players_proto_rawDescData = protoimpl.X.CompressGZIP(file_players_proto_rawDescData)
	})
	return file_players_proto_rawDescData
}

var file_players_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_players_proto_goTypes = []interface{}{
	(*Player)(nil),        // 0: envelop.Player
	(*PlayerList)(nil),    // 1: envelop.PlayerList
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_players_proto_depIdxs = []int32{
	0, // 0: envelop.PlayerList.players:type_name -> envelop.Player
	2, // 1: envelop.Players.ListPlayers:input_type -> google.protobuf.Empty
	2, // 2: envelop.Players.StreamPlayers:input_type -> google.protobuf.Empty
	1, // 3: envelop.Players.ListPlayers:output_type -> envelop.PlayerList
	1, // 4: envelop.Players.StreamPlayers:output_type -> envelop.PlayerList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_players_proto_init() }
func file_players_proto_init() {
	if File_players_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_players_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_players_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerList); i {
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
			RawDescriptor: file_players_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_players_proto_goTypes,
		DependencyIndexes: file_players_proto_depIdxs,
		MessageInfos:      file_players_proto_msgTypes,
	}.Build()
	File_players_proto = out.File
	file_players_proto_rawDesc = nil
	file_players_proto_goTypes = nil
	file_players_proto_depIdxs = nil
}
