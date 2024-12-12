# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto
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
    'com/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....admin.pruning.v30 import pruning_pb2 as com_dot_digitalasset_dot_canton_dot_admin_dot_pruning_dot_v30_dot_pruning__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nPcom/digitalasset/canton/mediator/admin/v30/mediator_administration_service.proto\x12*com.digitalasset.canton.mediator.admin.v30\x1a\x37\x63om/digitalasset/canton/admin/pruning/v30/pruning.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"l\n\x0fMediatorPruning\x1aH\n\x0cPruneRequest\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x1a\x0f\n\rPruneResponse2\xdc\t\n\x1dMediatorAdministrationService\x12\x9c\x01\n\x05Prune\x12H.com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneRequest\x1aI.com.digitalasset.canton.mediator.admin.v30.MediatorPruning.PruneResponse\x12\x8e\x01\n\x0bSetSchedule\x12>.com.digitalasset.canton.admin.pruning.v30.SetSchedule.Request\x1a?.com.digitalasset.canton.admin.pruning.v30.SetSchedule.Response\x12\x82\x01\n\x07SetCron\x12:.com.digitalasset.canton.admin.pruning.v30.SetCron.Request\x1a;.com.digitalasset.canton.admin.pruning.v30.SetCron.Response\x12\x97\x01\n\x0eSetMaxDuration\x12\x41.com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Request\x1a\x42.com.digitalasset.canton.admin.pruning.v30.SetMaxDuration.Response\x12\x91\x01\n\x0cSetRetention\x12?.com.digitalasset.canton.admin.pruning.v30.SetRetention.Request\x1a@.com.digitalasset.canton.admin.pruning.v30.SetRetention.Response\x12\x94\x01\n\rClearSchedule\x12@.com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Request\x1a\x41.com.digitalasset.canton.admin.pruning.v30.ClearSchedule.Response\x12\x8e\x01\n\x0bGetSchedule\x12>.com.digitalasset.canton.admin.pruning.v30.GetSchedule.Request\x1a?.com.digitalasset.canton.admin.pruning.v30.GetSchedule.Response\x12\xaf\x01\n\x16LocatePruningTimestamp\x12I.com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.Request\x1aJ.com.digitalasset.canton.admin.pruning.v30.LocatePruningTimestamp.ResponseB[ZYgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/mediator/admin/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.mediator.admin.v30.mediator_administration_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZYgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/mediator/admin/v30'
  _globals['_MEDIATORPRUNING']._serialized_start=218
  _globals['_MEDIATORPRUNING']._serialized_end=326
  _globals['_MEDIATORPRUNING_PRUNEREQUEST']._serialized_start=237
  _globals['_MEDIATORPRUNING_PRUNEREQUEST']._serialized_end=309
  _globals['_MEDIATORPRUNING_PRUNERESPONSE']._serialized_start=311
  _globals['_MEDIATORPRUNING_PRUNERESPONSE']._serialized_end=326
  _globals['_MEDIATORADMINISTRATIONSERVICE']._serialized_start=329
  _globals['_MEDIATORADMINISTRATIONSERVICE']._serialized_end=1573
# @@protoc_insertion_point(module_scope)
