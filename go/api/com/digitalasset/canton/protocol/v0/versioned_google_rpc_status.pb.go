// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: com/digitalasset/canton/protocol/v0/versioned_google_rpc_status.proto

package v0

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type VersionedStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VersionedStatus) Reset() {
	*x = VersionedStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionedStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionedStatus) ProtoMessage() {}

func (x *VersionedStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionedStatus.ProtoReflect.Descriptor instead.
func (*VersionedStatus) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescGZIP(), []int{0}
}

func (x *VersionedStatus) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescData = file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_goTypes = []interface{}{
	(*VersionedStatus)(nil), // 0: com.digitalasset.canton.protocol.v0.VersionedStatus
	(*status.Status)(nil),   // 1: google.rpc.Status
}
var file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.protocol.v0.VersionedStatus.status:type_name -> google.rpc.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_init() }
func file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_init() {
	if File_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionedStatus); i {
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
			RawDescriptor: file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto = out.File
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v0_versioned_google_rpc_status_proto_depIdxs = nil
}
