// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: identity/userapi/v1/userapi.proto

package userapiv1

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uun      string `protobuf:"bytes,1,opt,name=uun,proto3" json:"uun,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUun() string {
	if x != nil {
		return x.Uun
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{1}
}

type UserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserByIDRequest) Reset() {
	*x = UserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIDRequest) ProtoMessage() {}

func (x *UserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIDRequest.ProtoReflect.Descriptor instead.
func (*UserByIDRequest) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{2}
}

func (x *UserByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uun          string `protobuf:"bytes,2,opt,name=uun,proto3" json:"uun,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreationTime string `protobuf:"bytes,6,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Role         string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Status       string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserByIDResponse) Reset() {
	*x = UserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserByIDResponse) ProtoMessage() {}

func (x *UserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserByIDResponse.ProtoReflect.Descriptor instead.
func (*UserByIDResponse) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{3}
}

func (x *UserByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserByIDResponse) GetUun() string {
	if x != nil {
		return x.Uun
	}
	return ""
}

func (x *UserByIDResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserByIDResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserByIDResponse) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *UserByIDResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserByIDResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ChangePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	NewRole  string `protobuf:"bytes,2,opt,name=new_role,json=newRole,proto3" json:"new_role,omitempty"`
}

func (x *ChangePermissionRequest) Reset() {
	*x = ChangePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePermissionRequest) ProtoMessage() {}

func (x *ChangePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePermissionRequest.ProtoReflect.Descriptor instead.
func (*ChangePermissionRequest) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChangePermissionRequest) GetNewRole() string {
	if x != nil {
		return x.NewRole
	}
	return ""
}

type ChangePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangePermissionResponse) Reset() {
	*x = ChangePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePermissionResponse) ProtoMessage() {}

func (x *ChangePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePermissionResponse.ProtoReflect.Descriptor instead.
func (*ChangePermissionResponse) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{5}
}

type ChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ChangeStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeStatusResponse) Reset() {
	*x = ChangeStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_userapi_v1_userapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusResponse) ProtoMessage() {}

func (x *ChangeStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_userapi_v1_userapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusResponse.ProtoReflect.Descriptor instead.
func (*ChangeStatusResponse) Descriptor() ([]byte, []int) {
	return file_identity_userapi_v1_userapi_proto_rawDescGZIP(), []int{7}
}

var File_identity_userapi_v1_userapi_proto protoreflect.FileDescriptor

var file_identity_userapi_v1_userapi_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x75, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x75, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x75, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x65, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9b, 0x03, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6f, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x28, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_userapi_v1_userapi_proto_rawDescOnce sync.Once
	file_identity_userapi_v1_userapi_proto_rawDescData = file_identity_userapi_v1_userapi_proto_rawDesc
)

func file_identity_userapi_v1_userapi_proto_rawDescGZIP() []byte {
	file_identity_userapi_v1_userapi_proto_rawDescOnce.Do(func() {
		file_identity_userapi_v1_userapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_userapi_v1_userapi_proto_rawDescData)
	})
	return file_identity_userapi_v1_userapi_proto_rawDescData
}

var file_identity_userapi_v1_userapi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_identity_userapi_v1_userapi_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),        // 0: identity.userapi.v1.CreateUserRequest
	(*CreateUserResponse)(nil),       // 1: identity.userapi.v1.CreateUserResponse
	(*UserByIDRequest)(nil),          // 2: identity.userapi.v1.UserByIDRequest
	(*UserByIDResponse)(nil),         // 3: identity.userapi.v1.UserByIDResponse
	(*ChangePermissionRequest)(nil),  // 4: identity.userapi.v1.ChangePermissionRequest
	(*ChangePermissionResponse)(nil), // 5: identity.userapi.v1.ChangePermissionResponse
	(*ChangeStatusRequest)(nil),      // 6: identity.userapi.v1.ChangeStatusRequest
	(*ChangeStatusResponse)(nil),     // 7: identity.userapi.v1.ChangeStatusResponse
}
var file_identity_userapi_v1_userapi_proto_depIdxs = []int32{
	0, // 0: identity.userapi.v1.UserService.CreateUser:input_type -> identity.userapi.v1.CreateUserRequest
	2, // 1: identity.userapi.v1.UserService.UserByID:input_type -> identity.userapi.v1.UserByIDRequest
	4, // 2: identity.userapi.v1.UserService.ChangePermission:input_type -> identity.userapi.v1.ChangePermissionRequest
	6, // 3: identity.userapi.v1.UserService.ChangeStatus:input_type -> identity.userapi.v1.ChangeStatusRequest
	1, // 4: identity.userapi.v1.UserService.CreateUser:output_type -> identity.userapi.v1.CreateUserResponse
	3, // 5: identity.userapi.v1.UserService.UserByID:output_type -> identity.userapi.v1.UserByIDResponse
	5, // 6: identity.userapi.v1.UserService.ChangePermission:output_type -> identity.userapi.v1.ChangePermissionResponse
	7, // 7: identity.userapi.v1.UserService.ChangeStatus:output_type -> identity.userapi.v1.ChangeStatusResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identity_userapi_v1_userapi_proto_init() }
func file_identity_userapi_v1_userapi_proto_init() {
	if File_identity_userapi_v1_userapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_userapi_v1_userapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIDRequest); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserByIDResponse); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePermissionRequest); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePermissionResponse); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusRequest); i {
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
		file_identity_userapi_v1_userapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusResponse); i {
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
			RawDescriptor: file_identity_userapi_v1_userapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identity_userapi_v1_userapi_proto_goTypes,
		DependencyIndexes: file_identity_userapi_v1_userapi_proto_depIdxs,
		MessageInfos:      file_identity_userapi_v1_userapi_proto_msgTypes,
	}.Build()
	File_identity_userapi_v1_userapi_proto = out.File
	file_identity_userapi_v1_userapi_proto_rawDesc = nil
	file_identity_userapi_v1_userapi_proto_goTypes = nil
	file_identity_userapi_v1_userapi_proto_depIdxs = nil
}
