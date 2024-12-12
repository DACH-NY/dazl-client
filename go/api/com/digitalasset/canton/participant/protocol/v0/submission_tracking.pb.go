// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/participant/protocol/v0/submission_tracking.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubmissionTrackingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Tracking:
	//
	//	*SubmissionTrackingData_Transaction
	Tracking isSubmissionTrackingData_Tracking `protobuf_oneof:"tracking"`
}

func (x *SubmissionTrackingData) Reset() {
	*x = SubmissionTrackingData{}
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmissionTrackingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionTrackingData) ProtoMessage() {}

func (x *SubmissionTrackingData) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionTrackingData.ProtoReflect.Descriptor instead.
func (*SubmissionTrackingData) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescGZIP(), []int{0}
}

func (m *SubmissionTrackingData) GetTracking() isSubmissionTrackingData_Tracking {
	if m != nil {
		return m.Tracking
	}
	return nil
}

func (x *SubmissionTrackingData) GetTransaction() *TransactionSubmissionTrackingData {
	if x, ok := x.GetTracking().(*SubmissionTrackingData_Transaction); ok {
		return x.Transaction
	}
	return nil
}

type isSubmissionTrackingData_Tracking interface {
	isSubmissionTrackingData_Tracking()
}

type SubmissionTrackingData_Transaction struct {
	Transaction *TransactionSubmissionTrackingData `protobuf:"bytes,1,opt,name=transaction,proto3,oneof"`
}

func (*SubmissionTrackingData_Transaction) isSubmissionTrackingData_Tracking() {}

type TransactionSubmissionTrackingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompletionInfo *CompletionInfo                                   `protobuf:"bytes,1,opt,name=completion_info,json=completionInfo,proto3" json:"completion_info,omitempty"`
	RejectionCause *TransactionSubmissionTrackingData_RejectionCause `protobuf:"bytes,2,opt,name=rejection_cause,json=rejectionCause,proto3" json:"rejection_cause,omitempty"`
	DomainId       *wrapperspb.StringValue                           `protobuf:"bytes,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *TransactionSubmissionTrackingData) Reset() {
	*x = TransactionSubmissionTrackingData{}
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionSubmissionTrackingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSubmissionTrackingData) ProtoMessage() {}

func (x *TransactionSubmissionTrackingData) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSubmissionTrackingData.ProtoReflect.Descriptor instead.
func (*TransactionSubmissionTrackingData) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionSubmissionTrackingData) GetCompletionInfo() *CompletionInfo {
	if x != nil {
		return x.CompletionInfo
	}
	return nil
}

func (x *TransactionSubmissionTrackingData) GetRejectionCause() *TransactionSubmissionTrackingData_RejectionCause {
	if x != nil {
		return x.RejectionCause
	}
	return nil
}

func (x *TransactionSubmissionTrackingData) GetDomainId() *wrapperspb.StringValue {
	if x != nil {
		return x.DomainId
	}
	return nil
}

type TransactionSubmissionTrackingData_RejectionCause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Cause:
	//
	//	*TransactionSubmissionTrackingData_RejectionCause_Timeout
	//	*TransactionSubmissionTrackingData_RejectionCause_RejectionReasonTemplate
	Cause isTransactionSubmissionTrackingData_RejectionCause_Cause `protobuf_oneof:"cause"`
}

func (x *TransactionSubmissionTrackingData_RejectionCause) Reset() {
	*x = TransactionSubmissionTrackingData_RejectionCause{}
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionSubmissionTrackingData_RejectionCause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSubmissionTrackingData_RejectionCause) ProtoMessage() {}

func (x *TransactionSubmissionTrackingData_RejectionCause) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSubmissionTrackingData_RejectionCause.ProtoReflect.Descriptor instead.
func (*TransactionSubmissionTrackingData_RejectionCause) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescGZIP(), []int{1, 0}
}

func (m *TransactionSubmissionTrackingData_RejectionCause) GetCause() isTransactionSubmissionTrackingData_RejectionCause_Cause {
	if m != nil {
		return m.Cause
	}
	return nil
}

func (x *TransactionSubmissionTrackingData_RejectionCause) GetTimeout() *emptypb.Empty {
	if x, ok := x.GetCause().(*TransactionSubmissionTrackingData_RejectionCause_Timeout); ok {
		return x.Timeout
	}
	return nil
}

func (x *TransactionSubmissionTrackingData_RejectionCause) GetRejectionReasonTemplate() *CommandRejected_GrpcRejectionReasonTemplate {
	if x, ok := x.GetCause().(*TransactionSubmissionTrackingData_RejectionCause_RejectionReasonTemplate); ok {
		return x.RejectionReasonTemplate
	}
	return nil
}

type isTransactionSubmissionTrackingData_RejectionCause_Cause interface {
	isTransactionSubmissionTrackingData_RejectionCause_Cause()
}

type TransactionSubmissionTrackingData_RejectionCause_Timeout struct {
	Timeout *emptypb.Empty `protobuf:"bytes,1,opt,name=timeout,proto3,oneof"`
}

type TransactionSubmissionTrackingData_RejectionCause_RejectionReasonTemplate struct {
	RejectionReasonTemplate *CommandRejected_GrpcRejectionReasonTemplate `protobuf:"bytes,2,opt,name=rejection_reason_template,json=rejectionReasonTemplate,proto3,oneof"`
}

func (*TransactionSubmissionTrackingData_RejectionCause_Timeout) isTransactionSubmissionTrackingData_RejectionCause_Cause() {
}

func (*TransactionSubmissionTrackingData_RejectionCause_RejectionReasonTemplate) isTransactionSubmissionTrackingData_RejectionCause_Cause() {
}

var File_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDesc = []byte{
	0x0a, 0x49, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x30, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x1a, 0x47, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x76, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x52, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x22, 0xc2, 0x04, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x68, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x0e,
	0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x1a, 0xea, 0x01, 0x0a, 0x0e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x9a, 0x01, 0x0a, 0x19, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x5c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x48, 0x00, 0x52, 0x17, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescData = file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDesc
)

func file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescData)
	})
	return file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDescData
}

var file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_goTypes = []any{
	(*SubmissionTrackingData)(nil),                           // 0: com.digitalasset.canton.participant.protocol.v0.SubmissionTrackingData
	(*TransactionSubmissionTrackingData)(nil),                // 1: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData
	(*TransactionSubmissionTrackingData_RejectionCause)(nil), // 2: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.RejectionCause
	(*CompletionInfo)(nil),                                   // 3: com.digitalasset.canton.participant.protocol.v0.CompletionInfo
	(*wrapperspb.StringValue)(nil),                           // 4: google.protobuf.StringValue
	(*emptypb.Empty)(nil),                                    // 5: google.protobuf.Empty
	(*CommandRejected_GrpcRejectionReasonTemplate)(nil),      // 6: com.digitalasset.canton.participant.protocol.v0.CommandRejected.GrpcRejectionReasonTemplate
}
var file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.participant.protocol.v0.SubmissionTrackingData.transaction:type_name -> com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData
	3, // 1: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.completion_info:type_name -> com.digitalasset.canton.participant.protocol.v0.CompletionInfo
	2, // 2: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.rejection_cause:type_name -> com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.RejectionCause
	4, // 3: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.domain_id:type_name -> google.protobuf.StringValue
	5, // 4: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.RejectionCause.timeout:type_name -> google.protobuf.Empty
	6, // 5: com.digitalasset.canton.participant.protocol.v0.TransactionSubmissionTrackingData.RejectionCause.rejection_reason_template:type_name -> com.digitalasset.canton.participant.protocol.v0.CommandRejected.GrpcRejectionReasonTemplate
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_init() }
func file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_init() {
	if File_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto != nil {
		return
	}
	file_com_digitalasset_canton_participant_protocol_v0_ledger_sync_event_proto_init()
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[0].OneofWrappers = []any{
		(*SubmissionTrackingData_Transaction)(nil),
	}
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes[2].OneofWrappers = []any{
		(*TransactionSubmissionTrackingData_RejectionCause_Timeout)(nil),
		(*TransactionSubmissionTrackingData_RejectionCause_RejectionReasonTemplate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto = out.File
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_rawDesc = nil
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_goTypes = nil
	file_com_digitalasset_canton_participant_protocol_v0_submission_tracking_proto_depIdxs = nil
}
