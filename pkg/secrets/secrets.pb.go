// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/secrets/secrets.proto

package secrets

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SecretTypes int32

const (
	SecretTypes_SECRET   SecretTypes = 0
	SecretTypes_REGISTRY SecretTypes = 1
)

// Enum value maps for SecretTypes.
var (
	SecretTypes_name = map[int32]string{
		0: "SECRET",
		1: "REGISTRY",
	}
	SecretTypes_value = map[string]int32{
		"SECRET":   0,
		"REGISTRY": 1,
	}
)

func (x SecretTypes) Enum() *SecretTypes {
	p := new(SecretTypes)
	*p = x
	return p
}

func (x SecretTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecretTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_secrets_secrets_proto_enumTypes[0].Descriptor()
}

func (SecretTypes) Type() protoreflect.EnumType {
	return &file_pkg_secrets_secrets_proto_enumTypes[0]
}

func (x SecretTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecretTypes.Descriptor instead.
func (SecretTypes) EnumDescriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{0}
}

type SecretsDeleteBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *SecretsDeleteBucketRequest) Reset() {
	*x = SecretsDeleteBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDeleteBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDeleteBucketRequest) ProtoMessage() {}

func (x *SecretsDeleteBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDeleteBucketRequest.ProtoReflect.Descriptor instead.
func (*SecretsDeleteBucketRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *SecretsDeleteBucketRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type SecretsCreateBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *SecretsCreateBucketRequest) Reset() {
	*x = SecretsCreateBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsCreateBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsCreateBucketRequest) ProtoMessage() {}

func (x *SecretsCreateBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsCreateBucketRequest.ProtoReflect.Descriptor instead.
func (*SecretsCreateBucketRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *SecretsCreateBucketRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type SecretsCreateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pkey []byte `protobuf:"bytes,1,opt,name=pkey,proto3,oneof" json:"pkey,omitempty"`
}

func (x *SecretsCreateBucketResponse) Reset() {
	*x = SecretsCreateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsCreateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsCreateBucketResponse) ProtoMessage() {}

func (x *SecretsCreateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsCreateBucketResponse.ProtoReflect.Descriptor instead.
func (*SecretsCreateBucketResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *SecretsCreateBucketResponse) GetPkey() []byte {
	if x != nil {
		return x.Pkey
	}
	return nil
}

type SecretsStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string      `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string      `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Data      []byte       `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Stype     *SecretTypes `protobuf:"varint,4,opt,name=stype,proto3,enum=secrets.SecretTypes,oneof" json:"stype,omitempty"`
}

func (x *SecretsStoreRequest) Reset() {
	*x = SecretsStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsStoreRequest) ProtoMessage() {}

func (x *SecretsStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsStoreRequest.ProtoReflect.Descriptor instead.
func (*SecretsStoreRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *SecretsStoreRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsStoreRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SecretsStoreRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SecretsStoreRequest) GetStype() SecretTypes {
	if x != nil && x.Stype != nil {
		return *x.Stype
	}
	return SecretTypes_SECRET
}

type SecretsRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *SecretsRetrieveRequest) Reset() {
	*x = SecretsRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRetrieveRequest) ProtoMessage() {}

func (x *SecretsRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRetrieveRequest.ProtoReflect.Descriptor instead.
func (*SecretsRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{4}
}

func (x *SecretsRetrieveRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsRetrieveRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type SecretsDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string      `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Name      *string      `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Stype     *SecretTypes `protobuf:"varint,3,opt,name=stype,proto3,enum=secrets.SecretTypes,oneof" json:"stype,omitempty"`
}

func (x *SecretsDeleteRequest) Reset() {
	*x = SecretsDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDeleteRequest) ProtoMessage() {}

func (x *SecretsDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDeleteRequest.ProtoReflect.Descriptor instead.
func (*SecretsDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{5}
}

func (x *SecretsDeleteRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *SecretsDeleteRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SecretsDeleteRequest) GetStype() SecretTypes {
	if x != nil && x.Stype != nil {
		return *x.Stype
	}
	return SecretTypes_SECRET
}

type SecretsDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *int32 `protobuf:"varint,1,opt,name=count,proto3,oneof" json:"count,omitempty"`
}

func (x *SecretsDeleteResponse) Reset() {
	*x = SecretsDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDeleteResponse) ProtoMessage() {}

func (x *SecretsDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDeleteResponse.ProtoReflect.Descriptor instead.
func (*SecretsDeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{6}
}

func (x *SecretsDeleteResponse) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type SecretsRetrieveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *SecretsRetrieveResponse) Reset() {
	*x = SecretsRetrieveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRetrieveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRetrieveResponse) ProtoMessage() {}

func (x *SecretsRetrieveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRetrieveResponse.ProtoReflect.Descriptor instead.
func (*SecretsRetrieveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{7}
}

func (x *SecretsRetrieveResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string      `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	Stype     *SecretTypes `protobuf:"varint,2,opt,name=stype,proto3,enum=secrets.SecretTypes,oneof" json:"stype,omitempty"`
}

func (x *GetSecretsRequest) Reset() {
	*x = GetSecretsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsRequest) ProtoMessage() {}

func (x *GetSecretsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsRequest.ProtoReflect.Descriptor instead.
func (*GetSecretsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{8}
}

func (x *GetSecretsRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *GetSecretsRequest) GetStype() SecretTypes {
	if x != nil && x.Stype != nil {
		return *x.Stype
	}
	return SecretTypes_SECRET
}

type GetSecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets []*GetSecretsResponse_Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *GetSecretsResponse) Reset() {
	*x = GetSecretsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse) ProtoMessage() {}

func (x *GetSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{9}
}

func (x *GetSecretsResponse) GetSecrets() []*GetSecretsResponse_Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type GetSecretsDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets []*GetSecretsDataResponse_Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (x *GetSecretsDataResponse) Reset() {
	*x = GetSecretsDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsDataResponse) ProtoMessage() {}

func (x *GetSecretsDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsDataResponse.ProtoReflect.Descriptor instead.
func (*GetSecretsDataResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{10}
}

func (x *GetSecretsDataResponse) GetSecrets() []*GetSecretsDataResponse_Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type GetSecretsResponse_Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *GetSecretsResponse_Secret) Reset() {
	*x = GetSecretsResponse_Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsResponse_Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsResponse_Secret) ProtoMessage() {}

func (x *GetSecretsResponse_Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsResponse_Secret.ProtoReflect.Descriptor instead.
func (*GetSecretsResponse_Secret) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{9, 0}
}

func (x *GetSecretsResponse_Secret) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type GetSecretsDataResponse_Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *GetSecretsDataResponse_Secret) Reset() {
	*x = GetSecretsDataResponse_Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_secrets_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretsDataResponse_Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretsDataResponse_Secret) ProtoMessage() {}

func (x *GetSecretsDataResponse_Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_secrets_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretsDataResponse_Secret.ProtoReflect.Descriptor instead.
func (*GetSecretsDataResponse_Secret) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_secrets_proto_rawDescGZIP(), []int{10, 0}
}

func (x *GetSecretsDataResponse_Secret) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetSecretsDataResponse_Secret) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_secrets_secrets_proto protoreflect.FileDescriptor

var file_pkg_secrets_secrets_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x1a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x1a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0x3f, 0x0a, 0x1b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70,
	0x6b, 0x65, 0x79, 0x22, 0xc5, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x03, 0x52, 0x05, 0x73, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x16, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a,
	0x05, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x48, 0x02, 0x52, 0x05, 0x73, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3c, 0x0a, 0x15, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a,
	0x17, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x48, 0x01, 0x52, 0x05, 0x73, 0x74, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a,
	0x2a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52,
	0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x4c, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x10, 0x01, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f,
	0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_secrets_secrets_proto_rawDescOnce sync.Once
	file_pkg_secrets_secrets_proto_rawDescData = file_pkg_secrets_secrets_proto_rawDesc
)

func file_pkg_secrets_secrets_proto_rawDescGZIP() []byte {
	file_pkg_secrets_secrets_proto_rawDescOnce.Do(func() {
		file_pkg_secrets_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_secrets_secrets_proto_rawDescData)
	})
	return file_pkg_secrets_secrets_proto_rawDescData
}

var file_pkg_secrets_secrets_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_secrets_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_secrets_secrets_proto_goTypes = []interface{}{
	(SecretTypes)(0),                      // 0: secrets.SecretTypes
	(*SecretsDeleteBucketRequest)(nil),    // 1: secrets.SecretsDeleteBucketRequest
	(*SecretsCreateBucketRequest)(nil),    // 2: secrets.SecretsCreateBucketRequest
	(*SecretsCreateBucketResponse)(nil),   // 3: secrets.SecretsCreateBucketResponse
	(*SecretsStoreRequest)(nil),           // 4: secrets.SecretsStoreRequest
	(*SecretsRetrieveRequest)(nil),        // 5: secrets.SecretsRetrieveRequest
	(*SecretsDeleteRequest)(nil),          // 6: secrets.SecretsDeleteRequest
	(*SecretsDeleteResponse)(nil),         // 7: secrets.SecretsDeleteResponse
	(*SecretsRetrieveResponse)(nil),       // 8: secrets.SecretsRetrieveResponse
	(*GetSecretsRequest)(nil),             // 9: secrets.GetSecretsRequest
	(*GetSecretsResponse)(nil),            // 10: secrets.GetSecretsResponse
	(*GetSecretsDataResponse)(nil),        // 11: secrets.GetSecretsDataResponse
	(*GetSecretsResponse_Secret)(nil),     // 12: secrets.GetSecretsResponse.Secret
	(*GetSecretsDataResponse_Secret)(nil), // 13: secrets.GetSecretsDataResponse.Secret
}
var file_pkg_secrets_secrets_proto_depIdxs = []int32{
	0,  // 0: secrets.SecretsStoreRequest.stype:type_name -> secrets.SecretTypes
	0,  // 1: secrets.SecretsDeleteRequest.stype:type_name -> secrets.SecretTypes
	0,  // 2: secrets.GetSecretsRequest.stype:type_name -> secrets.SecretTypes
	12, // 3: secrets.GetSecretsResponse.secrets:type_name -> secrets.GetSecretsResponse.Secret
	13, // 4: secrets.GetSecretsDataResponse.secrets:type_name -> secrets.GetSecretsDataResponse.Secret
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_secrets_secrets_proto_init() }
func file_pkg_secrets_secrets_proto_init() {
	if File_pkg_secrets_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_secrets_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDeleteBucketRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsCreateBucketRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsCreateBucketResponse); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsStoreRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRetrieveRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDeleteRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDeleteResponse); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRetrieveResponse); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsRequest); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsDataResponse); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsResponse_Secret); i {
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
		file_pkg_secrets_secrets_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretsDataResponse_Secret); i {
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
	file_pkg_secrets_secrets_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[11].OneofWrappers = []interface{}{}
	file_pkg_secrets_secrets_proto_msgTypes[12].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_secrets_secrets_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_secrets_secrets_proto_goTypes,
		DependencyIndexes: file_pkg_secrets_secrets_proto_depIdxs,
		EnumInfos:         file_pkg_secrets_secrets_proto_enumTypes,
		MessageInfos:      file_pkg_secrets_secrets_proto_msgTypes,
	}.Build()
	File_pkg_secrets_secrets_proto = out.File
	file_pkg_secrets_secrets_proto_rawDesc = nil
	file_pkg_secrets_secrets_proto_goTypes = nil
	file_pkg_secrets_secrets_proto_depIdxs = nil
}
