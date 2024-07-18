# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/participant/admin/v0/party_name_management.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nHcom/digitalasset/canton/participant/admin/v0/party_name_management.proto\x12,com.digitalasset.canton.participant.admin.v0\"Z\n\x1aSetPartyDisplayNameRequest\x12\x19\n\x08party_id\x18\x01 \x01(\tR\x07partyId\x12!\n\x0c\x64isplay_name\x18\x02 \x01(\tR\x0b\x64isplayName\"\x1d\n\x1bSetPartyDisplayNameResponse2\xc9\x01\n\x1aPartyNameManagementService\x12\xaa\x01\n\x13setPartyDisplayName\x12H.com.digitalasset.canton.participant.admin.v0.SetPartyDisplayNameRequest\x1aI.com.digitalasset.canton.participant.admin.v0.SetPartyDisplayNameResponseB]Z[github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/participant/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.participant.admin.v0.party_name_management_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z[github.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/participant/admin/v0'
  _globals['_SETPARTYDISPLAYNAMEREQUEST']._serialized_start=122
  _globals['_SETPARTYDISPLAYNAMEREQUEST']._serialized_end=212
  _globals['_SETPARTYDISPLAYNAMERESPONSE']._serialized_start=214
  _globals['_SETPARTYDISPLAYNAMERESPONSE']._serialized_end=243
  _globals['_PARTYNAMEMANAGEMENTSERVICE']._serialized_start=246
  _globals['_PARTYNAMEMANAGEMENTSERVICE']._serialized_end=447
# @@protoc_insertion_point(module_scope)
