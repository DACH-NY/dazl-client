# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/topology/admin/v0/topology_manager_read_service.proto
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
from ....protocol.v0 import topology_ext_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__ext__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nMcom/digitalasset/canton/topology/admin/v0/topology_manager_read_service.proto\x12)com.digitalasset.canton.topology.admin.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a\x32\x63om/digitalasset/canton/protocol/v0/topology.proto\x1a\x34\x63om/digitalasset/canton/protocol/v1/sequencing.proto\x1a\x36\x63om/digitalasset/canton/protocol/v0/topology_ext.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xa0\x01\n\x1eListNamespaceDelegationRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12)\n\x10\x66ilter_namespace\x18\x02 \x01(\tR\x0f\x66ilterNamespace\"\xea\x02\n\x1dListNamespaceDelegationResult\x12i\n\x07results\x18\x01 \x03(\x0b\x32O.com.digitalasset.canton.topology.admin.v0.ListNamespaceDelegationResult.ResultR\x07results\x1a\xdd\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12L\n\x04item\x18\x02 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.NamespaceDelegationR\x04item\x12\x34\n\x16target_key_fingerprint\x18\x03 \x01(\tR\x14targetKeyFingerprint\"\x95\x01\n\x1fListIdentifierDelegationRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12\x1d\n\nfilter_uid\x18\x02 \x01(\tR\tfilterUid\"\xed\x02\n\x1eListIdentifierDelegationResult\x12j\n\x07results\x18\x01 \x03(\x0b\x32P.com.digitalasset.canton.topology.admin.v0.ListIdentifierDelegationResult.ResultR\x07results\x1a\xde\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12M\n\x04item\x18\x02 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.IdentifierDelegationR\x04item\x12\x34\n\x16target_key_fingerprint\x18\x03 \x01(\tR\x14targetKeyFingerprint\"\x93\x05\n\tBaseQuery\x12!\n\x0c\x66ilter_store\x18\x01 \x01(\tR\x0b\x66ilterStore\x12&\n\x0fuse_state_store\x18\x02 \x01(\x08R\ruseStateStore\x12S\n\toperation\x18\x03 \x01(\x0e\x32\x35.com.digitalasset.canton.protocol.v0.TopologyChangeOpR\toperation\x12)\n\x10\x66ilter_operation\x18\x04 \x01(\x08R\x0f\x66ilterOperation\x12\x38\n\x08snapshot\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00R\x08snapshot\x12\x37\n\nhead_state\x18\x06 \x01(\x0b\x32\x16.google.protobuf.EmptyH\x00R\theadState\x12V\n\x05range\x18\x07 \x01(\x0b\x32>.com.digitalasset.canton.topology.admin.v0.BaseQuery.TimeRangeH\x00R\x05range\x12*\n\x11\x66ilter_signed_key\x18\x08 \x01(\tR\x0f\x66ilterSignedKey\x12G\n\x10protocol_version\x18\t \x01(\x0b\x32\x1c.google.protobuf.StringValueR\x0fprotocolVersion\x1am\n\tTimeRange\x12.\n\x04\x66rom\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12\x30\n\x05until\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x05untilB\x0c\n\ntime_query\"\xfd\x02\n\nBaseResult\x12\x14\n\x05store\x18\x01 \x01(\tR\x05store\x12\x38\n\tsequenced\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tsequenced\x12\x39\n\nvalid_from\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tvalidFrom\x12;\n\x0bvalid_until\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nvalidUntil\x12S\n\toperation\x18\x04 \x01(\x0e\x32\x35.com.digitalasset.canton.protocol.v0.TopologyChangeOpR\toperation\x12\x1e\n\nserialized\x18\x05 \x01(\x0cR\nserialized\x12\x32\n\x15signed_by_fingerprint\x18\x06 \x01(\tR\x13signedByFingerprint\"\xb1\x02\n\x1cListPartyToParticipantResult\x12h\n\x07results\x18\x02 \x03(\x0b\x32N.com.digitalasset.canton.topology.admin.v0.ListPartyToParticipantResult.ResultR\x07results\x1a\xa6\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12K\n\x04item\x18\x02 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.PartyToParticipantR\x04item\"\x9f\x05\n\x1dListPartyToParticipantRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12!\n\x0c\x66ilter_party\x18\x02 \x01(\tR\x0b\x66ilterParty\x12-\n\x12\x66ilter_participant\x18\x03 \x01(\tR\x11\x66ilterParticipant\x12\x8a\x01\n\x13\x66ilter_request_side\x18\x04 \x01(\x0b\x32Z.com.digitalasset.canton.topology.admin.v0.ListPartyToParticipantRequest.FilterRequestSideR\x11\x66ilterRequestSide\x12\x86\x01\n\x11\x66ilter_permission\x18\x05 \x01(\x0b\x32Y.com.digitalasset.canton.topology.admin.v0.ListPartyToParticipantRequest.FilterPermissionR\x10\x66ilterPermission\x1a[\n\x11\x46ilterRequestSide\x12\x46\n\x05value\x18\x01 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x05value\x1a\x64\n\x10\x46ilterPermission\x12P\n\x05value\x18\x01 \x01(\x0e\x32:.com.digitalasset.canton.protocol.v0.ParticipantPermissionR\x05value\"\xb9\x03\n\x1cListOwnerToKeyMappingRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12\x31\n\x15\x66ilter_key_owner_type\x18\x02 \x01(\tR\x12\x66ilterKeyOwnerType\x12/\n\x14\x66ilter_key_owner_uid\x18\x03 \x01(\tR\x11\x66ilterKeyOwnerUid\x12\x86\x01\n\x12\x66ilter_key_purpose\x18\x04 \x01(\x0b\x32X.com.digitalasset.canton.topology.admin.v0.ListOwnerToKeyMappingRequest.FilterKeyPurposeR\x10\x66ilterKeyPurpose\x1aW\n\x10\x46ilterKeyPurpose\x12\x43\n\x05value\x18\x01 \x01(\x0e\x32-.com.digitalasset.canton.crypto.v0.KeyPurposeR\x05value\"\xd7\x02\n\x1bListOwnerToKeyMappingResult\x12g\n\x07results\x18\x01 \x03(\x0b\x32M.com.digitalasset.canton.topology.admin.v0.ListOwnerToKeyMappingResult.ResultR\x07results\x1a\xce\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12J\n\x04item\x18\x02 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.OwnerToKeyMappingR\x04item\x12\'\n\x0fkey_fingerprint\x18\x03 \x01(\tR\x0ekeyFingerprint\"\x99\x01\n#ListSignedLegalIdentityClaimRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12\x1d\n\nfilter_uid\x18\x02 \x01(\tR\tfilterUid\"\xc3\x02\n\"ListSignedLegalIdentityClaimResult\x12n\n\x07results\x18\x01 \x03(\x0b\x32T.com.digitalasset.canton.topology.admin.v0.ListSignedLegalIdentityClaimResult.ResultR\x07results\x1a\xac\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12Q\n\x04item\x18\x02 \x01(\x0b\x32=.com.digitalasset.canton.protocol.v0.SignedLegalIdentityClaimR\x04item\"\x9f\x01\n\x19ListVettedPackagesRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12-\n\x12\x66ilter_participant\x18\x02 \x01(\tR\x11\x66ilterParticipant\"\xa5\x02\n\x18ListVettedPackagesResult\x12\x64\n\x07results\x18\x01 \x03(\x0b\x32J.com.digitalasset.canton.topology.admin.v0.ListVettedPackagesResult.ResultR\x07results\x1a\xa2\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12G\n\x04item\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.protocol.v0.VettedPackagesR\x04item\"y\n\"ListDomainParametersChangesRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\"\x9c\x03\n!ListDomainParametersChangesResult\x12m\n\x07results\x18\x01 \x03(\x0b\x32S.com.digitalasset.canton.topology.admin.v0.ListDomainParametersChangesResult.ResultR\x07results\x1a\x87\x02\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12N\n\x02v0\x18\x02 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v0.DynamicDomainParametersH\x00R\x02v0\x12N\n\x02v1\x18\x03 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v1.DynamicDomainParametersH\x00R\x02v1B\x0c\n\nparameters\"\x1c\n\x1aListAvailableStoresRequest\"8\n\x19ListAvailableStoresResult\x12\x1b\n\tstore_ids\x18\x01 \x03(\tR\x08storeIds\"\xcc\x01\n!ListParticipantDomainStateRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12#\n\rfilter_domain\x18\x02 \x01(\tR\x0c\x66ilterDomain\x12-\n\x12\x66ilter_participant\x18\x03 \x01(\tR\x11\x66ilterParticipant\"\xb7\x02\n ListParticipantDomainStateResult\x12l\n\x07results\x18\x01 \x03(\x0b\x32R.com.digitalasset.canton.topology.admin.v0.ListParticipantDomainStateResult.ResultR\x07results\x1a\xa4\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12I\n\x04item\x18\x02 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.ParticipantStateR\x04item\"\xc3\x01\n\x1eListMediatorDomainStateRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\x12#\n\rfilter_domain\x18\x02 \x01(\tR\x0c\x66ilterDomain\x12\'\n\x0f\x66ilter_mediator\x18\x03 \x01(\tR\x0e\x66ilterMediator\"\xb4\x02\n\x1dListMediatorDomainStateResult\x12i\n\x07results\x18\x01 \x03(\x0b\x32O.com.digitalasset.canton.topology.admin.v0.ListMediatorDomainStateResult.ResultR\x07results\x1a\xa7\x01\n\x06Result\x12O\n\x07\x63ontext\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.topology.admin.v0.BaseResultR\x07\x63ontext\x12L\n\x04item\x18\x02 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.MediatorDomainStateR\x04item\"e\n\x0eListAllRequest\x12S\n\nbase_query\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.topology.admin.v0.BaseQueryR\tbaseQuery\"d\n\x0fListAllResponse\x12Q\n\x06result\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.TopologyTransactionsR\x06result2\x8c\x0f\n\x1aTopologyManagerReadService\x12\xa2\x01\n\x13ListAvailableStores\x12\x45.com.digitalasset.canton.topology.admin.v0.ListAvailableStoresRequest\x1a\x44.com.digitalasset.canton.topology.admin.v0.ListAvailableStoresResult\x12\xab\x01\n\x16ListPartyToParticipant\x12H.com.digitalasset.canton.topology.admin.v0.ListPartyToParticipantRequest\x1aG.com.digitalasset.canton.topology.admin.v0.ListPartyToParticipantResult\x12\xa8\x01\n\x15ListOwnerToKeyMapping\x12G.com.digitalasset.canton.topology.admin.v0.ListOwnerToKeyMappingRequest\x1a\x46.com.digitalasset.canton.topology.admin.v0.ListOwnerToKeyMappingResult\x12\xae\x01\n\x17ListNamespaceDelegation\x12I.com.digitalasset.canton.topology.admin.v0.ListNamespaceDelegationRequest\x1aH.com.digitalasset.canton.topology.admin.v0.ListNamespaceDelegationResult\x12\xb1\x01\n\x18ListIdentifierDelegation\x12J.com.digitalasset.canton.topology.admin.v0.ListIdentifierDelegationRequest\x1aI.com.digitalasset.canton.topology.admin.v0.ListIdentifierDelegationResult\x12\xbd\x01\n\x1cListSignedLegalIdentityClaim\x12N.com.digitalasset.canton.topology.admin.v0.ListSignedLegalIdentityClaimRequest\x1aM.com.digitalasset.canton.topology.admin.v0.ListSignedLegalIdentityClaimResult\x12\xb7\x01\n\x1aListParticipantDomainState\x12L.com.digitalasset.canton.topology.admin.v0.ListParticipantDomainStateRequest\x1aK.com.digitalasset.canton.topology.admin.v0.ListParticipantDomainStateResult\x12\xae\x01\n\x17ListMediatorDomainState\x12I.com.digitalasset.canton.topology.admin.v0.ListMediatorDomainStateRequest\x1aH.com.digitalasset.canton.topology.admin.v0.ListMediatorDomainStateResult\x12\x9f\x01\n\x12ListVettedPackages\x12\x44.com.digitalasset.canton.topology.admin.v0.ListVettedPackagesRequest\x1a\x43.com.digitalasset.canton.topology.admin.v0.ListVettedPackagesResult\x12\xba\x01\n\x1bListDomainParametersChanges\x12M.com.digitalasset.canton.topology.admin.v0.ListDomainParametersChangesRequest\x1aL.com.digitalasset.canton.topology.admin.v0.ListDomainParametersChangesResult\x12\x80\x01\n\x07ListAll\x12\x39.com.digitalasset.canton.topology.admin.v0.ListAllRequest\x1a:.com.digitalasset.canton.topology.admin.v0.ListAllResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.topology.admin.v0.topology_manager_read_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_LISTNAMESPACEDELEGATIONREQUEST']._serialized_start=483
  _globals['_LISTNAMESPACEDELEGATIONREQUEST']._serialized_end=643
  _globals['_LISTNAMESPACEDELEGATIONRESULT']._serialized_start=646
  _globals['_LISTNAMESPACEDELEGATIONRESULT']._serialized_end=1008
  _globals['_LISTNAMESPACEDELEGATIONRESULT_RESULT']._serialized_start=787
  _globals['_LISTNAMESPACEDELEGATIONRESULT_RESULT']._serialized_end=1008
  _globals['_LISTIDENTIFIERDELEGATIONREQUEST']._serialized_start=1011
  _globals['_LISTIDENTIFIERDELEGATIONREQUEST']._serialized_end=1160
  _globals['_LISTIDENTIFIERDELEGATIONRESULT']._serialized_start=1163
  _globals['_LISTIDENTIFIERDELEGATIONRESULT']._serialized_end=1528
  _globals['_LISTIDENTIFIERDELEGATIONRESULT_RESULT']._serialized_start=1306
  _globals['_LISTIDENTIFIERDELEGATIONRESULT_RESULT']._serialized_end=1528
  _globals['_BASEQUERY']._serialized_start=1531
  _globals['_BASEQUERY']._serialized_end=2190
  _globals['_BASEQUERY_TIMERANGE']._serialized_start=2067
  _globals['_BASEQUERY_TIMERANGE']._serialized_end=2176
  _globals['_BASERESULT']._serialized_start=2193
  _globals['_BASERESULT']._serialized_end=2574
  _globals['_LISTPARTYTOPARTICIPANTRESULT']._serialized_start=2577
  _globals['_LISTPARTYTOPARTICIPANTRESULT']._serialized_end=2882
  _globals['_LISTPARTYTOPARTICIPANTRESULT_RESULT']._serialized_start=2716
  _globals['_LISTPARTYTOPARTICIPANTRESULT_RESULT']._serialized_end=2882
  _globals['_LISTPARTYTOPARTICIPANTREQUEST']._serialized_start=2885
  _globals['_LISTPARTYTOPARTICIPANTREQUEST']._serialized_end=3556
  _globals['_LISTPARTYTOPARTICIPANTREQUEST_FILTERREQUESTSIDE']._serialized_start=3363
  _globals['_LISTPARTYTOPARTICIPANTREQUEST_FILTERREQUESTSIDE']._serialized_end=3454
  _globals['_LISTPARTYTOPARTICIPANTREQUEST_FILTERPERMISSION']._serialized_start=3456
  _globals['_LISTPARTYTOPARTICIPANTREQUEST_FILTERPERMISSION']._serialized_end=3556
  _globals['_LISTOWNERTOKEYMAPPINGREQUEST']._serialized_start=3559
  _globals['_LISTOWNERTOKEYMAPPINGREQUEST']._serialized_end=4000
  _globals['_LISTOWNERTOKEYMAPPINGREQUEST_FILTERKEYPURPOSE']._serialized_start=3913
  _globals['_LISTOWNERTOKEYMAPPINGREQUEST_FILTERKEYPURPOSE']._serialized_end=4000
  _globals['_LISTOWNERTOKEYMAPPINGRESULT']._serialized_start=4003
  _globals['_LISTOWNERTOKEYMAPPINGRESULT']._serialized_end=4346
  _globals['_LISTOWNERTOKEYMAPPINGRESULT_RESULT']._serialized_start=4140
  _globals['_LISTOWNERTOKEYMAPPINGRESULT_RESULT']._serialized_end=4346
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMREQUEST']._serialized_start=4349
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMREQUEST']._serialized_end=4502
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMRESULT']._serialized_start=4505
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMRESULT']._serialized_end=4828
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMRESULT_RESULT']._serialized_start=4656
  _globals['_LISTSIGNEDLEGALIDENTITYCLAIMRESULT_RESULT']._serialized_end=4828
  _globals['_LISTVETTEDPACKAGESREQUEST']._serialized_start=4831
  _globals['_LISTVETTEDPACKAGESREQUEST']._serialized_end=4990
  _globals['_LISTVETTEDPACKAGESRESULT']._serialized_start=4993
  _globals['_LISTVETTEDPACKAGESRESULT']._serialized_end=5286
  _globals['_LISTVETTEDPACKAGESRESULT_RESULT']._serialized_start=5124
  _globals['_LISTVETTEDPACKAGESRESULT_RESULT']._serialized_end=5286
  _globals['_LISTDOMAINPARAMETERSCHANGESREQUEST']._serialized_start=5288
  _globals['_LISTDOMAINPARAMETERSCHANGESREQUEST']._serialized_end=5409
  _globals['_LISTDOMAINPARAMETERSCHANGESRESULT']._serialized_start=5412
  _globals['_LISTDOMAINPARAMETERSCHANGESRESULT']._serialized_end=5824
  _globals['_LISTDOMAINPARAMETERSCHANGESRESULT_RESULT']._serialized_start=5561
  _globals['_LISTDOMAINPARAMETERSCHANGESRESULT_RESULT']._serialized_end=5824
  _globals['_LISTAVAILABLESTORESREQUEST']._serialized_start=5826
  _globals['_LISTAVAILABLESTORESREQUEST']._serialized_end=5854
  _globals['_LISTAVAILABLESTORESRESULT']._serialized_start=5856
  _globals['_LISTAVAILABLESTORESRESULT']._serialized_end=5912
  _globals['_LISTPARTICIPANTDOMAINSTATEREQUEST']._serialized_start=5915
  _globals['_LISTPARTICIPANTDOMAINSTATEREQUEST']._serialized_end=6119
  _globals['_LISTPARTICIPANTDOMAINSTATERESULT']._serialized_start=6122
  _globals['_LISTPARTICIPANTDOMAINSTATERESULT']._serialized_end=6433
  _globals['_LISTPARTICIPANTDOMAINSTATERESULT_RESULT']._serialized_start=6269
  _globals['_LISTPARTICIPANTDOMAINSTATERESULT_RESULT']._serialized_end=6433
  _globals['_LISTMEDIATORDOMAINSTATEREQUEST']._serialized_start=6436
  _globals['_LISTMEDIATORDOMAINSTATEREQUEST']._serialized_end=6631
  _globals['_LISTMEDIATORDOMAINSTATERESULT']._serialized_start=6634
  _globals['_LISTMEDIATORDOMAINSTATERESULT']._serialized_end=6942
  _globals['_LISTMEDIATORDOMAINSTATERESULT_RESULT']._serialized_start=6775
  _globals['_LISTMEDIATORDOMAINSTATERESULT_RESULT']._serialized_end=6942
  _globals['_LISTALLREQUEST']._serialized_start=6944
  _globals['_LISTALLREQUEST']._serialized_end=7045
  _globals['_LISTALLRESPONSE']._serialized_start=7047
  _globals['_LISTALLRESPONSE']._serialized_end=7147
  _globals['_TOPOLOGYMANAGERREADSERVICE']._serialized_start=7150
  _globals['_TOPOLOGYMANAGERREADSERVICE']._serialized_end=9082
# @@protoc_insertion_point(module_scope)