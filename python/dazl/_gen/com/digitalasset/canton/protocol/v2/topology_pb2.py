# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v2/topology.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from . import domain_params_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v2_dot_domain__params__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2com/digitalasset/canton/protocol/v2/topology.proto\x12#com.digitalasset.canton.protocol.v2\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x37\x63om/digitalasset/canton/protocol/v2/domain_params.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb6\x01\n\x14NamespaceDelegationX\x12\x1c\n\tnamespace\x18\x01 \x01(\tR\tnamespace\x12R\n\ntarget_key\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.crypto.v0.SigningPublicKeyR\ttargetKey\x12,\n\x12is_root_delegation\x18\x03 \x01(\x08R\x10isRootDelegation\"m\n\x15UnionspaceDefinitionX\x12\x1e\n\nunionspace\x18\x01 \x01(\tR\nunionspace\x12\x1c\n\tthreshold\x18\x02 \x01(\x05R\tthreshold\x12\x16\n\x06owners\x18\x03 \x03(\tR\x06owners\"\x98\x01\n\x15IdentifierDelegationX\x12+\n\x11unique_identifier\x18\x01 \x01(\tR\x10uniqueIdentifier\x12R\n\ntarget_key\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.crypto.v0.SigningPublicKeyR\ttargetKey\"\x93\x01\n\x12OwnerToKeyMappingX\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12M\n\x0bpublic_keys\x18\x02 \x03(\x0b\x32,.com.digitalasset.canton.crypto.v0.PublicKeyR\npublicKeys\x12\x16\n\x06\x64omain\x18\x03 \x01(\tR\x06\x64omain\"\xcb\x01\n\x17\x44omainTrustCertificateX\x12 \n\x0bparticipant\x18\x01 \x01(\tR\x0bparticipant\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12O\n%transfer_only_to_given_target_domains\x18\x03 \x01(\x08R transferOnlyToGivenTargetDomains\x12%\n\x0etarget_domains\x18\x04 \x03(\tR\rtargetDomains\"\x9b\x03\n\x1cParticipantDomainPermissionX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12 \n\x0bparticipant\x18\x02 \x01(\tR\x0bparticipant\x12[\n\npermission\x18\x03 \x01(\x0e\x32;.com.digitalasset.canton.protocol.v2.ParticipantPermissionXR\npermission\x12Q\n\x0btrust_level\x18\x04 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v2.TrustLevelXR\ntrustLevel\x12T\n\x06limits\x18\x05 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v2.ParticipantDomainLimitsR\x06limits\x12;\n\x0blogin_after\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nloginAfter\"Y\n\x13PartyHostingLimitsX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12\x14\n\x05party\x18\x02 \x01(\tR\x05party\x12\x14\n\x05quota\x18\x03 \x01(\rR\x05quota\"l\n\x0fVettedPackagesX\x12 \n\x0bparticipant\x18\x01 \x01(\tR\x0bparticipant\x12\x1f\n\x0bpackage_ids\x18\x02 \x03(\tR\npackageIds\x12\x16\n\x06\x64omain\x18\x03 \x01(\tR\x06\x64omain\"\x93\x03\n\x13PartyToParticipantX\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x12\x1c\n\tthreshold\x18\x02 \x01(\rR\tthreshold\x12o\n\x0cparticipants\x18\x03 \x03(\x0b\x32K.com.digitalasset.canton.protocol.v2.PartyToParticipantX.HostingParticipantR\x0cparticipants\x12)\n\x10group_addressing\x18\x04 \x01(\x08R\x0fgroupAddressing\x12\x16\n\x06\x64omain\x18\x05 \x01(\tR\x06\x64omain\x1a\x93\x01\n\x12HostingParticipant\x12 \n\x0bparticipant\x18\x01 \x01(\tR\x0bparticipant\x12[\n\npermission\x18\x02 \x01(\x0e\x32;.com.digitalasset.canton.protocol.v2.ParticipantPermissionXR\npermission\"t\n\x0c\x41uthorityOfX\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x12\x1c\n\tthreshold\x18\x02 \x01(\rR\tthreshold\x12\x18\n\x07parties\x18\x03 \x03(\tR\x07parties\x12\x16\n\x06\x64omain\x18\x04 \x01(\tR\x06\x64omain\"\x9c\x01\n\x16\x44omainParametersStateX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12j\n\x11\x64omain_parameters\x18\x02 \x01(\x0b\x32=.com.digitalasset.canton.protocol.v2.DynamicDomainParametersXR\x10\x64omainParameters\"\x98\x01\n\x14MediatorDomainStateX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12\x14\n\x05group\x18\x02 \x01(\rR\x05group\x12\x1c\n\tthreshold\x18\x03 \x01(\rR\tthreshold\x12\x16\n\x06\x61\x63tive\x18\x04 \x03(\tR\x06\x61\x63tive\x12\x1c\n\tobservers\x18\x05 \x03(\tR\tobservers\"\x83\x01\n\x15SequencerDomainStateX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12\x1c\n\tthreshold\x18\x02 \x01(\rR\tthreshold\x12\x16\n\x06\x61\x63tive\x18\x03 \x03(\tR\x06\x61\x63tive\x12\x1c\n\tobservers\x18\x04 \x03(\tR\tobservers\"\x86\x01\n\x19PurgeTopologyTransactionX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12Q\n\x08mappings\x18\x02 \x03(\x0b\x32\x35.com.digitalasset.canton.protocol.v2.TopologyMappingXR\x08mappings\"\x81\x01\n\x14TrafficControlStateX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12\x16\n\x06member\x18\x02 \x01(\tR\x06member\x12\x39\n\x19total_extra_traffic_limit\x18\x03 \x01(\x04R\x16totalExtraTrafficLimit\"\xa7\r\n\x10TopologyMappingX\x12n\n\x14namespace_delegation\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v2.NamespaceDelegationXH\x00R\x13namespaceDelegation\x12q\n\x15identifier_delegation\x18\x02 \x01(\x0b\x32:.com.digitalasset.canton.protocol.v2.IdentifierDelegationXH\x00R\x14identifierDelegation\x12q\n\x15unionspace_definition\x18\x03 \x01(\x0b\x32:.com.digitalasset.canton.protocol.v2.UnionspaceDefinitionXH\x00R\x14unionspaceDefinition\x12j\n\x14owner_to_key_mapping\x18\x04 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v2.OwnerToKeyMappingXH\x00R\x11ownerToKeyMapping\x12x\n\x18\x64omain_trust_certificate\x18\x05 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v2.DomainTrustCertificateXH\x00R\x16\x64omainTrustCertificate\x12z\n\x16participant_permission\x18\x06 \x01(\x0b\x32\x41.com.digitalasset.canton.protocol.v2.ParticipantDomainPermissionXH\x00R\x15participantPermission\x12l\n\x14party_hosting_limits\x18\x07 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v2.PartyHostingLimitsXH\x00R\x12partyHostingLimits\x12_\n\x0fvetted_packages\x18\x08 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v2.VettedPackagesXH\x00R\x0evettedPackages\x12l\n\x14party_to_participant\x18\t \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v2.PartyToParticipantXH\x00R\x12partyToParticipant\x12V\n\x0c\x61uthority_of\x18\n \x01(\x0b\x32\x31.com.digitalasset.canton.protocol.v2.AuthorityOfXH\x00R\x0b\x61uthorityOf\x12u\n\x17\x64omain_parameters_state\x18\x0b \x01(\x0b\x32;.com.digitalasset.canton.protocol.v2.DomainParametersStateXH\x00R\x15\x64omainParametersState\x12o\n\x15mediator_domain_state\x18\x0c \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v2.MediatorDomainStateXH\x00R\x13mediatorDomainState\x12r\n\x16sequencer_domain_state\x18\r \x01(\x0b\x32:.com.digitalasset.canton.protocol.v2.SequencerDomainStateXH\x00R\x14sequencerDomainState\x12n\n\x12purge_topology_txs\x18\x0e \x01(\x0b\x32>.com.digitalasset.canton.protocol.v2.PurgeTopologyTransactionXH\x00R\x10purgeTopologyTxs\x12o\n\x15traffic_control_state\x18\x0f \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v2.TrafficControlStateXH\x00R\x13trafficControlStateB\t\n\x07mapping\"\xd5\x01\n\x14TopologyTransactionX\x12T\n\toperation\x18\x01 \x01(\x0e\x32\x36.com.digitalasset.canton.protocol.v2.TopologyChangeOpXR\toperation\x12\x16\n\x06serial\x18\x02 \x01(\rR\x06serial\x12O\n\x07mapping\x18\x03 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v2.TopologyMappingXR\x07mapping\"\xa8\x01\n\x1aSignedTopologyTransactionX\x12 \n\x0btransaction\x18\x01 \x01(\x0cR\x0btransaction\x12L\n\nsignatures\x18\x02 \x03(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\nsignatures\x12\x1a\n\x08proposal\x18\x03 \x01(\x08R\x08proposal\"\xbf\x02\n\x1d\x41\x63\x63\x65ptedTopologyTransactionsX\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12n\n\x08\x61\x63\x63\x65pted\x18\x02 \x03(\x0b\x32R.com.digitalasset.canton.protocol.v2.AcceptedTopologyTransactionsX.AcceptedRequestR\x08\x61\x63\x63\x65pted\x1a\x95\x01\n\x0f\x41\x63\x63\x65ptedRequest\x12\x1d\n\nrequest_id\x18\x01 \x01(\tR\trequestId\x12\x63\n\x0ctransactions\x18\x02 \x03(\x0b\x32?.com.digitalasset.canton.protocol.v2.SignedTopologyTransactionXR\x0ctransactions\"\x89\x02\n#RegisterTopologyTransactionRequestX\x12!\n\x0crequested_by\x18\x01 \x01(\tR\x0brequestedBy\x12#\n\rrequested_for\x18\x02 \x01(\tR\x0crequestedFor\x12\x1d\n\nrequest_id\x18\x03 \x01(\tR\trequestId\x12\x63\n\x0ctransactions\x18\x04 \x03(\x0b\x32?.com.digitalasset.canton.protocol.v2.SignedTopologyTransactionXR\x0ctransactions\x12\x16\n\x06\x64omain\x18\x05 \x01(\tR\x06\x64omain\"\xeb\x03\n$RegisterTopologyTransactionResponseX\x12!\n\x0crequested_by\x18\x01 \x01(\tR\x0brequestedBy\x12#\n\rrequested_for\x18\x02 \x01(\tR\x0crequestedFor\x12\x1d\n\nrequest_id\x18\x03 \x01(\tR\trequestId\x12j\n\x07results\x18\x04 \x03(\x0b\x32P.com.digitalasset.canton.protocol.v2.RegisterTopologyTransactionResponseX.ResultR\x07results\x12\x16\n\x06\x64omain\x18\x05 \x01(\tR\x06\x64omain\x1a\xd7\x01\n\x06Result\x12l\n\x05state\x18\x01 \x01(\x0e\x32V.com.digitalasset.canton.protocol.v2.RegisterTopologyTransactionResponseX.Result.StateR\x05state\"_\n\x05State\x12\x11\n\rMISSING_STATE\x10\x00\x12\n\n\x06\x46\x41ILED\x10\x01\x12\x0c\n\x08REJECTED\x10\x02\x12\x0c\n\x08\x41\x43\x43\x45PTED\x10\x03\x12\r\n\tDUPLICATE\x10\x04\x12\x0c\n\x08OBSOLETE\x10\x05*,\n\x11TopologyChangeOpX\x12\x0b\n\x07Replace\x10\x00\x12\n\n\x06Remove\x10\x01*;\n\x0bTrustLevelX\x12\x15\n\x11MissingTrustLevel\x10\x00\x12\x0c\n\x08Ordinary\x10\x01\x12\x07\n\x03Vip\x10\x02*m\n\x16ParticipantPermissionX\x12 \n\x1cMissingParticipantPermission\x10\x00\x12\x0e\n\nSubmission\x10\x01\x12\x10\n\x0c\x43onfirmation\x10\x02\x12\x0f\n\x0bObservation\x10\x03\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v2.topology_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_TOPOLOGYCHANGEOPX']._serialized_start=6069
  _globals['_TOPOLOGYCHANGEOPX']._serialized_end=6113
  _globals['_TRUSTLEVELX']._serialized_start=6115
  _globals['_TRUSTLEVELX']._serialized_end=6174
  _globals['_PARTICIPANTPERMISSIONX']._serialized_start=6176
  _globals['_PARTICIPANTPERMISSIONX']._serialized_end=6285
  _globals['_NAMESPACEDELEGATIONX']._serialized_start=230
  _globals['_NAMESPACEDELEGATIONX']._serialized_end=412
  _globals['_UNIONSPACEDEFINITIONX']._serialized_start=414
  _globals['_UNIONSPACEDEFINITIONX']._serialized_end=523
  _globals['_IDENTIFIERDELEGATIONX']._serialized_start=526
  _globals['_IDENTIFIERDELEGATIONX']._serialized_end=678
  _globals['_OWNERTOKEYMAPPINGX']._serialized_start=681
  _globals['_OWNERTOKEYMAPPINGX']._serialized_end=828
  _globals['_DOMAINTRUSTCERTIFICATEX']._serialized_start=831
  _globals['_DOMAINTRUSTCERTIFICATEX']._serialized_end=1034
  _globals['_PARTICIPANTDOMAINPERMISSIONX']._serialized_start=1037
  _globals['_PARTICIPANTDOMAINPERMISSIONX']._serialized_end=1448
  _globals['_PARTYHOSTINGLIMITSX']._serialized_start=1450
  _globals['_PARTYHOSTINGLIMITSX']._serialized_end=1539
  _globals['_VETTEDPACKAGESX']._serialized_start=1541
  _globals['_VETTEDPACKAGESX']._serialized_end=1649
  _globals['_PARTYTOPARTICIPANTX']._serialized_start=1652
  _globals['_PARTYTOPARTICIPANTX']._serialized_end=2055
  _globals['_PARTYTOPARTICIPANTX_HOSTINGPARTICIPANT']._serialized_start=1908
  _globals['_PARTYTOPARTICIPANTX_HOSTINGPARTICIPANT']._serialized_end=2055
  _globals['_AUTHORITYOFX']._serialized_start=2057
  _globals['_AUTHORITYOFX']._serialized_end=2173
  _globals['_DOMAINPARAMETERSSTATEX']._serialized_start=2176
  _globals['_DOMAINPARAMETERSSTATEX']._serialized_end=2332
  _globals['_MEDIATORDOMAINSTATEX']._serialized_start=2335
  _globals['_MEDIATORDOMAINSTATEX']._serialized_end=2487
  _globals['_SEQUENCERDOMAINSTATEX']._serialized_start=2490
  _globals['_SEQUENCERDOMAINSTATEX']._serialized_end=2621
  _globals['_PURGETOPOLOGYTRANSACTIONX']._serialized_start=2624
  _globals['_PURGETOPOLOGYTRANSACTIONX']._serialized_end=2758
  _globals['_TRAFFICCONTROLSTATEX']._serialized_start=2761
  _globals['_TRAFFICCONTROLSTATEX']._serialized_end=2890
  _globals['_TOPOLOGYMAPPINGX']._serialized_start=2893
  _globals['_TOPOLOGYMAPPINGX']._serialized_end=4596
  _globals['_TOPOLOGYTRANSACTIONX']._serialized_start=4599
  _globals['_TOPOLOGYTRANSACTIONX']._serialized_end=4812
  _globals['_SIGNEDTOPOLOGYTRANSACTIONX']._serialized_start=4815
  _globals['_SIGNEDTOPOLOGYTRANSACTIONX']._serialized_end=4983
  _globals['_ACCEPTEDTOPOLOGYTRANSACTIONSX']._serialized_start=4986
  _globals['_ACCEPTEDTOPOLOGYTRANSACTIONSX']._serialized_end=5305
  _globals['_ACCEPTEDTOPOLOGYTRANSACTIONSX_ACCEPTEDREQUEST']._serialized_start=5156
  _globals['_ACCEPTEDTOPOLOGYTRANSACTIONSX_ACCEPTEDREQUEST']._serialized_end=5305
  _globals['_REGISTERTOPOLOGYTRANSACTIONREQUESTX']._serialized_start=5308
  _globals['_REGISTERTOPOLOGYTRANSACTIONREQUESTX']._serialized_end=5573
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX']._serialized_start=5576
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX']._serialized_end=6067
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX_RESULT']._serialized_start=5852
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX_RESULT']._serialized_end=6067
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX_RESULT_STATE']._serialized_start=5972
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSEX_RESULT_STATE']._serialized_end=6067
# @@protoc_insertion_point(module_scope)
