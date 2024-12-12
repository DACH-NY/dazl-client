# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/domain/admin/v0/sequencer_initialization_snapshot.proto
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
    'com/digitalasset/canton/domain/admin/v0/sequencer_initialization_snapshot.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import sequencer_administration_service_pb2 as com_dot_digitalasset_dot_canton_dot_domain_dot_admin_dot_v0_dot_sequencer__administration__service__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nOcom/digitalasset/canton/domain/admin/v0/sequencer_initialization_snapshot.proto\x12\'com.digitalasset.canton.domain.admin.v0\x1aNcom/digitalasset/canton/domain/admin/v0/sequencer_administration_service.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe6\x03\n\x11SequencerSnapshot\x12\x45\n\x10latest_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0flatestTimestamp\x12\x84\x01\n\x14head_member_counters\x18\x02 \x03(\x0b\x32R.com.digitalasset.canton.domain.admin.v0.SequencerSnapshot.HeadMemberCountersEntryR\x12headMemberCounters\x12W\n\x06status\x18\x03 \x01(\x0b\x32?.com.digitalasset.canton.domain.admin.v0.SequencerPruningStatusR\x06status\x12\x63\n\nadditional\x18\x04 \x01(\x0b\x32\x43.com.digitalasset.canton.domain.admin.v0.ImplementationSpecificInfoR\nadditional\x1a\x45\n\x17HeadMemberCountersEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x03R\x05value:\x02\x38\x01\"a\n\x1aImplementationSpecificInfo\x12/\n\x13implementation_name\x18\x01 \x01(\tR\x12implementationName\x12\x12\n\x04info\x18\x02 \x01(\x0cR\x04infoBXZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.admin.v0.sequencer_initialization_snapshot_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/admin/v0'
  _globals['_SEQUENCERSNAPSHOT_HEADMEMBERCOUNTERSENTRY']._loaded_options = None
  _globals['_SEQUENCERSNAPSHOT_HEADMEMBERCOUNTERSENTRY']._serialized_options = b'8\001'
  _globals['_SEQUENCERSNAPSHOT']._serialized_start=238
  _globals['_SEQUENCERSNAPSHOT']._serialized_end=724
  _globals['_SEQUENCERSNAPSHOT_HEADMEMBERCOUNTERSENTRY']._serialized_start=655
  _globals['_SEQUENCERSNAPSHOT_HEADMEMBERCOUNTERSENTRY']._serialized_end=724
  _globals['_IMPLEMENTATIONSPECIFICINFO']._serialized_start=726
  _globals['_IMPLEMENTATIONSPECIFICINFO']._serialized_end=823
# @@protoc_insertion_point(module_scope)
