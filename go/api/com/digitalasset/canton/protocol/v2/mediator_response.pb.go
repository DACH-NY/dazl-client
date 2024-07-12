// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/digitalasset/canton/protocol/v2/mediator_response.proto

package v2

import (
	v1 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v1"
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

type MediatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Sender            string                 `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	LocalVerdict      *v1.LocalVerdict       `protobuf:"bytes,4,opt,name=local_verdict,json=localVerdict,proto3" json:"local_verdict,omitempty"`
	RootHash          []byte                 `protobuf:"bytes,5,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`                            // May be empty for Malformed verdicts if the participant cannot determine the root hash.
	ConfirmingParties []string               `protobuf:"bytes,6,rep,name=confirming_parties,json=confirmingParties,proto3" json:"confirming_parties,omitempty"` // Empty iff the verdict is malformed.
	DomainId          string                 `protobuf:"bytes,7,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	ViewPosition      *ViewPosition          `protobuf:"bytes,8,opt,name=view_position,json=viewPosition,proto3" json:"view_position,omitempty"` // Added view_position. May be empty for Malformed verdicts
}

func (x *MediatorResponse) Reset() {
	*x = MediatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediatorResponse) ProtoMessage() {}

func (x *MediatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediatorResponse.ProtoReflect.Descriptor instead.
func (*MediatorResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescGZIP(), []int{0}
}

func (x *MediatorResponse) GetRequestId() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *MediatorResponse) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MediatorResponse) GetLocalVerdict() *v1.LocalVerdict {
	if x != nil {
		return x.LocalVerdict
	}
	return nil
}

func (x *MediatorResponse) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

func (x *MediatorResponse) GetConfirmingParties() []string {
	if x != nil {
		return x.ConfirmingParties
	}
	return nil
}

func (x *MediatorResponse) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *MediatorResponse) GetViewPosition() *ViewPosition {
	if x != nil {
		return x.ViewPosition
	}
	return nil
}

// New message
type ViewPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position []*MerkleSeqIndex `protobuf:"bytes,1,rep,name=position,proto3" json:"position,omitempty"`
}

func (x *ViewPosition) Reset() {
	*x = ViewPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewPosition) ProtoMessage() {}

func (x *ViewPosition) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewPosition.ProtoReflect.Descriptor instead.
func (*ViewPosition) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescGZIP(), []int{1}
}

func (x *ViewPosition) GetPosition() []*MerkleSeqIndex {
	if x != nil {
		return x.Position
	}
	return nil
}

// New message
type MerkleSeqIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRight []bool `protobuf:"varint,1,rep,packed,name=is_right,json=isRight,proto3" json:"is_right,omitempty"`
}

func (x *MerkleSeqIndex) Reset() {
	*x = MerkleSeqIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleSeqIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleSeqIndex) ProtoMessage() {}

func (x *MerkleSeqIndex) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleSeqIndex.ProtoReflect.Descriptor instead.
func (*MerkleSeqIndex) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescGZIP(), []int{2}
}

func (x *MerkleSeqIndex) GetIsRight() []bool {
	if x != nil {
		return x.IsRight
	}
	return nil
}

var File_com_digitalasset_canton_protocol_v2_mediator_response_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x32, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x03, 0x0a, 0x10, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x64, 0x69,
	0x63, 0x74, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0d, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x5f, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e,
	0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0e, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x53, 0x65, 0x71, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescData = file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_digitalasset_canton_protocol_v2_mediator_response_proto_goTypes = []interface{}{
	(*MediatorResponse)(nil),      // 0: com.digitalasset.canton.protocol.v2.MediatorResponse
	(*ViewPosition)(nil),          // 1: com.digitalasset.canton.protocol.v2.ViewPosition
	(*MerkleSeqIndex)(nil),        // 2: com.digitalasset.canton.protocol.v2.MerkleSeqIndex
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*v1.LocalVerdict)(nil),       // 4: com.digitalasset.canton.protocol.v1.LocalVerdict
}
var file_com_digitalasset_canton_protocol_v2_mediator_response_proto_depIdxs = []int32{
	3, // 0: com.digitalasset.canton.protocol.v2.MediatorResponse.request_id:type_name -> google.protobuf.Timestamp
	4, // 1: com.digitalasset.canton.protocol.v2.MediatorResponse.local_verdict:type_name -> com.digitalasset.canton.protocol.v1.LocalVerdict
	1, // 2: com.digitalasset.canton.protocol.v2.MediatorResponse.view_position:type_name -> com.digitalasset.canton.protocol.v2.ViewPosition
	2, // 3: com.digitalasset.canton.protocol.v2.ViewPosition.position:type_name -> com.digitalasset.canton.protocol.v2.MerkleSeqIndex
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v2_mediator_response_proto_init() }
func file_com_digitalasset_canton_protocol_v2_mediator_response_proto_init() {
	if File_com_digitalasset_canton_protocol_v2_mediator_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediatorResponse); i {
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
		file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewPosition); i {
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
		file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleSeqIndex); i {
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
			RawDescriptor: file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v2_mediator_response_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v2_mediator_response_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v2_mediator_response_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v2_mediator_response_proto = out.File
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v2_mediator_response_proto_depIdxs = nil
}