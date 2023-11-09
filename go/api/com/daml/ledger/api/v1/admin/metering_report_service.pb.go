// Copyright (c) 2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/daml/ledger/api/v1/admin/metering_report_service.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// Authorized if and only if the authenticated user is a participant admin.
type GetMeteringReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The from timestamp (inclusive).
	// Required.
	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// The to timestamp (exclusive).
	// If not provided, the server will default to its current time.
	To *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// If set to a non-empty value, then the report will only be generated for that application.
	// Optional.
	ApplicationId string `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *GetMeteringReportRequest) Reset() {
	*x = GetMeteringReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeteringReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeteringReportRequest) ProtoMessage() {}

func (x *GetMeteringReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeteringReportRequest.ProtoReflect.Descriptor instead.
func (*GetMeteringReportRequest) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMeteringReportRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetMeteringReportRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *GetMeteringReportRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type GetMeteringReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual request that was executed.
	Request *GetMeteringReportRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// The time at which the report was computed.
	ReportGenerationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=report_generation_time,json=reportGenerationTime,proto3" json:"report_generation_time,omitempty"`
	// The metering report json.  For a JSON Schema definition of the JSon see:
	// https://github.com/digital-asset/daml/blob/main/ledger-api/grpc-definitions/metering-report-schema.json
	MeteringReportJson *structpb.Struct `protobuf:"bytes,4,opt,name=metering_report_json,json=meteringReportJson,proto3" json:"metering_report_json,omitempty"`
}

func (x *GetMeteringReportResponse) Reset() {
	*x = GetMeteringReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeteringReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeteringReportResponse) ProtoMessage() {}

func (x *GetMeteringReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeteringReportResponse.ProtoReflect.Descriptor instead.
func (*GetMeteringReportResponse) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMeteringReportResponse) GetRequest() *GetMeteringReportRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GetMeteringReportResponse) GetReportGenerationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReportGenerationTime
	}
	return nil
}

func (x *GetMeteringReportResponse) GetMeteringReportJson() *structpb.Struct {
	if x != nil {
		return x.MeteringReportJson
	}
	return nil
}

var File_com_daml_ledger_api_v1_admin_metering_report_service_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa4, 0x02, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x16, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x14, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x12, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x12, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x32, 0x9e, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xa4, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x42, 0x18, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xaa, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x2e,
	0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescData = file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDesc
)

func file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDescData
}

var file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_daml_ledger_api_v1_admin_metering_report_service_proto_goTypes = []interface{}{
	(*GetMeteringReportRequest)(nil),  // 0: com.daml.ledger.api.v1.admin.GetMeteringReportRequest
	(*GetMeteringReportResponse)(nil), // 1: com.daml.ledger.api.v1.admin.GetMeteringReportResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
	(*structpb.Struct)(nil),           // 3: google.protobuf.Struct
}
var file_com_daml_ledger_api_v1_admin_metering_report_service_proto_depIdxs = []int32{
	2, // 0: com.daml.ledger.api.v1.admin.GetMeteringReportRequest.from:type_name -> google.protobuf.Timestamp
	2, // 1: com.daml.ledger.api.v1.admin.GetMeteringReportRequest.to:type_name -> google.protobuf.Timestamp
	0, // 2: com.daml.ledger.api.v1.admin.GetMeteringReportResponse.request:type_name -> com.daml.ledger.api.v1.admin.GetMeteringReportRequest
	2, // 3: com.daml.ledger.api.v1.admin.GetMeteringReportResponse.report_generation_time:type_name -> google.protobuf.Timestamp
	3, // 4: com.daml.ledger.api.v1.admin.GetMeteringReportResponse.metering_report_json:type_name -> google.protobuf.Struct
	0, // 5: com.daml.ledger.api.v1.admin.MeteringReportService.GetMeteringReport:input_type -> com.daml.ledger.api.v1.admin.GetMeteringReportRequest
	1, // 6: com.daml.ledger.api.v1.admin.MeteringReportService.GetMeteringReport:output_type -> com.daml.ledger.api.v1.admin.GetMeteringReportResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_admin_metering_report_service_proto_init() }
func file_com_daml_ledger_api_v1_admin_metering_report_service_proto_init() {
	if File_com_daml_ledger_api_v1_admin_metering_report_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeteringReportRequest); i {
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
		file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeteringReportResponse); i {
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
			RawDescriptor: file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_daml_ledger_api_v1_admin_metering_report_service_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_admin_metering_report_service_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_admin_metering_report_service_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_admin_metering_report_service_proto = out.File
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_goTypes = nil
	file_com_daml_ledger_api_v1_admin_metering_report_service_proto_depIdxs = nil
}
