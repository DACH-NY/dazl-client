// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/daml/ledger/api/v2/admin/command_inspection_service.proto

package admin

import (
	v2 "github.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandState int32

const (
	CommandState_COMMAND_STATE_UNSPECIFIED CommandState = 0
	CommandState_COMMAND_STATE_PENDING     CommandState = 1
	CommandState_COMMAND_STATE_SUCCEEDED   CommandState = 2
	CommandState_COMMAND_STATE_FAILED      CommandState = 3
)

// Enum value maps for CommandState.
var (
	CommandState_name = map[int32]string{
		0: "COMMAND_STATE_UNSPECIFIED",
		1: "COMMAND_STATE_PENDING",
		2: "COMMAND_STATE_SUCCEEDED",
		3: "COMMAND_STATE_FAILED",
	}
	CommandState_value = map[string]int32{
		"COMMAND_STATE_UNSPECIFIED": 0,
		"COMMAND_STATE_PENDING":     1,
		"COMMAND_STATE_SUCCEEDED":   2,
		"COMMAND_STATE_FAILED":      3,
	}
)

func (x CommandState) Enum() *CommandState {
	p := new(CommandState)
	*p = x
	return p
}

func (x CommandState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandState) Descriptor() protoreflect.EnumDescriptor {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_enumTypes[0].Descriptor()
}

func (CommandState) Type() protoreflect.EnumType {
	return &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_enumTypes[0]
}

func (x CommandState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandState.Descriptor instead.
func (CommandState) EnumDescriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{0}
}

type GetCommandStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandIdPrefix string       `protobuf:"bytes,1,opt,name=command_id_prefix,json=commandIdPrefix,proto3" json:"command_id_prefix,omitempty"`
	State           CommandState `protobuf:"varint,2,opt,name=state,proto3,enum=com.daml.ledger.api.v2.admin.CommandState" json:"state,omitempty"`
	Limit           uint32       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetCommandStatusRequest) Reset() {
	*x = GetCommandStatusRequest{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommandStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommandStatusRequest) ProtoMessage() {}

func (x *GetCommandStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommandStatusRequest.ProtoReflect.Descriptor instead.
func (*GetCommandStatusRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCommandStatusRequest) GetCommandIdPrefix() string {
	if x != nil {
		return x.CommandIdPrefix
	}
	return ""
}

func (x *GetCommandStatusRequest) GetState() CommandState {
	if x != nil {
		return x.State
	}
	return CommandState_COMMAND_STATE_UNSPECIFIED
}

func (x *GetCommandStatusRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCommandStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandStatus []*CommandStatus `protobuf:"bytes,1,rep,name=command_status,json=commandStatus,proto3" json:"command_status,omitempty"`
}

func (x *GetCommandStatusResponse) Reset() {
	*x = GetCommandStatusResponse{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommandStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommandStatusResponse) ProtoMessage() {}

func (x *GetCommandStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommandStatusResponse.ProtoReflect.Descriptor instead.
func (*GetCommandStatusResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCommandStatusResponse) GetCommandStatus() []*CommandStatus {
	if x != nil {
		return x.CommandStatus
	}
	return nil
}

type CommandStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Started           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=started,proto3" json:"started,omitempty"`
	Completed         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=completed,proto3" json:"completed,omitempty"`
	Completion        *v2.Completion         `protobuf:"bytes,3,opt,name=completion,proto3" json:"completion,omitempty"`
	State             CommandState           `protobuf:"varint,4,opt,name=state,proto3,enum=com.daml.ledger.api.v2.admin.CommandState" json:"state,omitempty"`
	Commands          []*v2.Command          `protobuf:"bytes,5,rep,name=commands,proto3" json:"commands,omitempty"`
	RequestStatistics *RequestStatistics     `protobuf:"bytes,6,opt,name=request_statistics,json=requestStatistics,proto3" json:"request_statistics,omitempty"`
	Updates           *CommandUpdates        `protobuf:"bytes,7,opt,name=updates,proto3" json:"updates,omitempty"`
}

func (x *CommandStatus) Reset() {
	*x = CommandStatus{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandStatus) ProtoMessage() {}

func (x *CommandStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandStatus.ProtoReflect.Descriptor instead.
func (*CommandStatus) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{2}
}

func (x *CommandStatus) GetStarted() *timestamppb.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

func (x *CommandStatus) GetCompleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Completed
	}
	return nil
}

func (x *CommandStatus) GetCompletion() *v2.Completion {
	if x != nil {
		return x.Completion
	}
	return nil
}

func (x *CommandStatus) GetState() CommandState {
	if x != nil {
		return x.State
	}
	return CommandState_COMMAND_STATE_UNSPECIFIED
}

func (x *CommandStatus) GetCommands() []*v2.Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *CommandStatus) GetRequestStatistics() *RequestStatistics {
	if x != nil {
		return x.RequestStatistics
	}
	return nil
}

func (x *CommandStatus) GetUpdates() *CommandUpdates {
	if x != nil {
		return x.Updates
	}
	return nil
}

type RequestStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelopes   uint32 `protobuf:"varint,1,opt,name=envelopes,proto3" json:"envelopes,omitempty"`
	RequestSize uint32 `protobuf:"varint,2,opt,name=request_size,json=requestSize,proto3" json:"request_size,omitempty"`
	Recipients  uint32 `protobuf:"varint,3,opt,name=recipients,proto3" json:"recipients,omitempty"`
}

func (x *RequestStatistics) Reset() {
	*x = RequestStatistics{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestStatistics) ProtoMessage() {}

func (x *RequestStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestStatistics.ProtoReflect.Descriptor instead.
func (*RequestStatistics) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{3}
}

func (x *RequestStatistics) GetEnvelopes() uint32 {
	if x != nil {
		return x.Envelopes
	}
	return 0
}

func (x *RequestStatistics) GetRequestSize() uint32 {
	if x != nil {
		return x.RequestSize
	}
	return 0
}

func (x *RequestStatistics) GetRecipients() uint32 {
	if x != nil {
		return x.Recipients
	}
	return 0
}

type CommandUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created       []*Contract `protobuf:"bytes,1,rep,name=created,proto3" json:"created,omitempty"`
	Archived      []*Contract `protobuf:"bytes,2,rep,name=archived,proto3" json:"archived,omitempty"`
	Exercised     uint32      `protobuf:"varint,3,opt,name=exercised,proto3" json:"exercised,omitempty"`
	Fetched       uint32      `protobuf:"varint,4,opt,name=fetched,proto3" json:"fetched,omitempty"`
	LookedUpByKey uint32      `protobuf:"varint,5,opt,name=looked_up_by_key,json=lookedUpByKey,proto3" json:"looked_up_by_key,omitempty"`
}

func (x *CommandUpdates) Reset() {
	*x = CommandUpdates{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandUpdates) ProtoMessage() {}

func (x *CommandUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandUpdates.ProtoReflect.Descriptor instead.
func (*CommandUpdates) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{4}
}

func (x *CommandUpdates) GetCreated() []*Contract {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *CommandUpdates) GetArchived() []*Contract {
	if x != nil {
		return x.Archived
	}
	return nil
}

func (x *CommandUpdates) GetExercised() uint32 {
	if x != nil {
		return x.Exercised
	}
	return 0
}

func (x *CommandUpdates) GetFetched() uint32 {
	if x != nil {
		return x.Fetched
	}
	return 0
}

func (x *CommandUpdates) GetLookedUpByKey() uint32 {
	if x != nil {
		return x.LookedUpByKey
	}
	return 0
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId  *v2.Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ContractId  string         `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	ContractKey *v2.Value      `protobuf:"bytes,3,opt,name=contract_key,json=contractKey,proto3" json:"contract_key,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP(), []int{5}
}

func (x *Contract) GetTemplateId() *v2.Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *Contract) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *Contract) GetContractKey() *v2.Value {
	if x != nil {
		return x.ContractKey
	}
	return nil
}

var File_com_daml_ledger_api_v2_admin_command_inspection_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x25, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x6e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d,
	0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0xea, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x5e, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x74, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xf7, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x08, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x6c, 0x6f, 0x6f, 0x6b, 0x65, 0x64,
	0x5f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x6c, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x55, 0x70, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x22,
	0xb2, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x43, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x4b, 0x65, 0x79, 0x2a, 0x7f, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0x9e, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xae, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x4b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xaa, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x2e, 0x44,
	0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescData = file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDesc
)

func file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDescData
}

var file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_goTypes = []any{
	(CommandState)(0),                // 0: com.daml.ledger.api.v2.admin.CommandState
	(*GetCommandStatusRequest)(nil),  // 1: com.daml.ledger.api.v2.admin.GetCommandStatusRequest
	(*GetCommandStatusResponse)(nil), // 2: com.daml.ledger.api.v2.admin.GetCommandStatusResponse
	(*CommandStatus)(nil),            // 3: com.daml.ledger.api.v2.admin.CommandStatus
	(*RequestStatistics)(nil),        // 4: com.daml.ledger.api.v2.admin.RequestStatistics
	(*CommandUpdates)(nil),           // 5: com.daml.ledger.api.v2.admin.CommandUpdates
	(*Contract)(nil),                 // 6: com.daml.ledger.api.v2.admin.Contract
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*v2.Completion)(nil),            // 8: com.daml.ledger.api.v2.Completion
	(*v2.Command)(nil),               // 9: com.daml.ledger.api.v2.Command
	(*v2.Identifier)(nil),            // 10: com.daml.ledger.api.v2.Identifier
	(*v2.Value)(nil),                 // 11: com.daml.ledger.api.v2.Value
}
var file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_depIdxs = []int32{
	0,  // 0: com.daml.ledger.api.v2.admin.GetCommandStatusRequest.state:type_name -> com.daml.ledger.api.v2.admin.CommandState
	3,  // 1: com.daml.ledger.api.v2.admin.GetCommandStatusResponse.command_status:type_name -> com.daml.ledger.api.v2.admin.CommandStatus
	7,  // 2: com.daml.ledger.api.v2.admin.CommandStatus.started:type_name -> google.protobuf.Timestamp
	7,  // 3: com.daml.ledger.api.v2.admin.CommandStatus.completed:type_name -> google.protobuf.Timestamp
	8,  // 4: com.daml.ledger.api.v2.admin.CommandStatus.completion:type_name -> com.daml.ledger.api.v2.Completion
	0,  // 5: com.daml.ledger.api.v2.admin.CommandStatus.state:type_name -> com.daml.ledger.api.v2.admin.CommandState
	9,  // 6: com.daml.ledger.api.v2.admin.CommandStatus.commands:type_name -> com.daml.ledger.api.v2.Command
	4,  // 7: com.daml.ledger.api.v2.admin.CommandStatus.request_statistics:type_name -> com.daml.ledger.api.v2.admin.RequestStatistics
	5,  // 8: com.daml.ledger.api.v2.admin.CommandStatus.updates:type_name -> com.daml.ledger.api.v2.admin.CommandUpdates
	6,  // 9: com.daml.ledger.api.v2.admin.CommandUpdates.created:type_name -> com.daml.ledger.api.v2.admin.Contract
	6,  // 10: com.daml.ledger.api.v2.admin.CommandUpdates.archived:type_name -> com.daml.ledger.api.v2.admin.Contract
	10, // 11: com.daml.ledger.api.v2.admin.Contract.template_id:type_name -> com.daml.ledger.api.v2.Identifier
	11, // 12: com.daml.ledger.api.v2.admin.Contract.contract_key:type_name -> com.daml.ledger.api.v2.Value
	1,  // 13: com.daml.ledger.api.v2.admin.CommandInspectionService.GetCommandStatus:input_type -> com.daml.ledger.api.v2.admin.GetCommandStatusRequest
	2,  // 14: com.daml.ledger.api.v2.admin.CommandInspectionService.GetCommandStatus:output_type -> com.daml.ledger.api.v2.admin.GetCommandStatusResponse
	14, // [14:15] is the sub-list for method output_type
	13, // [13:14] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_init() }
func file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_init() {
	if File_com_daml_ledger_api_v2_admin_command_inspection_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_depIdxs,
		EnumInfos:         file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_enumTypes,
		MessageInfos:      file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v2_admin_command_inspection_service_proto = out.File
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_goTypes = nil
	file_com_daml_ledger_api_v2_admin_command_inspection_service_proto_depIdxs = nil
}