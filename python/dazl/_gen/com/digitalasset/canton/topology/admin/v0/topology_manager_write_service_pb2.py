# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/topology/admin/v0/topology_manager_write_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2
from ....protocol.v0 import topology_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__pb2
from ....protocol.v1 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_sequencing__pb2
from ....protocol.v2 import domain_params_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v2_dot_domain__params__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nNcom/digitalasset/canton/topology/admin/v0/topology_manager_write_service.proto\x12)com.digitalasset.canton.topology.admin.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a\x32\x63om/digitalasset/canton/protocol/v0/topology.proto\x1a\x34\x63om/digitalasset/canton/protocol/v1/sequencing.proto\x1a\x37\x63om/digitalasset/canton/protocol/v2/domain_params.proto\"6\n\x14\x41uthorizationSuccess\x12\x1e\n\nserialized\x18\x01 \x01(\x0cR\nserialized\"\x11\n\x0f\x41\x64\x64itionSuccess\"C\n!SignedTopologyTransactionAddition\x12\x1e\n\nserialized\x18\x01 \x01(\x0cR\nserialized\"\xcd\x01\n\x11\x41uthorizationData\x12M\n\x06\x63hange\x18\x01 \x01(\x0e\x32\x35.com.digitalasset.canton.protocol.v0.TopologyChangeOpR\x06\x63hange\x12\x1b\n\tsigned_by\x18\x02 \x01(\tR\x08signedBy\x12)\n\x10replace_existing\x18\x03 \x01(\x08R\x0freplaceExisting\x12!\n\x0c\x66orce_change\x18\x04 \x01(\x08R\x0b\x66orceChange\"\x95\x02\n NamespaceDelegationAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x1c\n\tnamespace\x18\x02 \x01(\tR\tnamespace\x12\x41\n\x1d\x66ingerprint_of_authorized_key\x18\x03 \x01(\tR\x1a\x66ingerprintOfAuthorizedKey\x12,\n\x12is_root_delegation\x18\x04 \x01(\x08R\x10isRootDelegation\"\xea\x01\n!IdentifierDelegationAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x1e\n\nidentifier\x18\x02 \x01(\tR\nidentifier\x12\x41\n\x1d\x66ingerprint_of_authorized_key\x18\x03 \x01(\tR\x1a\x66ingerprintOfAuthorizedKey\"\xdf\x02\n\x1fPartyToParticipantAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x44\n\x04side\x18\x02 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x14\n\x05party\x18\x03 \x01(\tR\x05party\x12 \n\x0bparticipant\x18\x04 \x01(\tR\x0bparticipant\x12Z\n\npermission\x18\x05 \x01(\x0e\x32:.com.digitalasset.canton.protocol.v0.ParticipantPermissionR\npermission\"\x9f\x02\n\x1eOwnerToKeyMappingAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x1b\n\tkey_owner\x18\x02 \x01(\tR\x08keyOwner\x12,\n\x12\x66ingerprint_of_key\x18\x03 \x01(\tR\x10\x66ingerprintOfKey\x12N\n\x0bkey_purpose\x18\x04 \x01(\x0e\x32-.com.digitalasset.canton.crypto.v0.KeyPurposeR\nkeyPurpose\"\xb7\x03\n#ParticipantDomainStateAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x44\n\x04side\x18\x02 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x16\n\x06\x64omain\x18\x03 \x01(\tR\x06\x64omain\x12 \n\x0bparticipant\x18\x04 \x01(\tR\x0bparticipant\x12Z\n\npermission\x18\x05 \x01(\x0e\x32:.com.digitalasset.canton.protocol.v0.ParticipantPermissionR\npermission\x12P\n\x0btrust_level\x18\x06 \x01(\x0e\x32/.com.digitalasset.canton.protocol.v0.TrustLevelR\ntrustLevel\"\x80\x02\n MediatorDomainStateAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x44\n\x04side\x18\x02 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x16\n\x06\x64omain\x18\x03 \x01(\tR\x06\x64omain\x12\x1a\n\x08mediator\x18\x04 \x01(\tR\x08mediator\"\xc4\x01\n\x1bVettedPackagesAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12 \n\x0bparticipant\x18\x02 \x01(\tR\x0bparticipant\x12\x1f\n\x0bpackage_ids\x18\x03 \x03(\tR\npackageIds\"\xde\x03\n#DomainParametersChangeAuthorization\x12\x62\n\rauthorization\x18\x01 \x01(\x0b\x32<.com.digitalasset.canton.topology.admin.v0.AuthorizationDataR\rauthorization\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x63\n\rparameters_v0\x18\x03 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v0.DynamicDomainParametersH\x00R\x0cparametersV0\x12\x63\n\rparameters_v1\x18\x04 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v1.DynamicDomainParametersH\x00R\x0cparametersV1\x12\x63\n\rparameters_v2\x18\x05 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v2.DynamicDomainParametersH\x00R\x0cparametersV2B\x0c\n\nparameters2\xbe\x0c\n\x1bTopologyManagerWriteService\x12\xaa\x01\n\x1b\x41uthorizePartyToParticipant\x12J.com.digitalasset.canton.topology.admin.v0.PartyToParticipantAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xa8\x01\n\x1a\x41uthorizeOwnerToKeyMapping\x12I.com.digitalasset.canton.topology.admin.v0.OwnerToKeyMappingAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xac\x01\n\x1c\x41uthorizeNamespaceDelegation\x12K.com.digitalasset.canton.topology.admin.v0.NamespaceDelegationAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xae\x01\n\x1d\x41uthorizeIdentifierDelegation\x12L.com.digitalasset.canton.topology.admin.v0.IdentifierDelegationAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xa2\x01\n\x17\x41uthorizeVettedPackages\x12\x46.com.digitalasset.canton.topology.admin.v0.VettedPackagesAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xb2\x01\n\x1f\x41uthorizeDomainParametersChange\x12N.com.digitalasset.canton.topology.admin.v0.DomainParametersChangeAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xb2\x01\n\x1f\x41uthorizeParticipantDomainState\x12N.com.digitalasset.canton.topology.admin.v0.ParticipantDomainStateAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xac\x01\n\x1c\x41uthorizeMediatorDomainState\x12K.com.digitalasset.canton.topology.admin.v0.MediatorDomainStateAuthorization\x1a?.com.digitalasset.canton.topology.admin.v0.AuthorizationSuccess\x12\xa8\x01\n\x1c\x41\x64\x64SignedTopologyTransaction\x12L.com.digitalasset.canton.topology.admin.v0.SignedTopologyTransactionAddition\x1a:.com.digitalasset.canton.topology.admin.v0.AdditionSuccessBZZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.topology.admin.v0.topology_manager_write_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v0'
  _globals['_AUTHORIZATIONSUCCESS']._serialized_start=390
  _globals['_AUTHORIZATIONSUCCESS']._serialized_end=444
  _globals['_ADDITIONSUCCESS']._serialized_start=446
  _globals['_ADDITIONSUCCESS']._serialized_end=463
  _globals['_SIGNEDTOPOLOGYTRANSACTIONADDITION']._serialized_start=465
  _globals['_SIGNEDTOPOLOGYTRANSACTIONADDITION']._serialized_end=532
  _globals['_AUTHORIZATIONDATA']._serialized_start=535
  _globals['_AUTHORIZATIONDATA']._serialized_end=740
  _globals['_NAMESPACEDELEGATIONAUTHORIZATION']._serialized_start=743
  _globals['_NAMESPACEDELEGATIONAUTHORIZATION']._serialized_end=1020
  _globals['_IDENTIFIERDELEGATIONAUTHORIZATION']._serialized_start=1023
  _globals['_IDENTIFIERDELEGATIONAUTHORIZATION']._serialized_end=1257
  _globals['_PARTYTOPARTICIPANTAUTHORIZATION']._serialized_start=1260
  _globals['_PARTYTOPARTICIPANTAUTHORIZATION']._serialized_end=1611
  _globals['_OWNERTOKEYMAPPINGAUTHORIZATION']._serialized_start=1614
  _globals['_OWNERTOKEYMAPPINGAUTHORIZATION']._serialized_end=1901
  _globals['_PARTICIPANTDOMAINSTATEAUTHORIZATION']._serialized_start=1904
  _globals['_PARTICIPANTDOMAINSTATEAUTHORIZATION']._serialized_end=2343
  _globals['_MEDIATORDOMAINSTATEAUTHORIZATION']._serialized_start=2346
  _globals['_MEDIATORDOMAINSTATEAUTHORIZATION']._serialized_end=2602
  _globals['_VETTEDPACKAGESAUTHORIZATION']._serialized_start=2605
  _globals['_VETTEDPACKAGESAUTHORIZATION']._serialized_end=2801
  _globals['_DOMAINPARAMETERSCHANGEAUTHORIZATION']._serialized_start=2804
  _globals['_DOMAINPARAMETERSCHANGEAUTHORIZATION']._serialized_end=3282
  _globals['_TOPOLOGYMANAGERWRITESERVICE']._serialized_start=3285
  _globals['_TOPOLOGYMANAGERWRITESERVICE']._serialized_end=4883
# @@protoc_insertion_point(module_scope)
