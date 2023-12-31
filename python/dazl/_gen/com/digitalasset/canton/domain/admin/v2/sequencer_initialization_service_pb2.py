# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/domain/admin/v2/sequencer_initialization_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....protocol.v1 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_sequencing__pb2
from ....protocol.v0 import topology_ext_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__ext__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nNcom/digitalasset/canton/domain/admin/v2/sequencer_initialization_service.proto\x12\'com.digitalasset.canton.domain.admin.v2\x1a\x34\x63om/digitalasset/canton/protocol/v1/sequencing.proto\x1a\x36\x63om/digitalasset/canton/protocol/v0/topology_ext.proto\"\x8a\x02\n\x1aInitializeSequencerRequest\x12\x66\n\x11topology_snapshot\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.TopologyTransactionsR\x10topologySnapshot\x12h\n\x11\x64omain_parameters\x18\x02 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v1.StaticDomainParametersR\x10\x64omainParameters\x12\x1a\n\x08snapshot\x18\x03 \x01(\x0cR\x08snapshot\"=\n\x1bInitializeSequencerResponse\x12\x1e\n\nreplicated\x18\x01 \x01(\x08R\nreplicated\"\x98\x02\n\x0bInitRequest\x12\x1b\n\tdomain_id\x18\x01 \x01(\tR\x08\x64omainId\x12\x66\n\x11topology_snapshot\x18\x02 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.TopologyTransactionsR\x10topologySnapshot\x12h\n\x11\x64omain_parameters\x18\x04 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v1.StaticDomainParametersR\x10\x64omainParameters\x12\x1a\n\x08snapshot\x18\x03 \x01(\x0cR\x08snapshot2\xba\x01\n\x1eSequencerInitializationService\x12\x97\x01\n\nInitialize\x12\x43.com.digitalasset.canton.domain.admin.v2.InitializeSequencerRequest\x1a\x44.com.digitalasset.canton.domain.admin.v2.InitializeSequencerResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.admin.v2.sequencer_initialization_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_INITIALIZESEQUENCERREQUEST']._serialized_start=234
  _globals['_INITIALIZESEQUENCERREQUEST']._serialized_end=500
  _globals['_INITIALIZESEQUENCERRESPONSE']._serialized_start=502
  _globals['_INITIALIZESEQUENCERRESPONSE']._serialized_end=563
  _globals['_INITREQUEST']._serialized_start=566
  _globals['_INITREQUEST']._serialized_end=846
  _globals['_SEQUENCERINITIALIZATIONSERVICE']._serialized_start=849
  _globals['_SEQUENCERINITIALIZATIONSERVICE']._serialized_end=1035
# @@protoc_insertion_point(module_scope)
