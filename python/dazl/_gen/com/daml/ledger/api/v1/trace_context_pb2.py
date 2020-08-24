# Copyright (c) 2017-2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/trace_context.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='com/daml/ledger/api/v1/trace_context.proto',
  package='com.daml.ledger.api.v1',
  syntax='proto3',
  serialized_options=b'\n\026com.daml.ledger.api.v1B\026TraceContextOuterClass\252\002\026Com.Daml.Ledger.Api.V1',
  serialized_pb=b'\n*com/daml/ledger/api/v1/trace_context.proto\x12\x16\x63om.daml.ledger.api.v1\x1a\x1egoogle/protobuf/wrappers.proto\"\x9b\x01\n\x0cTraceContext\x12\x19\n\rtrace_id_high\x18\x01 \x01(\x04\x42\x02\x30\x01\x12\x14\n\x08trace_id\x18\x02 \x01(\x04\x42\x02\x30\x01\x12\x13\n\x07span_id\x18\x03 \x01(\x04\x42\x02\x30\x01\x12\x34\n\x0eparent_span_id\x18\x04 \x01(\x0b\x32\x1c.google.protobuf.UInt64Value\x12\x0f\n\x07sampled\x18\x05 \x01(\x08\x42I\n\x16\x63om.daml.ledger.api.v1B\x16TraceContextOuterClass\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,])




_TRACECONTEXT = _descriptor.Descriptor(
  name='TraceContext',
  full_name='com.daml.ledger.api.v1.TraceContext',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='trace_id_high', full_name='com.daml.ledger.api.v1.TraceContext.trace_id_high', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'0\001', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='trace_id', full_name='com.daml.ledger.api.v1.TraceContext.trace_id', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'0\001', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='span_id', full_name='com.daml.ledger.api.v1.TraceContext.span_id', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'0\001', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_span_id', full_name='com.daml.ledger.api.v1.TraceContext.parent_span_id', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sampled', full_name='com.daml.ledger.api.v1.TraceContext.sampled', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=103,
  serialized_end=258,
)

_TRACECONTEXT.fields_by_name['parent_span_id'].message_type = google_dot_protobuf_dot_wrappers__pb2._UINT64VALUE
DESCRIPTOR.message_types_by_name['TraceContext'] = _TRACECONTEXT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TraceContext = _reflection.GeneratedProtocolMessageType('TraceContext', (_message.Message,), {
  'DESCRIPTOR' : _TRACECONTEXT,
  '__module__' : 'com.daml.ledger.api.v1.trace_context_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.TraceContext)
  })
_sym_db.RegisterMessage(TraceContext)


DESCRIPTOR._options = None
_TRACECONTEXT.fields_by_name['trace_id_high']._options = None
_TRACECONTEXT.fields_by_name['trace_id']._options = None
_TRACECONTEXT.fields_by_name['span_id']._options = None
# @@protoc_insertion_point(module_scope)
