# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/ledger_configuration_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n9com/daml/ledger/api/v1/ledger_configuration_service.proto\x12\x16\x63om.daml.ledger.api.v1\x1a\x1egoogle/protobuf/duration.proto\"<\n\x1dGetLedgerConfigurationRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\"\x80\x01\n\x1eGetLedgerConfigurationResponse\x12^\n\x14ledger_configuration\x18\x01 \x01(\x0b\x32+.com.daml.ledger.api.v1.LedgerConfigurationR\x13ledgerConfiguration\"z\n\x13LedgerConfiguration\x12W\n\x1amax_deduplication_duration\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x18maxDeduplicationDurationJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\x32\xa8\x01\n\x1aLedgerConfigurationService\x12\x89\x01\n\x16GetLedgerConfiguration\x12\x35.com.daml.ledger.api.v1.GetLedgerConfigurationRequest\x1a\x36.com.daml.ledger.api.v1.GetLedgerConfigurationResponse0\x01\x42\x9e\x01\n\x16\x63om.daml.ledger.api.v1B$LedgerConfigurationServiceOuterClassZEgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.ledger_configuration_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026com.daml.ledger.api.v1B$LedgerConfigurationServiceOuterClassZEgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1\252\002\026Com.Daml.Ledger.Api.V1'
  _globals['_GETLEDGERCONFIGURATIONREQUEST']._serialized_start=117
  _globals['_GETLEDGERCONFIGURATIONREQUEST']._serialized_end=177
  _globals['_GETLEDGERCONFIGURATIONRESPONSE']._serialized_start=180
  _globals['_GETLEDGERCONFIGURATIONRESPONSE']._serialized_end=308
  _globals['_LEDGERCONFIGURATION']._serialized_start=310
  _globals['_LEDGERCONFIGURATION']._serialized_end=432
  _globals['_LEDGERCONFIGURATIONSERVICE']._serialized_start=435
  _globals['_LEDGERCONFIGURATIONSERVICE']._serialized_end=603
# @@protoc_insertion_point(module_scope)
