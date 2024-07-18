# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/ledger_identity_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4com/daml/ledger/api/v1/ledger_identity_service.proto\x12\x16\x63om.daml.ledger.api.v1\"\x1e\n\x18GetLedgerIdentityRequest:\x02\x18\x01\"@\n\x19GetLedgerIdentityResponse\x12\x1f\n\tledger_id\x18\x01 \x01(\tB\x02\x18\x01R\x08ledgerId:\x02\x18\x01\x32\x9b\x01\n\x15LedgerIdentityService\x12}\n\x11GetLedgerIdentity\x12\x30.com.daml.ledger.api.v1.GetLedgerIdentityRequest\x1a\x31.com.daml.ledger.api.v1.GetLedgerIdentityResponse\"\x03\x88\x02\x01\x1a\x03\x88\x02\x01\x42\x99\x01\n\x16\x63om.daml.ledger.api.v1B\x1fLedgerIdentityServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.ledger_identity_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.daml.ledger.api.v1B\037LedgerIdentityServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\252\002\026Com.Daml.Ledger.Api.V1'
  _globals['_GETLEDGERIDENTITYREQUEST']._loaded_options = None
  _globals['_GETLEDGERIDENTITYREQUEST']._serialized_options = b'\030\001'
  _globals['_GETLEDGERIDENTITYRESPONSE'].fields_by_name['ledger_id']._loaded_options = None
  _globals['_GETLEDGERIDENTITYRESPONSE'].fields_by_name['ledger_id']._serialized_options = b'\030\001'
  _globals['_GETLEDGERIDENTITYRESPONSE']._loaded_options = None
  _globals['_GETLEDGERIDENTITYRESPONSE']._serialized_options = b'\030\001'
  _globals['_LEDGERIDENTITYSERVICE']._loaded_options = None
  _globals['_LEDGERIDENTITYSERVICE']._serialized_options = b'\210\002\001'
  _globals['_LEDGERIDENTITYSERVICE'].methods_by_name['GetLedgerIdentity']._loaded_options = None
  _globals['_LEDGERIDENTITYSERVICE'].methods_by_name['GetLedgerIdentity']._serialized_options = b'\210\002\001'
  _globals['_GETLEDGERIDENTITYREQUEST']._serialized_start=80
  _globals['_GETLEDGERIDENTITYREQUEST']._serialized_end=110
  _globals['_GETLEDGERIDENTITYRESPONSE']._serialized_start=112
  _globals['_GETLEDGERIDENTITYRESPONSE']._serialized_end=176
  _globals['_LEDGERIDENTITYSERVICE']._serialized_start=179
  _globals['_LEDGERIDENTITYSERVICE']._serialized_end=334
# @@protoc_insertion_point(module_scope)
