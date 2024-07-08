# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/admin/object_meta.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.com/daml/ledger/api/v1/admin/object_meta.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\"\xd4\x01\n\nObjectMeta\x12)\n\x10resource_version\x18\x06 \x01(\tR\x0fresourceVersion\x12[\n\x0b\x61nnotations\x18\x0c \x03(\x0b\x32\x39.com.daml.ledger.api.v1.admin.ObjectMeta.AnnotationsEntryR\x0b\x61nnotations\x1a>\n\x10\x41nnotationsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\xa0\x01\n\x1c\x63om.daml.ledger.api.v1.adminB\x14ObjectMetaOuterClassZKgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3')



_OBJECTMETA = DESCRIPTOR.message_types_by_name['ObjectMeta']
_OBJECTMETA_ANNOTATIONSENTRY = _OBJECTMETA.nested_types_by_name['AnnotationsEntry']
ObjectMeta = _reflection.GeneratedProtocolMessageType('ObjectMeta', (_message.Message,), {

  'AnnotationsEntry' : _reflection.GeneratedProtocolMessageType('AnnotationsEntry', (_message.Message,), {
    'DESCRIPTOR' : _OBJECTMETA_ANNOTATIONSENTRY,
    '__module__' : 'com.daml.ledger.api.v1.admin.object_meta_pb2'
    # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.ObjectMeta.AnnotationsEntry)
    })
  ,
  'DESCRIPTOR' : _OBJECTMETA,
  '__module__' : 'com.daml.ledger.api.v1.admin.object_meta_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.ObjectMeta)
  })
_sym_db.RegisterMessage(ObjectMeta)
_sym_db.RegisterMessage(ObjectMeta.AnnotationsEntry)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\034com.daml.ledger.api.v1.adminB\024ObjectMetaOuterClassZKgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1/admin\252\002\034Com.Daml.Ledger.Api.V1.Admin'
  _OBJECTMETA_ANNOTATIONSENTRY._options = None
  _OBJECTMETA_ANNOTATIONSENTRY._serialized_options = b'8\001'
  _OBJECTMETA._serialized_start=81
  _OBJECTMETA._serialized_end=293
  _OBJECTMETA_ANNOTATIONSENTRY._serialized_start=231
  _OBJECTMETA_ANNOTATIONSENTRY._serialized_end=293
# @@protoc_insertion_point(module_scope)
