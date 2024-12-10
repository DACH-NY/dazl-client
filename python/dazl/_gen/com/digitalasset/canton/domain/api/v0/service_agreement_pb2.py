# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/domain/api/v0/service_agreement.proto
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
    'com/digitalasset/canton/domain/api/v0/service_agreement.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n=com/digitalasset/canton/domain/api/v0/service_agreement.proto\x12%com.digitalasset.canton.domain.api.v0\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\"\x1c\n\x1aGetServiceAgreementRequest\"r\n\x1bGetServiceAgreementResponse\x12S\n\tagreement\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.ServiceAgreementR\tagreementBVZTgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.api.v0.service_agreement_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZTgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0'
  _globals['_GETSERVICEAGREEMENTREQUEST']._serialized_start=158
  _globals['_GETSERVICEAGREEMENTREQUEST']._serialized_end=186
  _globals['_GETSERVICEAGREEMENTRESPONSE']._serialized_start=188
  _globals['_GETSERVICEAGREEMENTRESPONSE']._serialized_end=302
# @@protoc_insertion_point(module_scope)
