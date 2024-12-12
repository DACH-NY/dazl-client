// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/protocol/v30/common.proto

package v30

import (
	v30 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/crypto/v30"
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

type ViewType int32

const (
	ViewType_VIEW_TYPE_UNSPECIFIED  ViewType = 0
	ViewType_VIEW_TYPE_TRANSACTION  ViewType = 1
	ViewType_VIEW_TYPE_TRANSFER_OUT ViewType = 2
	ViewType_VIEW_TYPE_TRANSFER_IN  ViewType = 3
)

// Enum value maps for ViewType.
var (
	ViewType_name = map[int32]string{
		0: "VIEW_TYPE_UNSPECIFIED",
		1: "VIEW_TYPE_TRANSACTION",
		2: "VIEW_TYPE_TRANSFER_OUT",
		3: "VIEW_TYPE_TRANSFER_IN",
	}
	ViewType_value = map[string]int32{
		"VIEW_TYPE_UNSPECIFIED":  0,
		"VIEW_TYPE_TRANSACTION":  1,
		"VIEW_TYPE_TRANSFER_OUT": 2,
		"VIEW_TYPE_TRANSFER_IN":  3,
	}
)

func (x ViewType) Enum() *ViewType {
	p := new(ViewType)
	*p = x
	return p
}

func (x ViewType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ViewType) Descriptor() protoreflect.EnumDescriptor {
	return file_com_digitalasset_canton_protocol_v30_common_proto_enumTypes[0].Descriptor()
}

func (ViewType) Type() protoreflect.EnumType {
	return &file_com_digitalasset_canton_protocol_v30_common_proto_enumTypes[0]
}

func (x ViewType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ViewType.Descriptor instead.
func (ViewType) EnumDescriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v30_common_proto_rawDescGZIP(), []int{0}
}

type DriverContractMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractSalt *v30.Salt `protobuf:"bytes,1,opt,name=contract_salt,json=contractSalt,proto3" json:"contract_salt,omitempty"`
}

func (x *DriverContractMetadata) Reset() {
	*x = DriverContractMetadata{}
	mi := &file_com_digitalasset_canton_protocol_v30_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DriverContractMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverContractMetadata) ProtoMessage() {}

func (x *DriverContractMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v30_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverContractMetadata.ProtoReflect.Descriptor instead.
func (*DriverContractMetadata) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v30_common_proto_rawDescGZIP(), []int{0}
}

func (x *DriverContractMetadata) GetContractSalt() *v30.Salt {
	if x != nil {
		return x.ContractSalt
	}
	return nil
}

var File_com_digitalasset_canton_protocol_v30_common_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v30_common_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x30, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x30, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x16, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x33, 0x30,
	0x2e, 0x53, 0x61, 0x6c, 0x74, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53,
	0x61, 0x6c, 0x74, 0x2a, 0x77, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x56, 0x49,
	0x45, 0x57, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x55, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f,
	0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x33, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v30_common_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v30_common_proto_rawDescData = file_com_digitalasset_canton_protocol_v30_common_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v30_common_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v30_common_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v30_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v30_common_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v30_common_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v30_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_digitalasset_canton_protocol_v30_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_protocol_v30_common_proto_goTypes = []any{
	(ViewType)(0),                  // 0: com.digitalasset.canton.protocol.v30.ViewType
	(*DriverContractMetadata)(nil), // 1: com.digitalasset.canton.protocol.v30.DriverContractMetadata
	(*v30.Salt)(nil),               // 2: com.digitalasset.canton.crypto.v30.Salt
}
var file_com_digitalasset_canton_protocol_v30_common_proto_depIdxs = []int32{
	2, // 0: com.digitalasset.canton.protocol.v30.DriverContractMetadata.contract_salt:type_name -> com.digitalasset.canton.crypto.v30.Salt
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v30_common_proto_init() }
func file_com_digitalasset_canton_protocol_v30_common_proto_init() {
	if File_com_digitalasset_canton_protocol_v30_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_protocol_v30_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v30_common_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v30_common_proto_depIdxs,
		EnumInfos:         file_com_digitalasset_canton_protocol_v30_common_proto_enumTypes,
		MessageInfos:      file_com_digitalasset_canton_protocol_v30_common_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v30_common_proto = out.File
	file_com_digitalasset_canton_protocol_v30_common_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v30_common_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v30_common_proto_depIdxs = nil
}
