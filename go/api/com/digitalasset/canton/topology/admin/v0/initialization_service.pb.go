// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: com/digitalasset/canton/topology/admin/v0/initialization_service.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InitIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier  string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Instance    string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *InitIdRequest) Reset() {
	*x = InitIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitIdRequest) ProtoMessage() {}

func (x *InitIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitIdRequest.ProtoReflect.Descriptor instead.
func (*InitIdRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitIdRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *InitIdRequest) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *InitIdRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type InitIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueIdentifier string `protobuf:"bytes,1,opt,name=unique_identifier,json=uniqueIdentifier,proto3" json:"unique_identifier,omitempty"`
	Instance         string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *InitIdResponse) Reset() {
	*x = InitIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitIdResponse) ProtoMessage() {}

func (x *InitIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitIdResponse.ProtoReflect.Descriptor instead.
func (*InitIdResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescGZIP(), []int{1}
}

func (x *InitIdResponse) GetUniqueIdentifier() string {
	if x != nil {
		return x.UniqueIdentifier
	}
	return ""
}

func (x *InitIdResponse) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type GetIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Initialized      bool   `protobuf:"varint,1,opt,name=initialized,proto3" json:"initialized,omitempty"`
	UniqueIdentifier string `protobuf:"bytes,2,opt,name=unique_identifier,json=uniqueIdentifier,proto3" json:"unique_identifier,omitempty"`
	Instance         string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *GetIdResponse) Reset() {
	*x = GetIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdResponse) ProtoMessage() {}

func (x *GetIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdResponse.ProtoReflect.Descriptor instead.
func (*GetIdResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdResponse) GetInitialized() bool {
	if x != nil {
		return x.Initialized
	}
	return false
}

func (x *GetIdResponse) GetUniqueIdentifier() string {
	if x != nil {
		return x.UniqueIdentifier
	}
	return ""
}

func (x *GetIdResponse) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

var File_com_digitalasset_canton_topology_admin_v0_initialization_service_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x30, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x0d, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x59, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x7a, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2b,
	0x0a, 0x11, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xb4, 0x02, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7d, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x38, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x5a,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescData = file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDesc
)

func file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescData)
	})
	return file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDescData
}

var file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_goTypes = []interface{}{
	(*InitIdRequest)(nil),         // 0: com.digitalasset.canton.topology.admin.v0.InitIdRequest
	(*InitIdResponse)(nil),        // 1: com.digitalasset.canton.topology.admin.v0.InitIdResponse
	(*GetIdResponse)(nil),         // 2: com.digitalasset.canton.topology.admin.v0.GetIdResponse
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_depIdxs = []int32{
	0, // 0: com.digitalasset.canton.topology.admin.v0.InitializationService.InitId:input_type -> com.digitalasset.canton.topology.admin.v0.InitIdRequest
	3, // 1: com.digitalasset.canton.topology.admin.v0.InitializationService.GetId:input_type -> google.protobuf.Empty
	3, // 2: com.digitalasset.canton.topology.admin.v0.InitializationService.CurrentTime:input_type -> google.protobuf.Empty
	1, // 3: com.digitalasset.canton.topology.admin.v0.InitializationService.InitId:output_type -> com.digitalasset.canton.topology.admin.v0.InitIdResponse
	2, // 4: com.digitalasset.canton.topology.admin.v0.InitializationService.GetId:output_type -> com.digitalasset.canton.topology.admin.v0.GetIdResponse
	4, // 5: com.digitalasset.canton.topology.admin.v0.InitializationService.CurrentTime:output_type -> google.protobuf.Timestamp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_init() }
func file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_init() {
	if File_com_digitalasset_canton_topology_admin_v0_initialization_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitIdRequest); i {
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
		file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitIdResponse); i {
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
		file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdResponse); i {
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
			RawDescriptor: file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_topology_admin_v0_initialization_service_proto = out.File
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_rawDesc = nil
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_goTypes = nil
	file_com_digitalasset_canton_topology_admin_v0_initialization_service_proto_depIdxs = nil
}
