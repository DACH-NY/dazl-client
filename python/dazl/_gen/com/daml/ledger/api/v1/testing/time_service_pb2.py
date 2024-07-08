# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/testing/time_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1com/daml/ledger/api/v1/testing/time_service.proto\x12\x1e\x63om.daml.ledger.api.v1.testing\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"-\n\x0eGetTimeRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\"P\n\x0fGetTimeResponse\x12=\n\x0c\x63urrent_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0b\x63urrentTime\"\xa3\x01\n\x0eSetTimeRequest\x12\x1b\n\tledger_id\x18\x01 \x01(\tR\x08ledgerId\x12=\n\x0c\x63urrent_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0b\x63urrentTime\x12\x35\n\x08new_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07newTime2\xce\x01\n\x0bTimeService\x12l\n\x07GetTime\x12..com.daml.ledger.api.v1.testing.GetTimeRequest\x1a/.com.daml.ledger.api.v1.testing.GetTimeResponse0\x01\x12Q\n\x07SetTime\x12..com.daml.ledger.api.v1.testing.SetTimeRequest\x1a\x16.google.protobuf.EmptyB\xa7\x01\n\x1e\x63om.daml.ledger.api.v1.testingB\x15TimeServiceOuterClassZMgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1/testing\xaa\x02\x1e\x43om.Daml.Ledger.Api.V1.Testingb\x06proto3')



_GETTIMEREQUEST = DESCRIPTOR.message_types_by_name['GetTimeRequest']
_GETTIMERESPONSE = DESCRIPTOR.message_types_by_name['GetTimeResponse']
_SETTIMEREQUEST = DESCRIPTOR.message_types_by_name['SetTimeRequest']
GetTimeRequest = _reflection.GeneratedProtocolMessageType('GetTimeRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMEREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.GetTimeRequest)
  })
_sym_db.RegisterMessage(GetTimeRequest)

GetTimeResponse = _reflection.GeneratedProtocolMessageType('GetTimeResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMERESPONSE,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.GetTimeResponse)
  })
_sym_db.RegisterMessage(GetTimeResponse)

SetTimeRequest = _reflection.GeneratedProtocolMessageType('SetTimeRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMEREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.SetTimeRequest)
  })
_sym_db.RegisterMessage(SetTimeRequest)

_TIMESERVICE = DESCRIPTOR.services_by_name['TimeService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\036com.daml.ledger.api.v1.testingB\025TimeServiceOuterClassZMgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1/testing\252\002\036Com.Daml.Ledger.Api.V1.Testing'
  _GETTIMEREQUEST._serialized_start=147
  _GETTIMEREQUEST._serialized_end=192
  _GETTIMERESPONSE._serialized_start=194
  _GETTIMERESPONSE._serialized_end=274
  _SETTIMEREQUEST._serialized_start=277
  _SETTIMEREQUEST._serialized_end=440
  _TIMESERVICE._serialized_start=443
  _TIMESERVICE._serialized_end=649
# @@protoc_insertion_point(module_scope)
