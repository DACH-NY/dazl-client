# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v30/signed_content.proto
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
    'com/digitalasset/canton/protocol/v30/signed_content.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v30 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v30_dot_crypto__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n9com/digitalasset/canton/protocol/v30/signed_content.proto\x12$com.digitalasset.canton.protocol.v30\x1a/com/digitalasset/canton/crypto/v30/crypto.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xeb\x01\n\rSignedContent\x12\x35\n\x07\x63ontent\x18\x01 \x01(\x0b\x32\x1b.google.protobuf.BytesValueR\x07\x63ontent\x12M\n\nsignatures\x18\x02 \x03(\x0b\x32-.com.digitalasset.canton.crypto.v30.SignatureR\nsignatures\x12T\n\x18timestamp_of_signing_key\x18\x03 \x01(\x0b\x32\x1b.google.protobuf.Int64ValueR\x15timestampOfSigningKeyBUZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v30.signed_content_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30'
  _globals['_SIGNEDCONTENT']._serialized_start=181
  _globals['_SIGNEDCONTENT']._serialized_end=416
# @@protoc_insertion_point(module_scope)