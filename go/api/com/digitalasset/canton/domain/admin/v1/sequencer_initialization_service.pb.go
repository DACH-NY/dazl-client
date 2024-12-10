// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/domain/admin/v1/sequencer_initialization_service.proto

package v1

import (
	v01 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/admin/v0"
	v0 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0"
	v1 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v1"
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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId         string                     `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	TopologySnapshot *v0.TopologyTransactions   `protobuf:"bytes,2,opt,name=topology_snapshot,json=topologySnapshot,proto3" json:"topology_snapshot,omitempty"`
	DomainParameters *v1.StaticDomainParameters `protobuf:"bytes,4,opt,name=domain_parameters,json=domainParameters,proto3" json:"domain_parameters,omitempty"`
	Snapshot         *v01.SequencerSnapshot     `protobuf:"bytes,3,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *InitRequest) GetTopologySnapshot() *v0.TopologyTransactions {
	if x != nil {
		return x.TopologySnapshot
	}
	return nil
}

func (x *InitRequest) GetDomainParameters() *v1.StaticDomainParameters {
	if x != nil {
		return x.DomainParameters
	}
	return nil
}

func (x *InitRequest) GetSnapshot() *v01.SequencerSnapshot {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

var File_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x4f, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x30, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x0b, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x11, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x68, 0x0a,
	0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x56, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42,
	0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c,
	0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescData = file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDesc
)

func file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescData)
	})
	return file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDescData
}

var file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_goTypes = []any{
	(*InitRequest)(nil),               // 0: com.digitalasset.canton.domain.admin.v1.InitRequest
	(*v0.TopologyTransactions)(nil),   // 1: com.digitalasset.canton.protocol.v0.TopologyTransactions
	(*v1.StaticDomainParameters)(nil), // 2: com.digitalasset.canton.protocol.v1.StaticDomainParameters
	(*v01.SequencerSnapshot)(nil),     // 3: com.digitalasset.canton.domain.admin.v0.SequencerSnapshot
}
var file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.domain.admin.v1.InitRequest.topology_snapshot:type_name -> com.digitalasset.canton.protocol.v0.TopologyTransactions
	2, // 1: com.digitalasset.canton.domain.admin.v1.InitRequest.domain_parameters:type_name -> com.digitalasset.canton.protocol.v1.StaticDomainParameters
	3, // 2: com.digitalasset.canton.domain.admin.v1.InitRequest.snapshot:type_name -> com.digitalasset.canton.domain.admin.v0.SequencerSnapshot
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_init()
}
func file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_init() {
	if File_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InitRequest); i {
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
			RawDescriptor: file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto = out.File
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_rawDesc = nil
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_goTypes = nil
	file_com_digitalasset_canton_domain_admin_v1_sequencer_initialization_service_proto_depIdxs = nil
}
