// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: external/users/v1/social.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequesterId   string                 `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	RecipientId   string                 `protobuf:"bytes,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddFriendRequest) Reset() {
	*x = AddFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendRequest) ProtoMessage() {}

func (x *AddFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[0]
	if x != nil {
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
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{0}
}

func (x *AddFriendRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *AddFriendRequest) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type AcceptFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipientId   string                 `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	RequesterId   string                 `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AcceptFriendRequest) Reset() {
	*x = AcceptFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptFriendRequest) ProtoMessage() {}

func (x *AcceptFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptFriendRequest.ProtoReflect.Descriptor instead.
func (*AcceptFriendRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptFriendRequest) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *AcceptFriendRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

type RejectFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecipientId   string                 `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	RequesterId   string                 `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RejectFriendRequest) Reset() {
	*x = RejectFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectFriendRequest) ProtoMessage() {}

func (x *RejectFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectFriendRequest.ProtoReflect.Descriptor instead.
func (*RejectFriendRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{2}
}

func (x *RejectFriendRequest) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *RejectFriendRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

type RemoveFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequesterId   string                 `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	FriendId      string                 `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFriendRequest) Reset() {
	*x = RemoveFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFriendRequest) ProtoMessage() {}

func (x *RemoveFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFriendRequest.ProtoReflect.Descriptor instead.
func (*RemoveFriendRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFriendRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *RemoveFriendRequest) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

type ListFriendsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFriendsRequest) Reset() {
	*x = ListFriendsRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendsRequest) ProtoMessage() {}

func (x *ListFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendsRequest.ProtoReflect.Descriptor instead.
func (*ListFriendsRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{4}
}

func (x *ListFriendsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BlockFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId      string                 `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlockFriendRequest) Reset() {
	*x = BlockFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFriendRequest) ProtoMessage() {}

func (x *BlockFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFriendRequest.ProtoReflect.Descriptor instead.
func (*BlockFriendRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{5}
}

func (x *BlockFriendRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BlockFriendRequest) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

type UnblockFriendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId      string                 `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnblockFriendRequest) Reset() {
	*x = UnblockFriendRequest{}
	mi := &file_external_users_v1_social_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnblockFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockFriendRequest) ProtoMessage() {}

func (x *UnblockFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_users_v1_social_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockFriendRequest.ProtoReflect.Descriptor instead.
func (*UnblockFriendRequest) Descriptor() ([]byte, []int) {
	return file_external_users_v1_social_proto_rawDescGZIP(), []int{6}
}

func (x *UnblockFriendRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UnblockFriendRequest) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

var File_external_users_v1_social_proto protoreflect.FileDescriptor

var file_external_users_v1_social_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x5b, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b,
	0x0a, 0x13, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4a, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x4c, 0x0a,
	0x14, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x32, 0xfc, 0x03, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x47, 0x0a, 0x0d, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_external_users_v1_social_proto_rawDescOnce sync.Once
	file_external_users_v1_social_proto_rawDescData []byte
)

func file_external_users_v1_social_proto_rawDescGZIP() []byte {
	file_external_users_v1_social_proto_rawDescOnce.Do(func() {
		file_external_users_v1_social_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_external_users_v1_social_proto_rawDesc), len(file_external_users_v1_social_proto_rawDesc)))
	})
	return file_external_users_v1_social_proto_rawDescData
}

var file_external_users_v1_social_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_external_users_v1_social_proto_goTypes = []any{
	(*AddFriendRequest)(nil),     // 0: users.v1.AddFriendRequest
	(*AcceptFriendRequest)(nil),  // 1: users.v1.AcceptFriendRequest
	(*RejectFriendRequest)(nil),  // 2: users.v1.RejectFriendRequest
	(*RemoveFriendRequest)(nil),  // 3: users.v1.RemoveFriendRequest
	(*ListFriendsRequest)(nil),   // 4: users.v1.ListFriendsRequest
	(*BlockFriendRequest)(nil),   // 5: users.v1.BlockFriendRequest
	(*UnblockFriendRequest)(nil), // 6: users.v1.UnblockFriendRequest
	(*emptypb.Empty)(nil),        // 7: google.protobuf.Empty
	(*FriendsList)(nil),          // 8: users.v1.FriendsList
}
var file_external_users_v1_social_proto_depIdxs = []int32{
	0, // 0: users.v1.UsersSocialService.AddFriend:input_type -> users.v1.AddFriendRequest
	1, // 1: users.v1.UsersSocialService.AcceptFriend:input_type -> users.v1.AcceptFriendRequest
	2, // 2: users.v1.UsersSocialService.RejectFriend:input_type -> users.v1.RejectFriendRequest
	3, // 3: users.v1.UsersSocialService.RemoveFriend:input_type -> users.v1.RemoveFriendRequest
	4, // 4: users.v1.UsersSocialService.ListFriends:input_type -> users.v1.ListFriendsRequest
	5, // 5: users.v1.UsersSocialService.BlockFriend:input_type -> users.v1.BlockFriendRequest
	6, // 6: users.v1.UsersSocialService.UnblockFriend:input_type -> users.v1.UnblockFriendRequest
	7, // 7: users.v1.UsersSocialService.AddFriend:output_type -> google.protobuf.Empty
	7, // 8: users.v1.UsersSocialService.AcceptFriend:output_type -> google.protobuf.Empty
	7, // 9: users.v1.UsersSocialService.RejectFriend:output_type -> google.protobuf.Empty
	7, // 10: users.v1.UsersSocialService.RemoveFriend:output_type -> google.protobuf.Empty
	8, // 11: users.v1.UsersSocialService.ListFriends:output_type -> users.v1.FriendsList
	7, // 12: users.v1.UsersSocialService.BlockFriend:output_type -> google.protobuf.Empty
	7, // 13: users.v1.UsersSocialService.UnblockFriend:output_type -> google.protobuf.Empty
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_users_v1_social_proto_init() }
func file_external_users_v1_social_proto_init() {
	if File_external_users_v1_social_proto != nil {
		return
	}
	file_external_users_v1_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_external_users_v1_social_proto_rawDesc), len(file_external_users_v1_social_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_users_v1_social_proto_goTypes,
		DependencyIndexes: file_external_users_v1_social_proto_depIdxs,
		MessageInfos:      file_external_users_v1_social_proto_msgTypes,
	}.Build()
	File_external_users_v1_social_proto = out.File
	file_external_users_v1_social_proto_goTypes = nil
	file_external_users_v1_social_proto_depIdxs = nil
}
