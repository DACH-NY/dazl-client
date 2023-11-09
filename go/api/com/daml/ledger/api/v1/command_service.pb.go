// Copyright (c) 2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/daml/ledger/api/v1/command_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// These commands are atomic, and will become transactions.
type SubmitAndWaitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The commands to be submitted.
	// Required
	Commands *Commands `protobuf:"bytes,1,opt,name=commands,proto3" json:"commands,omitempty"`
}

func (x *SubmitAndWaitRequest) Reset() {
	*x = SubmitAndWaitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitAndWaitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAndWaitRequest) ProtoMessage() {}

func (x *SubmitAndWaitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAndWaitRequest.ProtoReflect.Descriptor instead.
func (*SubmitAndWaitRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_command_service_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitAndWaitRequest) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type SubmitAndWaitForTransactionIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the transaction that resulted from the submitted command.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// The format of this field is described in ``ledger_offset.proto``.
	// Optional
	CompletionOffset string `protobuf:"bytes,2,opt,name=completion_offset,json=completionOffset,proto3" json:"completion_offset,omitempty"`
}

func (x *SubmitAndWaitForTransactionIdResponse) Reset() {
	*x = SubmitAndWaitForTransactionIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitAndWaitForTransactionIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAndWaitForTransactionIdResponse) ProtoMessage() {}

func (x *SubmitAndWaitForTransactionIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAndWaitForTransactionIdResponse.ProtoReflect.Descriptor instead.
func (*SubmitAndWaitForTransactionIdResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_command_service_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitAndWaitForTransactionIdResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *SubmitAndWaitForTransactionIdResponse) GetCompletionOffset() string {
	if x != nil {
		return x.CompletionOffset
	}
	return ""
}

type SubmitAndWaitForTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The flat transaction that resulted from the submitted command.
	// Required
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// The format of this field is described in ``ledger_offset.proto``.
	// Optional
	CompletionOffset string `protobuf:"bytes,2,opt,name=completion_offset,json=completionOffset,proto3" json:"completion_offset,omitempty"`
}

func (x *SubmitAndWaitForTransactionResponse) Reset() {
	*x = SubmitAndWaitForTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitAndWaitForTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAndWaitForTransactionResponse) ProtoMessage() {}

func (x *SubmitAndWaitForTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAndWaitForTransactionResponse.ProtoReflect.Descriptor instead.
func (*SubmitAndWaitForTransactionResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_command_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitAndWaitForTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *SubmitAndWaitForTransactionResponse) GetCompletionOffset() string {
	if x != nil {
		return x.CompletionOffset
	}
	return ""
}

type SubmitAndWaitForTransactionTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction tree that resulted from the submitted command.
	// Required
	Transaction *TransactionTree `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// The format of this field is described in ``ledger_offset.proto``.
	// Optional
	CompletionOffset string `protobuf:"bytes,2,opt,name=completion_offset,json=completionOffset,proto3" json:"completion_offset,omitempty"`
}

func (x *SubmitAndWaitForTransactionTreeResponse) Reset() {
	*x = SubmitAndWaitForTransactionTreeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitAndWaitForTransactionTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitAndWaitForTransactionTreeResponse) ProtoMessage() {}

func (x *SubmitAndWaitForTransactionTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_command_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitAndWaitForTransactionTreeResponse.ProtoReflect.Descriptor instead.
func (*SubmitAndWaitForTransactionTreeResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_command_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitAndWaitForTransactionTreeResponse) GetTransaction() *TransactionTree {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *SubmitAndWaitForTransactionTreeResponse) GetCompletionOffset() string {
	if x != nil {
		return x.CompletionOffset
	}
	return ""
}

var File_com_daml_ledger_api_v1_command_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_command_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c,
	0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x64, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x7b, 0x0a, 0x25, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x23, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x27, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0x94, 0x04, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x8c, 0x01, 0x0a, 0x1d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57,
	0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x88, 0x01, 0x0a, 0x1b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61,
	0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x1f,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41,
	0x6e, 0x64, 0x57, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x64,
	0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x92,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61,
	0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x6d,
	0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_command_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_command_service_proto_rawDescData = file_com_daml_ledger_api_v1_command_service_proto_rawDesc
)

func file_com_daml_ledger_api_v1_command_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_command_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_command_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_command_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_command_service_proto_rawDescData
}

var file_com_daml_ledger_api_v1_command_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_daml_ledger_api_v1_command_service_proto_goTypes = []interface{}{
	(*SubmitAndWaitRequest)(nil),                    // 0: com.daml.ledger.api.v1.SubmitAndWaitRequest
	(*SubmitAndWaitForTransactionIdResponse)(nil),   // 1: com.daml.ledger.api.v1.SubmitAndWaitForTransactionIdResponse
	(*SubmitAndWaitForTransactionResponse)(nil),     // 2: com.daml.ledger.api.v1.SubmitAndWaitForTransactionResponse
	(*SubmitAndWaitForTransactionTreeResponse)(nil), // 3: com.daml.ledger.api.v1.SubmitAndWaitForTransactionTreeResponse
	(*Commands)(nil),                                // 4: com.daml.ledger.api.v1.Commands
	(*Transaction)(nil),                             // 5: com.daml.ledger.api.v1.Transaction
	(*TransactionTree)(nil),                         // 6: com.daml.ledger.api.v1.TransactionTree
	(*emptypb.Empty)(nil),                           // 7: google.protobuf.Empty
}
var file_com_daml_ledger_api_v1_command_service_proto_depIdxs = []int32{
	4, // 0: com.daml.ledger.api.v1.SubmitAndWaitRequest.commands:type_name -> com.daml.ledger.api.v1.Commands
	5, // 1: com.daml.ledger.api.v1.SubmitAndWaitForTransactionResponse.transaction:type_name -> com.daml.ledger.api.v1.Transaction
	6, // 2: com.daml.ledger.api.v1.SubmitAndWaitForTransactionTreeResponse.transaction:type_name -> com.daml.ledger.api.v1.TransactionTree
	0, // 3: com.daml.ledger.api.v1.CommandService.SubmitAndWait:input_type -> com.daml.ledger.api.v1.SubmitAndWaitRequest
	0, // 4: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransactionId:input_type -> com.daml.ledger.api.v1.SubmitAndWaitRequest
	0, // 5: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransaction:input_type -> com.daml.ledger.api.v1.SubmitAndWaitRequest
	0, // 6: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransactionTree:input_type -> com.daml.ledger.api.v1.SubmitAndWaitRequest
	7, // 7: com.daml.ledger.api.v1.CommandService.SubmitAndWait:output_type -> google.protobuf.Empty
	1, // 8: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransactionId:output_type -> com.daml.ledger.api.v1.SubmitAndWaitForTransactionIdResponse
	2, // 9: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransaction:output_type -> com.daml.ledger.api.v1.SubmitAndWaitForTransactionResponse
	3, // 10: com.daml.ledger.api.v1.CommandService.SubmitAndWaitForTransactionTree:output_type -> com.daml.ledger.api.v1.SubmitAndWaitForTransactionTreeResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_command_service_proto_init() }
func file_com_daml_ledger_api_v1_command_service_proto_init() {
	if File_com_daml_ledger_api_v1_command_service_proto != nil {
		return
	}
	file_com_daml_ledger_api_v1_commands_proto_init()
	file_com_daml_ledger_api_v1_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_command_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitAndWaitRequest); i {
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
		file_com_daml_ledger_api_v1_command_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitAndWaitForTransactionIdResponse); i {
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
		file_com_daml_ledger_api_v1_command_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitAndWaitForTransactionResponse); i {
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
		file_com_daml_ledger_api_v1_command_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitAndWaitForTransactionTreeResponse); i {
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
			RawDescriptor: file_com_daml_ledger_api_v1_command_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v1_command_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_command_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_command_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_command_service_proto = out.File
	file_com_daml_ledger_api_v1_command_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_command_service_proto_goTypes = nil
	file_com_daml_ledger_api_v1_command_service_proto_depIdxs = nil
}
