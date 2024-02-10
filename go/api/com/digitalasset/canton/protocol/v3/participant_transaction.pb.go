// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/digitalasset/canton/protocol/v3/participant_transaction.proto

package v3

import (
	v0 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/crypto/v0"
	v01 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0"
	v1 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v1"
	v2 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v2"
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

// Compared to v2: action_description changed type from v1.ActionDescription to v3.ActionDescription
type ViewParticipantData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salt                           *v0.Salt                                 `protobuf:"bytes,1,opt,name=salt,proto3" json:"salt,omitempty"`
	CoreInputs                     []*v1.InputContract                      `protobuf:"bytes,2,rep,name=core_inputs,json=coreInputs,proto3" json:"core_inputs,omitempty"`
	CreatedCore                    []*v1.CreatedContract                    `protobuf:"bytes,3,rep,name=created_core,json=createdCore,proto3" json:"created_core,omitempty"`
	CreatedInSubviewArchivedInCore []string                                 `protobuf:"bytes,4,rep,name=created_in_subview_archived_in_core,json=createdInSubviewArchivedInCore,proto3" json:"created_in_subview_archived_in_core,omitempty"` // ids of contracts created in a subview and archived in the core
	ResolvedKeys                   []*v01.ViewParticipantData_ResolvedKey   `protobuf:"bytes,5,rep,name=resolved_keys,json=resolvedKeys,proto3" json:"resolved_keys,omitempty"`
	ActionDescription              *v2.ActionDescription                    `protobuf:"bytes,6,opt,name=action_description,json=actionDescription,proto3" json:"action_description,omitempty"`
	RollbackContext                *v01.ViewParticipantData_RollbackContext `protobuf:"bytes,7,opt,name=rollback_context,json=rollbackContext,proto3" json:"rollback_context,omitempty"` // optional; defaults to the empty RollbackContext if omitted.
}

func (x *ViewParticipantData) Reset() {
	*x = ViewParticipantData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewParticipantData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewParticipantData) ProtoMessage() {}

func (x *ViewParticipantData) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewParticipantData.ProtoReflect.Descriptor instead.
func (*ViewParticipantData) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *ViewParticipantData) GetSalt() *v0.Salt {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *ViewParticipantData) GetCoreInputs() []*v1.InputContract {
	if x != nil {
		return x.CoreInputs
	}
	return nil
}

func (x *ViewParticipantData) GetCreatedCore() []*v1.CreatedContract {
	if x != nil {
		return x.CreatedCore
	}
	return nil
}

func (x *ViewParticipantData) GetCreatedInSubviewArchivedInCore() []string {
	if x != nil {
		return x.CreatedInSubviewArchivedInCore
	}
	return nil
}

func (x *ViewParticipantData) GetResolvedKeys() []*v01.ViewParticipantData_ResolvedKey {
	if x != nil {
		return x.ResolvedKeys
	}
	return nil
}

func (x *ViewParticipantData) GetActionDescription() *v2.ActionDescription {
	if x != nil {
		return x.ActionDescription
	}
	return nil
}

func (x *ViewParticipantData) GetRollbackContext() *v01.ViewParticipantData_RollbackContext {
	if x != nil {
		return x.RollbackContext
	}
	return nil
}

var File_com_digitalasset_canton_protocol_v3_participant_transaction_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x05, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x61, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x61, 0x6c, 0x74,
	0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x0a, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x72, 0x65, 0x12, 0x4b, 0x0a, 0x23, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x1e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x53, 0x75, 0x62, 0x76,
	0x69, 0x65, 0x77, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x43, 0x6f, 0x72,
	0x65, 0x12, 0x69, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x65, 0x0a, 0x12,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x73, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x30, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescData = file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_goTypes = []interface{}{
	(*ViewParticipantData)(nil),                     // 0: com.digitalasset.canton.protocol.v3.ViewParticipantData
	(*v0.Salt)(nil),                                 // 1: com.digitalasset.canton.crypto.v0.Salt
	(*v1.InputContract)(nil),                        // 2: com.digitalasset.canton.protocol.v1.InputContract
	(*v1.CreatedContract)(nil),                      // 3: com.digitalasset.canton.protocol.v1.CreatedContract
	(*v01.ViewParticipantData_ResolvedKey)(nil),     // 4: com.digitalasset.canton.protocol.v0.ViewParticipantData.ResolvedKey
	(*v2.ActionDescription)(nil),                    // 5: com.digitalasset.canton.protocol.v2.ActionDescription
	(*v01.ViewParticipantData_RollbackContext)(nil), // 6: com.digitalasset.canton.protocol.v0.ViewParticipantData.RollbackContext
}
var file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.protocol.v3.ViewParticipantData.salt:type_name -> com.digitalasset.canton.crypto.v0.Salt
	2, // 1: com.digitalasset.canton.protocol.v3.ViewParticipantData.core_inputs:type_name -> com.digitalasset.canton.protocol.v1.InputContract
	3, // 2: com.digitalasset.canton.protocol.v3.ViewParticipantData.created_core:type_name -> com.digitalasset.canton.protocol.v1.CreatedContract
	4, // 3: com.digitalasset.canton.protocol.v3.ViewParticipantData.resolved_keys:type_name -> com.digitalasset.canton.protocol.v0.ViewParticipantData.ResolvedKey
	5, // 4: com.digitalasset.canton.protocol.v3.ViewParticipantData.action_description:type_name -> com.digitalasset.canton.protocol.v2.ActionDescription
	6, // 5: com.digitalasset.canton.protocol.v3.ViewParticipantData.rollback_context:type_name -> com.digitalasset.canton.protocol.v0.ViewParticipantData.RollbackContext
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_init() }
func file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_init() {
	if File_com_digitalasset_canton_protocol_v3_participant_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewParticipantData); i {
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
			RawDescriptor: file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v3_participant_transaction_proto = out.File
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v3_participant_transaction_proto_depIdxs = nil
}
