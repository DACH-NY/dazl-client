// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/protocol/v3/synchronization.proto

package v3

import (
	v0 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v0"
	v1 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v1"
	v2 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v2"
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
	SomeEnvelopeContent isEnvelopeContent_SomeEnvelopeContent `protobuf_oneof:"some_envelope_content"`
}

func (x *EnvelopeContent) Reset() {
	*x = EnvelopeContent{}
	mi := &file_com_digitalasset_canton_protocol_v3_synchronization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnvelopeContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvelopeContent) ProtoMessage() {}

func (x *EnvelopeContent) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v3_synchronization_proto_msgTypes[0]
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
	return file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescGZIP(), []int{0}
}

func (m *EnvelopeContent) GetSomeEnvelopeContent() isEnvelopeContent_SomeEnvelopeContent {
	if m != nil {
		return m.SomeEnvelopeContent
	}
	return nil
}

func (x *EnvelopeContent) GetInformeeMessage() *v1.InformeeMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_InformeeMessage); ok {
		return x.InformeeMessage
	}
	return nil
}

func (x *EnvelopeContent) GetSignedMessage() *v0.SignedProtocolMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_SignedMessage); ok {
		return x.SignedMessage
	}
	return nil
}

func (x *EnvelopeContent) GetEncryptedViewMessage() *v2.EncryptedViewMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_EncryptedViewMessage); ok {
		return x.EncryptedViewMessage
	}
	return nil
}

func (x *EnvelopeContent) GetDomainTopologyTransactionMessage() *v1.DomainTopologyTransactionMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_DomainTopologyTransactionMessage); ok {
		return x.DomainTopologyTransactionMessage
	}
	return nil
}

func (x *EnvelopeContent) GetTransferOutMediatorMessage() *v1.TransferOutMediatorMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_TransferOutMediatorMessage); ok {
		return x.TransferOutMediatorMessage
	}
	return nil
}

func (x *EnvelopeContent) GetTransferInMediatorMessage() *v1.TransferInMediatorMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_TransferInMediatorMessage); ok {
		return x.TransferInMediatorMessage
	}
	return nil
}

func (x *EnvelopeContent) GetRootHashMessage() *v0.RootHashMessage {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RootHashMessage); ok {
		return x.RootHashMessage
	}
	return nil
}

func (x *EnvelopeContent) GetRegisterTopologyTransactionRequest() *v0.RegisterTopologyTransactionRequest {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RegisterTopologyTransactionRequest); ok {
		return x.RegisterTopologyTransactionRequest
	}
	return nil
}

func (x *EnvelopeContent) GetRegisterTopologyTransactionResponse() *v1.RegisterTopologyTransactionResponse {
	if x, ok := x.GetSomeEnvelopeContent().(*EnvelopeContent_RegisterTopologyTransactionResponse); ok {
		return x.RegisterTopologyTransactionResponse
	}
	return nil
}

type isEnvelopeContent_SomeEnvelopeContent interface {
	isEnvelopeContent_SomeEnvelopeContent()
}

type EnvelopeContent_InformeeMessage struct {
	InformeeMessage *v1.InformeeMessage `protobuf:"bytes,1,opt,name=informee_message,json=informeeMessage,proto3,oneof"`
}

type EnvelopeContent_SignedMessage struct {
	SignedMessage *v0.SignedProtocolMessage `protobuf:"bytes,2,opt,name=signed_message,json=signedMessage,proto3,oneof"`
}

type EnvelopeContent_EncryptedViewMessage struct {
	EncryptedViewMessage *v2.EncryptedViewMessage `protobuf:"bytes,3,opt,name=encrypted_view_message,json=encryptedViewMessage,proto3,oneof"`
}

type EnvelopeContent_DomainTopologyTransactionMessage struct {
	DomainTopologyTransactionMessage *v1.DomainTopologyTransactionMessage `protobuf:"bytes,5,opt,name=domain_topology_transaction_message,json=domainTopologyTransactionMessage,proto3,oneof"`
}

type EnvelopeContent_TransferOutMediatorMessage struct {
	TransferOutMediatorMessage *v1.TransferOutMediatorMessage `protobuf:"bytes,6,opt,name=transfer_out_mediator_message,json=transferOutMediatorMessage,proto3,oneof"`
}

type EnvelopeContent_TransferInMediatorMessage struct {
	TransferInMediatorMessage *v1.TransferInMediatorMessage `protobuf:"bytes,7,opt,name=transfer_in_mediator_message,json=transferInMediatorMessage,proto3,oneof"`
}

type EnvelopeContent_RootHashMessage struct {
	RootHashMessage *v0.RootHashMessage `protobuf:"bytes,8,opt,name=root_hash_message,json=rootHashMessage,proto3,oneof"`
}

type EnvelopeContent_RegisterTopologyTransactionRequest struct {
	RegisterTopologyTransactionRequest *v0.RegisterTopologyTransactionRequest `protobuf:"bytes,9,opt,name=register_topology_transaction_request,json=registerTopologyTransactionRequest,proto3,oneof"`
}

type EnvelopeContent_RegisterTopologyTransactionResponse struct {
	RegisterTopologyTransactionResponse *v1.RegisterTopologyTransactionResponse `protobuf:"bytes,10,opt,name=register_topology_transaction_response,json=registerTopologyTransactionResponse,proto3,oneof"`
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

var File_com_digitalasset_canton_protocol_v3_synchronization_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33,
	0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x30, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x09, 0x0a,
	0x0f, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x61, 0x0a, 0x10, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x30, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x71, 0x0a, 0x16, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x56, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x23,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e,
	0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x1c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x49, 0x6e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x62, 0x0a, 0x11, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30,
	0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x25, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x22,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x9f, 0x01, 0x0a, 0x26, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x23, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x54, 0x5a,
	0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescData = file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v3_synchronization_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_digitalasset_canton_protocol_v3_synchronization_proto_goTypes = []any{
	(*EnvelopeContent)(nil),                        // 0: com.digitalasset.canton.protocol.v3.EnvelopeContent
	(*v1.InformeeMessage)(nil),                     // 1: com.digitalasset.canton.protocol.v1.InformeeMessage
	(*v0.SignedProtocolMessage)(nil),               // 2: com.digitalasset.canton.protocol.v0.SignedProtocolMessage
	(*v2.EncryptedViewMessage)(nil),                // 3: com.digitalasset.canton.protocol.v2.EncryptedViewMessage
	(*v1.DomainTopologyTransactionMessage)(nil),    // 4: com.digitalasset.canton.protocol.v1.DomainTopologyTransactionMessage
	(*v1.TransferOutMediatorMessage)(nil),          // 5: com.digitalasset.canton.protocol.v1.TransferOutMediatorMessage
	(*v1.TransferInMediatorMessage)(nil),           // 6: com.digitalasset.canton.protocol.v1.TransferInMediatorMessage
	(*v0.RootHashMessage)(nil),                     // 7: com.digitalasset.canton.protocol.v0.RootHashMessage
	(*v0.RegisterTopologyTransactionRequest)(nil),  // 8: com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionRequest
	(*v1.RegisterTopologyTransactionResponse)(nil), // 9: com.digitalasset.canton.protocol.v1.RegisterTopologyTransactionResponse
}
var file_com_digitalasset_canton_protocol_v3_synchronization_proto_depIdxs = []int32{
	1, // 0: com.digitalasset.canton.protocol.v3.EnvelopeContent.informee_message:type_name -> com.digitalasset.canton.protocol.v1.InformeeMessage
	2, // 1: com.digitalasset.canton.protocol.v3.EnvelopeContent.signed_message:type_name -> com.digitalasset.canton.protocol.v0.SignedProtocolMessage
	3, // 2: com.digitalasset.canton.protocol.v3.EnvelopeContent.encrypted_view_message:type_name -> com.digitalasset.canton.protocol.v2.EncryptedViewMessage
	4, // 3: com.digitalasset.canton.protocol.v3.EnvelopeContent.domain_topology_transaction_message:type_name -> com.digitalasset.canton.protocol.v1.DomainTopologyTransactionMessage
	5, // 4: com.digitalasset.canton.protocol.v3.EnvelopeContent.transfer_out_mediator_message:type_name -> com.digitalasset.canton.protocol.v1.TransferOutMediatorMessage
	6, // 5: com.digitalasset.canton.protocol.v3.EnvelopeContent.transfer_in_mediator_message:type_name -> com.digitalasset.canton.protocol.v1.TransferInMediatorMessage
	7, // 6: com.digitalasset.canton.protocol.v3.EnvelopeContent.root_hash_message:type_name -> com.digitalasset.canton.protocol.v0.RootHashMessage
	8, // 7: com.digitalasset.canton.protocol.v3.EnvelopeContent.register_topology_transaction_request:type_name -> com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionRequest
	9, // 8: com.digitalasset.canton.protocol.v3.EnvelopeContent.register_topology_transaction_response:type_name -> com.digitalasset.canton.protocol.v1.RegisterTopologyTransactionResponse
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v3_synchronization_proto_init() }
func file_com_digitalasset_canton_protocol_v3_synchronization_proto_init() {
	if File_com_digitalasset_canton_protocol_v3_synchronization_proto != nil {
		return
	}
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_msgTypes[0].OneofWrappers = []any{
		(*EnvelopeContent_InformeeMessage)(nil),
		(*EnvelopeContent_SignedMessage)(nil),
		(*EnvelopeContent_EncryptedViewMessage)(nil),
		(*EnvelopeContent_DomainTopologyTransactionMessage)(nil),
		(*EnvelopeContent_TransferOutMediatorMessage)(nil),
		(*EnvelopeContent_TransferInMediatorMessage)(nil),
		(*EnvelopeContent_RootHashMessage)(nil),
		(*EnvelopeContent_RegisterTopologyTransactionRequest)(nil),
		(*EnvelopeContent_RegisterTopologyTransactionResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v3_synchronization_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v3_synchronization_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v3_synchronization_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v3_synchronization_proto = out.File
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v3_synchronization_proto_depIdxs = nil
}
