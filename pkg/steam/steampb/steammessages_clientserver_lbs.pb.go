// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: steammessages_clientserver_lbs.proto

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

type CMsgClientLBSSetScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId             *uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	LeaderboardId     *int32  `protobuf:"varint,2,opt,name=leaderboard_id,json=leaderboardId" json:"leaderboard_id,omitempty"`
	Score             *int32  `protobuf:"varint,3,opt,name=score" json:"score,omitempty"`
	Details           []byte  `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
	UploadScoreMethod *int32  `protobuf:"varint,5,opt,name=upload_score_method,json=uploadScoreMethod" json:"upload_score_method,omitempty"`
}

func (x *CMsgClientLBSSetScore) Reset() {
	*x = CMsgClientLBSSetScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSSetScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSSetScore) ProtoMessage() {}

func (x *CMsgClientLBSSetScore) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSSetScore.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSSetScore) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgClientLBSSetScore) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgClientLBSSetScore) GetLeaderboardId() int32 {
	if x != nil && x.LeaderboardId != nil {
		return *x.LeaderboardId
	}
	return 0
}

func (x *CMsgClientLBSSetScore) GetScore() int32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *CMsgClientLBSSetScore) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *CMsgClientLBSSetScore) GetUploadScoreMethod() int32 {
	if x != nil && x.UploadScoreMethod != nil {
		return *x.UploadScoreMethod
	}
	return 0
}

type CMsgClientLBSSetScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult               *int32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	LeaderboardEntryCount *int32 `protobuf:"varint,2,opt,name=leaderboard_entry_count,json=leaderboardEntryCount" json:"leaderboard_entry_count,omitempty"`
	ScoreChanged          *bool  `protobuf:"varint,3,opt,name=score_changed,json=scoreChanged" json:"score_changed,omitempty"`
	GlobalRankPrevious    *int32 `protobuf:"varint,4,opt,name=global_rank_previous,json=globalRankPrevious" json:"global_rank_previous,omitempty"`
	GlobalRankNew         *int32 `protobuf:"varint,5,opt,name=global_rank_new,json=globalRankNew" json:"global_rank_new,omitempty"`
}

// Default values for CMsgClientLBSSetScoreResponse fields.
const (
	Default_CMsgClientLBSSetScoreResponse_Eresult = int32(2)
)

func (x *CMsgClientLBSSetScoreResponse) Reset() {
	*x = CMsgClientLBSSetScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSSetScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSSetScoreResponse) ProtoMessage() {}

func (x *CMsgClientLBSSetScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSSetScoreResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSSetScoreResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgClientLBSSetScoreResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientLBSSetScoreResponse_Eresult
}

func (x *CMsgClientLBSSetScoreResponse) GetLeaderboardEntryCount() int32 {
	if x != nil && x.LeaderboardEntryCount != nil {
		return *x.LeaderboardEntryCount
	}
	return 0
}

func (x *CMsgClientLBSSetScoreResponse) GetScoreChanged() bool {
	if x != nil && x.ScoreChanged != nil {
		return *x.ScoreChanged
	}
	return false
}

func (x *CMsgClientLBSSetScoreResponse) GetGlobalRankPrevious() int32 {
	if x != nil && x.GlobalRankPrevious != nil {
		return *x.GlobalRankPrevious
	}
	return 0
}

func (x *CMsgClientLBSSetScoreResponse) GetGlobalRankNew() int32 {
	if x != nil && x.GlobalRankNew != nil {
		return *x.GlobalRankNew
	}
	return 0
}

type CMsgClientLBSSetUGC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId         *uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	LeaderboardId *int32  `protobuf:"varint,2,opt,name=leaderboard_id,json=leaderboardId" json:"leaderboard_id,omitempty"`
	UgcId         *uint64 `protobuf:"fixed64,3,opt,name=ugc_id,json=ugcId" json:"ugc_id,omitempty"`
}

func (x *CMsgClientLBSSetUGC) Reset() {
	*x = CMsgClientLBSSetUGC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSSetUGC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSSetUGC) ProtoMessage() {}

func (x *CMsgClientLBSSetUGC) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSSetUGC.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSSetUGC) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgClientLBSSetUGC) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgClientLBSSetUGC) GetLeaderboardId() int32 {
	if x != nil && x.LeaderboardId != nil {
		return *x.LeaderboardId
	}
	return 0
}

func (x *CMsgClientLBSSetUGC) GetUgcId() uint64 {
	if x != nil && x.UgcId != nil {
		return *x.UgcId
	}
	return 0
}

type CMsgClientLBSSetUGCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult *int32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
}

// Default values for CMsgClientLBSSetUGCResponse fields.
const (
	Default_CMsgClientLBSSetUGCResponse_Eresult = int32(2)
)

func (x *CMsgClientLBSSetUGCResponse) Reset() {
	*x = CMsgClientLBSSetUGCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSSetUGCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSSetUGCResponse) ProtoMessage() {}

func (x *CMsgClientLBSSetUGCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSSetUGCResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSSetUGCResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgClientLBSSetUGCResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientLBSSetUGCResponse_Eresult
}

type CMsgClientLBSFindOrCreateLB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId                  *uint32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	LeaderboardSortMethod  *int32  `protobuf:"varint,2,opt,name=leaderboard_sort_method,json=leaderboardSortMethod" json:"leaderboard_sort_method,omitempty"`
	LeaderboardDisplayType *int32  `protobuf:"varint,3,opt,name=leaderboard_display_type,json=leaderboardDisplayType" json:"leaderboard_display_type,omitempty"`
	CreateIfNotFound       *bool   `protobuf:"varint,4,opt,name=create_if_not_found,json=createIfNotFound" json:"create_if_not_found,omitempty"`
	LeaderboardName        *string `protobuf:"bytes,5,opt,name=leaderboard_name,json=leaderboardName" json:"leaderboard_name,omitempty"`
}

func (x *CMsgClientLBSFindOrCreateLB) Reset() {
	*x = CMsgClientLBSFindOrCreateLB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSFindOrCreateLB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSFindOrCreateLB) ProtoMessage() {}

func (x *CMsgClientLBSFindOrCreateLB) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSFindOrCreateLB.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSFindOrCreateLB) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{4}
}

func (x *CMsgClientLBSFindOrCreateLB) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgClientLBSFindOrCreateLB) GetLeaderboardSortMethod() int32 {
	if x != nil && x.LeaderboardSortMethod != nil {
		return *x.LeaderboardSortMethod
	}
	return 0
}

func (x *CMsgClientLBSFindOrCreateLB) GetLeaderboardDisplayType() int32 {
	if x != nil && x.LeaderboardDisplayType != nil {
		return *x.LeaderboardDisplayType
	}
	return 0
}

func (x *CMsgClientLBSFindOrCreateLB) GetCreateIfNotFound() bool {
	if x != nil && x.CreateIfNotFound != nil {
		return *x.CreateIfNotFound
	}
	return false
}

func (x *CMsgClientLBSFindOrCreateLB) GetLeaderboardName() string {
	if x != nil && x.LeaderboardName != nil {
		return *x.LeaderboardName
	}
	return ""
}

type CMsgClientLBSFindOrCreateLBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult                *int32  `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	LeaderboardId          *int32  `protobuf:"varint,2,opt,name=leaderboard_id,json=leaderboardId" json:"leaderboard_id,omitempty"`
	LeaderboardEntryCount  *int32  `protobuf:"varint,3,opt,name=leaderboard_entry_count,json=leaderboardEntryCount" json:"leaderboard_entry_count,omitempty"`
	LeaderboardSortMethod  *int32  `protobuf:"varint,4,opt,name=leaderboard_sort_method,json=leaderboardSortMethod,def=0" json:"leaderboard_sort_method,omitempty"`
	LeaderboardDisplayType *int32  `protobuf:"varint,5,opt,name=leaderboard_display_type,json=leaderboardDisplayType,def=0" json:"leaderboard_display_type,omitempty"`
	LeaderboardName        *string `protobuf:"bytes,6,opt,name=leaderboard_name,json=leaderboardName" json:"leaderboard_name,omitempty"`
}

// Default values for CMsgClientLBSFindOrCreateLBResponse fields.
const (
	Default_CMsgClientLBSFindOrCreateLBResponse_Eresult                = int32(2)
	Default_CMsgClientLBSFindOrCreateLBResponse_LeaderboardSortMethod  = int32(0)
	Default_CMsgClientLBSFindOrCreateLBResponse_LeaderboardDisplayType = int32(0)
)

func (x *CMsgClientLBSFindOrCreateLBResponse) Reset() {
	*x = CMsgClientLBSFindOrCreateLBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSFindOrCreateLBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSFindOrCreateLBResponse) ProtoMessage() {}

func (x *CMsgClientLBSFindOrCreateLBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSFindOrCreateLBResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSFindOrCreateLBResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{5}
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientLBSFindOrCreateLBResponse_Eresult
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetLeaderboardId() int32 {
	if x != nil && x.LeaderboardId != nil {
		return *x.LeaderboardId
	}
	return 0
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetLeaderboardEntryCount() int32 {
	if x != nil && x.LeaderboardEntryCount != nil {
		return *x.LeaderboardEntryCount
	}
	return 0
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetLeaderboardSortMethod() int32 {
	if x != nil && x.LeaderboardSortMethod != nil {
		return *x.LeaderboardSortMethod
	}
	return Default_CMsgClientLBSFindOrCreateLBResponse_LeaderboardSortMethod
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetLeaderboardDisplayType() int32 {
	if x != nil && x.LeaderboardDisplayType != nil {
		return *x.LeaderboardDisplayType
	}
	return Default_CMsgClientLBSFindOrCreateLBResponse_LeaderboardDisplayType
}

func (x *CMsgClientLBSFindOrCreateLBResponse) GetLeaderboardName() string {
	if x != nil && x.LeaderboardName != nil {
		return *x.LeaderboardName
	}
	return ""
}

type CMsgClientLBSGetLBEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId                  *int32   `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	LeaderboardId          *int32   `protobuf:"varint,2,opt,name=leaderboard_id,json=leaderboardId" json:"leaderboard_id,omitempty"`
	RangeStart             *int32   `protobuf:"varint,3,opt,name=range_start,json=rangeStart" json:"range_start,omitempty"`
	RangeEnd               *int32   `protobuf:"varint,4,opt,name=range_end,json=rangeEnd" json:"range_end,omitempty"`
	LeaderboardDataRequest *int32   `protobuf:"varint,5,opt,name=leaderboard_data_request,json=leaderboardDataRequest" json:"leaderboard_data_request,omitempty"`
	Steamids               []uint64 `protobuf:"fixed64,6,rep,name=steamids" json:"steamids,omitempty"`
}

func (x *CMsgClientLBSGetLBEntries) Reset() {
	*x = CMsgClientLBSGetLBEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSGetLBEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSGetLBEntries) ProtoMessage() {}

func (x *CMsgClientLBSGetLBEntries) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSGetLBEntries.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSGetLBEntries) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{6}
}

func (x *CMsgClientLBSGetLBEntries) GetAppId() int32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntries) GetLeaderboardId() int32 {
	if x != nil && x.LeaderboardId != nil {
		return *x.LeaderboardId
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntries) GetRangeStart() int32 {
	if x != nil && x.RangeStart != nil {
		return *x.RangeStart
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntries) GetRangeEnd() int32 {
	if x != nil && x.RangeEnd != nil {
		return *x.RangeEnd
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntries) GetLeaderboardDataRequest() int32 {
	if x != nil && x.LeaderboardDataRequest != nil {
		return *x.LeaderboardDataRequest
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntries) GetSteamids() []uint64 {
	if x != nil {
		return x.Steamids
	}
	return nil
}

type CMsgClientLBSGetLBEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult               *int32                                     `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	LeaderboardEntryCount *int32                                     `protobuf:"varint,2,opt,name=leaderboard_entry_count,json=leaderboardEntryCount" json:"leaderboard_entry_count,omitempty"`
	Entries               []*CMsgClientLBSGetLBEntriesResponse_Entry `protobuf:"bytes,3,rep,name=entries" json:"entries,omitempty"`
}

// Default values for CMsgClientLBSGetLBEntriesResponse fields.
const (
	Default_CMsgClientLBSGetLBEntriesResponse_Eresult = int32(2)
)

func (x *CMsgClientLBSGetLBEntriesResponse) Reset() {
	*x = CMsgClientLBSGetLBEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSGetLBEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSGetLBEntriesResponse) ProtoMessage() {}

func (x *CMsgClientLBSGetLBEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSGetLBEntriesResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSGetLBEntriesResponse) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{7}
}

func (x *CMsgClientLBSGetLBEntriesResponse) GetEresult() int32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return Default_CMsgClientLBSGetLBEntriesResponse_Eresult
}

func (x *CMsgClientLBSGetLBEntriesResponse) GetLeaderboardEntryCount() int32 {
	if x != nil && x.LeaderboardEntryCount != nil {
		return *x.LeaderboardEntryCount
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntriesResponse) GetEntries() []*CMsgClientLBSGetLBEntriesResponse_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type CMsgClientLBSGetLBEntriesResponse_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamIdUser *uint64 `protobuf:"fixed64,1,opt,name=steam_id_user,json=steamIdUser" json:"steam_id_user,omitempty"`
	GlobalRank  *int32  `protobuf:"varint,2,opt,name=global_rank,json=globalRank" json:"global_rank,omitempty"`
	Score       *int32  `protobuf:"varint,3,opt,name=score" json:"score,omitempty"`
	Details     []byte  `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
	UgcId       *uint64 `protobuf:"fixed64,5,opt,name=ugc_id,json=ugcId" json:"ugc_id,omitempty"`
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) Reset() {
	*x = CMsgClientLBSGetLBEntriesResponse_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_lbs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientLBSGetLBEntriesResponse_Entry) ProtoMessage() {}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_lbs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientLBSGetLBEntriesResponse_Entry.ProtoReflect.Descriptor instead.
func (*CMsgClientLBSGetLBEntriesResponse_Entry) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_lbs_proto_rawDescGZIP(), []int{7, 0}
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) GetSteamIdUser() uint64 {
	if x != nil && x.SteamIdUser != nil {
		return *x.SteamIdUser
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) GetGlobalRank() int32 {
	if x != nil && x.GlobalRank != nil {
		return *x.GlobalRank
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) GetScore() int32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *CMsgClientLBSGetLBEntriesResponse_Entry) GetUgcId() uint64 {
	if x != nil && x.UgcId != nil {
		return *x.UgcId
	}
	return 0
}

var File_steammessages_clientserver_lbs_proto protoreflect.FileDescriptor

var file_steammessages_clientserver_lbs_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x62, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb5, 0x01, 0x0a, 0x15, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x42, 0x53, 0x53, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xf3, 0x01, 0x0a, 0x1d, 0x43, 0x4d, 0x73,
	0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x53, 0x65, 0x74, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07,
	0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x77, 0x22, 0x6a,
	0x0a, 0x13, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x53,
	0x65, 0x74, 0x55, 0x47, 0x43, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x75, 0x67, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x06, 0x52, 0x05, 0x75, 0x67, 0x63, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x1b, 0x43, 0x4d,
	0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x53, 0x65, 0x74, 0x55, 0x47,
	0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x65,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x80, 0x02, 0x0a, 0x1b, 0x43, 0x4d, 0x73, 0x67, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x42, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x17, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x74,
	0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc4, 0x02, 0x0a, 0x23, 0x43, 0x4d,
	0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x17, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01,
	0x30, 0x52, 0x15, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x6f,
	0x72, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3b, 0x0a, 0x18, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x16, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xed, 0x01, 0x0a, 0x19, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x42, 0x53, 0x47, 0x65, 0x74, 0x4c, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x06, 0x52, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x73,
	0x22, 0xd2, 0x02, 0x0a, 0x21, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x42, 0x53, 0x47, 0x65, 0x74, 0x4c, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x65, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x43,
	0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x42, 0x53, 0x47, 0x65, 0x74, 0x4c,
	0x42, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a,
	0x93, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x75, 0x67, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05,
	0x75, 0x67, 0x63, 0x49, 0x64, 0x42, 0x35, 0x48, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x6e, 0x6f, 0x37, 0x37, 0x32, 0x2f,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x80, 0x01, 0x00,
}

var (
	file_steammessages_clientserver_lbs_proto_rawDescOnce sync.Once
	file_steammessages_clientserver_lbs_proto_rawDescData = file_steammessages_clientserver_lbs_proto_rawDesc
)

func file_steammessages_clientserver_lbs_proto_rawDescGZIP() []byte {
	file_steammessages_clientserver_lbs_proto_rawDescOnce.Do(func() {
		file_steammessages_clientserver_lbs_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_clientserver_lbs_proto_rawDescData)
	})
	return file_steammessages_clientserver_lbs_proto_rawDescData
}

var file_steammessages_clientserver_lbs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_steammessages_clientserver_lbs_proto_goTypes = []interface{}{
	(*CMsgClientLBSSetScore)(nil),                   // 0: CMsgClientLBSSetScore
	(*CMsgClientLBSSetScoreResponse)(nil),           // 1: CMsgClientLBSSetScoreResponse
	(*CMsgClientLBSSetUGC)(nil),                     // 2: CMsgClientLBSSetUGC
	(*CMsgClientLBSSetUGCResponse)(nil),             // 3: CMsgClientLBSSetUGCResponse
	(*CMsgClientLBSFindOrCreateLB)(nil),             // 4: CMsgClientLBSFindOrCreateLB
	(*CMsgClientLBSFindOrCreateLBResponse)(nil),     // 5: CMsgClientLBSFindOrCreateLBResponse
	(*CMsgClientLBSGetLBEntries)(nil),               // 6: CMsgClientLBSGetLBEntries
	(*CMsgClientLBSGetLBEntriesResponse)(nil),       // 7: CMsgClientLBSGetLBEntriesResponse
	(*CMsgClientLBSGetLBEntriesResponse_Entry)(nil), // 8: CMsgClientLBSGetLBEntriesResponse.Entry
}
var file_steammessages_clientserver_lbs_proto_depIdxs = []int32{
	8, // 0: CMsgClientLBSGetLBEntriesResponse.entries:type_name -> CMsgClientLBSGetLBEntriesResponse.Entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_steammessages_clientserver_lbs_proto_init() }
func file_steammessages_clientserver_lbs_proto_init() {
	if File_steammessages_clientserver_lbs_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_clientserver_lbs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSSetScore); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSSetScoreResponse); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSSetUGC); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSSetUGCResponse); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSFindOrCreateLB); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSFindOrCreateLBResponse); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSGetLBEntries); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSGetLBEntriesResponse); i {
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
		file_steammessages_clientserver_lbs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientLBSGetLBEntriesResponse_Entry); i {
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
			RawDescriptor: file_steammessages_clientserver_lbs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_clientserver_lbs_proto_goTypes,
		DependencyIndexes: file_steammessages_clientserver_lbs_proto_depIdxs,
		MessageInfos:      file_steammessages_clientserver_lbs_proto_msgTypes,
	}.Build()
	File_steammessages_clientserver_lbs_proto = out.File
	file_steammessages_clientserver_lbs_proto_rawDesc = nil
	file_steammessages_clientserver_lbs_proto_goTypes = nil
	file_steammessages_clientserver_lbs_proto_depIdxs = nil
}
