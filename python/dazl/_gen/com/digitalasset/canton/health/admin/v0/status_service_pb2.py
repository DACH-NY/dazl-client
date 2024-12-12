# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/health/admin/v0/status_service.proto
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
    'com/digitalasset/canton/health/admin/v0/status_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n<com/digitalasset/canton/health/admin/v0/status_service.proto\x12\'com.digitalasset.canton.health.admin.v0\x1a\x1egoogle/protobuf/duration.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\"i\n\x13TopologyQueueStatus\x12\x18\n\x07manager\x18\x01 \x01(\rR\x07manager\x12\x1e\n\ndispatcher\x18\x02 \x01(\rR\ndispatcher\x12\x18\n\x07\x63lients\x18\x03 \x01(\rR\x07\x63lients\"\x8a\n\n\nNodeStatus\x12m\n\x0fnot_initialized\x18\x01 \x01(\x0b\x32\x42.com.digitalasset.canton.health.admin.v0.NodeStatus.NotInitializedH\x00R\x0enotInitialized\x12V\n\x07success\x18\x02 \x01(\x0b\x32:.com.digitalasset.canton.health.admin.v0.NodeStatus.StatusH\x00R\x07success\x1a\xdc\x03\n\x06Status\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x31\n\x06uptime\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x06uptime\x12[\n\x05ports\x18\x03 \x03(\x0b\x32\x45.com.digitalasset.canton.health.admin.v0.NodeStatus.Status.PortsEntryR\x05ports\x12\x14\n\x05\x65xtra\x18\x04 \x01(\x0cR\x05\x65xtra\x12\x16\n\x06\x61\x63tive\x18\x05 \x01(\x08R\x06\x61\x63tive\x12\x65\n\x0ftopology_queues\x18\x06 \x01(\x0b\x32<.com.digitalasset.canton.health.admin.v0.TopologyQueueStatusR\x0etopologyQueues\x12\x63\n\ncomponents\x18\x07 \x03(\x0b\x32\x43.com.digitalasset.canton.health.admin.v0.NodeStatus.ComponentStatusR\ncomponents\x1a\x38\n\nPortsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x05R\x05value:\x02\x38\x01\x1a\x9f\x04\n\x0f\x43omponentStatus\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12`\n\x02ok\x18\x02 \x01(\x0b\x32N.com.digitalasset.canton.health.admin.v0.NodeStatus.ComponentStatus.StatusDataH\x00R\x02ok\x12l\n\x08\x64\x65graded\x18\x03 \x01(\x0b\x32N.com.digitalasset.canton.health.admin.v0.NodeStatus.ComponentStatus.StatusDataH\x00R\x08\x64\x65graded\x12h\n\x06\x66\x61iled\x18\x04 \x01(\x0b\x32N.com.digitalasset.canton.health.admin.v0.NodeStatus.ComponentStatus.StatusDataH\x00R\x06\x66\x61iled\x12\x66\n\x05\x66\x61tal\x18\x05 \x01(\x0b\x32N.com.digitalasset.canton.health.admin.v0.NodeStatus.ComponentStatus.StatusDataH\x00R\x05\x66\x61tal\x1aL\n\nStatusData\x12>\n\x0b\x64\x65scription\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.StringValueR\x0b\x64\x65scriptionB\x08\n\x06status\x1a(\n\x0eNotInitialized\x12\x16\n\x06\x61\x63tive\x18\x01 \x01(\x08R\x06\x61\x63tiveB\n\n\x08response\"O\n\x11HealthDumpRequest\x12:\n\tchunkSize\x18\x01 \x01(\x0b\x32\x1c.google.protobuf.UInt32ValueR\tchunkSize\"\'\n\x0fHealthDumpChunk\x12\x14\n\x05\x63hunk\x18\x01 \x01(\x0cR\x05\x63hunk\"\xa7\x01\n\x10\x44omainStatusInfo\x12\x35\n\x16\x63onnected_participants\x18\x01 \x03(\tR\x15\x63onnectedParticipants\x12\\\n\tsequencer\x18\x02 \x01(\x0b\x32>.com.digitalasset.canton.health.admin.v0.SequencerHealthStatusR\tsequencer\"\xf1\x01\n\x15ParticipantStatusInfo\x12{\n\x11\x63onnected_domains\x18\x01 \x03(\x0b\x32N.com.digitalasset.canton.health.admin.v0.ParticipantStatusInfo.ConnectedDomainR\x10\x63onnectedDomains\x12\x16\n\x06\x61\x63tive\x18\x02 \x01(\x08R\x06\x61\x63tive\x1a\x43\n\x0f\x43onnectedDomain\x12\x16\n\x06\x64omain\x18\x01 \x01(\tR\x06\x64omain\x12\x18\n\x07healthy\x18\x02 \x01(\x08R\x07healthy\"\x9c\x02\n\x13SequencerNodeStatus\x12\x35\n\x16\x63onnected_participants\x18\x01 \x03(\tR\x15\x63onnectedParticipants\x12\\\n\tsequencer\x18\x02 \x01(\x0b\x32>.com.digitalasset.canton.health.admin.v0.SequencerHealthStatusR\tsequencer\x12\x1b\n\tdomain_id\x18\x03 \x01(\tR\x08\x64omainId\x12S\n\x05\x61\x64min\x18\x04 \x01(\x0b\x32=.com.digitalasset.canton.health.admin.v0.SequencerAdminStatusR\x05\x61\x64min\"g\n\x15SequencerHealthStatus\x12\x16\n\x06\x61\x63tive\x18\x01 \x01(\x08R\x06\x61\x63tive\x12\x36\n\x07\x64\x65tails\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValueR\x07\x64\x65tails\"J\n\x14SequencerAdminStatus\x12\x32\n\x15\x61\x63\x63\x65pts_admin_changes\x18\x01 \x01(\x08R\x13\x61\x63\x63\x65ptsAdminChanges\"1\n\x12MediatorNodeStatus\x12\x1b\n\tdomain_id\x18\x01 \x01(\tR\x08\x64omainId\"*\n\x12SetLogLevelRequest\x12\x14\n\x05level\x18\x01 \x01(\tR\x05level\"\x15\n\x13SetLogLevelResponse\"\x16\n\x14GetLastErrorsRequest\"\xb3\x01\n\x15GetLastErrorsResponse\x12\\\n\x06\x65rrors\x18\x01 \x03(\x0b\x32\x44.com.digitalasset.canton.health.admin.v0.GetLastErrorsResponse.ErrorR\x06\x65rrors\x1a<\n\x05\x45rror\x12\x19\n\x08trace_id\x18\x01 \x01(\tR\x07traceId\x12\x18\n\x07message\x18\x02 \x01(\tR\x07message\"5\n\x18GetLastErrorTraceRequest\x12\x19\n\x08trace_id\x18\x01 \x01(\tR\x07traceId\"7\n\x19GetLastErrorTraceResponse\x12\x1a\n\x08messages\x18\x01 \x03(\tR\x08messages2\xa2\x05\n\rStatusService\x12U\n\x06Status\x12\x16.google.protobuf.Empty\x1a\x33.com.digitalasset.canton.health.admin.v0.NodeStatus\x12\x84\x01\n\nHealthDump\x12:.com.digitalasset.canton.health.admin.v0.HealthDumpRequest\x1a\x38.com.digitalasset.canton.health.admin.v0.HealthDumpChunk0\x01\x12\x88\x01\n\x0bSetLogLevel\x12;.com.digitalasset.canton.health.admin.v0.SetLogLevelRequest\x1a<.com.digitalasset.canton.health.admin.v0.SetLogLevelResponse\x12\x8e\x01\n\rGetLastErrors\x12=.com.digitalasset.canton.health.admin.v0.GetLastErrorsRequest\x1a>.com.digitalasset.canton.health.admin.v0.GetLastErrorsResponse\x12\x96\x01\n\rGetErrorTrace\x12\x41.com.digitalasset.canton.health.admin.v0.GetLastErrorTraceRequest\x1a\x42.com.digitalasset.canton.health.admin.v0.GetLastErrorTraceResponseBXZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/health/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.health.admin.v0.status_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZVgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/health/admin/v0'
  _globals['_NODESTATUS_STATUS_PORTSENTRY']._loaded_options = None
  _globals['_NODESTATUS_STATUS_PORTSENTRY']._serialized_options = b'8\001'
  _globals['_TOPOLOGYQUEUESTATUS']._serialized_start=198
  _globals['_TOPOLOGYQUEUESTATUS']._serialized_end=303
  _globals['_NODESTATUS']._serialized_start=306
  _globals['_NODESTATUS']._serialized_end=1596
  _globals['_NODESTATUS_STATUS']._serialized_start=520
  _globals['_NODESTATUS_STATUS']._serialized_end=996
  _globals['_NODESTATUS_STATUS_PORTSENTRY']._serialized_start=940
  _globals['_NODESTATUS_STATUS_PORTSENTRY']._serialized_end=996
  _globals['_NODESTATUS_COMPONENTSTATUS']._serialized_start=999
  _globals['_NODESTATUS_COMPONENTSTATUS']._serialized_end=1542
  _globals['_NODESTATUS_COMPONENTSTATUS_STATUSDATA']._serialized_start=1456
  _globals['_NODESTATUS_COMPONENTSTATUS_STATUSDATA']._serialized_end=1532
  _globals['_NODESTATUS_NOTINITIALIZED']._serialized_start=1544
  _globals['_NODESTATUS_NOTINITIALIZED']._serialized_end=1584
  _globals['_HEALTHDUMPREQUEST']._serialized_start=1598
  _globals['_HEALTHDUMPREQUEST']._serialized_end=1677
  _globals['_HEALTHDUMPCHUNK']._serialized_start=1679
  _globals['_HEALTHDUMPCHUNK']._serialized_end=1718
  _globals['_DOMAINSTATUSINFO']._serialized_start=1721
  _globals['_DOMAINSTATUSINFO']._serialized_end=1888
  _globals['_PARTICIPANTSTATUSINFO']._serialized_start=1891
  _globals['_PARTICIPANTSTATUSINFO']._serialized_end=2132
  _globals['_PARTICIPANTSTATUSINFO_CONNECTEDDOMAIN']._serialized_start=2065
  _globals['_PARTICIPANTSTATUSINFO_CONNECTEDDOMAIN']._serialized_end=2132
  _globals['_SEQUENCERNODESTATUS']._serialized_start=2135
  _globals['_SEQUENCERNODESTATUS']._serialized_end=2419
  _globals['_SEQUENCERHEALTHSTATUS']._serialized_start=2421
  _globals['_SEQUENCERHEALTHSTATUS']._serialized_end=2524
  _globals['_SEQUENCERADMINSTATUS']._serialized_start=2526
  _globals['_SEQUENCERADMINSTATUS']._serialized_end=2600
  _globals['_MEDIATORNODESTATUS']._serialized_start=2602
  _globals['_MEDIATORNODESTATUS']._serialized_end=2651
  _globals['_SETLOGLEVELREQUEST']._serialized_start=2653
  _globals['_SETLOGLEVELREQUEST']._serialized_end=2695
  _globals['_SETLOGLEVELRESPONSE']._serialized_start=2697
  _globals['_SETLOGLEVELRESPONSE']._serialized_end=2718
  _globals['_GETLASTERRORSREQUEST']._serialized_start=2720
  _globals['_GETLASTERRORSREQUEST']._serialized_end=2742
  _globals['_GETLASTERRORSRESPONSE']._serialized_start=2745
  _globals['_GETLASTERRORSRESPONSE']._serialized_end=2924
  _globals['_GETLASTERRORSRESPONSE_ERROR']._serialized_start=2864
  _globals['_GETLASTERRORSRESPONSE_ERROR']._serialized_end=2924
  _globals['_GETLASTERRORTRACEREQUEST']._serialized_start=2926
  _globals['_GETLASTERRORTRACEREQUEST']._serialized_end=2979
  _globals['_GETLASTERRORTRACERESPONSE']._serialized_start=2981
  _globals['_GETLASTERRORTRACERESPONSE']._serialized_end=3036
  _globals['_STATUSSERVICE']._serialized_start=3039
  _globals['_STATUSSERVICE']._serialized_end=3713
# @@protoc_insertion_point(module_scope)
