// Copyright (c) 2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/daml/ledger/api/v1/admin/config_management_service.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type GetTimeModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTimeModelRequest) Reset() {
	*x = GetTimeModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeModelRequest) ProtoMessage() {}

func (x *GetTimeModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeModelRequest.ProtoReflect.Descriptor instead.
func (*GetTimeModelRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP(), []int{0}
}

type GetTimeModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The current configuration generation. The generation is a monotonically increasing
	// integer that is incremented on each change. Used when setting the time model.
	ConfigurationGeneration int64 `protobuf:"varint,1,opt,name=configuration_generation,json=configurationGeneration,proto3" json:"configuration_generation,omitempty"`
	// The current ledger time model.
	TimeModel *TimeModel `protobuf:"bytes,2,opt,name=time_model,json=timeModel,proto3" json:"time_model,omitempty"`
}

func (x *GetTimeModelResponse) Reset() {
	*x = GetTimeModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeModelResponse) ProtoMessage() {}

func (x *GetTimeModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeModelResponse.ProtoReflect.Descriptor instead.
func (*GetTimeModelResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTimeModelResponse) GetConfigurationGeneration() int64 {
	if x != nil {
		return x.ConfigurationGeneration
	}
	return 0
}

func (x *GetTimeModelResponse) GetTimeModel() *TimeModel {
	if x != nil {
		return x.TimeModel
	}
	return nil
}

type SetTimeModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Submission identifier used for tracking the request and to reject
	// duplicate submissions.
	// Required.
	SubmissionId string `protobuf:"bytes,1,opt,name=submission_id,json=submissionId,proto3" json:"submission_id,omitempty"`
	// Deadline for the configuration change after which the change is rejected.
	MaximumRecordTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=maximum_record_time,json=maximumRecordTime,proto3" json:"maximum_record_time,omitempty"`
	// The current configuration generation which we're submitting the change against.
	// This is used to perform a compare-and-swap of the configuration to
	// safeguard against concurrent modifications.
	// Required.
	ConfigurationGeneration int64 `protobuf:"varint,3,opt,name=configuration_generation,json=configurationGeneration,proto3" json:"configuration_generation,omitempty"`
	// The new time model that replaces the current one.
	// Required.
	NewTimeModel *TimeModel `protobuf:"bytes,4,opt,name=new_time_model,json=newTimeModel,proto3" json:"new_time_model,omitempty"`
}

func (x *SetTimeModelRequest) Reset() {
	*x = SetTimeModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTimeModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTimeModelRequest) ProtoMessage() {}

func (x *SetTimeModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTimeModelRequest.ProtoReflect.Descriptor instead.
func (*SetTimeModelRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetTimeModelRequest) GetSubmissionId() string {
	if x != nil {
		return x.SubmissionId
	}
	return ""
}

func (x *SetTimeModelRequest) GetMaximumRecordTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MaximumRecordTime
	}
	return nil
}

func (x *SetTimeModelRequest) GetConfigurationGeneration() int64 {
	if x != nil {
		return x.ConfigurationGeneration
	}
	return 0
}

func (x *SetTimeModelRequest) GetNewTimeModel() *TimeModel {
	if x != nil {
		return x.NewTimeModel
	}
	return nil
}

type SetTimeModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The configuration generation of the committed time model.
	ConfigurationGeneration int64 `protobuf:"varint,1,opt,name=configuration_generation,json=configurationGeneration,proto3" json:"configuration_generation,omitempty"`
}

func (x *SetTimeModelResponse) Reset() {
	*x = SetTimeModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTimeModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTimeModelResponse) ProtoMessage() {}

func (x *SetTimeModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTimeModelResponse.ProtoReflect.Descriptor instead.
func (*SetTimeModelResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetTimeModelResponse) GetConfigurationGeneration() int64 {
	if x != nil {
		return x.ConfigurationGeneration
	}
	return 0
}

type TimeModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The expected average latency of a transaction, i.e., the average time
	// from submitting the transaction to a [[WriteService]] and the transaction
	// being assigned a record time.
	// Required.
	AvgTransactionLatency *durationpb.Duration `protobuf:"bytes,4,opt,name=avg_transaction_latency,json=avgTransactionLatency,proto3" json:"avg_transaction_latency,omitempty"`
	// The minimimum skew between ledger time and record time: lt_TX >= rt_TX - minSkew
	// Required.
	MinSkew *durationpb.Duration `protobuf:"bytes,5,opt,name=min_skew,json=minSkew,proto3" json:"min_skew,omitempty"`
	// The maximum skew between ledger time and record time: lt_TX <= rt_TX + maxSkew
	// Required.
	MaxSkew *durationpb.Duration `protobuf:"bytes,6,opt,name=max_skew,json=maxSkew,proto3" json:"max_skew,omitempty"`
}

func (x *TimeModel) Reset() {
	*x = TimeModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeModel) ProtoMessage() {}

func (x *TimeModel) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeModel.ProtoReflect.Descriptor instead.
func (*TimeModel) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP(), []int{4}
}

func (x *TimeModel) GetAvgTransactionLatency() *durationpb.Duration {
	if x != nil {
		return x.AvgTransactionLatency
	}
	return nil
}

func (x *TimeModel) GetMinSkew() *durationpb.Duration {
	if x != nil {
		return x.MinSkew
	}
	return nil
}

func (x *TimeModel) GetMaxSkew() *durationpb.Duration {
	if x != nil {
		return x.MaxSkew
	}
	return nil
}

var File_com_daml_ledger_api_v1_admin_config_management_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x90, 0x02, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4a, 0x0a,
	0x13, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x51, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x51, 0x0a, 0x17, 0x61, 0x76, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x15, 0x61, 0x76, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x6b, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x6b, 0x65, 0x77, 0x12, 0x34, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x6b, 0x65, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x53,
	0x6b, 0x65, 0x77, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x32, 0x87, 0x02, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x75, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xad, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x42, 0x21, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61,
	0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0xaa, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescData = file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDesc
)

func file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDescData
}

var file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_com_daml_ledger_api_v1_admin_config_management_service_proto_goTypes = []interface{}{
	(*GetTimeModelRequest)(nil),   // 0: com.daml.ledger.api.v1.admin.GetTimeModelRequest
	(*GetTimeModelResponse)(nil),  // 1: com.daml.ledger.api.v1.admin.GetTimeModelResponse
	(*SetTimeModelRequest)(nil),   // 2: com.daml.ledger.api.v1.admin.SetTimeModelRequest
	(*SetTimeModelResponse)(nil),  // 3: com.daml.ledger.api.v1.admin.SetTimeModelResponse
	(*TimeModel)(nil),             // 4: com.daml.ledger.api.v1.admin.TimeModel
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file_com_daml_ledger_api_v1_admin_config_management_service_proto_depIdxs = []int32{
	4, // 0: com.daml.ledger.api.v1.admin.GetTimeModelResponse.time_model:type_name -> com.daml.ledger.api.v1.admin.TimeModel
	5, // 1: com.daml.ledger.api.v1.admin.SetTimeModelRequest.maximum_record_time:type_name -> google.protobuf.Timestamp
	4, // 2: com.daml.ledger.api.v1.admin.SetTimeModelRequest.new_time_model:type_name -> com.daml.ledger.api.v1.admin.TimeModel
	6, // 3: com.daml.ledger.api.v1.admin.TimeModel.avg_transaction_latency:type_name -> google.protobuf.Duration
	6, // 4: com.daml.ledger.api.v1.admin.TimeModel.min_skew:type_name -> google.protobuf.Duration
	6, // 5: com.daml.ledger.api.v1.admin.TimeModel.max_skew:type_name -> google.protobuf.Duration
	0, // 6: com.daml.ledger.api.v1.admin.ConfigManagementService.GetTimeModel:input_type -> com.daml.ledger.api.v1.admin.GetTimeModelRequest
	2, // 7: com.daml.ledger.api.v1.admin.ConfigManagementService.SetTimeModel:input_type -> com.daml.ledger.api.v1.admin.SetTimeModelRequest
	1, // 8: com.daml.ledger.api.v1.admin.ConfigManagementService.GetTimeModel:output_type -> com.daml.ledger.api.v1.admin.GetTimeModelResponse
	3, // 9: com.daml.ledger.api.v1.admin.ConfigManagementService.SetTimeModel:output_type -> com.daml.ledger.api.v1.admin.SetTimeModelResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_admin_config_management_service_proto_init() }
func file_com_daml_ledger_api_v1_admin_config_management_service_proto_init() {
	if File_com_daml_ledger_api_v1_admin_config_management_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeModelRequest); i {
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
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeModelResponse); i {
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
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTimeModelRequest); i {
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
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTimeModelResponse); i {
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
		file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeModel); i {
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
			RawDescriptor: file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v1_admin_config_management_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_admin_config_management_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_admin_config_management_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_admin_config_management_service_proto = out.File
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_goTypes = nil
	file_com_daml_ledger_api_v1_admin_config_management_service_proto_depIdxs = nil
}
