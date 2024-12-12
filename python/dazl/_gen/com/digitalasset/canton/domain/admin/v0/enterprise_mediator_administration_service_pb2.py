# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/domain/admin/v0/enterprise_mediator_administration_service.proto
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
    'com/digitalasset/canton/domain/admin/v0/enterprise_mediator_administration_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....pruning.admin.v0 import pruning_pb2 as com_dot_digitalasset_dot_canton_dot_pruning_dot_admin_dot_v0_dot_pruning__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nXcom/digitalasset/canton/domain/admin/v0/enterprise_mediator_administration_service.proto\x12\'com.digitalasset.canton.domain.admin.v0\x1a\x36\x63om/digitalasset/canton/pruning/admin/v0/pruning.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"R\n\x16MediatorPruningRequest\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp2\x9b\t\n\'EnterpriseMediatorAdministrationService\x12`\n\x05Prune\x12?.com.digitalasset.canton.domain.admin.v0.MediatorPruningRequest\x1a\x16.google.protobuf.Empty\x12\x8c\x01\n\x0bSetSchedule\x12=.com.digitalasset.canton.pruning.admin.v0.SetSchedule.Request\x1a>.com.digitalasset.canton.pruning.admin.v0.SetSchedule.Response\x12\x80\x01\n\x07SetCron\x12\x39.com.digitalasset.canton.pruning.admin.v0.SetCron.Request\x1a:.com.digitalasset.canton.pruning.admin.v0.SetCron.Response\x12\x95\x01\n\x0eSetMaxDuration\x12@.com.digitalasset.canton.pruning.admin.v0.SetMaxDuration.Request\x1a\x41.com.digitalasset.canton.pruning.admin.v0.SetMaxDuration.Response\x12\x8f\x01\n\x0cSetRetention\x12>.com.digitalasset.canton.pruning.admin.v0.SetRetention.Request\x1a?.com.digitalasset.canton.pruning.admin.v0.SetRetention.Response\x12\x92\x01\n\rClearSchedule\x12?.com.digitalasset.canton.pruning.admin.v0.ClearSchedule.Request\x1a@.com.digitalasset.canton.pruning.admin.v0.ClearSchedule.Response\x12\x8c\x01\n\x0bGetSchedule\x12=.com.digitalasset.canton.pruning.admin.v0.GetSchedule.Request\x1a>.com.digitalasset.canton.pruning.admin.v0.GetSchedule.Response\x12\xad\x01\n\x16LocatePruningTimestamp\x12H.com.digitalasset.canton.pruning.admin.v0.LocatePruningTimestamp.Request\x1aI.com.digitalasset.canton.pruning.admin.v0.LocatePruningTimestamp.ResponseBXZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.admin.v0.enterprise_mediator_administration_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/admin/v0'
  _globals['_MEDIATORPRUNINGREQUEST']._serialized_start=251
  _globals['_MEDIATORPRUNINGREQUEST']._serialized_end=333
  _globals['_ENTERPRISEMEDIATORADMINISTRATIONSERVICE']._serialized_start=336
  _globals['_ENTERPRISEMEDIATORADMINISTRATIONSERVICE']._serialized_end=1515
# @@protoc_insertion_point(module_scope)
