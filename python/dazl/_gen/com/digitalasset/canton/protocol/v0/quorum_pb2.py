# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v0/quorum.proto
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
    'com/digitalasset/canton/protocol/v0/quorum.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import topology_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0com/digitalasset/canton/protocol/v0/quorum.proto\x12#com.digitalasset.canton.protocol.v0\x1a\x32\x63om/digitalasset/canton/protocol/v0/topology.proto\"\x95\x01\n\x06Quorum\x12m\n\x16party_index_and_weight\x18\x01 \x03(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.PartyIndexAndWeightR\x13partyIndexAndWeight\x12\x1c\n\tthreshold\x18\x02 \x01(\x05R\tthreshold\"C\n\x13PartyIndexAndWeight\x12\x14\n\x05index\x18\x01 \x01(\x05R\x05index\x12\x16\n\x06weight\x18\x02 \x01(\x05R\x06weight\"\x85\x01\n\nTrustParty\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x12\x61\n\x14required_trust_level\x18\x02 \x01(\x0e\x32/.com.digitalasset.canton.protocol.v0.TrustLevelR\x12requiredTrustLevelBTZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.quorum_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v0'
  _globals['_QUORUM']._serialized_start=142
  _globals['_QUORUM']._serialized_end=291
  _globals['_PARTYINDEXANDWEIGHT']._serialized_start=293
  _globals['_PARTYINDEXANDWEIGHT']._serialized_end=360
  _globals['_TRUSTPARTY']._serialized_start=363
  _globals['_TRUSTPARTY']._serialized_end=496
# @@protoc_insertion_point(module_scope)
