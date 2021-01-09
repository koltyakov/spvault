// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spvault.proto

package proto

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

type Strategy int32

const (
	Strategy_addin      Strategy = 0
	Strategy_adfs       Strategy = 1
	Strategy_fba        Strategy = 2
	Strategy_saml       Strategy = 3
	Strategy_tmg        Strategy = 4
	Strategy_azurecert  Strategy = 5
	Strategy_azurecreds Strategy = 6
)

// Enum value maps for Strategy.
var (
	Strategy_name = map[int32]string{
		0: "addin",
		1: "adfs",
		2: "fba",
		3: "saml",
		4: "tmg",
		5: "azurecert",
		6: "azurecreds",
	}
	Strategy_value = map[string]int32{
		"addin":      0,
		"adfs":       1,
		"fba":        2,
		"saml":       3,
		"tmg":        4,
		"azurecert":  5,
		"azurecreds": 6,
	}
)

func (x Strategy) Enum() *Strategy {
	p := new(Strategy)
	*p = x
	return p
}

func (x Strategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Strategy) Descriptor() protoreflect.EnumDescriptor {
	return file_spvault_proto_enumTypes[0].Descriptor()
}

func (Strategy) Type() protoreflect.EnumType {
	return &file_spvault_proto_enumTypes[0]
}

func (x Strategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Strategy.Descriptor instead.
func (Strategy) EnumDescriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{0}
}

type TokenType int32

const (
	TokenType_Bearer TokenType = 0
	TokenType_Cookie TokenType = 1
	TokenType_Custom TokenType = 3
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "Bearer",
		1: "Cookie",
		3: "Custom",
	}
	TokenType_value = map[string]int32{
		"Bearer": 0,
		"Cookie": 1,
		"Custom": 3,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_spvault_proto_enumTypes[1].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_spvault_proto_enumTypes[1]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{1}
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteUrl     string   `protobuf:"bytes,1,opt,name=siteUrl,proto3" json:"siteUrl,omitempty"`
	Strategy    Strategy `protobuf:"varint,2,opt,name=strategy,proto3,enum=spvault.Strategy" json:"strategy,omitempty"`
	Credentials string   `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *AuthRequest) GetStrategy() Strategy {
	if x != nil {
		return x.Strategy
	}
	return Strategy_addin
}

func (x *AuthRequest) GetCredentials() string {
	if x != nil {
		return x.Credentials
	}
	return ""
}

type TokenAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultToken string `protobuf:"bytes,1,opt,name=vaultToken,proto3" json:"vaultToken,omitempty"`
}

func (x *TokenAuthRequest) Reset() {
	*x = TokenAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenAuthRequest) ProtoMessage() {}

func (x *TokenAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenAuthRequest.ProtoReflect.Descriptor instead.
func (*TokenAuthRequest) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{1}
}

func (x *TokenAuthRequest) GetVaultToken() string {
	if x != nil {
		return x.VaultToken
	}
	return ""
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken  string    `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	TokenType  TokenType `protobuf:"varint,2,opt,name=tokenType,proto3,enum=spvault.TokenType" json:"tokenType,omitempty"`
	Expiration int64     `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{2}
}

func (x *AuthReply) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *AuthReply) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_Bearer
}

func (x *AuthReply) GetExpiration() int64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

type RegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthRequest *AuthRequest `protobuf:"bytes,1,opt,name=authRequest,proto3" json:"authRequest,omitempty"`
	VaultToken  string       `protobuf:"bytes,2,opt,name=vaultToken,proto3" json:"vaultToken,omitempty"` // optional requires --experimental_allow_proto3_optional which is not completely supported with some gRPC generators like .Net Grpc.Tools
}

func (x *RegRequest) Reset() {
	*x = RegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRequest) ProtoMessage() {}

func (x *RegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRequest.ProtoReflect.Descriptor instead.
func (*RegRequest) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{3}
}

func (x *RegRequest) GetAuthRequest() *AuthRequest {
	if x != nil {
		return x.AuthRequest
	}
	return nil
}

func (x *RegRequest) GetVaultToken() string {
	if x != nil {
		return x.VaultToken
	}
	return ""
}

type RegReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultToken string `protobuf:"bytes,1,opt,name=vaultToken,proto3" json:"vaultToken,omitempty"`
}

func (x *RegReply) Reset() {
	*x = RegReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegReply) ProtoMessage() {}

func (x *RegReply) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegReply.ProtoReflect.Descriptor instead.
func (*RegReply) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{4}
}

func (x *RegReply) GetVaultToken() string {
	if x != nil {
		return x.VaultToken
	}
	return ""
}

type DeRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultToken string `protobuf:"bytes,1,opt,name=vaultToken,proto3" json:"vaultToken,omitempty"`
}

func (x *DeRegRequest) Reset() {
	*x = DeRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeRegRequest) ProtoMessage() {}

func (x *DeRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeRegRequest.ProtoReflect.Descriptor instead.
func (*DeRegRequest) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{5}
}

func (x *DeRegRequest) GetVaultToken() string {
	if x != nil {
		return x.VaultToken
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spvault_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_spvault_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_spvault_proto_rawDescGZIP(), []int{6}
}

var File_spvault_proto protoreflect.FileDescriptor

var file_spvault_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x78, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x22, 0x32, 0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7b, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x30, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x0c, 0x44, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x5a,
	0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x64,
	0x64, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x61, 0x64, 0x66, 0x73, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x66, 0x62, 0x61, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x61, 0x6d, 0x6c,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x74, 0x6d, 0x67, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x63, 0x65, 0x72, 0x74, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x63, 0x72, 0x65, 0x64, 0x73, 0x10, 0x06, 0x2a, 0x2f, 0x0a, 0x09, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x10, 0x03, 0x32, 0x83, 0x02, 0x0a, 0x05,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x73, 0x12, 0x14,
	0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x15, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x52, 0x65, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x2e, 0x44, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x73, 0x70, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x07, 0x53,
	0x50, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spvault_proto_rawDescOnce sync.Once
	file_spvault_proto_rawDescData = file_spvault_proto_rawDesc
)

func file_spvault_proto_rawDescGZIP() []byte {
	file_spvault_proto_rawDescOnce.Do(func() {
		file_spvault_proto_rawDescData = protoimpl.X.CompressGZIP(file_spvault_proto_rawDescData)
	})
	return file_spvault_proto_rawDescData
}

var file_spvault_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spvault_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spvault_proto_goTypes = []interface{}{
	(Strategy)(0),            // 0: spvault.Strategy
	(TokenType)(0),           // 1: spvault.TokenType
	(*AuthRequest)(nil),      // 2: spvault.AuthRequest
	(*TokenAuthRequest)(nil), // 3: spvault.TokenAuthRequest
	(*AuthReply)(nil),        // 4: spvault.AuthReply
	(*RegRequest)(nil),       // 5: spvault.RegRequest
	(*RegReply)(nil),         // 6: spvault.RegReply
	(*DeRegRequest)(nil),     // 7: spvault.DeRegRequest
	(*Empty)(nil),            // 8: spvault.Empty
}
var file_spvault_proto_depIdxs = []int32{
	0, // 0: spvault.AuthRequest.strategy:type_name -> spvault.Strategy
	1, // 1: spvault.AuthReply.tokenType:type_name -> spvault.TokenType
	2, // 2: spvault.RegRequest.authRequest:type_name -> spvault.AuthRequest
	2, // 3: spvault.Vault.AuthenticateWithCreds:input_type -> spvault.AuthRequest
	3, // 4: spvault.Vault.AuthenticateWithToken:input_type -> spvault.TokenAuthRequest
	5, // 5: spvault.Vault.Register:input_type -> spvault.RegRequest
	7, // 6: spvault.Vault.DeRegister:input_type -> spvault.DeRegRequest
	4, // 7: spvault.Vault.AuthenticateWithCreds:output_type -> spvault.AuthReply
	4, // 8: spvault.Vault.AuthenticateWithToken:output_type -> spvault.AuthReply
	6, // 9: spvault.Vault.Register:output_type -> spvault.RegReply
	8, // 10: spvault.Vault.DeRegister:output_type -> spvault.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spvault_proto_init() }
func file_spvault_proto_init() {
	if File_spvault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spvault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_spvault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenAuthRequest); i {
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
		file_spvault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_spvault_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegRequest); i {
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
		file_spvault_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegReply); i {
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
		file_spvault_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeRegRequest); i {
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
		file_spvault_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_spvault_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spvault_proto_goTypes,
		DependencyIndexes: file_spvault_proto_depIdxs,
		EnumInfos:         file_spvault_proto_enumTypes,
		MessageInfos:      file_spvault_proto_msgTypes,
	}.Build()
	File_spvault_proto = out.File
	file_spvault_proto_rawDesc = nil
	file_spvault_proto_goTypes = nil
	file_spvault_proto_depIdxs = nil
}
