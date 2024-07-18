# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/active_contracts_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import event_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_event__pb2
from . import transaction_filter_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_transaction__filter__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n5com/daml/ledger/api/v1/active_contracts_service.proto\x12\x16\x63om.daml.ledger.api.v1\x1a\"com/daml/ledger/api/v1/event.proto\x1a/com/daml/ledger/api/v1/transaction_filter.proto\"\xbf\x01\n\x19GetActiveContractsRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\x12\x41\n\x06\x66ilter\x18\x02 \x01(\x0b\x32).com.daml.ledger.api.v1.TransactionFilterR\x06\x66ilter\x12\x18\n\x07verbose\x18\x03 \x01(\x08R\x07verbose\x12(\n\x10\x61\x63tive_at_offset\x18\x04 \x01(\tR\x0e\x61\x63tiveAtOffset\"\xa6\x01\n\x1aGetActiveContractsResponse\x12\x16\n\x06offset\x18\x01 \x01(\tR\x06offset\x12\x1f\n\x0bworkflow_id\x18\x02 \x01(\tR\nworkflowId\x12O\n\x10\x61\x63tive_contracts\x18\x03 \x03(\x0b\x32$.com.daml.ledger.api.v1.CreatedEventR\x0f\x61\x63tiveContracts2\x97\x01\n\x16\x41\x63tiveContractsService\x12}\n\x12GetActiveContracts\x12\x31.com.daml.ledger.api.v1.GetActiveContractsRequest\x1a\x32.com.daml.ledger.api.v1.GetActiveContractsResponse0\x01\x42\x9a\x01\n\x16\x63om.daml.ledger.api.v1B ActiveContractsServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.active_contracts_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.daml.ledger.api.v1B ActiveContractsServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\252\002\026Com.Daml.Ledger.Api.V1'
  _globals['_GETACTIVECONTRACTSREQUEST']._serialized_start=167
  _globals['_GETACTIVECONTRACTSREQUEST']._serialized_end=358
  _globals['_GETACTIVECONTRACTSRESPONSE']._serialized_start=361
  _globals['_GETACTIVECONTRACTSRESPONSE']._serialized_end=527
  _globals['_ACTIVECONTRACTSSERVICE']._serialized_start=530
  _globals['_ACTIVECONTRACTSSERVICE']._serialized_end=681
# @@protoc_insertion_point(module_scope)
