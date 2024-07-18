# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/command_completion_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import completion_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_completion__pb2
from . import ledger_offset_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_ledger__offset__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n7com/daml/ledger/api/v1/command_completion_service.proto\x12\x16\x63om.daml.ledger.api.v1\x1a\'com/daml/ledger/api/v1/completion.proto\x1a*com/daml/ledger/api/v1/ledger_offset.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb5\x01\n\x17\x43ompletionStreamRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\x12%\n\x0e\x61pplication_id\x18\x02 \x01(\tR\rapplicationId\x12\x18\n\x07parties\x18\x03 \x03(\tR\x07parties\x12<\n\x06offset\x18\x04 \x01(\x0b\x32$.com.daml.ledger.api.v1.LedgerOffsetR\x06offset\"\xa4\x01\n\x18\x43ompletionStreamResponse\x12\x42\n\ncheckpoint\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.CheckpointR\ncheckpoint\x12\x44\n\x0b\x63ompletions\x18\x02 \x03(\x0b\x32\".com.daml.ledger.api.v1.CompletionR\x0b\x63ompletions\"\x87\x01\n\nCheckpoint\x12;\n\x0brecord_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nrecordTime\x12<\n\x06offset\x18\x02 \x01(\x0b\x32$.com.daml.ledger.api.v1.LedgerOffsetR\x06offset\"3\n\x14\x43ompletionEndRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\"U\n\x15\x43ompletionEndResponse\x12<\n\x06offset\x18\x01 \x01(\x0b\x32$.com.daml.ledger.api.v1.LedgerOffsetR\x06offset2\x81\x02\n\x18\x43ommandCompletionService\x12w\n\x10\x43ompletionStream\x12/.com.daml.ledger.api.v1.CompletionStreamRequest\x1a\x30.com.daml.ledger.api.v1.CompletionStreamResponse0\x01\x12l\n\rCompletionEnd\x12,.com.daml.ledger.api.v1.CompletionEndRequest\x1a-.com.daml.ledger.api.v1.CompletionEndResponseB\x9c\x01\n\x16\x63om.daml.ledger.api.v1B\"CommandCompletionServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.command_completion_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.daml.ledger.api.v1B\"CommandCompletionServiceOuterClassZEgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1\252\002\026Com.Daml.Ledger.Api.V1'
  _globals['_COMPLETIONSTREAMREQUEST']._serialized_start=202
  _globals['_COMPLETIONSTREAMREQUEST']._serialized_end=383
  _globals['_COMPLETIONSTREAMRESPONSE']._serialized_start=386
  _globals['_COMPLETIONSTREAMRESPONSE']._serialized_end=550
  _globals['_CHECKPOINT']._serialized_start=553
  _globals['_CHECKPOINT']._serialized_end=688
  _globals['_COMPLETIONENDREQUEST']._serialized_start=690
  _globals['_COMPLETIONENDREQUEST']._serialized_end=741
  _globals['_COMPLETIONENDRESPONSE']._serialized_start=743
  _globals['_COMPLETIONENDRESPONSE']._serialized_end=828
  _globals['_COMMANDCOMPLETIONSERVICE']._serialized_start=831
  _globals['_COMMANDCOMPLETIONSERVICE']._serialized_end=1088
# @@protoc_insertion_point(module_scope)
