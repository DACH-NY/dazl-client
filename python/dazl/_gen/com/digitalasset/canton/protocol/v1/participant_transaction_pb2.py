# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v1/participant_transaction.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'com/digitalasset/canton/protocol/v1/participant_transaction.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from ..v0 import common_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_common__pb2
from ..v0 import participant_transaction_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_participant__transaction__pb2
from ..v0 import topology_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__pb2
from . import common_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_common__pb2
from . import merkle_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_merkle__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nAcom/digitalasset/canton/protocol/v1/participant_transaction.proto\x12#com.digitalasset.canton.protocol.v1\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x30\x63om/digitalasset/canton/protocol/v0/common.proto\x1a\x41\x63om/digitalasset/canton/protocol/v0/participant_transaction.proto\x1a\x32\x63om/digitalasset/canton/protocol/v0/topology.proto\x1a\x30\x63om/digitalasset/canton/protocol/v1/common.proto\x1a\x30\x63om/digitalasset/canton/protocol/v1/merkle.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc2\x08\n\x11\x41\x63tionDescription\x12h\n\x06\x63reate\x18\x01 \x01(\x0b\x32N.com.digitalasset.canton.protocol.v0.ActionDescription.CreateActionDescriptionH\x00R\x06\x63reate\x12n\n\x08\x65xercise\x18\x02 \x01(\x0b\x32P.com.digitalasset.canton.protocol.v1.ActionDescription.ExerciseActionDescriptionH\x00R\x08\x65xercise\x12\x65\n\x05\x66\x65tch\x18\x03 \x01(\x0b\x32M.com.digitalasset.canton.protocol.v0.ActionDescription.FetchActionDescriptionH\x00R\x05\x66\x65tch\x12y\n\rlookup_by_key\x18\x04 \x01(\x0b\x32S.com.digitalasset.canton.protocol.v0.ActionDescription.LookupByKeyActionDescriptionH\x00R\x0blookupByKey\x1a\xb9\x02\n\x19\x45xerciseActionDescription\x12*\n\x11input_contract_id\x18\x01 \x01(\tR\x0finputContractId\x12\x16\n\x06\x63hoice\x18\x02 \x01(\tR\x06\x63hoice\x12!\n\x0c\x63hosen_value\x18\x03 \x01(\x0cR\x0b\x63hosenValue\x12\x16\n\x06\x61\x63tors\x18\x04 \x03(\tR\x06\x61\x63tors\x12\x15\n\x06\x62y_key\x18\x05 \x01(\x08R\x05\x62yKey\x12\x1b\n\tnode_seed\x18\x06 \x01(\x0cR\x08nodeSeed\x12\x18\n\x07version\x18\x07 \x01(\tR\x07version\x12\x16\n\x06\x66\x61iled\x18\x08 \x01(\x08R\x06\x66\x61iled\x12&\n\x0cinterface_id\x18\t \x01(\tH\x00R\x0binterfaceId\x88\x01\x01\x42\x0f\n\r_interface_id\x1a`\n\x1cLookupByKeyActionDescription\x12@\n\x03key\x18\x01 \x01(\x0b\x32..com.digitalasset.canton.protocol.v1.GlobalKeyR\x03key\x1a\xc3\x01\n\x16\x46\x65tchActionDescription\x12*\n\x11input_contract_id\x18\x01 \x01(\tR\x0finputContractId\x12\x16\n\x06\x61\x63tors\x18\x02 \x03(\tR\x06\x61\x63tors\x12\x15\n\x06\x62y_key\x18\x03 \x01(\x08R\x05\x62yKey\x12\x18\n\x07version\x18\x04 \x01(\tR\x07version\x12$\n\x0btemplate_id\x18\x05 \x01(\tH\x00R\ntemplateId\x88\x01\x01\x42\x0e\n\x0c_template_idB\r\n\x0b\x64\x65scription\"\x9c\x02\n\x08ViewNode\x12\\\n\x10view_common_data\x18\x01 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v1.BlindableNodeR\x0eviewCommonData\x12\x66\n\x15view_participant_data\x18\x02 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v1.BlindableNodeR\x13viewParticipantData\x12J\n\x08subviews\x18\x03 \x01(\x0b\x32..com.digitalasset.canton.protocol.v1.MerkleSeqR\x08subviews\"\xb8\x01\n\x0eViewCommonData\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12K\n\tinformees\x18\x02 \x03(\x0b\x32-.com.digitalasset.canton.protocol.v1.InformeeR\tinformees\x12\x1c\n\tthreshold\x18\x03 \x01(\x05R\tthreshold\"\x9b\x01\n\x08Informee\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x12\x16\n\x06weight\x18\x02 \x01(\x05R\x06weight\x12\x61\n\x14required_trust_level\x18\x03 \x01(\x0e\x32/.com.digitalasset.canton.protocol.v0.TrustLevelR\x12requiredTrustLevel\"\xf5\x03\n\x14\x45ncryptedViewMessage\x12\x1b\n\tview_tree\x18\x01 \x01(\x0cR\x08viewTree\x12\x62\n\x11\x65ncryption_scheme\x18\x02 \x01(\x0e\x32\x35.com.digitalasset.canton.crypto.v0.SymmetricKeySchemeR\x10\x65ncryptionScheme\x12t\n\x1fsubmitter_participant_signature\x18\x03 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\x1dsubmitterParticipantSignature\x12\x1b\n\tview_hash\x18\x04 \x01(\x0cR\x08viewHash\x12`\n\nrandomness\x18\x05 \x03(\x0b\x32@.com.digitalasset.canton.protocol.v1.ParticipantRandomnessLookupR\nrandomness\x12\x1b\n\tdomain_id\x18\x06 \x01(\tR\x08\x64omainId\x12J\n\tview_type\x18\x07 \x01(\x0e\x32-.com.digitalasset.canton.protocol.v0.ViewTypeR\x08viewType\"_\n\x1bParticipantRandomnessLookup\x12\x1e\n\nrandomness\x18\x01 \x01(\x0cR\nrandomness\x12 \n\x0b\x66ingerprint\x18\x02 \x01(\tR\x0b\x66ingerprint\"\x1c\n\x16ViewParticipantMessage:\x02\x18\x01\"\xa1\x01\n\x0fInformeeMessage\x12\x63\n\x12\x66ull_informee_tree\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v1.FullInformeeTreeR\x10\x66ullInformeeTree\x12)\n\x10protocol_version\x18\x02 \x01(\x05R\x0fprotocolVersion\"\x8e\x01\n\x18LightTransactionViewTree\x12K\n\x04tree\x18\x01 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v1.GenTransactionTreeR\x04tree\x12%\n\x0esubview_hashes\x18\x02 \x03(\x0cR\rsubviewHashes\"_\n\x10\x46ullInformeeTree\x12K\n\x04tree\x18\x01 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v1.GenTransactionTreeR\x04tree\"\xb3\x01\n\x0f\x43reatedContract\x12U\n\x08\x63ontract\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v1.SerializableContractR\x08\x63ontract\x12(\n\x10\x63onsumed_in_core\x18\x02 \x01(\x08R\x0e\x63onsumedInCore\x12\x1f\n\x0brolled_back\x18\x03 \x01(\x08R\nrolledBack\"\x82\x01\n\rInputContract\x12U\n\x08\x63ontract\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v1.SerializableContractR\x08\x63ontract\x12\x1a\n\x08\x63onsumed\x18\x02 \x01(\x08R\x08\x63onsumed\"\xb0\x03\n\x11SubmitterMetadata\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12\x15\n\x06\x61\x63t_as\x18\x02 \x03(\tR\x05\x61\x63tAs\x12%\n\x0e\x61pplication_id\x18\x03 \x01(\tR\rapplicationId\x12\x1d\n\ncommand_id\x18\x04 \x01(\tR\tcommandId\x12\x33\n\x15submitter_participant\x18\x05 \x01(\tR\x14submitterParticipant\x12#\n\rsubmission_id\x18\x06 \x01(\tR\x0csubmissionId\x12[\n\x0c\x64\x65\x64up_period\x18\x07 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.DeduplicationPeriodR\x0b\x64\x65\x64upPeriod\x12J\n\x13max_sequencing_time\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x11maxSequencingTime\"\xd8\x01\n\x0bResolvedKey\x12@\n\x03key\x18\x01 \x01(\x0b\x32..com.digitalasset.canton.protocol.v1.GlobalKeyR\x03key\x12!\n\x0b\x63ontract_id\x18\x02 \x01(\tH\x00R\ncontractId\x12V\n\x04\x66ree\x18\x03 \x01(\x0b\x32@.com.digitalasset.canton.protocol.v0.ViewParticipantData.FreeKeyH\x00R\x04\x66reeB\x0c\n\nresolutionBTZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v1.participant_transaction_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v1'
  _globals['_VIEWPARTICIPANTMESSAGE']._loaded_options = None
  _globals['_VIEWPARTICIPANTMESSAGE']._serialized_options = b'\030\001'
  _globals['_ACTIONDESCRIPTION']._serialized_start=457
  _globals['_ACTIONDESCRIPTION']._serialized_end=1547
  _globals['_ACTIONDESCRIPTION_EXERCISEACTIONDESCRIPTION']._serialized_start=923
  _globals['_ACTIONDESCRIPTION_EXERCISEACTIONDESCRIPTION']._serialized_end=1236
  _globals['_ACTIONDESCRIPTION_LOOKUPBYKEYACTIONDESCRIPTION']._serialized_start=1238
  _globals['_ACTIONDESCRIPTION_LOOKUPBYKEYACTIONDESCRIPTION']._serialized_end=1334
  _globals['_ACTIONDESCRIPTION_FETCHACTIONDESCRIPTION']._serialized_start=1337
  _globals['_ACTIONDESCRIPTION_FETCHACTIONDESCRIPTION']._serialized_end=1532
  _globals['_VIEWNODE']._serialized_start=1550
  _globals['_VIEWNODE']._serialized_end=1834
  _globals['_VIEWCOMMONDATA']._serialized_start=1837
  _globals['_VIEWCOMMONDATA']._serialized_end=2021
  _globals['_INFORMEE']._serialized_start=2024
  _globals['_INFORMEE']._serialized_end=2179
  _globals['_ENCRYPTEDVIEWMESSAGE']._serialized_start=2182
  _globals['_ENCRYPTEDVIEWMESSAGE']._serialized_end=2683
  _globals['_PARTICIPANTRANDOMNESSLOOKUP']._serialized_start=2685
  _globals['_PARTICIPANTRANDOMNESSLOOKUP']._serialized_end=2780
  _globals['_VIEWPARTICIPANTMESSAGE']._serialized_start=2782
  _globals['_VIEWPARTICIPANTMESSAGE']._serialized_end=2810
  _globals['_INFORMEEMESSAGE']._serialized_start=2813
  _globals['_INFORMEEMESSAGE']._serialized_end=2974
  _globals['_LIGHTTRANSACTIONVIEWTREE']._serialized_start=2977
  _globals['_LIGHTTRANSACTIONVIEWTREE']._serialized_end=3119
  _globals['_FULLINFORMEETREE']._serialized_start=3121
  _globals['_FULLINFORMEETREE']._serialized_end=3216
  _globals['_CREATEDCONTRACT']._serialized_start=3219
  _globals['_CREATEDCONTRACT']._serialized_end=3398
  _globals['_INPUTCONTRACT']._serialized_start=3401
  _globals['_INPUTCONTRACT']._serialized_end=3531
  _globals['_SUBMITTERMETADATA']._serialized_start=3534
  _globals['_SUBMITTERMETADATA']._serialized_end=3966
  _globals['_RESOLVEDKEY']._serialized_start=3969
  _globals['_RESOLVEDKEY']._serialized_end=4185
# @@protoc_insertion_point(module_scope)
