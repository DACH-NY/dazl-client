# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v30/common.proto
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
    'com/digitalasset/canton/protocol/v30/common.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v30 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v30_dot_crypto__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1com/digitalasset/canton/protocol/v30/common.proto\x12$com.digitalasset.canton.protocol.v30\x1a/com/digitalasset/canton/crypto/v30/crypto.proto\"g\n\x16\x44riverContractMetadata\x12M\n\rcontract_salt\x18\x01 \x01(\x0b\x32(.com.digitalasset.canton.crypto.v30.SaltR\x0c\x63ontractSalt*w\n\x08ViewType\x12\x19\n\x15VIEW_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15VIEW_TYPE_TRANSACTION\x10\x01\x12\x1a\n\x16VIEW_TYPE_TRANSFER_OUT\x10\x02\x12\x19\n\x15VIEW_TYPE_TRANSFER_IN\x10\x03\x42UZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v30.common_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30'
  _globals['_VIEWTYPE']._serialized_start=245
  _globals['_VIEWTYPE']._serialized_end=364
  _globals['_DRIVERCONTRACTMETADATA']._serialized_start=140
  _globals['_DRIVERCONTRACTMETADATA']._serialized_end=243
# @@protoc_insertion_point(module_scope)
