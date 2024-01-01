# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/topology.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from . import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2com/digitalasset/canton/protocol/v0/topology.proto\x12#com.digitalasset.canton.protocol.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\"\xc0\x02\n\x10ParticipantState\x12\x44\n\x04side\x18\x01 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12 \n\x0bparticipant\x18\x03 \x01(\tR\x0bparticipant\x12Z\n\npermission\x18\x04 \x01(\x0e\x32:.com.digitalasset.canton.protocol.v0.ParticipantPermissionR\npermission\x12P\n\x0btrust_level\x18\x05 \x01(\x0e\x32/.com.digitalasset.canton.protocol.v0.TrustLevelR\ntrustLevel\"\xee\x01\n\x12PartyToParticipant\x12\x44\n\x04side\x18\x01 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x14\n\x05party\x18\x02 \x01(\tR\x05party\x12 \n\x0bparticipant\x18\x03 \x01(\tR\x0bparticipant\x12Z\n\npermission\x18\x04 \x01(\x0e\x32:.com.digitalasset.canton.protocol.v0.ParticipantPermissionR\npermission\"\x8f\x01\n\x13MediatorDomainState\x12\x44\n\x04side\x18\x01 \x01(\x0e\x32\x30.com.digitalasset.canton.protocol.v0.RequestSideR\x04side\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x1a\n\x08mediator\x18\x03 \x01(\tR\x08mediator\"\xb5\x01\n\x13NamespaceDelegation\x12\x1c\n\tnamespace\x18\x01 \x01(\tR\tnamespace\x12R\n\ntarget_key\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.crypto.v0.SigningPublicKeyR\ttargetKey\x12,\n\x12is_root_delegation\x18\x03 \x01(\x08R\x10isRootDelegation\"\x97\x01\n\x14IdentifierDelegation\x12+\n\x11unique_identifier\x18\x01 \x01(\tR\x10uniqueIdentifier\x12R\n\ntarget_key\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.crypto.v0.SigningPublicKeyR\ttargetKey\"}\n\x11OwnerToKeyMapping\x12\x1b\n\tkey_owner\x18\x01 \x01(\tR\x08keyOwner\x12K\n\npublic_key\x18\x02 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.PublicKeyR\tpublicKey\"|\n\x18SignedLegalIdentityClaim\x12\x14\n\x05\x63laim\x18\x01 \x01(\x0cR\x05\x63laim\x12J\n\tsignature\x18\x02 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\tsignature\"l\n\x12LegalIdentityClaim\x12+\n\x11unique_identifier\x18\x01 \x01(\tR\x10uniqueIdentifier\x12\x1d\n\tx509_cert\x18\x02 \x01(\x0cH\x00R\x08x509CertB\n\n\x08\x65vidence\"S\n\x0eVettedPackages\x12 \n\x0bparticipant\x18\x01 \x01(\tR\x0bparticipant\x12\x1f\n\x0bpackage_ids\x18\x02 \x03(\tR\npackageIds\"\xf4\x07\n\x13TopologyStateUpdate\x12S\n\toperation\x18\x01 \x01(\x0e\x32\x35.com.digitalasset.canton.protocol.v0.TopologyChangeOpR\toperation\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\x12m\n\x14namespace_delegation\x18\x03 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.NamespaceDelegationH\x00R\x13namespaceDelegation\x12p\n\x15identifier_delegation\x18\x04 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.IdentifierDelegationH\x00R\x14identifierDelegation\x12i\n\x14owner_to_key_mapping\x18\x05 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.OwnerToKeyMappingH\x00R\x11ownerToKeyMapping\x12k\n\x14party_to_participant\x18\x06 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.PartyToParticipantH\x00R\x12partyToParticipant\x12~\n\x1bsigned_legal_identity_claim\x18\x07 \x01(\x0b\x32=.com.digitalasset.canton.protocol.v0.SignedLegalIdentityClaimH\x00R\x18signedLegalIdentityClaim\x12\x64\n\x11participant_state\x18\x08 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.ParticipantStateH\x00R\x10participantState\x12^\n\x0fvetted_packages\x18\t \x01(\x0b\x32\x33.com.digitalasset.canton.protocol.v0.VettedPackagesH\x00R\x0evettedPackages\x12n\n\x15mediator_domain_state\x18\n \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.MediatorDomainStateH\x00R\x13mediatorDomainStateB\t\n\x07mapping\"\x9b\x01\n\x16\x44omainParametersChange\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12i\n\x11\x64omain_parameters\x18\x02 \x01(\x0b\x32<.com.digitalasset.canton.protocol.v0.DynamicDomainParametersR\x10\x64omainParameters\"\xa1\x01\n\x1b\x44omainGovernanceTransaction\x12w\n\x18\x64omain_parameters_change\x18\x01 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v0.DomainParametersChangeH\x00R\x16\x64omainParametersChangeB\t\n\x07mapping\"\xf4\x01\n\x13TopologyTransaction\x12]\n\x0cstate_update\x18\x01 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.TopologyStateUpdateH\x00R\x0bstateUpdate\x12o\n\x11\x64omain_governance\x18\x02 \x01(\x0b\x32@.com.digitalasset.canton.protocol.v0.DomainGovernanceTransactionH\x00R\x10\x64omainGovernanceB\r\n\x0btransaction\"\xd0\x01\n\x19SignedTopologyTransaction\x12 \n\x0btransaction\x18\x01 \x01(\x0cR\x0btransaction\x12\x45\n\x03key\x18\x02 \x01(\x0b\x32\x33.com.digitalasset.canton.crypto.v0.SigningPublicKeyR\x03key\x12J\n\tsignature\x18\x03 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\tsignature\"\xaf\x01\n DomainTopologyTransactionMessage\x12J\n\tsignature\x18\x01 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\tsignature\x12\x1b\n\tdomain_id\x18\x02 \x01(\tR\x08\x64omainId\x12\"\n\x0ctransactions\x18\x03 \x03(\x0cR\x0ctransactions\"\xe7\x01\n\"RegisterTopologyTransactionRequest\x12!\n\x0crequested_by\x18\x01 \x01(\tR\x0brequestedBy\x12 \n\x0bparticipant\x18\x02 \x01(\tR\x0bparticipant\x12\x1d\n\nrequest_id\x18\x03 \x01(\tR\trequestId\x12@\n\x1csigned_topology_transactions\x18\x04 \x03(\x0cR\x1asignedTopologyTransactions\x12\x1b\n\tdomain_id\x18\x05 \x01(\tR\x08\x64omainId\"\xbf\x04\n#RegisterTopologyTransactionResponse\x12!\n\x0crequested_by\x18\x01 \x01(\tR\x0brequestedBy\x12 \n\x0bparticipant\x18\x02 \x01(\tR\x0bparticipant\x12\x1d\n\nrequest_id\x18\x03 \x01(\tR\trequestId\x12i\n\x07results\x18\x04 \x03(\x0b\x32O.com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionResponse.ResultR\x07results\x12\x1b\n\tdomain_id\x18\x05 \x01(\tR\x08\x64omainId\x1a\xab\x02\n\x06Result\x12\x1f\n\x0bunique_path\x18\x01 \x01(\tR\nuniquePath\x12k\n\x05state\x18\x02 \x01(\x0e\x32U.com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionResponse.Result.StateR\x05state\x12#\n\rerror_message\x18\x03 \x01(\tR\x0c\x65rrorMessage\"n\n\x05State\x12\x11\n\rMISSING_STATE\x10\x00\x12\r\n\tREQUESTED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x12\x0c\n\x08REJECTED\x10\x03\x12\x0c\n\x08\x41\x43\x43\x45PTED\x10\x04\x12\r\n\tDUPLICATE\x10\x05\x12\x0c\n\x08OBSOLETE\x10\x06*4\n\x10TopologyChangeOp\x12\x07\n\x03\x41\x64\x64\x10\x00\x12\n\n\x06Remove\x10\x01\x12\x0b\n\x07Replace\x10\x02*:\n\nTrustLevel\x12\x15\n\x11MissingTrustLevel\x10\x00\x12\x0c\n\x08Ordinary\x10\x01\x12\x07\n\x03Vip\x10\x02*z\n\x15ParticipantPermission\x12 \n\x1cMissingParticipantPermission\x10\x00\x12\x0e\n\nSubmission\x10\x01\x12\x10\n\x0c\x43onfirmation\x10\x02\x12\x0f\n\x0bObservation\x10\x03\x12\x0c\n\x08\x44isabled\x10\x04*A\n\x0bRequestSide\x12\x16\n\x12MissingRequestSide\x10\x00\x12\x08\n\x04\x42oth\x10\x01\x12\x08\n\x04\x46rom\x10\x02\x12\x06\n\x02To\x10\x03\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.topology_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_TOPOLOGYCHANGEOP']._serialized_start=4474
  _globals['_TOPOLOGYCHANGEOP']._serialized_end=4526
  _globals['_TRUSTLEVEL']._serialized_start=4528
  _globals['_TRUSTLEVEL']._serialized_end=4586
  _globals['_PARTICIPANTPERMISSION']._serialized_start=4588
  _globals['_PARTICIPANTPERMISSION']._serialized_end=4710
  _globals['_REQUESTSIDE']._serialized_start=4712
  _globals['_REQUESTSIDE']._serialized_end=4777
  _globals['_PARTICIPANTSTATE']._serialized_start=194
  _globals['_PARTICIPANTSTATE']._serialized_end=514
  _globals['_PARTYTOPARTICIPANT']._serialized_start=517
  _globals['_PARTYTOPARTICIPANT']._serialized_end=755
  _globals['_MEDIATORDOMAINSTATE']._serialized_start=758
  _globals['_MEDIATORDOMAINSTATE']._serialized_end=901
  _globals['_NAMESPACEDELEGATION']._serialized_start=904
  _globals['_NAMESPACEDELEGATION']._serialized_end=1085
  _globals['_IDENTIFIERDELEGATION']._serialized_start=1088
  _globals['_IDENTIFIERDELEGATION']._serialized_end=1239
  _globals['_OWNERTOKEYMAPPING']._serialized_start=1241
  _globals['_OWNERTOKEYMAPPING']._serialized_end=1366
  _globals['_SIGNEDLEGALIDENTITYCLAIM']._serialized_start=1368
  _globals['_SIGNEDLEGALIDENTITYCLAIM']._serialized_end=1492
  _globals['_LEGALIDENTITYCLAIM']._serialized_start=1494
  _globals['_LEGALIDENTITYCLAIM']._serialized_end=1602
  _globals['_VETTEDPACKAGES']._serialized_start=1604
  _globals['_VETTEDPACKAGES']._serialized_end=1687
  _globals['_TOPOLOGYSTATEUPDATE']._serialized_start=1690
  _globals['_TOPOLOGYSTATEUPDATE']._serialized_end=2702
  _globals['_DOMAINPARAMETERSCHANGE']._serialized_start=2705
  _globals['_DOMAINPARAMETERSCHANGE']._serialized_end=2860
  _globals['_DOMAINGOVERNANCETRANSACTION']._serialized_start=2863
  _globals['_DOMAINGOVERNANCETRANSACTION']._serialized_end=3024
  _globals['_TOPOLOGYTRANSACTION']._serialized_start=3027
  _globals['_TOPOLOGYTRANSACTION']._serialized_end=3271
  _globals['_SIGNEDTOPOLOGYTRANSACTION']._serialized_start=3274
  _globals['_SIGNEDTOPOLOGYTRANSACTION']._serialized_end=3482
  _globals['_DOMAINTOPOLOGYTRANSACTIONMESSAGE']._serialized_start=3485
  _globals['_DOMAINTOPOLOGYTRANSACTIONMESSAGE']._serialized_end=3660
  _globals['_REGISTERTOPOLOGYTRANSACTIONREQUEST']._serialized_start=3663
  _globals['_REGISTERTOPOLOGYTRANSACTIONREQUEST']._serialized_end=3894
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE']._serialized_start=3897
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE']._serialized_end=4472
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE_RESULT']._serialized_start=4173
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE_RESULT']._serialized_end=4472
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE_RESULT_STATE']._serialized_start=4362
  _globals['_REGISTERTOPOLOGYTRANSACTIONRESPONSE_RESULT_STATE']._serialized_end=4472
# @@protoc_insertion_point(module_scope)