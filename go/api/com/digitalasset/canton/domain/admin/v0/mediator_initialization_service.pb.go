// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: com/digitalasset/canton/domain/admin/v0/mediator_initialization_service.proto

package v0

import (
	v02 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/crypto/v0"
	v01 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0"
	v0 "github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type InitializeMediatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the domain identifier
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// the mediator identifier
	MediatorId string `protobuf:"bytes,2,opt,name=mediator_id,json=mediatorId,proto3" json:"mediator_id,omitempty"`
	// topology state required for startup
	CurrentIdentityState *v0.TopologyTransactions `protobuf:"bytes,3,opt,name=current_identity_state,json=currentIdentityState,proto3" json:"current_identity_state,omitempty"`
	// parameters for the domain (includes the protocol version which needs to match the protocol version the domain
	// manager is running)
	DomainParameters *v0.StaticDomainParameters `protobuf:"bytes,4,opt,name=domain_parameters,json=domainParameters,proto3" json:"domain_parameters,omitempty"`
	// how should the member connect to the domain sequencer
	SequencerConnection *v01.SequencerConnection `protobuf:"bytes,5,opt,name=sequencer_connection,json=sequencerConnection,proto3" json:"sequencer_connection,omitempty"`
	// Optional fingerprint of the signing key the mediator should use. This key should already exist and have been
	// authorized during domain bootstrap
	SigningKeyFingerprint *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=signing_key_fingerprint,json=signingKeyFingerprint,proto3" json:"signing_key_fingerprint,omitempty"`
}

func (x *InitializeMediatorRequest) Reset() {
	*x = InitializeMediatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeMediatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeMediatorRequest) ProtoMessage() {}

func (x *InitializeMediatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeMediatorRequest.ProtoReflect.Descriptor instead.
func (*InitializeMediatorRequest) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeMediatorRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *InitializeMediatorRequest) GetMediatorId() string {
	if x != nil {
		return x.MediatorId
	}
	return ""
}

func (x *InitializeMediatorRequest) GetCurrentIdentityState() *v0.TopologyTransactions {
	if x != nil {
		return x.CurrentIdentityState
	}
	return nil
}

func (x *InitializeMediatorRequest) GetDomainParameters() *v0.StaticDomainParameters {
	if x != nil {
		return x.DomainParameters
	}
	return nil
}

func (x *InitializeMediatorRequest) GetSequencerConnection() *v01.SequencerConnection {
	if x != nil {
		return x.SequencerConnection
	}
	return nil
}

func (x *InitializeMediatorRequest) GetSigningKeyFingerprint() *wrapperspb.StringValue {
	if x != nil {
		return x.SigningKeyFingerprint
	}
	return nil
}

type InitializeMediatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*InitializeMediatorResponse_Success_
	//	*InitializeMediatorResponse_Failure_
	Value isInitializeMediatorResponse_Value `protobuf_oneof:"value"`
}

func (x *InitializeMediatorResponse) Reset() {
	*x = InitializeMediatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeMediatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeMediatorResponse) ProtoMessage() {}

func (x *InitializeMediatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeMediatorResponse.ProtoReflect.Descriptor instead.
func (*InitializeMediatorResponse) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescGZIP(), []int{1}
}

func (m *InitializeMediatorResponse) GetValue() isInitializeMediatorResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *InitializeMediatorResponse) GetSuccess() *InitializeMediatorResponse_Success {
	if x, ok := x.GetValue().(*InitializeMediatorResponse_Success_); ok {
		return x.Success
	}
	return nil
}

func (x *InitializeMediatorResponse) GetFailure() *InitializeMediatorResponse_Failure {
	if x, ok := x.GetValue().(*InitializeMediatorResponse_Failure_); ok {
		return x.Failure
	}
	return nil
}

type isInitializeMediatorResponse_Value interface {
	isInitializeMediatorResponse_Value()
}

type InitializeMediatorResponse_Success_ struct {
	Success *InitializeMediatorResponse_Success `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type InitializeMediatorResponse_Failure_ struct {
	Failure *InitializeMediatorResponse_Failure `protobuf:"bytes,2,opt,name=failure,proto3,oneof"`
}

func (*InitializeMediatorResponse_Success_) isInitializeMediatorResponse_Value() {}

func (*InitializeMediatorResponse_Failure_) isInitializeMediatorResponse_Value() {}

type InitializeMediatorResponse_Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current signing key
	MediatorKey *v02.SigningPublicKey `protobuf:"bytes,1,opt,name=mediator_key,json=mediatorKey,proto3" json:"mediator_key,omitempty"`
}

func (x *InitializeMediatorResponse_Success) Reset() {
	*x = InitializeMediatorResponse_Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeMediatorResponse_Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeMediatorResponse_Success) ProtoMessage() {}

func (x *InitializeMediatorResponse_Success) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeMediatorResponse_Success.ProtoReflect.Descriptor instead.
func (*InitializeMediatorResponse_Success) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *InitializeMediatorResponse_Success) GetMediatorKey() *v02.SigningPublicKey {
	if x != nil {
		return x.MediatorKey
	}
	return nil
}

type InitializeMediatorResponse_Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reason that can be logged
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *InitializeMediatorResponse_Failure) Reset() {
	*x = InitializeMediatorResponse_Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeMediatorResponse_Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeMediatorResponse_Failure) ProtoMessage() {}

func (x *InitializeMediatorResponse_Failure) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeMediatorResponse_Failure.ProtoReflect.Descriptor instead.
func (*InitializeMediatorResponse_Failure) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *InitializeMediatorResponse_Failure) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x19, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x6f, 0x0a, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x6d, 0x0a, 0x14, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54,
	0x0a, 0x17, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x15, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x22, 0xfd, 0x02, 0x0a, 0x1a, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x67, 0x0a, 0x07,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x1a, 0x61, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x56, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x1a, 0x21, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0xb7, 0x01, 0x0a, 0x1d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x30, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x58,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x37, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescData = file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDesc
)

func file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescData)
	})
	return file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDescData
}

var file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_goTypes = []interface{}{
	(*InitializeMediatorRequest)(nil),          // 0: com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest
	(*InitializeMediatorResponse)(nil),         // 1: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse
	(*InitializeMediatorResponse_Success)(nil), // 2: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.Success
	(*InitializeMediatorResponse_Failure)(nil), // 3: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.Failure
	(*v0.TopologyTransactions)(nil),            // 4: com.digitalasset.canton.protocol.v0.TopologyTransactions
	(*v0.StaticDomainParameters)(nil),          // 5: com.digitalasset.canton.protocol.v0.StaticDomainParameters
	(*v01.SequencerConnection)(nil),            // 6: com.digitalasset.canton.domain.api.v0.SequencerConnection
	(*wrapperspb.StringValue)(nil),             // 7: google.protobuf.StringValue
	(*v02.SigningPublicKey)(nil),               // 8: com.digitalasset.canton.crypto.v0.SigningPublicKey
}
var file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_depIdxs = []int32{
	4, // 0: com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest.current_identity_state:type_name -> com.digitalasset.canton.protocol.v0.TopologyTransactions
	5, // 1: com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest.domain_parameters:type_name -> com.digitalasset.canton.protocol.v0.StaticDomainParameters
	6, // 2: com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest.sequencer_connection:type_name -> com.digitalasset.canton.domain.api.v0.SequencerConnection
	7, // 3: com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest.signing_key_fingerprint:type_name -> google.protobuf.StringValue
	2, // 4: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.success:type_name -> com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.Success
	3, // 5: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.failure:type_name -> com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.Failure
	8, // 6: com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse.Success.mediator_key:type_name -> com.digitalasset.canton.crypto.v0.SigningPublicKey
	0, // 7: com.digitalasset.canton.domain.admin.v0.MediatorInitializationService.Initialize:input_type -> com.digitalasset.canton.domain.admin.v0.InitializeMediatorRequest
	1, // 8: com.digitalasset.canton.domain.admin.v0.MediatorInitializationService.Initialize:output_type -> com.digitalasset.canton.domain.admin.v0.InitializeMediatorResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_init()
}
func file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_init() {
	if File_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeMediatorRequest); i {
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
		file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeMediatorResponse); i {
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
		file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeMediatorResponse_Success); i {
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
		file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeMediatorResponse_Failure); i {
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
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InitializeMediatorResponse_Success_)(nil),
		(*InitializeMediatorResponse_Failure_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto = out.File
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_rawDesc = nil
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_goTypes = nil
	file_com_digitalasset_canton_domain_admin_v0_mediator_initialization_service_proto_depIdxs = nil
}
