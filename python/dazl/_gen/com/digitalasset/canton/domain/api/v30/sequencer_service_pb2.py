# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/domain/api/v30/sequencer_service.proto
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
    'com/digitalasset/canton/domain/api/v30/sequencer_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....protocol.v30 import traffic_control_parameters_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v30_dot_traffic__control__parameters__pb2
from ....topology.admin.v30 import topology_ext_pb2 as com_dot_digitalasset_dot_canton_dot_topology_dot_admin_dot_v30_dot_topology__ext__pb2
from ....v30 import trace_context_pb2 as com_dot_digitalasset_dot_canton_dot_v30_dot_trace__context__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n>com/digitalasset/canton/domain/api/v30/sequencer_service.proto\x12&com.digitalasset.canton.domain.api.v30\x1a\x45\x63om/digitalasset/canton/protocol/v30/traffic_control_parameters.proto\x1a=com/digitalasset/canton/topology/admin/v30/topology_ext.proto\x1a/com/digitalasset/canton/v30/trace_context.proto\"W\n\x19SendAsyncVersionedRequest\x12:\n\x19signed_submission_request\x18\x01 \x01(\x0cR\x17signedSubmissionRequest\"\xf8\x01\n\x19TrafficControlErrorReason\x12]\n\x05\x65rror\x18\x01 \x01(\x0b\x32G.com.digitalasset.canton.domain.api.v30.TrafficControlErrorReason.ErrorR\x05\x65rror\x1a|\n\x05\x45rror\x12\x33\n\x14insufficient_traffic\x18\x01 \x01(\tH\x00R\x13insufficientTraffic\x12\x34\n\x15outdated_traffic_cost\x18\x02 \x01(\tH\x00R\x13outdatedTrafficCostB\x08\n\x06reason\"\xd5\x04\n\x1aSendAsyncVersionedResponse\x12^\n\x05\x65rror\x18\x01 \x01(\x0b\x32H.com.digitalasset.canton.domain.api.v30.SendAsyncVersionedResponse.ErrorR\x05\x65rror\x1a\xd6\x03\n\x05\x45rror\x12)\n\x0frequest_invalid\x18\x01 \x01(\tH\x00R\x0erequestInvalid\x12)\n\x0frequest_refused\x18\x02 \x01(\tH\x00R\x0erequestRefused\x12 \n\noverloaded\x18\x03 \x01(\tH\x00R\noverloaded\x12\'\n\x0esender_unknown\x18\x04 \x01(\tH\x00R\rsenderUnknown\x12%\n\rshutting_down\x18\x05 \x01(\tH\x00R\x0cshuttingDown\x12\"\n\x0bunavailable\x18\x06 \x01(\tH\x00R\x0bunavailable\x12/\n\x12unknown_recipients\x18\x07 \x01(\tH\x00R\x11unknownRecipients\x12\x1c\n\x08internal\x18\x08 \x01(\tH\x00R\x08internal\x12\x1a\n\x07generic\x18\t \x01(\tH\x00R\x07generic\x12l\n\x0ftraffic_control\x18\n \x01(\x0b\x32\x41.com.digitalasset.canton.domain.api.v30.TrafficControlErrorReasonH\x00R\x0etrafficControlB\x08\n\x06reason\"G\n\x13SubscriptionRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x18\n\x07\x63ounter\x18\x02 \x01(\x03R\x07\x63ounter\"\xa5\x01\n\x1dVersionedSubscriptionResponse\x12\x34\n\x16signed_sequenced_event\x18\x01 \x01(\x0cR\x14signedSequencedEvent\x12N\n\rtrace_context\x18\x02 \x01(\x0b\x32).com.digitalasset.canton.v30.TraceContextR\x0ctraceContext\"J\n\x12\x41\x63knowledgeRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x1c\n\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\"\x15\n\x13\x41\x63knowledgeResponse\"X\n\x18\x41\x63knowledgeSignedRequest\x12<\n\x1asigned_acknowledge_request\x18\x01 \x01(\x0cR\x18signedAcknowledgeRequest\"\x1b\n\x19\x41\x63knowledgeSignedResponse\"=\n#DownloadTopologyStateForInitRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\"\x9d\x01\n$DownloadTopologyStateForInitResponse\x12u\n\x15topology_transactions\x18\x01 \x01(\x0b\x32@.com.digitalasset.canton.topology.admin.v30.TopologyTransactionsR\x14topologyTransactions\"W\n\x1fGetTrafficStateForMemberRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x1c\n\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\"{\n GetTrafficStateForMemberResponse\x12W\n\rtraffic_state\x18\x01 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v30.TrafficStateR\x0ctrafficState\"\x8a\x01\n\x1aSequencedEventTrafficState\x12\x36\n\x17\x65xtra_traffic_remainder\x18\x02 \x01(\x04R\x15\x65xtraTrafficRemainder\x12\x34\n\x16\x65xtra_traffic_consumed\x18\x03 \x01(\x04R\x14\x65xtraTrafficConsumed2\xd6\x06\n\x10SequencerService\x12\x9b\x01\n\x12SendAsyncVersioned\x12\x41.com.digitalasset.canton.domain.api.v30.SendAsyncVersionedRequest\x1a\x42.com.digitalasset.canton.domain.api.v30.SendAsyncVersionedResponse\x12\x9a\x01\n\x12SubscribeVersioned\x12;.com.digitalasset.canton.domain.api.v30.SubscriptionRequest\x1a\x45.com.digitalasset.canton.domain.api.v30.VersionedSubscriptionResponse0\x01\x12\x98\x01\n\x11\x41\x63knowledgeSigned\x12@.com.digitalasset.canton.domain.api.v30.AcknowledgeSignedRequest\x1a\x41.com.digitalasset.canton.domain.api.v30.AcknowledgeSignedResponse\x12\xbb\x01\n\x1c\x44ownloadTopologyStateForInit\x12K.com.digitalasset.canton.domain.api.v30.DownloadTopologyStateForInitRequest\x1aL.com.digitalasset.canton.domain.api.v30.DownloadTopologyStateForInitResponse0\x01\x12\xad\x01\n\x18GetTrafficStateForMember\x12G.com.digitalasset.canton.domain.api.v30.GetTrafficStateForMemberRequest\x1aH.com.digitalasset.canton.domain.api.v30.GetTrafficStateForMemberResponseBWZUgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/api/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.api.v30.sequencer_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZUgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/domain/api/v30'
  _globals['_SENDASYNCVERSIONEDREQUEST']._serialized_start=289
  _globals['_SENDASYNCVERSIONEDREQUEST']._serialized_end=376
  _globals['_TRAFFICCONTROLERRORREASON']._serialized_start=379
  _globals['_TRAFFICCONTROLERRORREASON']._serialized_end=627
  _globals['_TRAFFICCONTROLERRORREASON_ERROR']._serialized_start=503
  _globals['_TRAFFICCONTROLERRORREASON_ERROR']._serialized_end=627
  _globals['_SENDASYNCVERSIONEDRESPONSE']._serialized_start=630
  _globals['_SENDASYNCVERSIONEDRESPONSE']._serialized_end=1227
  _globals['_SENDASYNCVERSIONEDRESPONSE_ERROR']._serialized_start=757
  _globals['_SENDASYNCVERSIONEDRESPONSE_ERROR']._serialized_end=1227
  _globals['_SUBSCRIPTIONREQUEST']._serialized_start=1229
  _globals['_SUBSCRIPTIONREQUEST']._serialized_end=1300
  _globals['_VERSIONEDSUBSCRIPTIONRESPONSE']._serialized_start=1303
  _globals['_VERSIONEDSUBSCRIPTIONRESPONSE']._serialized_end=1468
  _globals['_ACKNOWLEDGEREQUEST']._serialized_start=1470
  _globals['_ACKNOWLEDGEREQUEST']._serialized_end=1544
  _globals['_ACKNOWLEDGERESPONSE']._serialized_start=1546
  _globals['_ACKNOWLEDGERESPONSE']._serialized_end=1567
  _globals['_ACKNOWLEDGESIGNEDREQUEST']._serialized_start=1569
  _globals['_ACKNOWLEDGESIGNEDREQUEST']._serialized_end=1657
  _globals['_ACKNOWLEDGESIGNEDRESPONSE']._serialized_start=1659
  _globals['_ACKNOWLEDGESIGNEDRESPONSE']._serialized_end=1686
  _globals['_DOWNLOADTOPOLOGYSTATEFORINITREQUEST']._serialized_start=1688
  _globals['_DOWNLOADTOPOLOGYSTATEFORINITREQUEST']._serialized_end=1749
  _globals['_DOWNLOADTOPOLOGYSTATEFORINITRESPONSE']._serialized_start=1752
  _globals['_DOWNLOADTOPOLOGYSTATEFORINITRESPONSE']._serialized_end=1909
  _globals['_GETTRAFFICSTATEFORMEMBERREQUEST']._serialized_start=1911
  _globals['_GETTRAFFICSTATEFORMEMBERREQUEST']._serialized_end=1998
  _globals['_GETTRAFFICSTATEFORMEMBERRESPONSE']._serialized_start=2000
  _globals['_GETTRAFFICSTATEFORMEMBERRESPONSE']._serialized_end=2123
  _globals['_SEQUENCEDEVENTTRAFFICSTATE']._serialized_start=2126
  _globals['_SEQUENCEDEVENTTRAFFICSTATE']._serialized_end=2264
  _globals['_SEQUENCERSERVICE']._serialized_start=2267
  _globals['_SEQUENCERSERVICE']._serialized_end=3121
# @@protoc_insertion_point(module_scope)
