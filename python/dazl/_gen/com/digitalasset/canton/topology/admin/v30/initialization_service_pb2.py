# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/topology/admin/v30/initialization_service.proto
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
    'com/digitalasset/canton/topology/admin/v30/initialization_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import topology_ext_pb2 as com_dot_digitalasset_dot_canton_dot_topology_dot_admin_dot_v30_dot_topology__ext__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nGcom/digitalasset/canton/topology/admin/v30/initialization_service.proto\x12*com.digitalasset.canton.topology.admin.v30\x1a=com/digitalasset/canton/topology/admin/v30/topology_ext.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"<\n\rInitIdRequest\x12+\n\x11unique_identifier\x18\x01 \x01(\tR\x10uniqueIdentifier\"\x10\n\x0eInitIdResponse\"^\n\rGetIdResponse\x12 \n\x0binitialized\x18\x01 \x01(\x08R\x0binitialized\x12+\n\x11unique_identifier\x18\x02 \x01(\tR\x10uniqueIdentifier\"\"\n GetOnboardingTransactionsRequest\"\x89\x01\n!GetOnboardingTransactionsResponse\x12\x64\n\x0ctransactions\x18\x01 \x01(\x0b\x32@.com.digitalasset.canton.topology.admin.v30.TopologyTransactionsR\x0ctransactions\"\x0e\n\x0cGetIdRequest\"\x14\n\x12\x43urrentTimeRequest\"T\n\x13\x43urrentTimeResponse\x12=\n\x0c\x63urrent_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0b\x63urrentTime2\xea\x04\n\x1dIdentityInitializationService\x12\x7f\n\x06InitId\x12\x39.com.digitalasset.canton.topology.admin.v30.InitIdRequest\x1a:.com.digitalasset.canton.topology.admin.v30.InitIdResponse\x12\xb8\x01\n\x19GetOnboardingTransactions\x12L.com.digitalasset.canton.topology.admin.v30.GetOnboardingTransactionsRequest\x1aM.com.digitalasset.canton.topology.admin.v30.GetOnboardingTransactionsResponse\x12|\n\x05GetId\x12\x38.com.digitalasset.canton.topology.admin.v30.GetIdRequest\x1a\x39.com.digitalasset.canton.topology.admin.v30.GetIdResponse\x12\x8e\x01\n\x0b\x43urrentTime\x12>.com.digitalasset.canton.topology.admin.v30.CurrentTimeRequest\x1a?.com.digitalasset.canton.topology.admin.v30.CurrentTimeResponseB[ZYgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/topology/admin/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.topology.admin.v30.initialization_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZYgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/topology/admin/v30'
  _globals['_INITIDREQUEST']._serialized_start=215
  _globals['_INITIDREQUEST']._serialized_end=275
  _globals['_INITIDRESPONSE']._serialized_start=277
  _globals['_INITIDRESPONSE']._serialized_end=293
  _globals['_GETIDRESPONSE']._serialized_start=295
  _globals['_GETIDRESPONSE']._serialized_end=389
  _globals['_GETONBOARDINGTRANSACTIONSREQUEST']._serialized_start=391
  _globals['_GETONBOARDINGTRANSACTIONSREQUEST']._serialized_end=425
  _globals['_GETONBOARDINGTRANSACTIONSRESPONSE']._serialized_start=428
  _globals['_GETONBOARDINGTRANSACTIONSRESPONSE']._serialized_end=565
  _globals['_GETIDREQUEST']._serialized_start=567
  _globals['_GETIDREQUEST']._serialized_end=581
  _globals['_CURRENTTIMEREQUEST']._serialized_start=583
  _globals['_CURRENTTIMEREQUEST']._serialized_end=603
  _globals['_CURRENTTIMERESPONSE']._serialized_start=605
  _globals['_CURRENTTIMERESPONSE']._serialized_end=689
  _globals['_IDENTITYINITIALIZATIONSERVICE']._serialized_start=692
  _globals['_IDENTITYINITIALIZATIONSERVICE']._serialized_end=1310
# @@protoc_insertion_point(module_scope)
