// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: api/friend.proto

package service

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

type AddFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FriendId int64  `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	ApplyMsg string `protobuf:"bytes,3,opt,name=apply_msg,json=applyMsg,proto3" json:"apply_msg,omitempty"`
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendRequest.ProtoReflect.Descriptor instead.
func (*AddFriendRequest) Descriptor() ([]byte, []int) {
	return file_api_friend_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddFriendRequest) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *AddFriendRequest) GetApplyMsg() string {
	if x != nil {
		return x.ApplyMsg
	}
	return ""
}

type DelFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FriendId int64 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *DelFriendRequest) Reset() {
	*x = DelFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelFriendRequest) ProtoMessage() {}

func (x *DelFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelFriendRequest.ProtoReflect.Descriptor instead.
func (*DelFriendRequest) Descriptor() ([]byte, []int) {
	return file_api_friend_proto_rawDescGZIP(), []int{1}
}

func (x *DelFriendRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DelFriendRequest) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

type AgreeApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FriendId int64 `protobuf:"varint,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Agree    bool  `protobuf:"varint,3,opt,name=agree,proto3" json:"agree,omitempty"`
}

func (x *AgreeApplyRequest) Reset() {
	*x = AgreeApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreeApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreeApplyRequest) ProtoMessage() {}

func (x *AgreeApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreeApplyRequest.ProtoReflect.Descriptor instead.
func (*AgreeApplyRequest) Descriptor() ([]byte, []int) {
	return file_api_friend_proto_rawDescGZIP(), []int{2}
}

func (x *AgreeApplyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgreeApplyRequest) GetFriendId() int64 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *AgreeApplyRequest) GetAgree() bool {
	if x != nil {
		return x.Agree
	}
	return false
}

var File_api_friend_proto protoreflect.FileDescriptor

var file_api_friend_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x10, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x4d, 0x73, 0x67, 0x22, 0x3f, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x11,
	0x41, 0x67, 0x72, 0x65, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x67, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61,
	0x67, 0x72, 0x65, 0x65, 0x32, 0xc8, 0x01, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x67, 0x72, 0x65, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x20, 0x5a, 0x1e, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_friend_proto_rawDescOnce sync.Once
	file_api_friend_proto_rawDescData = file_api_friend_proto_rawDesc
)

func file_api_friend_proto_rawDescGZIP() []byte {
	file_api_friend_proto_rawDescOnce.Do(func() {
		file_api_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_friend_proto_rawDescData)
	})
	return file_api_friend_proto_rawDescData
}

var file_api_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_friend_proto_goTypes = []interface{}{
	(*AddFriendRequest)(nil),  // 0: service.AddFriendRequest
	(*DelFriendRequest)(nil),  // 1: service.DelFriendRequest
	(*AgreeApplyRequest)(nil), // 2: service.AgreeApplyRequest
	(*Response)(nil),          // 3: service.Response
}
var file_api_friend_proto_depIdxs = []int32{
	0, // 0: service.FriendService.AddFriend:input_type -> service.AddFriendRequest
	1, // 1: service.FriendService.DelFriend:input_type -> service.DelFriendRequest
	2, // 2: service.FriendService.AgreeApply:input_type -> service.AgreeApplyRequest
	3, // 3: service.FriendService.AddFriend:output_type -> service.Response
	3, // 4: service.FriendService.DelFriend:output_type -> service.Response
	3, // 5: service.FriendService.AgreeApply:output_type -> service.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_friend_proto_init() }
func file_api_friend_proto_init() {
	if File_api_friend_proto != nil {
		return
	}
	file_api_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendRequest); i {
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
		file_api_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelFriendRequest); i {
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
		file_api_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreeApplyRequest); i {
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
			RawDescriptor: file_api_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_friend_proto_goTypes,
		DependencyIndexes: file_api_friend_proto_depIdxs,
		MessageInfos:      file_api_friend_proto_msgTypes,
	}.Build()
	File_api_friend_proto = out.File
	file_api_friend_proto_rawDesc = nil
	file_api_friend_proto_goTypes = nil
	file_api_friend_proto_depIdxs = nil
}
