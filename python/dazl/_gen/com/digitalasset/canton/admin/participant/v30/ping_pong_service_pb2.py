# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/admin/participant/v30/ping_pong_service.proto
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
    'com/digitalasset/canton/admin/participant/v30/ping_pong_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nEcom/digitalasset/canton/admin/participant/v30/ping_pong_service.proto\x12-com.digitalasset.canton.admin.participant.v30\x1a\x1egoogle/protobuf/duration.proto\"\xef\x01\n\x0bPingRequest\x12%\n\x0etarget_parties\x18\x01 \x03(\tR\rtargetParties\x12\x1e\n\nvalidators\x18\x02 \x03(\tR\nvalidators\x12\x33\n\x07timeout\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x07timeout\x12\x16\n\x06levels\x18\x04 \x01(\rR\x06levels\x12\x1b\n\tdomain_id\x18\x05 \x01(\tR\x08\x64omainId\x12\x1f\n\x0bworkflow_id\x18\x06 \x01(\tR\nworkflowId\x12\x0e\n\x02id\x18\x07 \x01(\tR\x02id\"H\n\x0bPingSuccess\x12\x1b\n\tping_time\x18\x01 \x01(\x04R\x08pingTime\x12\x1c\n\tresponder\x18\x02 \x01(\tR\tresponder\"%\n\x0bPingFailure\x12\x16\n\x06reason\x18\x01 \x01(\tR\x06reason\"\xca\x01\n\x0cPingResponse\x12V\n\x07success\x18\x01 \x01(\x0b\x32:.com.digitalasset.canton.admin.participant.v30.PingSuccessH\x00R\x07success\x12V\n\x07\x66\x61ilure\x18\x02 \x01(\x0b\x32:.com.digitalasset.canton.admin.participant.v30.PingFailureH\x00R\x07\x66\x61ilureB\n\n\x08response2\x8e\x01\n\x0bPingService\x12\x7f\n\x04ping\x12:.com.digitalasset.canton.admin.participant.v30.PingRequest\x1a;.com.digitalasset.canton.admin.participant.v30.PingResponseB^Z\\github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/admin/participant/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.admin.participant.v30.ping_pong_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\\github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/admin/participant/v30'
  _globals['_PINGREQUEST']._serialized_start=153
  _globals['_PINGREQUEST']._serialized_end=392
  _globals['_PINGSUCCESS']._serialized_start=394
  _globals['_PINGSUCCESS']._serialized_end=466
  _globals['_PINGFAILURE']._serialized_start=468
  _globals['_PINGFAILURE']._serialized_end=505
  _globals['_PINGRESPONSE']._serialized_start=508
  _globals['_PINGRESPONSE']._serialized_end=710
  _globals['_PINGSERVICE']._serialized_start=713
  _globals['_PINGSERVICE']._serialized_end=855
# @@protoc_insertion_point(module_scope)