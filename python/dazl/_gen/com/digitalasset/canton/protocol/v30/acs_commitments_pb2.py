# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v30/acs_commitments.proto
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
    'com/digitalasset/canton/protocol/v30/acs_commitments.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n:com/digitalasset/canton/protocol/v30/acs_commitments.proto\x12$com.digitalasset.canton.protocol.v30\"\x86\x02\n\rAcsCommitment\x12\x1b\n\tdomain_id\x18\x01 \x01(\tR\x08\x64omainId\x12\x36\n\x17sending_participant_uid\x18\x02 \x01(\tR\x15sendingParticipantUid\x12\x36\n\x17\x63ounter_participant_uid\x18\x03 \x01(\tR\x15\x63ounterParticipantUid\x12%\n\x0e\x66rom_exclusive\x18\x04 \x01(\x03R\rfromExclusive\x12!\n\x0cto_inclusive\x18\x05 \x01(\x03R\x0btoInclusive\x12\x1e\n\ncommitment\x18\x06 \x01(\x0cR\ncommitmentBUZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v30.acs_commitments_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30'
  _globals['_ACSCOMMITMENT']._serialized_start=101
  _globals['_ACSCOMMITMENT']._serialized_end=363
# @@protoc_insertion_point(module_scope)
