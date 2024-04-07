// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: careerhub/userinfo_service/suggester/suggester_grpc/userinfo.proto

package suggester_grpc

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

type GetConditionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conditions []*Condition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *GetConditionsResponse) Reset() {
	*x = GetConditionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConditionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConditionsResponse) ProtoMessage() {}

func (x *GetConditionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConditionsResponse.ProtoReflect.Descriptor instead.
func (*GetConditionsResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *GetConditionsResponse) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ConditionId   string `protobuf:"bytes,2,opt,name=conditionId,proto3" json:"conditionId,omitempty"`
	ConditionName string `protobuf:"bytes,3,opt,name=conditionName,proto3" json:"conditionName,omitempty"`
	Query         *Query `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Condition) GetConditionId() string {
	if x != nil {
		return x.ConditionId
	}
	return ""
}

func (x *Condition) GetConditionName() string {
	if x != nil {
		return x.ConditionName
	}
	return ""
}

func (x *Condition) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	SkillNames []*Skill    `protobuf:"bytes,2,rep,name=skillNames,proto3" json:"skillNames,omitempty"`
	MinCareer  *int32      `protobuf:"varint,3,opt,name=minCareer,proto3,oneof" json:"minCareer,omitempty"`
	MaxCareer  *int32      `protobuf:"varint,4,opt,name=maxCareer,proto3,oneof" json:"maxCareer,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Query) GetSkillNames() []*Skill {
	if x != nil {
		return x.SkillNames
	}
	return nil
}

func (x *Query) GetMinCareer() int32 {
	if x != nil && x.MinCareer != nil {
		return *x.MinCareer
	}
	return 0
}

func (x *Query) GetMaxCareer() int32 {
	if x != nil && x.MaxCareer != nil {
		return *x.MaxCareer
	}
	return 0
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Or []string `protobuf:"bytes,1,rep,name=Or,proto3" json:"Or,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP(), []int{3}
}

func (x *Skill) GetOr() []string {
	if x != nil {
		return x.Or
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site         string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP(), []int{4}
}

func (x *Category) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Category) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

var File_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto protoreflect.FileDescriptor

var file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDesc = []byte{
	0x0a, 0x42, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x68, 0x75, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x22, 0x90, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x50, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x43,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x69, 0x6e,
	0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x61, 0x78, 0x43, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x4f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x72, 0x22, 0x42, 0x0a,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0x75, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x69, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x40, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68,
	0x75, 0x62, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x63, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescOnce sync.Once
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescData = file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDesc
)

func file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescGZIP() []byte {
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescOnce.Do(func() {
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescData)
	})
	return file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDescData
}

var file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_goTypes = []interface{}{
	(*GetConditionsResponse)(nil), // 0: careerhub.userinfo_service.suggester_grpc.GetConditionsResponse
	(*Condition)(nil),             // 1: careerhub.userinfo_service.suggester_grpc.Condition
	(*Query)(nil),                 // 2: careerhub.userinfo_service.suggester_grpc.Query
	(*Skill)(nil),                 // 3: careerhub.userinfo_service.suggester_grpc.Skill
	(*Category)(nil),              // 4: careerhub.userinfo_service.suggester_grpc.Category
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_depIdxs = []int32{
	1, // 0: careerhub.userinfo_service.suggester_grpc.GetConditionsResponse.conditions:type_name -> careerhub.userinfo_service.suggester_grpc.Condition
	2, // 1: careerhub.userinfo_service.suggester_grpc.Condition.query:type_name -> careerhub.userinfo_service.suggester_grpc.Query
	4, // 2: careerhub.userinfo_service.suggester_grpc.Query.categories:type_name -> careerhub.userinfo_service.suggester_grpc.Category
	3, // 3: careerhub.userinfo_service.suggester_grpc.Query.skillNames:type_name -> careerhub.userinfo_service.suggester_grpc.Skill
	5, // 4: careerhub.userinfo_service.suggester_grpc.Userinfo.GetConditions:input_type -> google.protobuf.Empty
	0, // 5: careerhub.userinfo_service.suggester_grpc.Userinfo.GetConditions:output_type -> careerhub.userinfo_service.suggester_grpc.GetConditionsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_init() }
func file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_init() {
	if File_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConditionsResponse); i {
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
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_goTypes,
		DependencyIndexes: file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_depIdxs,
		MessageInfos:      file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_msgTypes,
	}.Build()
	File_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto = out.File
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_rawDesc = nil
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_goTypes = nil
	file_careerhub_userinfo_service_suggester_suggester_grpc_userinfo_proto_depIdxs = nil
}
