// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/time/admin/v0/time_tracker_config.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeProofRequestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialRetryDelay  *durationpb.Duration `protobuf:"bytes,1,opt,name=initialRetryDelay,proto3" json:"initialRetryDelay,omitempty"`
	MaxRetryDelay      *durationpb.Duration `protobuf:"bytes,2,opt,name=maxRetryDelay,proto3" json:"maxRetryDelay,omitempty"`
	MaxSequencingDelay *durationpb.Duration `protobuf:"bytes,3,opt,name=maxSequencingDelay,proto3" json:"maxSequencingDelay,omitempty"`
}

func (x *TimeProofRequestConfig) Reset() {
	*x = TimeProofRequestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeProofRequestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeProofRequestConfig) ProtoMessage() {}

func (x *TimeProofRequestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeProofRequestConfig.ProtoReflect.Descriptor instead.
func (*TimeProofRequestConfig) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescGZIP(), []int{0}
}

func (x *TimeProofRequestConfig) GetInitialRetryDelay() *durationpb.Duration {
	if x != nil {
		return x.InitialRetryDelay
	}
	return nil
}

func (x *TimeProofRequestConfig) GetMaxRetryDelay() *durationpb.Duration {
	if x != nil {
		return x.MaxRetryDelay
	}
	return nil
}

func (x *TimeProofRequestConfig) GetMaxSequencingDelay() *durationpb.Duration {
	if x != nil {
		return x.MaxSequencingDelay
	}
	return nil
}

type DomainTimeTrackerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObservationLatency     *durationpb.Duration    `protobuf:"bytes,1,opt,name=observationLatency,proto3" json:"observationLatency,omitempty"`
	PatienceDuration       *durationpb.Duration    `protobuf:"bytes,2,opt,name=patienceDuration,proto3" json:"patienceDuration,omitempty"`
	MinObservationDuration *durationpb.Duration    `protobuf:"bytes,3,opt,name=minObservationDuration,proto3" json:"minObservationDuration,omitempty"`
	TimeProofRequest       *TimeProofRequestConfig `protobuf:"bytes,4,opt,name=timeProofRequest,proto3" json:"timeProofRequest,omitempty"`
}

func (x *DomainTimeTrackerConfig) Reset() {
	*x = DomainTimeTrackerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainTimeTrackerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainTimeTrackerConfig) ProtoMessage() {}

func (x *DomainTimeTrackerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainTimeTrackerConfig.ProtoReflect.Descriptor instead.
func (*DomainTimeTrackerConfig) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescGZIP(), []int{1}
}

func (x *DomainTimeTrackerConfig) GetObservationLatency() *durationpb.Duration {
	if x != nil {
		return x.ObservationLatency
	}
	return nil
}

func (x *DomainTimeTrackerConfig) GetPatienceDuration() *durationpb.Duration {
	if x != nil {
		return x.PatienceDuration
	}
	return nil
}

func (x *DomainTimeTrackerConfig) GetMinObservationDuration() *durationpb.Duration {
	if x != nil {
		return x.MinObservationDuration
	}
	return nil
}

func (x *DomainTimeTrackerConfig) GetTimeProofRequest() *TimeProofRequestConfig {
	if x != nil {
		return x.TimeProofRequest
	}
	return nil
}

var File_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x16, 0x54, 0x69, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x3f, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x49, 0x0a,
	0x12, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0xe9, 0x02, 0x0a, 0x17, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x45, 0x0a, 0x10, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x16, 0x6d, 0x69, 0x6e, 0x4f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x16, 0x6d, 0x69, 0x6e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a, 0x10, 0x74, 0x69, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescData = file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDesc
)

func file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescData)
	})
	return file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDescData
}

var file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_goTypes = []any{
	(*TimeProofRequestConfig)(nil),  // 0: com.digitalasset.canton.time.admin.v0.TimeProofRequestConfig
	(*DomainTimeTrackerConfig)(nil), // 1: com.digitalasset.canton.time.admin.v0.DomainTimeTrackerConfig
	(*durationpb.Duration)(nil),     // 2: google.protobuf.Duration
}
var file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_depIdxs = []int32{
	2, // 0: com.digitalasset.canton.time.admin.v0.TimeProofRequestConfig.initialRetryDelay:type_name -> google.protobuf.Duration
	2, // 1: com.digitalasset.canton.time.admin.v0.TimeProofRequestConfig.maxRetryDelay:type_name -> google.protobuf.Duration
	2, // 2: com.digitalasset.canton.time.admin.v0.TimeProofRequestConfig.maxSequencingDelay:type_name -> google.protobuf.Duration
	2, // 3: com.digitalasset.canton.time.admin.v0.DomainTimeTrackerConfig.observationLatency:type_name -> google.protobuf.Duration
	2, // 4: com.digitalasset.canton.time.admin.v0.DomainTimeTrackerConfig.patienceDuration:type_name -> google.protobuf.Duration
	2, // 5: com.digitalasset.canton.time.admin.v0.DomainTimeTrackerConfig.minObservationDuration:type_name -> google.protobuf.Duration
	0, // 6: com.digitalasset.canton.time.admin.v0.DomainTimeTrackerConfig.timeProofRequest:type_name -> com.digitalasset.canton.time.admin.v0.TimeProofRequestConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_init() }
func file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_init() {
	if File_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TimeProofRequestConfig); i {
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
		file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DomainTimeTrackerConfig); i {
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
			RawDescriptor: file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto = out.File
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_rawDesc = nil
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_goTypes = nil
	file_com_digitalasset_canton_time_admin_v0_time_tracker_config_proto_depIdxs = nil
}
