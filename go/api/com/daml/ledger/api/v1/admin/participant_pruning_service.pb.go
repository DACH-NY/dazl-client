// Copyright (c) 2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/daml/ledger/api/v1/admin/participant_pruning_service.proto

package admin

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

type PruneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Inclusive offset up to which the ledger is to be pruned.
	// By default the following data is pruned:
	//   1. All normal and divulged contracts that have been archived before
	//   `prune_up_to`.
	//   2. All transaction events and completions before `prune_up_to`
	PruneUpTo string `protobuf:"bytes,1,opt,name=prune_up_to,json=pruneUpTo,proto3" json:"prune_up_to,omitempty"`
	// Unique submission identifier.
	// Optional, defaults to a random identifier, used for logging.
	SubmissionId string `protobuf:"bytes,2,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	// Prune all immediately and retroactively divulged contracts created before `prune_up_to`
	// independent of whether they were archived before `prune_up_to`. Useful to avoid leaking
	// storage on participant nodes that can see a divulged contract but not its archival.
	//
	// Application developers SHOULD write their Daml applications
	// such that they do not rely on divulged contracts; i.e., no warnings from
	// using divulged contracts as inputs to transactions are emitted.
	//
	// Participant node operators SHOULD set the `prune_all_divulged_contracts` flag to avoid leaking
	// storage due to accumulating unarchived divulged contracts PROVIDED that:
	//   1. no application using this participant node relies on divulgence OR
	//   2. divulged contracts on which applications rely have been re-divulged after the `prune_up_to` offset.
	PruneAllDivulgedContracts bool `protobuf:"varint,3,opt,name=prune_all_divulged_contracts,json=pruneAllDivulgedContracts,proto3" json:"prune_all_divulged_contracts,omitempty"`
}

func (x *PruneRequest) Reset() {
	*x = PruneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneRequest) ProtoMessage() {}

func (x *PruneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneRequest.ProtoReflect.Descriptor instead.
func (*PruneRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescGZIP(), []int{0}
}

func (x *PruneRequest) GetPruneUpTo() string {
	if x != nil {
		return x.PruneUpTo
	}
	return ""
}

func (x *PruneRequest) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

func (x *PruneRequest) GetPruneAllDivulgedContracts() bool {
	if x != nil {
		return x.PruneAllDivulgedContracts
	}
	return false
}

type PruneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PruneResponse) Reset() {
	*x = PruneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PruneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PruneResponse) ProtoMessage() {}

func (x *PruneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PruneResponse.ProtoReflect.Descriptor instead.
func (*PruneResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescGZIP(), []int{1}
}

var File_com_daml_ledger_api_v1_admin_participant_pruning_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x94,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0b, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x5f, 0x75, 0x70, 0x5f, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x55, 0x70, 0x54, 0x6f, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x64, 0x69, 0x76, 0x75, 0x6c, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x70, 0x72, 0x75, 0x6e,
	0x65, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x76, 0x75, 0x6c, 0x67, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7d, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x05, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xaf, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x23, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x4b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xaa, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x2e, 0x44,
	0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescData = file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDesc
)

func file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDescData
}

var file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_goTypes = []interface{}{
	(*PruneRequest)(nil),  // 0: com.daml.ledger.api.v1.admin.PruneRequest
	(*PruneResponse)(nil), // 1: com.daml.ledger.api.v1.admin.PruneResponse
}
var file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_depIdxs = []int32{
	0, // 0: com.daml.ledger.api.v1.admin.ParticipantPruningService.Prune:input_type -> com.daml.ledger.api.v1.admin.PruneRequest
	1, // 1: com.daml.ledger.api.v1.admin.ParticipantPruningService.Prune:output_type -> com.daml.ledger.api.v1.admin.PruneResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_init() }
func file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_init() {
	if File_com_daml_ledger_api_v1_admin_participant_pruning_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneRequest); i {
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
		file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PruneResponse); i {
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
			RawDescriptor: file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_admin_participant_pruning_service_proto = out.File
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_goTypes = nil
	file_com_daml_ledger_api_v1_admin_participant_pruning_service_proto_depIdxs = nil
}
