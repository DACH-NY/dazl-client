// Copyright (c) 2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: com/daml/ledger/api/v1/version_service.proto

package v1

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

type GetLedgerApiVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
}

func (x *GetLedgerApiVersionRequest) Reset() {
	*x = GetLedgerApiVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_version_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgerApiVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgerApiVersionRequest) ProtoMessage() {}

func (x *GetLedgerApiVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_version_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgerApiVersionRequest.ProtoReflect.Descriptor instead.
func (*GetLedgerApiVersionRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_version_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLedgerApiVersionRequest) GetLedgerId() string {
	if x != nil {
		return x.LedgerId
	}
	return ""
}

type GetLedgerApiVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the ledger API
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetLedgerApiVersionResponse) Reset() {
	*x = GetLedgerApiVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_version_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgerApiVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgerApiVersionResponse) ProtoMessage() {}

func (x *GetLedgerApiVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_version_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgerApiVersionResponse.ProtoReflect.Descriptor instead.
func (*GetLedgerApiVersionResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_version_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLedgerApiVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_com_daml_ledger_api_v1_version_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_version_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x37, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x90, 0x01, 0x0a, 0x0e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x82, 0x01,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x67,
	0x6f, 0x2f, 0x76, 0x37, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x6d, 0x2e,
	0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_version_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_version_service_proto_rawDescData = file_com_daml_ledger_api_v1_version_service_proto_rawDesc
)

func file_com_daml_ledger_api_v1_version_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_version_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_version_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_version_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_version_service_proto_rawDescData
}

var file_com_daml_ledger_api_v1_version_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_daml_ledger_api_v1_version_service_proto_goTypes = []interface{}{
	(*GetLedgerApiVersionRequest)(nil),  // 0: com.daml.ledger.api.v1.GetLedgerApiVersionRequest
	(*GetLedgerApiVersionResponse)(nil), // 1: com.daml.ledger.api.v1.GetLedgerApiVersionResponse
}
var file_com_daml_ledger_api_v1_version_service_proto_depIdxs = []int32{
	0, // 0: com.daml.ledger.api.v1.VersionService.GetLedgerApiVersion:input_type -> com.daml.ledger.api.v1.GetLedgerApiVersionRequest
	1, // 1: com.daml.ledger.api.v1.VersionService.GetLedgerApiVersion:output_type -> com.daml.ledger.api.v1.GetLedgerApiVersionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_version_service_proto_init() }
func file_com_daml_ledger_api_v1_version_service_proto_init() {
	if File_com_daml_ledger_api_v1_version_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_version_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedgerApiVersionRequest); i {
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
		file_com_daml_ledger_api_v1_version_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedgerApiVersionResponse); i {
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
			RawDescriptor: file_com_daml_ledger_api_v1_version_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v1_version_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_version_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_version_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_version_service_proto = out.File
	file_com_daml_ledger_api_v1_version_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_version_service_proto_goTypes = nil
	file_com_daml_ledger_api_v1_version_service_proto_depIdxs = nil
}