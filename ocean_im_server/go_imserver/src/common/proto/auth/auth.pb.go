// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: auth.proto

package oceanImAuth

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

type UserRegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	PassWord string `protobuf:"bytes,4,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birth    string `protobuf:"bytes,6,opt,name=birth,proto3" json:"birth,omitempty"`
	HeadImg  string `protobuf:"bytes,7,opt,name=headImg,proto3" json:"headImg,omitempty"`
	Gender   int32  `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Ex       string `protobuf:"bytes,9,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *UserRegisterReq) Reset() {
	*x = UserRegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterReq) ProtoMessage() {}

func (x *UserRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterReq.ProtoReflect.Descriptor instead.
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserRegisterReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserRegisterReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserRegisterReq) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *UserRegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterReq) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *UserRegisterReq) GetHeadImg() string {
	if x != nil {
		return x.HeadImg
	}
	return ""
}

func (x *UserRegisterReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserRegisterReq) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

type UserLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	PassWord string `protobuf:"bytes,4,opt,name=passWord,proto3" json:"passWord,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLoginReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserLoginReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserLoginReq) GetPassWord() string {
	if x != nil {
		return x.PassWord
	}
	return ""
}

func (x *UserLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform int32  `protobuf:"varint,1,opt,name=platform,proto3" json:"platform,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserTokenReq) Reset() {
	*x = UserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenReq) ProtoMessage() {}

func (x *UserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenReq.ProtoReflect.Descriptor instead.
func (*UserTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserTokenReq) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *UserTokenReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserRegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrCode int32  `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *UserRegisterResp) Reset() {
	*x = UserRegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResp) ProtoMessage() {}

func (x *UserRegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResp.ProtoReflect.Descriptor instead.
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegisterResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UserRegisterResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *UserRegisterResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type UserLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrCode int32  `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *UserLoginResp) Reset() {
	*x = UserLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResp) ProtoMessage() {}

func (x *UserLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResp.ProtoReflect.Descriptor instead.
func (*UserLoginResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserLoginResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UserLoginResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *UserLoginResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type UserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode     int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg      string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Token       string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredTime int64  `protobuf:"varint,4,opt,name=expiredTime,proto3" json:"expiredTime,omitempty"`
}

func (x *UserTokenResp) Reset() {
	*x = UserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenResp) ProtoMessage() {}

func (x *UserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenResp.ProtoReflect.Descriptor instead.
func (*UserTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserTokenResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *UserTokenResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *UserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserTokenResp) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x22, 0xe5, 0x01, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x57,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x57,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65,
	0x78, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x42,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x5e, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x22, 0x5b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22,
	0x79, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xdb, 0x01, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x42, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x19, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x49,
	0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f,
	0x63, 0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x49, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(*UserRegisterReq)(nil),  // 0: oceanImAuth.UserRegisterReq
	(*UserLoginReq)(nil),     // 1: oceanImAuth.UserLoginReq
	(*UserTokenReq)(nil),     // 2: oceanImAuth.UserTokenReq
	(*UserRegisterResp)(nil), // 3: oceanImAuth.UserRegisterResp
	(*UserLoginResp)(nil),    // 4: oceanImAuth.UserLoginResp
	(*UserTokenResp)(nil),    // 5: oceanImAuth.UserTokenResp
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: oceanImAuth.Auth.UserLogin:input_type -> oceanImAuth.UserLoginReq
	0, // 1: oceanImAuth.Auth.UserRegister:input_type -> oceanImAuth.UserRegisterReq
	2, // 2: oceanImAuth.Auth.UserToken:input_type -> oceanImAuth.UserTokenReq
	4, // 3: oceanImAuth.Auth.UserLogin:output_type -> oceanImAuth.UserLoginResp
	3, // 4: oceanImAuth.Auth.UserRegister:output_type -> oceanImAuth.UserRegisterResp
	5, // 5: oceanImAuth.Auth.UserToken:output_type -> oceanImAuth.UserTokenResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterReq); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReq); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenReq); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResp); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResp); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenResp); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
