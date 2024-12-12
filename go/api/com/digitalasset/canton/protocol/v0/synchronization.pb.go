// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/protocol/v0/synchronization.proto

package v0

import (
	v0 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/crypto/v0"
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

type SignedProtocolMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature *v0.Signature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are assignable to SomeSignedProtocolMessage:
	//	*SignedProtocolMessage_MediatorResponse
	//	*SignedProtocolMessage_TransactionResult
	//	*SignedProtocolMessage_MalformedMediatorRequestResult
	//	*SignedProtocolMessage_TransferResult
	//	*SignedProtocolMessage_AcsCommitment
	SomeSignedProtocolMessage isSignedProtocolMessage_SomeSignedProtocolMessage `protobuf_oneof:"some_signed_protocol_message"`
}

func (x *SignedProtocolMessage) Reset() {
	*x = SignedProtocolMessage{}
	mi := &file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignedProtocolMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedProtocolMessage) ProtoMessage() {}

func (x *SignedProtocolMessage) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedProtocolMessage.ProtoReflect.Descriptor instead.
func (*SignedProtocolMessage) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescGZIP(), []int{0}
}

func (x *SignedProtocolMessage) GetSignature() *v0.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *SignedProtocolMessage) GetSomeSignedProtocolMessage() isSignedProtocolMessage_SomeSignedProtocolMessage {
	if m != nil {
		return m.SomeSignedProtocolMessage
	}
	return nil
}

func (x *SignedProtocolMessage) GetMediatorResponse() []byte {
	if x, ok := x.GetSomeSignedProtocolMessage().(*SignedProtocolMessage_MediatorResponse); ok {
		return x.MediatorResponse
	}
	return nil
}

func (x *SignedProtocolMessage) GetTransactionResult() []byte {
	if x, ok := x.GetSomeSignedProtocolMessage().(*SignedProtocolMessage_TransactionResult); ok {
		return x.TransactionResult
	}
	return nil
}

func (x *SignedProtocolMessage) GetMalformedMediatorRequestResult() []byte {
	if x, ok := x.GetSomeSignedProtocolMessage().(*SignedProtocolMessage_MalformedMediatorRequestResult); ok {
		return x.MalformedMediatorRequestResult
	}
	return nil
}

func (x *SignedProtocolMessage) GetTransferResult() []byte {
	if x, ok := x.GetSomeSignedProtocolMessage().(*SignedProtocolMessage_TransferResult); ok {
		return x.TransferResult
	}
	return nil
}

func (x *SignedProtocolMessage) GetAcsCommitment() []byte {
	if x, ok := x.GetSomeSignedProtocolMessage().(*SignedProtocolMessage_AcsCommitment); ok {
		return x.AcsCommitment
	}
	return nil
}

type isSignedProtocolMessage_SomeSignedProtocolMessage interface {
	isSignedProtocolMessage_SomeSignedProtocolMessage()
}

type SignedProtocolMessage_MediatorResponse struct {
	MediatorResponse []byte `protobuf:"bytes,2,opt,name=mediator_response,json=mediatorResponse,proto3,oneof"`
}

type SignedProtocolMessage_TransactionResult struct {
	TransactionResult []byte `protobuf:"bytes,3,opt,name=transaction_result,json=transactionResult,proto3,oneof"`
}

type SignedProtocolMessage_MalformedMediatorRequestResult struct {
	MalformedMediatorRequestResult []byte `protobuf:"bytes,4,opt,name=malformed_mediator_request_result,json=malformedMediatorRequestResult,proto3,oneof"`
}

type SignedProtocolMessage_TransferResult struct {
	TransferResult []byte `protobuf:"bytes,5,opt,name=transfer_result,json=transferResult,proto3,oneof"`
}

type SignedProtocolMessage_AcsCommitment struct {
	AcsCommitment []byte `protobuf:"bytes,6,opt,name=acs_commitment,json=acsCommitment,proto3,oneof"`
}

func (*SignedProtocolMessage_MediatorResponse) isSignedProtocolMessage_SomeSignedProtocolMessage() {}

func (*SignedProtocolMessage_TransactionResult) isSignedProtocolMessage_SomeSignedProtocolMessage() {}

func (*SignedProtocolMessage_MalformedMediatorRequestResult) isSignedProtocolMessage_SomeSignedProtocolMessage() {
}

func (*SignedProtocolMessage_TransferResult) isSignedProtocolMessage_SomeSignedProtocolMessage() {}

func (*SignedProtocolMessage_AcsCommitment) isSignedProtocolMessage_SomeSignedProtocolMessage() {}

type EnvelopeContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SomeEnvelopeContent:
	//	*EnvelopeContent_InformeeMessage
	//	*EnvelopeContent_SignedMessage
	//	*EnvelopeContent_EncryptedViewMessage
	//	*EnvelopeContent_DomainTopologyTransactionMessage
	//	*EnvelopeContent_TransferOutMediatorMessage
	//	*EnvelopeContent_TransferInMediatorMessage
	//	*EnvelopeContent_RootHashMessage
	//	*EnvelopeContent_RegisterTopologyTransactionRequest
	//	*EnvelopeContent_RegisterTopologyTransactionResponse
	//	*EnvelopeContent_CausalityMessage
	SomeEnvelopeContent isEnvelopeContent_SomeEnvelopeContent `protobuf_oneof:"some_envelope_content"`
}

func (x *EnvelopeContent) Reset() {
	*x = EnvelopeContent{}
	mi := &file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnvelopeContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvelopeContent) ProtoMessage() {}

func (x *EnvelopeContent) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvelopeContent.ProtoReflect.Descriptor instead.
func (*EnvelopeContent) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescGZIP(), []int{1}
}

func (m *EnvelopeContent) GetSomeEnvelopeContent() isEnvelopeContent_SomeEnvelopeContent {
	if m != nil {
		return m.SomeEnvelopeContent
	}
	return nil
}

func (x *EnvelopeContent) GetInformeeMessage() *InformeeMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_InformeeMessage); ok {
		return x.InformeeMessage
	}
	return nil
}

func (x *EnvelopeContent) GetSignedMessage() *SignedProtocolMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_SignedMessage); ok {
		return x.SignedMessage
	}
	return nil
}

func (x *EnvelopeContent) GetEncryptedViewMessage() *EncryptedViewMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_EncryptedViewMessage); ok {
		return x.EncryptedViewMessage
	}
	return nil
}

func (x *EnvelopeContent) GetDomainTopologyTransactionMessage() *DomainTopologyTransactionMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_DomainTopologyTransactionMessage); ok {
		return x.DomainTopologyTransactionMessage
	}
	return nil
}

func (x *EnvelopeContent) GetTransferOutMediatorMessage() *TransferOutMediatorMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_TransferOutMediatorMessage); ok {
		return x.TransferOutMediatorMessage
	}
	return nil
}

func (x *EnvelopeContent) GetTransferInMediatorMessage() *TransferInMediatorMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_TransferInMediatorMessage); ok {
		return x.TransferInMediatorMessage
	}
	return nil
}

func (x *EnvelopeContent) GetRootHashMessage() *RootHashMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RootHashMessage); ok {
		return x.RootHashMessage
	}
	return nil
}

func (x *EnvelopeContent) GetRegisterTopologyTransactionRequest() *RegisterTopologyTransactionRequest {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RegisterTopologyTransactionRequest); ok {
		return x.RegisterTopologyTransactionRequest
	}
	return nil
}

func (x *EnvelopeContent) GetRegisterTopologyTransactionResponse() *RegisterTopologyTransactionResponse {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RegisterTopologyTransactionResponse); ok {
		return x.RegisterTopologyTransactionResponse
	}
	return nil
}

func (x *EnvelopeContent) GetCausalityMessage() *CausalityMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_CausalityMessage); ok {
		return x.CausalityMessage
	}
	return nil
}

type isEnvelopeContent_SomeEnvelopeContent interface {
	isEnvelopeContent_SomeEnvelopeContent()
}

type EnvelopeContent_InformeeMessage struct {
	InformeeMessage *InformeeMessage `protobuf:"bytes,1,opt,name=informee_message,json=informeeMessage,proto3,oneof"`
}

type EnvelopeContent_SignedMessage struct {
	SignedMessage *SignedProtocolMessage `protobuf:"bytes,2,opt,name=signed_message,json=signedMessage,proto3,oneof"`
}

type EnvelopeContent_EncryptedViewMessage struct {
	EncryptedViewMessage *EncryptedViewMessage `protobuf:"bytes,3,opt,name=encrypted_view_message,json=encryptedViewMessage,proto3,oneof"`
}

type EnvelopeContent_DomainTopologyTransactionMessage struct {
	DomainTopologyTransactionMessage *DomainTopologyTransactionMessage `protobuf:"bytes,5,opt,name=domain_topology_transaction_message,json=domainTopologyTransactionMessage,proto3,oneof"`
}

type EnvelopeContent_TransferOutMediatorMessage struct {
	TransferOutMediatorMessage *TransferOutMediatorMessage `protobuf:"bytes,6,opt,name=transfer_out_mediator_message,json=transferOutMediatorMessage,proto3,oneof"`
}

type EnvelopeContent_TransferInMediatorMessage struct {
	TransferInMediatorMessage *TransferInMediatorMessage `protobuf:"bytes,7,opt,name=transfer_in_mediator_message,json=transferInMediatorMessage,proto3,oneof"`
}

type EnvelopeContent_RootHashMessage struct {
	RootHashMessage *RootHashMessage `protobuf:"bytes,8,opt,name=root_hash_message,json=rootHashMessage,proto3,oneof"`
}

type EnvelopeContent_RegisterTopologyTransactionRequest struct {
	RegisterTopologyTransactionRequest *RegisterTopologyTransactionRequest `protobuf:"bytes,9,opt,name=register_topology_transaction_request,json=registerTopologyTransactionRequest,proto3,oneof"`
}

type EnvelopeContent_RegisterTopologyTransactionResponse struct {
	RegisterTopologyTransactionResponse *RegisterTopologyTransactionResponse `protobuf:"bytes,10,opt,name=register_topology_transaction_response,json=registerTopologyTransactionResponse,proto3,oneof"`
}

type EnvelopeContent_CausalityMessage struct {
	CausalityMessage *CausalityMessage `protobuf:"bytes,11,opt,name=causality_message,json=causalityMessage,proto3,oneof"`
}

func (*EnvelopeContent_InformeeMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_SignedMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_EncryptedViewMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_DomainTopologyTransactionMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_TransferOutMediatorMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_TransferInMediatorMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_RootHashMessage) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_RegisterTopologyTransactionRequest) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_RegisterTopologyTransactionResponse) isEnvelopeContent_SomeEnvelopeContent() {}

func (*EnvelopeContent_CausalityMessage) isEnvelopeContent_SomeEnvelopeContent() {}

var File_com_digitalasset_canton_protocol_v0_synchronization_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2f, 0x76, 0x30, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a,
	0x15, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2d, 0x0a, 0x11, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x10, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x4b, 0x0a, 0x21, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x1e, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x29, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x61, 0x63,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x63, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x1e, 0x0a, 0x1c, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x94, 0x0a, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x10, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x71, 0x0a, 0x16, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56,
	0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x14, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x23, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x1d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x19, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x25, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x22, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x9f, 0x01, 0x0a, 0x26, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x23, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x63,
	0x61, 0x75, 0x73, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x61, 0x75,
	0x73, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x63, 0x61, 0x75, 0x73, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x17, 0x0a, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescData = file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_digitalasset_canton_protocol_v0_synchronization_proto_goTypes = []any{
	(*SignedProtocolMessage)(nil),               // 0: com.digitalasset.canton.protocol.v0.SignedProtocolMessage
	(*EnvelopeContent)(nil),                     // 1: com.digitalasset.canton.protocol.v0.EnvelopeContent
	(*v0.Signature)(nil),                        // 2: com.digitalasset.canton.crypto.v0.Signature
	(*InformeeMessage)(nil),                     // 3: com.digitalasset.canton.protocol.v0.InformeeMessage
	(*EncryptedViewMessage)(nil),                // 4: com.digitalasset.canton.protocol.v0.EncryptedViewMessage
	(*DomainTopologyTransactionMessage)(nil),    // 5: com.digitalasset.canton.protocol.v0.DomainTopologyTransactionMessage
	(*TransferOutMediatorMessage)(nil),          // 6: com.digitalasset.canton.protocol.v0.TransferOutMediatorMessage
	(*TransferInMediatorMessage)(nil),           // 7: com.digitalasset.canton.protocol.v0.TransferInMediatorMessage
	(*RootHashMessage)(nil),                     // 8: com.digitalasset.canton.protocol.v0.RootHashMessage
	(*RegisterTopologyTransactionRequest)(nil),  // 9: com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionRequest
	(*RegisterTopologyTransactionResponse)(nil), // 10: com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionResponse
	(*CausalityMessage)(nil),                    // 11: com.digitalasset.canton.protocol.v0.CausalityMessage
}
var file_com_digitalasset_canton_protocol_v0_synchronization_proto_depIdxs = []int32{
	2,  // 0: com.digitalasset.canton.protocol.v0.SignedProtocolMessage.signature:type_name -> com.digitalasset.canton.crypto.v0.Signature
	3,  // 1: com.digitalasset.canton.protocol.v0.EnvelopeContent.informee_message:type_name -> com.digitalasset.canton.protocol.v0.InformeeMessage
	0,  // 2: com.digitalasset.canton.protocol.v0.EnvelopeContent.signed_message:type_name -> com.digitalasset.canton.protocol.v0.SignedProtocolMessage
	4,  // 3: com.digitalasset.canton.protocol.v0.EnvelopeContent.encrypted_view_message:type_name -> com.digitalasset.canton.protocol.v0.EncryptedViewMessage
	5,  // 4: com.digitalasset.canton.protocol.v0.EnvelopeContent.domain_topology_transaction_message:type_name -> com.digitalasset.canton.protocol.v0.DomainTopologyTransactionMessage
	6,  // 5: com.digitalasset.canton.protocol.v0.EnvelopeContent.transfer_out_mediator_message:type_name -> com.digitalasset.canton.protocol.v0.TransferOutMediatorMessage
	7,  // 6: com.digitalasset.canton.protocol.v0.EnvelopeContent.transfer_in_mediator_message:type_name -> com.digitalasset.canton.protocol.v0.TransferInMediatorMessage
	8,  // 7: com.digitalasset.canton.protocol.v0.EnvelopeContent.root_hash_message:type_name -> com.digitalasset.canton.protocol.v0.RootHashMessage
	9,  // 8: com.digitalasset.canton.protocol.v0.EnvelopeContent.register_topology_transaction_request:type_name -> com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionRequest
	10, // 9: com.digitalasset.canton.protocol.v0.EnvelopeContent.register_topology_transaction_response:type_name -> com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionResponse
	11, // 10: com.digitalasset.canton.protocol.v0.EnvelopeContent.causality_message:type_name -> com.digitalasset.canton.protocol.v0.CausalityMessage
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v0_synchronization_proto_init() }
func file_com_digitalasset_canton_protocol_v0_synchronization_proto_init() {
	if File_com_digitalasset_canton_protocol_v0_synchronization_proto != nil {
		return
	}
	file_com_digitalasset_canton_protocol_v0_causality_proto_init()
	file_com_digitalasset_canton_protocol_v0_participant_transaction_proto_init()
	file_com_digitalasset_canton_protocol_v0_participant_transfer_proto_init()
	file_com_digitalasset_canton_protocol_v0_topology_proto_init()
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[0].OneofWrappers = []any{
		(*SignedProtocolMessage_MediatorResponse)(nil),
		(*SignedProtocolMessage_TransactionResult)(nil),
		(*SignedProtocolMessage_MalformedMediatorRequestResult)(nil),
		(*SignedProtocolMessage_TransferResult)(nil),
		(*SignedProtocolMessage_AcsCommitment)(nil),
	}
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes[1].OneofWrappers = []any{
		(*EnvelopeContent_InformeeMessage)(nil),
		(*EnvelopeContent_SignedMessage)(nil),
		(*EnvelopeContent_EncryptedViewMessage)(nil),
		(*EnvelopeContent_DomainTopologyTransactionMessage)(nil),
		(*EnvelopeContent_TransferOutMediatorMessage)(nil),
		(*EnvelopeContent_TransferInMediatorMessage)(nil),
		(*EnvelopeContent_RootHashMessage)(nil),
		(*EnvelopeContent_RegisterTopologyTransactionRequest)(nil),
		(*EnvelopeContent_RegisterTopologyTransactionResponse)(nil),
		(*EnvelopeContent_CausalityMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v0_synchronization_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v0_synchronization_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v0_synchronization_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v0_synchronization_proto = out.File
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v0_synchronization_proto_depIdxs = nil
}
