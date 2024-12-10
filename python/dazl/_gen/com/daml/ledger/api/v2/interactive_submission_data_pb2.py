# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/daml/ledger/api/v2/interactive_submission_data.proto
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
    'com/daml/ledger/api/v2/interactive_submission_data.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import value_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v2_dot_value__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n8com/daml/ledger/api/v2/interactive_submission_data.proto\x12\x16\x63om.daml.ledger.api.v2\x1a\"com/daml/ledger/api/v2/value.proto\"\x9e\x01\n\x13PreparedTransaction\x12I\n\x0btransaction\x18\x01 \x01(\x0b\x32\'.com.daml.ledger.api.v2.DamlTransactionR\x0btransaction\x12<\n\x08metadata\x18\x02 \x01(\x0b\x32 .com.daml.ledger.api.v2.MetadataR\x08metadata\"\xef\t\n\x08Metadata\x12U\n\x0esubmitter_info\x18\x01 \x01(\x0b\x32..com.daml.ledger.api.v2.Metadata.SubmitterInfoR\rsubmitterInfo\x12\'\n\x0fsubmission_seed\x18\x02 \x01(\x0cR\x0esubmissionSeed\x12\'\n\x0fsubmission_time\x18\x03 \x01(\x04R\x0esubmissionTime\x12#\n\rused_packages\x18\x04 \x03(\tR\x0cusedPackages\x12H\n\nnode_seeds\x18\x05 \x03(\x0b\x32).com.daml.ledger.api.v2.Metadata.NodeSeedR\tnodeSeeds\x12\x64\n\x12global_key_mapping\x18\x06 \x03(\x0b\x32\x36.com.daml.ledger.api.v2.Metadata.GlobalKeyMappingEntryR\x10globalKeyMapping\x12\x66\n\x10\x64isclosed_events\x18\x07 \x03(\x0b\x32;.com.daml.ledger.api.v2.Metadata.ProcessedDisclosedContractR\x0f\x64isclosedEvents\x12\x37\n\x15ledger_effective_time\x18\x08 \x01(\x04H\x00R\x13ledgerEffectiveTime\x88\x01\x01\x12$\n\x0bworkflow_id\x18\t \x01(\tH\x01R\nworkflowId\x88\x01\x01\x12:\n\x19interpretation_time_nanos\x18\n \x01(\x04R\x17interpretationTimeNanos\x12\x1b\n\tdomain_id\x18\x0b \x01(\tR\x08\x64omainId\x12%\n\x0emediator_group\x18\x0c \x01(\rR\rmediatorGroup\x12%\n\x0etransaction_id\x18\r \x01(\tR\rtransactionId\x1a^\n\rSubmitterInfo\x12\x15\n\x06\x61\x63t_as\x18\x01 \x03(\tR\x05\x61\x63tAs\x12\x17\n\x07read_as\x18\x02 \x03(\tR\x06readAs\x12\x1d\n\ncommand_id\x18\x03 \x01(\tR\tcommandId\x1a\x37\n\x08NodeSeed\x12\x17\n\x07node_id\x18\x01 \x01(\x05R\x06nodeId\x12\x12\n\x04seed\x18\x02 \x01(\x0cR\x04seed\x1a\x90\x01\n\x15GlobalKeyMappingEntry\x12\x33\n\x03key\x18\x01 \x01(\x0b\x32!.com.daml.ledger.api.v2.GlobalKeyR\x03key\x12\x38\n\x05value\x18\x02 \x01(\x0b\x32\x1d.com.daml.ledger.api.v2.ValueH\x00R\x05value\x88\x01\x01\x42\x08\n\x06_value\x1a\xa0\x01\n\x1aProcessedDisclosedContract\x12:\n\x08\x63ontract\x18\x01 \x01(\x0b\x32\x1e.com.daml.ledger.api.v2.CreateR\x08\x63ontract\x12\x1d\n\ncreated_at\x18\x02 \x01(\x04R\tcreatedAt\x12\'\n\x0f\x64river_metadata\x18\x03 \x01(\x0cR\x0e\x64riverMetadataB\x18\n\x16_ledger_effective_timeB\x0e\n\x0c_workflow_id\"u\n\x0f\x44\x61mlTransaction\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12\x14\n\x05roots\x18\x02 \x03(\tR\x05roots\x12\x32\n\x05nodes\x18\x03 \x03(\x0b\x32\x1c.com.daml.ledger.api.v2.NodeR\x05nodes\"\x85\x0c\n\x04Node\x12\x1d\n\x07version\x18\x01 \x01(\tH\x01R\x07version\x88\x01\x01\x12\x17\n\x07node_id\x18\x02 \x01(\tR\x06nodeId\x12\x38\n\x06\x63reate\x18\x03 \x01(\x0b\x32\x1e.com.daml.ledger.api.v2.CreateH\x00R\x06\x63reate\x12:\n\x05\x66\x65tch\x18\x04 \x01(\x0b\x32\".com.daml.ledger.api.v2.Node.FetchH\x00R\x05\x66\x65tch\x12\x43\n\x08\x65xercise\x18\x05 \x01(\x0b\x32%.com.daml.ledger.api.v2.Node.ExerciseH\x00R\x08\x65xercise\x12\x43\n\x08rollback\x18\x06 \x01(\x0b\x32%.com.daml.ledger.api.v2.Node.RollbackH\x00R\x08rollback\x12N\n\rlookup_by_key\x18\x07 \x01(\x0b\x32(.com.daml.ledger.api.v2.Node.LookupByKeyH\x00R\x0blookupByKey\x1a\xf8\x02\n\x05\x46\x65tch\x12\x1f\n\x0b\x63ontract_id\x18\x01 \x01(\tR\ncontractId\x12!\n\x0cpackage_name\x18\x02 \x01(\tR\x0bpackageName\x12\x43\n\x0btemplate_id\x18\x03 \x01(\x0b\x32\".com.daml.ledger.api.v2.IdentifierR\ntemplateId\x12 \n\x0bsignatories\x18\x04 \x03(\tR\x0bsignatories\x12\"\n\x0cstakeholders\x18\x05 \x03(\tR\x0cstakeholders\x12%\n\x0e\x61\x63ting_parties\x18\x06 \x03(\tR\ractingParties\x12\x62\n\x14key_with_maintainers\x18\x07 \x01(\x0b\x32\x30.com.daml.ledger.api.v2.GlobalKeyWithMaintainersR\x12keyWithMaintainers\x12\x15\n\x06\x62y_key\x18\x08 \x01(\x08R\x05\x62yKey\x1a\xc6\x03\n\x08\x45xercise\x12\x38\n\x05\x66\x65tch\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v2.Node.FetchR\x05\x66\x65tch\x12\x45\n\x0cinterface_id\x18\x02 \x01(\x0b\x32\".com.daml.ledger.api.v2.IdentifierR\x0binterfaceId\x12\x1b\n\tchoice_id\x18\x03 \x01(\tR\x08\x63hoiceId\x12@\n\x0c\x63hosen_value\x18\x04 \x01(\x0b\x32\x1d.com.daml.ledger.api.v2.ValueR\x0b\x63hosenValue\x12\x1c\n\tconsuming\x18\x05 \x01(\x08R\tconsuming\x12\x1a\n\x08\x63hildren\x18\x06 \x03(\tR\x08\x63hildren\x12\x46\n\x0f\x65xercise_result\x18\x07 \x01(\x0b\x32\x1d.com.daml.ledger.api.v2.ValueR\x0e\x65xerciseResult\x12)\n\x10\x63hoice_observers\x18\x08 \x03(\tR\x0f\x63hoiceObservers\x12-\n\x12\x63hoice_authorizers\x18\t \x03(\tR\x11\x63hoiceAuthorizers\x1a&\n\x08Rollback\x12\x1a\n\x08\x63hildren\x18\x01 \x03(\tR\x08\x63hildren\x1a\xef\x01\n\x0bLookupByKey\x12!\n\x0cpackage_name\x18\x01 \x01(\tR\x0bpackageName\x12\x43\n\x0btemplate_id\x18\x02 \x01(\x0b\x32\".com.daml.ledger.api.v2.IdentifierR\ntemplateId\x12\x42\n\x03key\x18\x03 \x01(\x0b\x32\x30.com.daml.ledger.api.v2.GlobalKeyWithMaintainersR\x03key\x12$\n\x0b\x63ontract_id\x18\x04 \x01(\tH\x00R\ncontractId\x88\x01\x01\x42\x0e\n\x0c_contract_idB\x0b\n\tnode_typeB\n\n\x08_version\"\xc5\x03\n\x06\x43reate\x12\x1f\n\x0b\x63ontract_id\x18\x01 \x01(\tR\ncontractId\x12!\n\x0cpackage_name\x18\x02 \x01(\tR\x0bpackageName\x12\x43\n\x0btemplate_id\x18\x03 \x01(\x0b\x32\".com.daml.ledger.api.v2.IdentifierR\ntemplateId\x12\x39\n\x08\x61rgument\x18\x04 \x01(\x0b\x32\x1d.com.daml.ledger.api.v2.ValueR\x08\x61rgument\x12 \n\x0bsignatories\x18\x05 \x03(\tR\x0bsignatories\x12\"\n\x0cstakeholders\x18\x06 \x03(\tR\x0cstakeholders\x12,\n\x0fpackage_version\x18\x07 \x01(\tH\x00R\x0epackageVersion\x88\x01\x01\x12o\n\x1bglobal_key_with_maintainers\x18\x08 \x01(\x0b\x32\x30.com.daml.ledger.api.v2.GlobalKeyWithMaintainersR\x18globalKeyWithMaintainersB\x12\n\x10_package_version\"\xb8\x01\n\tGlobalKey\x12\x43\n\x0btemplate_id\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v2.IdentifierR\ntemplateId\x12!\n\x0cpackage_name\x18\x02 \x01(\tR\x0bpackageName\x12/\n\x03key\x18\x03 \x01(\x0b\x32\x1d.com.daml.ledger.api.v2.ValueR\x03key\x12\x12\n\x04hash\x18\x04 \x01(\x0cR\x04hash\"q\n\x18GlobalKeyWithMaintainers\x12\x33\n\x03key\x18\x01 \x01(\x0b\x32!.com.daml.ledger.api.v2.GlobalKeyR\x03key\x12 \n\x0bmaintainers\x18\x02 \x03(\tR\x0bmaintainersB\x9d\x01\n\x16\x63om.daml.ledger.api.v2B#InteractiveSubmissionDataOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2\xaa\x02\x16\x43om.Daml.Ledger.Api.V2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v2.interactive_submission_data_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.daml.ledger.api.v2B#InteractiveSubmissionDataOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2\252\002\026Com.Daml.Ledger.Api.V2'
  _globals['_PREPAREDTRANSACTION']._serialized_start=121
  _globals['_PREPAREDTRANSACTION']._serialized_end=279
  _globals['_METADATA']._serialized_start=282
  _globals['_METADATA']._serialized_end=1545
  _globals['_METADATA_SUBMITTERINFO']._serialized_start=1042
  _globals['_METADATA_SUBMITTERINFO']._serialized_end=1136
  _globals['_METADATA_NODESEED']._serialized_start=1138
  _globals['_METADATA_NODESEED']._serialized_end=1193
  _globals['_METADATA_GLOBALKEYMAPPINGENTRY']._serialized_start=1196
  _globals['_METADATA_GLOBALKEYMAPPINGENTRY']._serialized_end=1340
  _globals['_METADATA_PROCESSEDDISCLOSEDCONTRACT']._serialized_start=1343
  _globals['_METADATA_PROCESSEDDISCLOSEDCONTRACT']._serialized_end=1503
  _globals['_DAMLTRANSACTION']._serialized_start=1547
  _globals['_DAMLTRANSACTION']._serialized_end=1664
  _globals['_NODE']._serialized_start=1667
  _globals['_NODE']._serialized_end=3208
  _globals['_NODE_FETCH']._serialized_start=2068
  _globals['_NODE_FETCH']._serialized_end=2444
  _globals['_NODE_EXERCISE']._serialized_start=2447
  _globals['_NODE_EXERCISE']._serialized_end=2901
  _globals['_NODE_ROLLBACK']._serialized_start=2903
  _globals['_NODE_ROLLBACK']._serialized_end=2941
  _globals['_NODE_LOOKUPBYKEY']._serialized_start=2944
  _globals['_NODE_LOOKUPBYKEY']._serialized_end=3183
  _globals['_CREATE']._serialized_start=3211
  _globals['_CREATE']._serialized_end=3664
  _globals['_GLOBALKEY']._serialized_start=3667
  _globals['_GLOBALKEY']._serialized_end=3851
  _globals['_GLOBALKEYWITHMAINTAINERS']._serialized_start=3853
  _globals['_GLOBALKEYWITHMAINTAINERS']._serialized_end=3966
# @@protoc_insertion_point(module_scope)
