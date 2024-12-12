// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto

package v30

import (
	v30 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/admin/pruning/v30"
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

type MediatorPruning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediatorPruning) Reset() {
	*x = MediatorPruning{}
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediatorPruning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediatorPruning) ProtoMessage() {}

func (x *MediatorPruning) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediatorPruning.ProtoReflect.Descriptor instead.
func (*MediatorPruning) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescGZIP(), []int{0}
}

type MediatorPruning_PruneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MediatorPruning_PruneRequest) Reset() {
	*x = MediatorPruning_PruneRequest{}
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediatorPruning_PruneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediatorPruning_PruneRequest) ProtoMessage() {}

func (x *MediatorPruning_PruneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediatorPruning_PruneRequest.ProtoReflect.Descriptor instead.
func (*MediatorPruning_PruneRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MediatorPruning_PruneRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type MediatorPruning_PruneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediatorPruning_PruneResponse) Reset() {
	*x = MediatorPruning_PruneResponse{}
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediatorPruning_PruneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediatorPruning_PruneResponse) ProtoMessage() {}

func (x *MediatorPruning_PruneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediatorPruning_PruneResponse.ProtoReflect.Descriptor instead.
func (*MediatorPruning_PruneResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescGZIP(), []int{0, 1}
}

var File_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDesc = []byte{
	0x0a, 0x50, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x30, 0x2f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x30, 0x1a, 0x37,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x30, 0x2f, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x48, 0x0a, 0x0c, 0x50,
	0x72, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdc, 0x09, 0x0a, 0x1d, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x75,
	0x6e, 0x65, 0x12, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x30, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x72, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x74,
	0x43, 0x72, 0x6f, 0x6e, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30,
	0x2e, 0x53, 0x65, 0x74, 0x43, 0x72, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74,
	0x43, 0x72, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x97, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x61, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0d,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x40, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0xaf, 0x01, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x49,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x33, 0x30, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5b, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x33, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescData = file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDesc
)

func file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescData)
	})
	return file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDescData
}

var file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_goTypes = []any{
	(*MediatorPruning)(nil),                     // 0: com.digitalasset.canton.mediator.admin.v30.MediatorPruning
	(*MediatorPruning_PruneRequest)(nil),        // 1: com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneRequest
	(*MediatorPruning_PruneResponse)(nil),       // 2: com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneResponse
	(*timestamppb.Timestamp)(nil),               // 3: google.protobuf.Timestamp
	(*v30.SetSchedule_Request)(nil),             // 4: com.digitalasset.canton.admin.pruning.v30.SetSchedule.Request
	(*v30.SetCron_Request)(nil),                 // 5: com.digitalasset.canton.admin.pruning.v30.SetCron.Request
	(*v30.SetMaxDuration_Request)(nil),          // 6: com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Request
	(*v30.SetRetention_Request)(nil),            // 7: com.digitalasset.canton.admin.pruning.v30.SetRetention.Request
	(*v30.ClearSchedule_Request)(nil),           // 8: com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Request
	(*v30.GetSchedule_Request)(nil),             // 9: com.digitalasset.canton.admin.pruning.v30.GetSchedule.Request
	(*v30.LocatePruningTimestamp_Request)(nil),  // 10: com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.Request
	(*v30.SetSchedule_Response)(nil),            // 11: com.digitalasset.canton.admin.pruning.v30.SetSchedule.Response
	(*v30.SetCron_Response)(nil),                // 12: com.digitalasset.canton.admin.pruning.v30.SetCron.Response
	(*v30.SetMaxDuration_Response)(nil),         // 13: com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Response
	(*v30.SetRetention_Response)(nil),           // 14: com.digitalasset.canton.admin.pruning.v30.SetRetention.Response
	(*v30.ClearSchedule_Response)(nil),          // 15: com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Response
	(*v30.GetSchedule_Response)(nil),            // 16: com.digitalasset.canton.admin.pruning.v30.GetSchedule.Response
	(*v30.LocatePruningTimestamp_Response)(nil), // 17: com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.Response
}
var file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_depIdxs = []int32{
	3,  // 0: com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneRequest.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 1: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.Prune:input_type -> com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneRequest
	4,  // 2: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetSchedule:input_type -> com.digitalasset.canton.admin.pruning.v30.SetSchedule.Request
	5,  // 3: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetCron:input_type -> com.digitalasset.canton.admin.pruning.v30.SetCron.Request
	6,  // 4: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetMaxDuration:input_type -> com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Request
	7,  // 5: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetRetention:input_type -> com.digitalasset.canton.admin.pruning.v30.SetRetention.Request
	8,  // 6: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.ClearSchedule:input_type -> com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Request
	9,  // 7: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.GetSchedule:input_type -> com.digitalasset.canton.admin.pruning.v30.GetSchedule.Request
	10, // 8: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.LocatePruningTimestamp:input_type -> com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.Request
	2,  // 9: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.Prune:output_type -> com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneResponse
	11, // 10: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetSchedule:output_type -> com.digitalasset.canton.admin.pruning.v30.SetSchedule.Response
	12, // 11: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetCron:output_type -> com.digitalasset.canton.admin.pruning.v30.SetCron.Response
	13, // 12: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetMaxDuration:output_type -> com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Response
	14, // 13: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.SetRetention:output_type -> com.digitalasset.canton.admin.pruning.v30.SetRetention.Response
	15, // 14: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.ClearSchedule:output_type -> com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Response
	16, // 15: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.GetSchedule:output_type -> com.digitalasset.canton.admin.pruning.v30.GetSchedule.Response
	17, // 16: com.digitalasset.canton.mediator.admin.v30.MediatorAdministrationService.LocatePruningTimestamp:output_type -> com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.Response
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() {
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_init()
}
func file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_init() {
	if File_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto = out.File
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_rawDesc = nil
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_goTypes = nil
	file_com_digitalasset_canton_mediator_admin_v30_mediator_administration_service_proto_depIdxs = nil
}