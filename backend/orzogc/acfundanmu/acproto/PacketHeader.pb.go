// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: PacketHeader.proto

package acproto

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

type PacketHeader_Flags int32

const (
	PacketHeader_kDirUpstream   PacketHeader_Flags = 0
	PacketHeader_kDirDownstream PacketHeader_Flags = 1
	PacketHeader_kDirMask       PacketHeader_Flags = 1
)

// Enum value maps for PacketHeader_Flags.
var (
	PacketHeader_Flags_name = map[int32]string{
		0: "kDirUpstream",
		1: "kDirDownstream",
		// Duplicate value: 1: "kDirMask",
	}
	PacketHeader_Flags_value = map[string]int32{
		"kDirUpstream":   0,
		"kDirDownstream": 1,
		"kDirMask":       1,
	}
)

func (x PacketHeader_Flags) Enum() *PacketHeader_Flags {
	p := new(PacketHeader_Flags)
	*p = x
	return p
}

func (x PacketHeader_Flags) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketHeader_Flags) Descriptor() protoreflect.EnumDescriptor {
	return file_PacketHeader_proto_enumTypes[0].Descriptor()
}

func (PacketHeader_Flags) Type() protoreflect.EnumType {
	return &file_PacketHeader_proto_enumTypes[0]
}

func (x PacketHeader_Flags) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketHeader_Flags.Descriptor instead.
func (PacketHeader_Flags) EnumDescriptor() ([]byte, []int) {
	return file_PacketHeader_proto_rawDescGZIP(), []int{0, 0}
}

type PacketHeader_EncodingType int32

const (
	PacketHeader_kEncodingNone PacketHeader_EncodingType = 0
	PacketHeader_kEncodingLz4  PacketHeader_EncodingType = 1
)

// Enum value maps for PacketHeader_EncodingType.
var (
	PacketHeader_EncodingType_name = map[int32]string{
		0: "kEncodingNone",
		1: "kEncodingLz4",
	}
	PacketHeader_EncodingType_value = map[string]int32{
		"kEncodingNone": 0,
		"kEncodingLz4":  1,
	}
)

func (x PacketHeader_EncodingType) Enum() *PacketHeader_EncodingType {
	p := new(PacketHeader_EncodingType)
	*p = x
	return p
}

func (x PacketHeader_EncodingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketHeader_EncodingType) Descriptor() protoreflect.EnumDescriptor {
	return file_PacketHeader_proto_enumTypes[1].Descriptor()
}

func (PacketHeader_EncodingType) Type() protoreflect.EnumType {
	return &file_PacketHeader_proto_enumTypes[1]
}

func (x PacketHeader_EncodingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketHeader_EncodingType.Descriptor instead.
func (PacketHeader_EncodingType) EnumDescriptor() ([]byte, []int) {
	return file_PacketHeader_proto_rawDescGZIP(), []int{0, 1}
}

type PacketHeader_EncryptionMode int32

const (
	PacketHeader_kEncryptionNone         PacketHeader_EncryptionMode = 0
	PacketHeader_kEncryptionServiceToken PacketHeader_EncryptionMode = 1
	PacketHeader_kEncryptionSessionKey   PacketHeader_EncryptionMode = 2
)

// Enum value maps for PacketHeader_EncryptionMode.
var (
	PacketHeader_EncryptionMode_name = map[int32]string{
		0: "kEncryptionNone",
		1: "kEncryptionServiceToken",
		2: "kEncryptionSessionKey",
	}
	PacketHeader_EncryptionMode_value = map[string]int32{
		"kEncryptionNone":         0,
		"kEncryptionServiceToken": 1,
		"kEncryptionSessionKey":   2,
	}
)

func (x PacketHeader_EncryptionMode) Enum() *PacketHeader_EncryptionMode {
	p := new(PacketHeader_EncryptionMode)
	*p = x
	return p
}

func (x PacketHeader_EncryptionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketHeader_EncryptionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_PacketHeader_proto_enumTypes[2].Descriptor()
}

func (PacketHeader_EncryptionMode) Type() protoreflect.EnumType {
	return &file_PacketHeader_proto_enumTypes[2]
}

func (x PacketHeader_EncryptionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketHeader_EncryptionMode.Descriptor instead.
func (PacketHeader_EncryptionMode) EnumDescriptor() ([]byte, []int) {
	return file_PacketHeader_proto_rawDescGZIP(), []int{0, 2}
}

type PacketHeader_Feature int32

const (
	PacketHeader_kReserve     PacketHeader_Feature = 0
	PacketHeader_kCompressLz4 PacketHeader_Feature = 1
)

// Enum value maps for PacketHeader_Feature.
var (
	PacketHeader_Feature_name = map[int32]string{
		0: "kReserve",
		1: "kCompressLz4",
	}
	PacketHeader_Feature_value = map[string]int32{
		"kReserve":     0,
		"kCompressLz4": 1,
	}
)

func (x PacketHeader_Feature) Enum() *PacketHeader_Feature {
	p := new(PacketHeader_Feature)
	*p = x
	return p
}

func (x PacketHeader_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketHeader_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_PacketHeader_proto_enumTypes[3].Descriptor()
}

func (PacketHeader_Feature) Type() protoreflect.EnumType {
	return &file_PacketHeader_proto_enumTypes[3]
}

func (x PacketHeader_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketHeader_Feature.Descriptor instead.
func (PacketHeader_Feature) EnumDescriptor() ([]byte, []int) {
	return file_PacketHeader_proto_rawDescGZIP(), []int{0, 3}
}

type PacketHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId             int32                       `protobuf:"varint,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Uid               int64                       `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	InstanceId        int64                       `protobuf:"varint,3,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	Flags             uint32                      `protobuf:"varint,5,opt,name=flags,proto3" json:"flags,omitempty"`
	EncodingType      PacketHeader_EncodingType   `protobuf:"varint,6,opt,name=encodingType,proto3,enum=AcFunDanmu.PacketHeader_EncodingType" json:"encodingType,omitempty"`
	DecodedPayloadLen int32                       `protobuf:"varint,7,opt,name=decodedPayloadLen,proto3" json:"decodedPayloadLen,omitempty"`
	EncryptionMode    PacketHeader_EncryptionMode `protobuf:"varint,8,opt,name=encryptionMode,proto3,enum=AcFunDanmu.PacketHeader_EncryptionMode" json:"encryptionMode,omitempty"`
	TokenInfo         *TokenInfo                  `protobuf:"bytes,9,opt,name=tokenInfo,proto3" json:"tokenInfo,omitempty"`
	SeqId             int64                       `protobuf:"varint,10,opt,name=seqId,proto3" json:"seqId,omitempty"`
	Features          []PacketHeader_Feature      `protobuf:"varint,11,rep,packed,name=features,proto3,enum=AcFunDanmu.PacketHeader_Feature" json:"features,omitempty"`
	Kpn               string                      `protobuf:"bytes,12,opt,name=kpn,proto3" json:"kpn,omitempty"`
}

func (x *PacketHeader) Reset() {
	*x = PacketHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PacketHeader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketHeader) ProtoMessage() {}

func (x *PacketHeader) ProtoReflect() protoreflect.Message {
	mi := &file_PacketHeader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketHeader.ProtoReflect.Descriptor instead.
func (*PacketHeader) Descriptor() ([]byte, []int) {
	return file_PacketHeader_proto_rawDescGZIP(), []int{0}
}

func (x *PacketHeader) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *PacketHeader) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PacketHeader) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *PacketHeader) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *PacketHeader) GetEncodingType() PacketHeader_EncodingType {
	if x != nil {
		return x.EncodingType
	}
	return PacketHeader_kEncodingNone
}

func (x *PacketHeader) GetDecodedPayloadLen() int32 {
	if x != nil {
		return x.DecodedPayloadLen
	}
	return 0
}

func (x *PacketHeader) GetEncryptionMode() PacketHeader_EncryptionMode {
	if x != nil {
		return x.EncryptionMode
	}
	return PacketHeader_kEncryptionNone
}

func (x *PacketHeader) GetTokenInfo() *TokenInfo {
	if x != nil {
		return x.TokenInfo
	}
	return nil
}

func (x *PacketHeader) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *PacketHeader) GetFeatures() []PacketHeader_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *PacketHeader) GetKpn() string {
	if x != nil {
		return x.Kpn
	}
	return ""
}

var File_PacketHeader_proto protoreflect.FileDescriptor

var file_PacketHeader_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x1a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x05, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x49, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x65, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x65, 0x6e, 0x12, 0x4f, 0x0a, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x70, 0x6e, 0x22, 0x3f, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x10,
	0x0a, 0x0c, 0x6b, 0x44, 0x69, 0x72, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x6b, 0x44, 0x69, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x44, 0x69, 0x72, 0x4d, 0x61, 0x73, 0x6b,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x33, 0x0a, 0x0c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x6b, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x7a, 0x34, 0x10, 0x01, 0x22, 0x5d, 0x0a, 0x0e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x6b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x6b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x10, 0x02, 0x22, 0x29, 0x0a, 0x07, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x7a, 0x34, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PacketHeader_proto_rawDescOnce sync.Once
	file_PacketHeader_proto_rawDescData = file_PacketHeader_proto_rawDesc
)

func file_PacketHeader_proto_rawDescGZIP() []byte {
	file_PacketHeader_proto_rawDescOnce.Do(func() {
		file_PacketHeader_proto_rawDescData = protoimpl.X.CompressGZIP(file_PacketHeader_proto_rawDescData)
	})
	return file_PacketHeader_proto_rawDescData
}

var file_PacketHeader_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_PacketHeader_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PacketHeader_proto_goTypes = []interface{}{
	(PacketHeader_Flags)(0),          // 0: AcFunDanmu.PacketHeader.Flags
	(PacketHeader_EncodingType)(0),   // 1: AcFunDanmu.PacketHeader.EncodingType
	(PacketHeader_EncryptionMode)(0), // 2: AcFunDanmu.PacketHeader.EncryptionMode
	(PacketHeader_Feature)(0),        // 3: AcFunDanmu.PacketHeader.Feature
	(*PacketHeader)(nil),             // 4: AcFunDanmu.PacketHeader
	(*TokenInfo)(nil),                // 5: AcFunDanmu.TokenInfo
}
var file_PacketHeader_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.PacketHeader.encodingType:type_name -> AcFunDanmu.PacketHeader.EncodingType
	2, // 1: AcFunDanmu.PacketHeader.encryptionMode:type_name -> AcFunDanmu.PacketHeader.EncryptionMode
	5, // 2: AcFunDanmu.PacketHeader.tokenInfo:type_name -> AcFunDanmu.TokenInfo
	3, // 3: AcFunDanmu.PacketHeader.features:type_name -> AcFunDanmu.PacketHeader.Feature
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_PacketHeader_proto_init() }
func file_PacketHeader_proto_init() {
	if File_PacketHeader_proto != nil {
		return
	}
	file_TokenInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PacketHeader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketHeader); i {
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
			RawDescriptor: file_PacketHeader_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PacketHeader_proto_goTypes,
		DependencyIndexes: file_PacketHeader_proto_depIdxs,
		EnumInfos:         file_PacketHeader_proto_enumTypes,
		MessageInfos:      file_PacketHeader_proto_msgTypes,
	}.Build()
	File_PacketHeader_proto = out.File
	file_PacketHeader_proto_rawDesc = nil
	file_PacketHeader_proto_goTypes = nil
	file_PacketHeader_proto_depIdxs = nil
}