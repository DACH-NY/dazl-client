// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/daml/ledger/api/v2/command_completion_service.proto

package v2

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

type CompletionStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId  string   `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Parties        []string `protobuf:"bytes,2,rep,name=parties,proto3" json:"parties,omitempty"`
	BeginExclusive int64    `protobuf:"varint,3,opt,name=begin_exclusive,json=beginExclusive,proto3" json:"begin_exclusive,omitempty"`
}

func (x *CompletionStreamRequest) Reset() {
	*x = CompletionStreamRequest{}
	mi := &file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletionStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionStreamRequest) ProtoMessage() {}

func (x *CompletionStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionStreamRequest.ProtoReflect.Descriptor instead.
func (*CompletionStreamRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompletionStreamRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *CompletionStreamRequest) GetParties() []string {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *CompletionStreamRequest) GetBeginExclusive() int64 {
	if x != nil {
		return x.BeginExclusive
	}
	return 0
}

type CompletionStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CompletionResponse:
	//
	//	*CompletionStreamResponse_Completion
	//	*CompletionStreamResponse_OffsetCheckpoint
	CompletionResponse isCompletionStreamResponse_CompletionResponse `protobuf_oneof:"completion_response"`
}

func (x *CompletionStreamResponse) Reset() {
	*x = CompletionStreamResponse{}
	mi := &file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletionStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionStreamResponse) ProtoMessage() {}

func (x *CompletionStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionStreamResponse.ProtoReflect.Descriptor instead.
func (*CompletionStreamResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescGZIP(), []int{1}
}

func (m *CompletionStreamResponse) GetCompletionResponse() isCompletionStreamResponse_CompletionResponse {
	if m != nil {
		return m.CompletionResponse
	}
	return nil
}

func (x *CompletionStreamResponse) GetCompletion() *Completion {
	if x, ok := x.GetCompletionResponse().(*CompletionStreamResponse_Completion); ok {
		return x.Completion
	}
	return nil
}

func (x *CompletionStreamResponse) GetOffsetCheckpoint() *OffsetCheckpoint {
	if x, ok := x.GetCompletionResponse().(*CompletionStreamResponse_OffsetCheckpoint); ok {
		return x.OffsetCheckpoint
	}
	return nil
}

type isCompletionStreamResponse_CompletionResponse interface {
	isCompletionStreamResponse_CompletionResponse()
}

type CompletionStreamResponse_Completion struct {
	Completion *Completion `protobuf:"bytes,1,opt,name=completion,proto3,oneof"`
}

type CompletionStreamResponse_OffsetCheckpoint struct {
	OffsetCheckpoint *OffsetCheckpoint `protobuf:"bytes,2,opt,name=offset_checkpoint,json=offsetCheckpoint,proto3,oneof"`
}

func (*CompletionStreamResponse_Completion) isCompletionStreamResponse_CompletionResponse() {}

func (*CompletionStreamResponse_OffsetCheckpoint) isCompletionStreamResponse_CompletionResponse() {}

var File_com_daml_ledger_api_v2_command_completion_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v2_command_completion_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65,
	0x22, 0xd0, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x11, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x15, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x93, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x77, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x9c, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x42, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75,
	0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d,
	0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0xaa,
	0x02, 0x16, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescData = file_com_daml_ledger_api_v2_command_completion_service_proto_rawDesc
)

func file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v2_command_completion_service_proto_rawDescData
}

var file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_daml_ledger_api_v2_command_completion_service_proto_goTypes = []any{
	(*CompletionStreamRequest)(nil),  // 0: com.daml.ledger.api.v2.CompletionStreamRequest
	(*CompletionStreamResponse)(nil), // 1: com.daml.ledger.api.v2.CompletionStreamResponse
	(*Completion)(nil),               // 2: com.daml.ledger.api.v2.Completion
	(*OffsetCheckpoint)(nil),         // 3: com.daml.ledger.api.v2.OffsetCheckpoint
}
var file_com_daml_ledger_api_v2_command_completion_service_proto_depIdxs = []int32{
	2, // 0: com.daml.ledger.api.v2.CompletionStreamResponse.completion:type_name -> com.daml.ledger.api.v2.Completion
	3, // 1: com.daml.ledger.api.v2.CompletionStreamResponse.offset_checkpoint:type_name -> com.daml.ledger.api.v2.OffsetCheckpoint
	0, // 2: com.daml.ledger.api.v2.CommandCompletionService.CompletionStream:input_type -> com.daml.ledger.api.v2.CompletionStreamRequest
	1, // 3: com.daml.ledger.api.v2.CommandCompletionService.CompletionStream:output_type -> com.daml.ledger.api.v2.CompletionStreamResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v2_command_completion_service_proto_init() }
func file_com_daml_ledger_api_v2_command_completion_service_proto_init() {
	if File_com_daml_ledger_api_v2_command_completion_service_proto != nil {
		return
	}
	file_com_daml_ledger_api_v2_completion_proto_init()
	file_com_daml_ledger_api_v2_offset_checkpoint_proto_init()
	file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes[1].OneofWrappers = []any{
		(*CompletionStreamResponse_Completion)(nil),
		(*CompletionStreamResponse_OffsetCheckpoint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v2_command_completion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v2_command_completion_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v2_command_completion_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v2_command_completion_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v2_command_completion_service_proto = out.File
	file_com_daml_ledger_api_v2_command_completion_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v2_command_completion_service_proto_goTypes = nil
	file_com_daml_ledger_api_v2_command_completion_service_proto_depIdxs = nil
}
