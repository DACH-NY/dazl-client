# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/admin/participant_pruning_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n>com/daml/ledger/api/v1/admin/participant_pruning_service.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\"\x94\x01\n\x0cPruneRequest\x12\x1e\n\x0bprune_up_to\x18\x01 \x01(\tR\tpruneUpTo\x12#\n\rsubmission_id\x18\x02 \x01(\tR\x0csubmissionId\x12?\n\x1cprune_all_divulged_contracts\x18\x03 \x01(\x08R\x19pruneAllDivulgedContracts\"\x0f\n\rPruneResponse2}\n\x19ParticipantPruningService\x12`\n\x05Prune\x12*.com.daml.ledger.api.v1.admin.PruneRequest\x1a+.com.daml.ledger.api.v1.admin.PruneResponseB\xaf\x01\n\x1c\x63om.daml.ledger.api.v1.adminB#ParticipantPruningServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.admin.participant_pruning_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\034com.daml.ledger.api.v1.adminB#ParticipantPruningServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\252\002\034Com.Daml.Ledger.Api.V1.Admin'
  _globals['_PRUNEREQUEST']._serialized_start=97
  _globals['_PRUNEREQUEST']._serialized_end=245
  _globals['_PRUNERESPONSE']._serialized_start=247
  _globals['_PRUNERESPONSE']._serialized_end=262
  _globals['_PARTICIPANTPRUNINGSERVICE']._serialized_start=264
  _globals['_PARTICIPANTPRUNINGSERVICE']._serialized_end=389
# @@protoc_insertion_point(module_scope)
