# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/admin/config_management_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n<com/daml/ledger/api/v1/admin/config_management_service.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x15\n\x13GetTimeModelRequest\"\x99\x01\n\x14GetTimeModelResponse\x12\x39\n\x18\x63onfiguration_generation\x18\x01 \x01(\x03R\x17\x63onfigurationGeneration\x12\x46\n\ntime_model\x18\x02 \x01(\x0b\x32\'.com.daml.ledger.api.v1.admin.TimeModelR\ttimeModel\"\x90\x02\n\x13SetTimeModelRequest\x12#\n\rsubmission_id\x18\x01 \x01(\tR\x0csubmissionId\x12J\n\x13maximum_record_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x11maximumRecordTime\x12\x39\n\x18\x63onfiguration_generation\x18\x03 \x01(\x03R\x17\x63onfigurationGeneration\x12M\n\x0enew_time_model\x18\x04 \x01(\x0b\x32\'.com.daml.ledger.api.v1.admin.TimeModelR\x0cnewTimeModel\"Q\n\x14SetTimeModelResponse\x12\x39\n\x18\x63onfiguration_generation\x18\x01 \x01(\x03R\x17\x63onfigurationGeneration\"\xdc\x01\n\tTimeModel\x12Q\n\x17\x61vg_transaction_latency\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x15\x61vgTransactionLatency\x12\x34\n\x08min_skew\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x07minSkew\x12\x34\n\x08max_skew\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x07maxSkewJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\x32\x87\x02\n\x17\x43onfigManagementService\x12u\n\x0cGetTimeModel\x12\x31.com.daml.ledger.api.v1.admin.GetTimeModelRequest\x1a\x32.com.daml.ledger.api.v1.admin.GetTimeModelResponse\x12u\n\x0cSetTimeModel\x12\x31.com.daml.ledger.api.v1.admin.SetTimeModelRequest\x1a\x32.com.daml.ledger.api.v1.admin.SetTimeModelResponseB\xad\x01\n\x1c\x63om.daml.ledger.api.v1.adminB!ConfigManagementServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.admin.config_management_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\034com.daml.ledger.api.v1.adminB!ConfigManagementServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\252\002\034Com.Daml.Ledger.Api.V1.Admin'
  _globals['_GETTIMEMODELREQUEST']._serialized_start=159
  _globals['_GETTIMEMODELREQUEST']._serialized_end=180
  _globals['_GETTIMEMODELRESPONSE']._serialized_start=183
  _globals['_GETTIMEMODELRESPONSE']._serialized_end=336
  _globals['_SETTIMEMODELREQUEST']._serialized_start=339
  _globals['_SETTIMEMODELREQUEST']._serialized_end=611
  _globals['_SETTIMEMODELRESPONSE']._serialized_start=613
  _globals['_SETTIMEMODELRESPONSE']._serialized_end=694
  _globals['_TIMEMODEL']._serialized_start=697
  _globals['_TIMEMODEL']._serialized_end=917
  _globals['_CONFIGMANAGEMENTSERVICE']._serialized_start=920
  _globals['_CONFIGMANAGEMENTSERVICE']._serialized_end=1183
# @@protoc_insertion_point(module_scope)
