// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: com/digitalasset/canton/domain/admin/v0/enterprise_sequencer_connection_service.proto

package v0

import (
	v0 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0"
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

type GetConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConnectionRequest) Reset() {
	*x = GetConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionRequest) ProtoMessage() {}

func (x *GetConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetConnectionRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescGZIP(), []int{0}
}

type GetConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequencerConnections    []*v0.SequencerConnection `protobuf:"bytes,1,rep,name=sequencer_connections,json=sequencerConnections,proto3" json:"sequencer_connections,omitempty"`
	SequencerTrustThreshold uint32                    `protobuf:"varint,2,opt,name=sequencerTrustThreshold,proto3" json:"sequencerTrustThreshold,omitempty"`
}

func (x *GetConnectionResponse) Reset() {
	*x = GetConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionResponse) ProtoMessage() {}

func (x *GetConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionResponse.ProtoReflect.Descriptor instead.
func (*GetConnectionResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetConnectionResponse) GetSequencerConnections() []*v0.SequencerConnection {
	if x != nil {
		return x.SequencerConnections
	}
	return nil
}

func (x *GetConnectionResponse) GetSequencerTrustThreshold() uint32 {
	if x != nil {
		return x.SequencerTrustThreshold
	}
	return 0
}

type SetConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SequencerConnections    []*v0.SequencerConnection `protobuf:"bytes,1,rep,name=sequencer_connections,json=sequencerConnections,proto3" json:"sequencer_connections,omitempty"`
	SequencerTrustThreshold uint32                    `protobuf:"varint,2,opt,name=sequencerTrustThreshold,proto3" json:"sequencerTrustThreshold,omitempty"`
}

func (x *SetConnectionRequest) Reset() {
	*x = SetConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConnectionRequest) ProtoMessage() {}

func (x *SetConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConnectionRequest.ProtoReflect.Descriptor instead.
func (*SetConnectionRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetConnectionRequest) GetSequencerConnections() []*v0.SequencerConnection {
	if x != nil {
		return x.SequencerConnections
	}
	return nil
}

func (x *SetConnectionRequest) GetSequencerTrustThreshold() uint32 {
	if x != nil {
		return x.SequencerTrustThreshold
	}
	return 0
}

var File_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDesc = []byte{
	0x0a, 0x55, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30,
	0x1a, 0x40, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6f, 0x0a, 0x15, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x17, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0xc1, 0x01, 0x0a,
	0x14, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x15, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x14, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0x54, 0x72, 0x75, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x72, 0x54, 0x72, 0x75, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x32, 0x9f, 0x02, 0x0a, 0x24, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x53, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64,
	0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescData = file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDesc
)

func file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescData)
	})
	return file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDescData
}

var file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_goTypes = []interface{}{
	(*GetConnectionRequest)(nil),   // 0: com.digitalasset.canton.domain.admin.v0.GetConnectionRequest
	(*GetConnectionResponse)(nil),  // 1: com.digitalasset.canton.domain.admin.v0.GetConnectionResponse
	(*SetConnectionRequest)(nil),   // 2: com.digitalasset.canton.domain.admin.v0.SetConnectionRequest
	(*v0.SequencerConnection)(nil), // 3: com.digitalasset.canton.domain.api.v0.SequencerConnection
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_depIdxs = []int32{
	3, // 0: com.digitalasset.canton.domain.admin.v0.GetConnectionResponse.sequencer_connections:type_name -> com.digitalasset.canton.domain.api.v0.SequencerConnection
	3, // 1: com.digitalasset.canton.domain.admin.v0.SetConnectionRequest.sequencer_connections:type_name -> com.digitalasset.canton.domain.api.v0.SequencerConnection
	0, // 2: com.digitalasset.canton.domain.admin.v0.EnterpriseSequencerConnectionService.GetConnection:input_type -> com.digitalasset.canton.domain.admin.v0.GetConnectionRequest
	2, // 3: com.digitalasset.canton.domain.admin.v0.EnterpriseSequencerConnectionService.SetConnection:input_type -> com.digitalasset.canton.domain.admin.v0.SetConnectionRequest
	1, // 4: com.digitalasset.canton.domain.admin.v0.EnterpriseSequencerConnectionService.GetConnection:output_type -> com.digitalasset.canton.domain.admin.v0.GetConnectionResponse
	4, // 5: com.digitalasset.canton.domain.admin.v0.EnterpriseSequencerConnectionService.SetConnection:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_init()
}
func file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_init() {
	if File_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionRequest); i {
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
		file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionResponse); i {
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
		file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConnectionRequest); i {
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
			RawDescriptor: file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto = out.File
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_rawDesc = nil
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_goTypes = nil
	file_com_digitalasset_canton_domain_admin_v0_enterprise_sequencer_connection_service_proto_depIdxs = nil
}
