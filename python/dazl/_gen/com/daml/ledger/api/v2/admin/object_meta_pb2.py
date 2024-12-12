# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/daml/ledger/api/v2/admin/object_meta.proto
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
    'com/daml/ledger/api/v2/admin/object_meta.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.com/daml/ledger/api/v2/admin/object_meta.proto\x12\x1c\x63om.daml.ledger.api.v2.admin\"\xd4\x01\n\nObjectMeta\x12)\n\x10resource_version\x18\x06 \x01(\tR\x0fresourceVersion\x12[\n\x0b\x61nnotations\x18\x0c \x03(\x0b\x32\x39.com.daml.ledger.api.v2.admin.ObjectMeta.AnnotationsEntryR\x0b\x61nnotations\x1a>\n\x10\x41nnotationsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\xa0\x01\n\x1c\x63om.daml.ledger.api.v2.adminB\x14ObjectMetaOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V2.Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v2.admin.object_meta_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\034com.daml.ledger.api.v2.adminB\024ObjectMetaOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2/admin\252\002\034Com.Daml.Ledger.Api.V2.Admin'
  _globals['_OBJECTMETA_ANNOTATIONSENTRY']._loaded_options = None
  _globals['_OBJECTMETA_ANNOTATIONSENTRY']._serialized_options = b'8\001'
  _globals['_OBJECTMETA']._serialized_start=81
  _globals['_OBJECTMETA']._serialized_end=293
  _globals['_OBJECTMETA_ANNOTATIONSENTRY']._serialized_start=231
  _globals['_OBJECTMETA_ANNOTATIONSENTRY']._serialized_end=293
# @@protoc_insertion_point(module_scope)