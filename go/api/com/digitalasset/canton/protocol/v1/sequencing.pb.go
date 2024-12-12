// Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: com/digitalasset/canton/protocol/v1/sequencing.proto

package v1

import (
	v0 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/crypto/v0"
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

type StaticDomainParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueContractKeys           bool                     `protobuf:"varint,1,opt,name=unique_contract_keys,json=uniqueContractKeys,proto3" json:"unique_contract_keys,omitempty"`
	RequiredSigningKeySchemes    []v0.SigningKeyScheme    `protobuf:"varint,2,rep,packed,name=required_signing_key_schemes,json=requiredSigningKeySchemes,proto3,enum=com.digitalasset.canton.crypto.v0.SigningKeyScheme" json:"required_signing_key_schemes,omitempty"`
	RequiredEncryptionKeySchemes []v0.EncryptionKeyScheme `protobuf:"varint,3,rep,packed,name=required_encryption_key_schemes,json=requiredEncryptionKeySchemes,proto3,enum=com.digitalasset.canton.crypto.v0.EncryptionKeyScheme" json:"required_encryption_key_schemes,omitempty"`
	RequiredSymmetricKeySchemes  []v0.SymmetricKeyScheme  `protobuf:"varint,4,rep,packed,name=required_symmetric_key_schemes,json=requiredSymmetricKeySchemes,proto3,enum=com.digitalasset.canton.crypto.v0.SymmetricKeyScheme" json:"required_symmetric_key_schemes,omitempty"`
	RequiredHashAlgorithms       []v0.HashAlgorithm       `protobuf:"varint,5,rep,packed,name=required_hash_algorithms,json=requiredHashAlgorithms,proto3,enum=com.digitalasset.canton.crypto.v0.HashAlgorithm" json:"required_hash_algorithms,omitempty"`
	RequiredCryptoKeyFormats     []v0.CryptoKeyFormat     `protobuf:"varint,6,rep,packed,name=required_crypto_key_formats,json=requiredCryptoKeyFormats,proto3,enum=com.digitalasset.canton.crypto.v0.CryptoKeyFormat" json:"required_crypto_key_formats,omitempty"`
	ProtocolVersion              int32                    `protobuf:"varint,7,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
}

func (x *StaticDomainParameters) Reset() {
	*x = StaticDomainParameters{}
	mi := &file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaticDomainParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticDomainParameters) ProtoMessage() {}

func (x *StaticDomainParameters) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticDomainParameters.ProtoReflect.Descriptor instead.
func (*StaticDomainParameters) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescGZIP(), []int{0}
}

func (x *StaticDomainParameters) GetUniqueContractKeys() bool {
	if x != nil {
		return x.UniqueContractKeys
	}
	return false
}

func (x *StaticDomainParameters) GetRequiredSigningKeySchemes() []v0.SigningKeyScheme {
	if x != nil {
		return x.RequiredSigningKeySchemes
	}
	return nil
}

func (x *StaticDomainParameters) GetRequiredEncryptionKeySchemes() []v0.EncryptionKeyScheme {
	if x != nil {
		return x.RequiredEncryptionKeySchemes
	}
	return nil
}

func (x *StaticDomainParameters) GetRequiredSymmetricKeySchemes() []v0.SymmetricKeyScheme {
	if x != nil {
		return x.RequiredSymmetricKeySchemes
	}
	return nil
}

func (x *StaticDomainParameters) GetRequiredHashAlgorithms() []v0.HashAlgorithm {
	if x != nil {
		return x.RequiredHashAlgorithms
	}
	return nil
}

func (x *StaticDomainParameters) GetRequiredCryptoKeyFormats() []v0.CryptoKeyFormat {
	if x != nil {
		return x.RequiredCryptoKeyFormats
	}
	return nil
}

func (x *StaticDomainParameters) GetProtocolVersion() int32 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

type DynamicDomainParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParticipantResponseTimeout    *durationpb.Duration `protobuf:"bytes,1,opt,name=participant_response_timeout,json=participantResponseTimeout,proto3" json:"participant_response_timeout,omitempty"`
	MediatorReactionTimeout       *durationpb.Duration `protobuf:"bytes,2,opt,name=mediator_reaction_timeout,json=mediatorReactionTimeout,proto3" json:"mediator_reaction_timeout,omitempty"`
	TransferExclusivityTimeout    *durationpb.Duration `protobuf:"bytes,3,opt,name=transfer_exclusivity_timeout,json=transferExclusivityTimeout,proto3" json:"transfer_exclusivity_timeout,omitempty"`
	TopologyChangeDelay           *durationpb.Duration `protobuf:"bytes,4,opt,name=topology_change_delay,json=topologyChangeDelay,proto3" json:"topology_change_delay,omitempty"`
	LedgerTimeRecordTimeTolerance *durationpb.Duration `protobuf:"bytes,5,opt,name=ledger_time_record_time_tolerance,json=ledgerTimeRecordTimeTolerance,proto3" json:"ledger_time_record_time_tolerance,omitempty"`
	ReconciliationInterval        *durationpb.Duration `protobuf:"bytes,6,opt,name=reconciliation_interval,json=reconciliationInterval,proto3" json:"reconciliation_interval,omitempty"`
	MediatorDeduplicationTimeout  *durationpb.Duration `protobuf:"bytes,7,opt,name=mediator_deduplication_timeout,json=mediatorDeduplicationTimeout,proto3" json:"mediator_deduplication_timeout,omitempty"`
	MaxRatePerParticipant         uint32               `protobuf:"varint,8,opt,name=max_rate_per_participant,json=maxRatePerParticipant,proto3" json:"max_rate_per_participant,omitempty"`
	MaxRequestSize                uint32               `protobuf:"varint,9,opt,name=max_request_size,json=maxRequestSize,proto3" json:"max_request_size,omitempty"`
}

func (x *DynamicDomainParameters) Reset() {
	*x = DynamicDomainParameters{}
	mi := &file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DynamicDomainParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicDomainParameters) ProtoMessage() {}

func (x *DynamicDomainParameters) ProtoReflect() protoreflect.Message {
	mi := &file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicDomainParameters.ProtoReflect.Descriptor instead.
func (*DynamicDomainParameters) Descriptor() ([]byte, []int) {
	return file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicDomainParameters) GetParticipantResponseTimeout() *durationpb.Duration {
	if x != nil {
		return x.ParticipantResponseTimeout
	}
	return nil
}

func (x *DynamicDomainParameters) GetMediatorReactionTimeout() *durationpb.Duration {
	if x != nil {
		return x.MediatorReactionTimeout
	}
	return nil
}

func (x *DynamicDomainParameters) GetTransferExclusivityTimeout() *durationpb.Duration {
	if x != nil {
		return x.TransferExclusivityTimeout
	}
	return nil
}

func (x *DynamicDomainParameters) GetTopologyChangeDelay() *durationpb.Duration {
	if x != nil {
		return x.TopologyChangeDelay
	}
	return nil
}

func (x *DynamicDomainParameters) GetLedgerTimeRecordTimeTolerance() *durationpb.Duration {
	if x != nil {
		return x.LedgerTimeRecordTimeTolerance
	}
	return nil
}

func (x *DynamicDomainParameters) GetReconciliationInterval() *durationpb.Duration {
	if x != nil {
		return x.ReconciliationInterval
	}
	return nil
}

func (x *DynamicDomainParameters) GetMediatorDeduplicationTimeout() *durationpb.Duration {
	if x != nil {
		return x.MediatorDeduplicationTimeout
	}
	return nil
}

func (x *DynamicDomainParameters) GetMaxRatePerParticipant() uint32 {
	if x != nil {
		return x.MaxRatePerParticipant
	}
	return 0
}

func (x *DynamicDomainParameters) GetMaxRequestSize() uint32 {
	if x != nil {
		return x.MaxRequestSize
	}
	return 0
}

var File_com_digitalasset_canton_protocol_v1_sequencing_proto protoreflect.FileDescriptor

var file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x30, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x05, 0x0a, 0x16,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x74, 0x0a, 0x1c, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x76, 0x30, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x52, 0x19, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x7d,
	0x0a, 0x1f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x1c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x7a, 0x0a,
	0x1e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x1b, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x6a, 0x0a, 0x18, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x16, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x71, 0x0a, 0x1b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x63, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x30, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x18,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xf6, 0x05, 0x0a, 0x17, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x5b, 0x0a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x1a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x55, 0x0a, 0x19,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x5b, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x4d, 0x0a, 0x15, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x63, 0x0a, 0x21, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x6c, 0x65, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x6c, 0x65, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x5f, 0x0a, 0x1e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1c, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x6d, 0x61, 0x78,
	0x52, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x54, 0x5a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x38, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f,
	0x63, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescOnce sync.Once
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescData = file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDesc
)

func file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescGZIP() []byte {
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescOnce.Do(func() {
		file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescData)
	})
	return file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDescData
}

var file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_digitalasset_canton_protocol_v1_sequencing_proto_goTypes = []any{
	(*StaticDomainParameters)(nil),  // 0: com.digitalasset.canton.protocol.v1.StaticDomainParameters
	(*DynamicDomainParameters)(nil), // 1: com.digitalasset.canton.protocol.v1.DynamicDomainParameters
	(v0.SigningKeyScheme)(0),        // 2: com.digitalasset.canton.crypto.v0.SigningKeyScheme
	(v0.EncryptionKeyScheme)(0),     // 3: com.digitalasset.canton.crypto.v0.EncryptionKeyScheme
	(v0.SymmetricKeyScheme)(0),      // 4: com.digitalasset.canton.crypto.v0.SymmetricKeyScheme
	(v0.HashAlgorithm)(0),           // 5: com.digitalasset.canton.crypto.v0.HashAlgorithm
	(v0.CryptoKeyFormat)(0),         // 6: com.digitalasset.canton.crypto.v0.CryptoKeyFormat
	(*durationpb.Duration)(nil),     // 7: google.protobuf.Duration
}
var file_com_digitalasset_canton_protocol_v1_sequencing_proto_depIdxs = []int32{
	2,  // 0: com.digitalasset.canton.protocol.v1.StaticDomainParameters.required_signing_key_schemes:type_name -> com.digitalasset.canton.crypto.v0.SigningKeyScheme
	3,  // 1: com.digitalasset.canton.protocol.v1.StaticDomainParameters.required_encryption_key_schemes:type_name -> com.digitalasset.canton.crypto.v0.EncryptionKeyScheme
	4,  // 2: com.digitalasset.canton.protocol.v1.StaticDomainParameters.required_symmetric_key_schemes:type_name -> com.digitalasset.canton.crypto.v0.SymmetricKeyScheme
	5,  // 3: com.digitalasset.canton.protocol.v1.StaticDomainParameters.required_hash_algorithms:type_name -> com.digitalasset.canton.crypto.v0.HashAlgorithm
	6,  // 4: com.digitalasset.canton.protocol.v1.StaticDomainParameters.required_crypto_key_formats:type_name -> com.digitalasset.canton.crypto.v0.CryptoKeyFormat
	7,  // 5: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.participant_response_timeout:type_name -> google.protobuf.Duration
	7,  // 6: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.mediator_reaction_timeout:type_name -> google.protobuf.Duration
	7,  // 7: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.transfer_exclusivity_timeout:type_name -> google.protobuf.Duration
	7,  // 8: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.topology_change_delay:type_name -> google.protobuf.Duration
	7,  // 9: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.ledger_time_record_time_tolerance:type_name -> google.protobuf.Duration
	7,  // 10: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.reconciliation_interval:type_name -> google.protobuf.Duration
	7,  // 11: com.digitalasset.canton.protocol.v1.DynamicDomainParameters.mediator_deduplication_timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_com_digitalasset_canton_protocol_v1_sequencing_proto_init() }
func file_com_digitalasset_canton_protocol_v1_sequencing_proto_init() {
	if File_com_digitalasset_canton_protocol_v1_sequencing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_digitalasset_canton_protocol_v1_sequencing_proto_goTypes,
		DependencyIndexes: file_com_digitalasset_canton_protocol_v1_sequencing_proto_depIdxs,
		MessageInfos:      file_com_digitalasset_canton_protocol_v1_sequencing_proto_msgTypes,
	}.Build()
	File_com_digitalasset_canton_protocol_v1_sequencing_proto = out.File
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_rawDesc = nil
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_goTypes = nil
	file_com_digitalasset_canton_protocol_v1_sequencing_proto_depIdxs = nil
}
