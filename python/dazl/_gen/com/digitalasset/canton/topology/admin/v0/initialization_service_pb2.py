# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/topology/admin/v0/initialization_service.proto
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
    'com/digitalasset/canton/topology/admin/v0/initialization_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nFcom/digitalasset/canton/topology/admin/v0/initialization_service.proto\x12)com.digitalasset.canton.topology.admin.v0\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"m\n\rInitIdRequest\x12\x1e\n\nidentifier\x18\x01 \x01(\tR\nidentifier\x12 \n\x0b\x66ingerprint\x18\x02 \x01(\tR\x0b\x66ingerprint\x12\x1a\n\x08instance\x18\x03 \x01(\tR\x08instance\"Y\n\x0eInitIdResponse\x12+\n\x11unique_identifier\x18\x01 \x01(\tR\x10uniqueIdentifier\x12\x1a\n\x08instance\x18\x02 \x01(\tR\x08instance\"z\n\rGetIdResponse\x12 \n\x0binitialized\x18\x01 \x01(\x08R\x0binitialized\x12+\n\x11unique_identifier\x18\x02 \x01(\tR\x10uniqueIdentifier\x12\x1a\n\x08instance\x18\x03 \x01(\tR\x08instance2\xb4\x02\n\x15InitializationService\x12}\n\x06InitId\x12\x38.com.digitalasset.canton.topology.admin.v0.InitIdRequest\x1a\x39.com.digitalasset.canton.topology.admin.v0.InitIdResponse\x12Y\n\x05GetId\x12\x16.google.protobuf.Empty\x1a\x38.com.digitalasset.canton.topology.admin.v0.GetIdResponse\x12\x41\n\x0b\x43urrentTime\x12\x16.google.protobuf.Empty\x1a\x1a.google.protobuf.TimestampBZZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.topology.admin.v0.initialization_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v0'
  _globals['_INITIDREQUEST']._serialized_start=179
  _globals['_INITIDREQUEST']._serialized_end=288
  _globals['_INITIDRESPONSE']._serialized_start=290
  _globals['_INITIDRESPONSE']._serialized_end=379
  _globals['_GETIDRESPONSE']._serialized_start=381
  _globals['_GETIDRESPONSE']._serialized_end=503
  _globals['_INITIALIZATIONSERVICE']._serialized_start=506
  _globals['_INITIALIZATIONSERVICE']._serialized_end=814
# @@protoc_insertion_point(module_scope)
