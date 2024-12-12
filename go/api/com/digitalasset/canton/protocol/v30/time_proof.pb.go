// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/time/v30/time_proof.proto

package v30

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

type TimeProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *PossiblyIgnoredSequencedEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *TimeProof) Reset() {
	*x = TimeProof{}
	mi := &file_com_digitalasset_canton_time_v30_time_proof_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeProof) ProtoMessage() {}

func (x *TimeProof) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_time_v30_time_proof_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeProof.ProtoReflect.Descriptor instead.
func (*TimeProof) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescGZIP(), []int{0}
}

func (x *TimeProof) GetEvent() *PossiblyIgnoredSequencedEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_com_digitalasset_canton_time_v30_time_proof_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_time_v30_time_proof_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x33, 0x30, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x33, 0x30, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x30, 0x2f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x59, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x30, 0x2e,
	0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x79, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescData = file_com_digitalasset_canton_time_v30_time_proof_proto_rawDesc
)

func file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescData)
	})
	return file_com_digitalasset_canton_time_v30_time_proof_proto_rawDescData
}

var file_com_digitalasset_canton_time_v30_time_proof_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_time_v30_time_proof_proto_goTypes = []any{
	(*TimeProof)(nil),                     // 0: com.digitalasset.canton.time.v30.TimeProof
	(*PossiblyIgnoredSequencedEvent)(nil), // 1: com.digitalasset.canton.protocol.v30.PossiblyIgnoredSequencedEvent
}
var file_com_digitalasset_canton_time_v30_time_proof_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.time.v30.TimeProof.event:type_name -> com.digitalasset.canton.protocol.v30.PossiblyIgnoredSequencedEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_time_v30_time_proof_proto_init() }
func file_com_digitalasset_canton_time_v30_time_proof_proto_init() {
	if File_com_digitalasset_canton_time_v30_time_proof_proto != nil {
		return
	}
	file_com_digitalasset_canton_protocol_v30_sequencing_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_time_v30_time_proof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_time_v30_time_proof_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_time_v30_time_proof_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_time_v30_time_proof_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_time_v30_time_proof_proto = out.File
	file_com_digitalasset_canton_time_v30_time_proof_proto_rawDesc = nil
	file_com_digitalasset_canton_time_v30_time_proof_proto_goTypes = nil
	file_com_digitalasset_canton_time_v30_time_proof_proto_depIdxs = nil
}
