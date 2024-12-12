// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/daml/ledger/api/v2/update_service.proto

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

type GetUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeginExclusive int64              `protobuf:"varint,1,opt,name=begin_exclusive,json=beginExclusive,proto3" json:"begin_exclusive,omitempty"`
	EndInclusive   *int64             `protobuf:"varint,2,opt,name=end_inclusive,json=endInclusive,proto3,oneof" json:"end_inclusive,omitempty"`
	Filter         *TransactionFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Verbose        bool               `protobuf:"varint,4,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *GetUpdatesRequest) Reset() {
	*x = GetUpdatesRequest{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatesRequest) ProtoMessage() {}

func (x *GetUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatesRequest.ProtoReflect.Descriptor instead.
func (*GetUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUpdatesRequest) GetBeginExclusive() int64 {
	if x != nil {
		return x.BeginExclusive
	}
	return 0
}

func (x *GetUpdatesRequest) GetEndInclusive() int64 {
	if x != nil && x.EndInclusive != nil {
		return *x.EndInclusive
	}
	return 0
}

func (x *GetUpdatesRequest) GetFilter() *TransactionFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *GetUpdatesRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

type GetUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Update:
	//	*GetUpdatesResponse_Transaction
	//	*GetUpdatesResponse_Reassignment
	//	*GetUpdatesResponse_OffsetCheckpoint
	Update isGetUpdatesResponse_Update `protobuf_oneof:"update"`
}

func (x *GetUpdatesResponse) Reset() {
	*x = GetUpdatesResponse{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatesResponse) ProtoMessage() {}

func (x *GetUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatesResponse.ProtoReflect.Descriptor instead.
func (*GetUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{1}
}

func (m *GetUpdatesResponse) GetUpdate() isGetUpdatesResponse_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *GetUpdatesResponse) GetTransaction() *Transaction {
	if x, ok := x.GetUpdate().(*GetUpdatesResponse_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (x *GetUpdatesResponse) GetReassignment() *Reassignment {
	if x, ok := x.GetUpdate().(*GetUpdatesResponse_Reassignment); ok {
		return x.Reassignment
	}
	return nil
}

func (x *GetUpdatesResponse) GetOffsetCheckpoint() *OffsetCheckpoint {
	if x, ok := x.GetUpdate().(*GetUpdatesResponse_OffsetCheckpoint); ok {
		return x.OffsetCheckpoint
	}
	return nil
}

type isGetUpdatesResponse_Update interface {
	isGetUpdatesResponse_Update()
}

type GetUpdatesResponse_Transaction struct {
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3,oneof"`
}

type GetUpdatesResponse_Reassignment struct {
	Reassignment *Reassignment `protobuf:"bytes,2,opt,name=reassignment,proto3,oneof"`
}

type GetUpdatesResponse_OffsetCheckpoint struct {
	OffsetCheckpoint *OffsetCheckpoint `protobuf:"bytes,3,opt,name=offset_checkpoint,json=offsetCheckpoint,proto3,oneof"`
}

func (*GetUpdatesResponse_Transaction) isGetUpdatesResponse_Update() {}

func (*GetUpdatesResponse_Reassignment) isGetUpdatesResponse_Update() {}

func (*GetUpdatesResponse_OffsetCheckpoint) isGetUpdatesResponse_Update() {}

type GetUpdateTreesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Update:
	//	*GetUpdateTreesResponse_TransactionTree
	//	*GetUpdateTreesResponse_Reassignment
	//	*GetUpdateTreesResponse_OffsetCheckpoint
	Update isGetUpdateTreesResponse_Update `protobuf_oneof:"update"`
}

func (x *GetUpdateTreesResponse) Reset() {
	*x = GetUpdateTreesResponse{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUpdateTreesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdateTreesResponse) ProtoMessage() {}

func (x *GetUpdateTreesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdateTreesResponse.ProtoReflect.Descriptor instead.
func (*GetUpdateTreesResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{2}
}

func (m *GetUpdateTreesResponse) GetUpdate() isGetUpdateTreesResponse_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *GetUpdateTreesResponse) GetTransactionTree() *TransactionTree {
	if x, ok := x.GetUpdate().(*GetUpdateTreesResponse_TransactionTree); ok {
		return x.TransactionTree
	}
	return nil
}

func (x *GetUpdateTreesResponse) GetReassignment() *Reassignment {
	if x, ok := x.GetUpdate().(*GetUpdateTreesResponse_Reassignment); ok {
		return x.Reassignment
	}
	return nil
}

func (x *GetUpdateTreesResponse) GetOffsetCheckpoint() *OffsetCheckpoint {
	if x, ok := x.GetUpdate().(*GetUpdateTreesResponse_OffsetCheckpoint); ok {
		return x.OffsetCheckpoint
	}
	return nil
}

type isGetUpdateTreesResponse_Update interface {
	isGetUpdateTreesResponse_Update()
}

type GetUpdateTreesResponse_TransactionTree struct {
	TransactionTree *TransactionTree `protobuf:"bytes,1,opt,name=transaction_tree,json=transactionTree,proto3,oneof"`
}

type GetUpdateTreesResponse_Reassignment struct {
	Reassignment *Reassignment `protobuf:"bytes,2,opt,name=reassignment,proto3,oneof"`
}

type GetUpdateTreesResponse_OffsetCheckpoint struct {
	OffsetCheckpoint *OffsetCheckpoint `protobuf:"bytes,3,opt,name=offset_checkpoint,json=offsetCheckpoint,proto3,oneof"`
}

func (*GetUpdateTreesResponse_TransactionTree) isGetUpdateTreesResponse_Update() {}

func (*GetUpdateTreesResponse_Reassignment) isGetUpdateTreesResponse_Update() {}

func (*GetUpdateTreesResponse_OffsetCheckpoint) isGetUpdateTreesResponse_Update() {}

type GetTransactionByEventIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId           string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	RequestingParties []string `protobuf:"bytes,2,rep,name=requesting_parties,json=requestingParties,proto3" json:"requesting_parties,omitempty"`
}

func (x *GetTransactionByEventIdRequest) Reset() {
	*x = GetTransactionByEventIdRequest{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionByEventIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionByEventIdRequest) ProtoMessage() {}

func (x *GetTransactionByEventIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionByEventIdRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionByEventIdRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionByEventIdRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *GetTransactionByEventIdRequest) GetRequestingParties() []string {
	if x != nil {
		return x.RequestingParties
	}
	return nil
}

type GetTransactionByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateId          string   `protobuf:"bytes,1,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	RequestingParties []string `protobuf:"bytes,2,rep,name=requesting_parties,json=requestingParties,proto3" json:"requesting_parties,omitempty"`
}

func (x *GetTransactionByIdRequest) Reset() {
	*x = GetTransactionByIdRequest{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionByIdRequest) ProtoMessage() {}

func (x *GetTransactionByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionByIdRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionByIdRequest) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

func (x *GetTransactionByIdRequest) GetRequestingParties() []string {
	if x != nil {
		return x.RequestingParties
	}
	return nil
}

type GetTransactionTreeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *TransactionTree `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTransactionTreeResponse) Reset() {
	*x = GetTransactionTreeResponse{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionTreeResponse) ProtoMessage() {}

func (x *GetTransactionTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionTreeResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionTreeResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTransactionTreeResponse) GetTransaction() *TransactionTree {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v2_update_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_com_daml_ledger_api_v2_update_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v2_update_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x6e,
	0x64, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x57, 0x0a, 0x11, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x65,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65,
	0x48, 0x00, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x72, 0x65, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x57, 0x0a, 0x11, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x6a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x67,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x72, 0x65, 0x65, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xef, 0x05, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6d, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x65, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x65, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x89, 0x01, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65,
	0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x91, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c,
	0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0xaa, 0x02,
	0x16, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v2_update_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v2_update_service_proto_rawDescData = file_com_daml_ledger_api_v2_update_service_proto_rawDesc
)

func file_com_daml_ledger_api_v2_update_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v2_update_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v2_update_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v2_update_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v2_update_service_proto_rawDescData
}

var file_com_daml_ledger_api_v2_update_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_com_daml_ledger_api_v2_update_service_proto_goTypes = []any{
	(*GetUpdatesRequest)(nil),              // 0: com.daml.ledger.api.v2.GetUpdatesRequest
	(*GetUpdatesResponse)(nil),             // 1: com.daml.ledger.api.v2.GetUpdatesResponse
	(*GetUpdateTreesResponse)(nil),         // 2: com.daml.ledger.api.v2.GetUpdateTreesResponse
	(*GetTransactionByEventIdRequest)(nil), // 3: com.daml.ledger.api.v2.GetTransactionByEventIdRequest
	(*GetTransactionByIdRequest)(nil),      // 4: com.daml.ledger.api.v2.GetTransactionByIdRequest
	(*GetTransactionTreeResponse)(nil),     // 5: com.daml.ledger.api.v2.GetTransactionTreeResponse
	(*GetTransactionResponse)(nil),         // 6: com.daml.ledger.api.v2.GetTransactionResponse
	(*TransactionFilter)(nil),              // 7: com.daml.ledger.api.v2.TransactionFilter
	(*Transaction)(nil),                    // 8: com.daml.ledger.api.v2.Transaction
	(*Reassignment)(nil),                   // 9: com.daml.ledger.api.v2.Reassignment
	(*OffsetCheckpoint)(nil),               // 10: com.daml.ledger.api.v2.OffsetCheckpoint
	(*TransactionTree)(nil),                // 11: com.daml.ledger.api.v2.TransactionTree
}
var file_com_daml_ledger_api_v2_update_service_proto_depIdxs = []int32{
	7,  // 0: com.daml.ledger.api.v2.GetUpdatesRequest.filter:type_name -> com.daml.ledger.api.v2.TransactionFilter
	8,  // 1: com.daml.ledger.api.v2.GetUpdatesResponse.transaction:type_name -> com.daml.ledger.api.v2.Transaction
	9,  // 2: com.daml.ledger.api.v2.GetUpdatesResponse.reassignment:type_name -> com.daml.ledger.api.v2.Reassignment
	10, // 3: com.daml.ledger.api.v2.GetUpdatesResponse.offset_checkpoint:type_name -> com.daml.ledger.api.v2.OffsetCheckpoint
	11, // 4: com.daml.ledger.api.v2.GetUpdateTreesResponse.transaction_tree:type_name -> com.daml.ledger.api.v2.TransactionTree
	9,  // 5: com.daml.ledger.api.v2.GetUpdateTreesResponse.reassignment:type_name -> com.daml.ledger.api.v2.Reassignment
	10, // 6: com.daml.ledger.api.v2.GetUpdateTreesResponse.offset_checkpoint:type_name -> com.daml.ledger.api.v2.OffsetCheckpoint
	11, // 7: com.daml.ledger.api.v2.GetTransactionTreeResponse.transaction:type_name -> com.daml.ledger.api.v2.TransactionTree
	8,  // 8: com.daml.ledger.api.v2.GetTransactionResponse.transaction:type_name -> com.daml.ledger.api.v2.Transaction
	0,  // 9: com.daml.ledger.api.v2.UpdateService.GetUpdates:input_type -> com.daml.ledger.api.v2.GetUpdatesRequest
	0,  // 10: com.daml.ledger.api.v2.UpdateService.GetUpdateTrees:input_type -> com.daml.ledger.api.v2.GetUpdatesRequest
	3,  // 11: com.daml.ledger.api.v2.UpdateService.GetTransactionTreeByEventId:input_type -> com.daml.ledger.api.v2.GetTransactionByEventIdRequest
	4,  // 12: com.daml.ledger.api.v2.UpdateService.GetTransactionTreeById:input_type -> com.daml.ledger.api.v2.GetTransactionByIdRequest
	3,  // 13: com.daml.ledger.api.v2.UpdateService.GetTransactionByEventId:input_type -> com.daml.ledger.api.v2.GetTransactionByEventIdRequest
	4,  // 14: com.daml.ledger.api.v2.UpdateService.GetTransactionById:input_type -> com.daml.ledger.api.v2.GetTransactionByIdRequest
	1,  // 15: com.daml.ledger.api.v2.UpdateService.GetUpdates:output_type -> com.daml.ledger.api.v2.GetUpdatesResponse
	2,  // 16: com.daml.ledger.api.v2.UpdateService.GetUpdateTrees:output_type -> com.daml.ledger.api.v2.GetUpdateTreesResponse
	5,  // 17: com.daml.ledger.api.v2.UpdateService.GetTransactionTreeByEventId:output_type -> com.daml.ledger.api.v2.GetTransactionTreeResponse
	5,  // 18: com.daml.ledger.api.v2.UpdateService.GetTransactionTreeById:output_type -> com.daml.ledger.api.v2.GetTransactionTreeResponse
	6,  // 19: com.daml.ledger.api.v2.UpdateService.GetTransactionByEventId:output_type -> com.daml.ledger.api.v2.GetTransactionResponse
	6,  // 20: com.daml.ledger.api.v2.UpdateService.GetTransactionById:output_type -> com.daml.ledger.api.v2.GetTransactionResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v2_update_service_proto_init() }
func file_com_daml_ledger_api_v2_update_service_proto_init() {
	if File_com_daml_ledger_api_v2_update_service_proto != nil {
		return
	}
	file_com_daml_ledger_api_v2_offset_checkpoint_proto_init()
	file_com_daml_ledger_api_v2_reassignment_proto_init()
	file_com_daml_ledger_api_v2_transaction_proto_init()
	file_com_daml_ledger_api_v2_transaction_filter_proto_init()
	file_com_daml_ledger_api_v2_update_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_com_daml_ledger_api_v2_update_service_proto_msgTypes[1].OneofWrappers = []any{
		(*GetUpdatesResponse_Transaction)(nil),
		(*GetUpdatesResponse_Reassignment)(nil),
		(*GetUpdatesResponse_OffsetCheckpoint)(nil),
	}
	file_com_daml_ledger_api_v2_update_service_proto_msgTypes[2].OneofWrappers = []any{
		(*GetUpdateTreesResponse_TransactionTree)(nil),
		(*GetUpdateTreesResponse_Reassignment)(nil),
		(*GetUpdateTreesResponse_OffsetCheckpoint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v2_update_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v2_update_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v2_update_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v2_update_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v2_update_service_proto = out.File
	file_com_daml_ledger_api_v2_update_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v2_update_service_proto_goTypes = nil
	file_com_daml_ledger_api_v2_update_service_proto_depIdxs = nil
}
