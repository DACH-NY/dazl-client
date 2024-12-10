# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/domain/api/v0/sequencer_service.proto
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
    'com/digitalasset/canton/domain/api/v0/sequencer_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2
from ....v0 import trace_context_pb2 as com_dot_digitalasset_dot_canton_dot_v0_dot_trace__context__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n=com/digitalasset/canton/domain/api/v0/sequencer_service.proto\x12%com.digitalasset.canton.domain.api.v0\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a.com/digitalasset/canton/v0/trace_context.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"W\n\x19SendAsyncVersionedRequest\x12:\n\x19signed_submission_request\x18\x01 \x01(\x0cR\x17signedSubmissionRequest\"Y\n(SendAsyncUnauthenticatedVersionedRequest\x12-\n\x12submission_request\x18\x01 \x01(\x0cR\x11submissionRequest\"\x9a\x03\n\x11SendAsyncResponse\x12T\n\x05\x65rror\x18\x01 \x01(\x0b\x32>.com.digitalasset.canton.domain.api.v0.SendAsyncResponse.ErrorR\x05\x65rror\x1a\xae\x02\n\x05\x45rror\x12)\n\x0frequest_invalid\x18\x01 \x01(\tH\x00R\x0erequestInvalid\x12)\n\x0frequest_refused\x18\x02 \x01(\tH\x00R\x0erequestRefused\x12 \n\noverloaded\x18\x03 \x01(\tH\x00R\noverloaded\x12\'\n\x0esender_unknown\x18\x04 \x01(\tH\x00R\rsenderUnknown\x12%\n\rshutting_down\x18\x05 \x01(\tH\x00R\x0cshuttingDown\x12\"\n\x0bunavailable\x18\x06 \x01(\tH\x00R\x0bunavailable\x12/\n\x12unknown_recipients\x18\x07 \x01(\tH\x00R\x11unknownRecipientsB\x08\n\x06reason\"\xe0\x03\n\x17SendAsyncSignedResponse\x12Z\n\x05\x65rror\x18\x01 \x01(\x0b\x32\x44.com.digitalasset.canton.domain.api.v0.SendAsyncSignedResponse.ErrorR\x05\x65rror\x1a\xe8\x02\n\x05\x45rror\x12)\n\x0frequest_invalid\x18\x01 \x01(\tH\x00R\x0erequestInvalid\x12)\n\x0frequest_refused\x18\x02 \x01(\tH\x00R\x0erequestRefused\x12 \n\noverloaded\x18\x03 \x01(\tH\x00R\noverloaded\x12\'\n\x0esender_unknown\x18\x04 \x01(\tH\x00R\rsenderUnknown\x12%\n\rshutting_down\x18\x05 \x01(\tH\x00R\x0cshuttingDown\x12\"\n\x0bunavailable\x18\x06 \x01(\tH\x00R\x0bunavailable\x12/\n\x12unknown_recipients\x18\x07 \x01(\tH\x00R\x11unknownRecipients\x12\x1c\n\x08internal\x18\x08 \x01(\tH\x00R\x08internal\x12\x1a\n\x07generic\x18\t \x01(\tH\x00R\x07genericB\x08\n\x06reason\"G\n\x13SubscriptionRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x18\n\x07\x63ounter\x18\x02 \x01(\x03R\x07\x63ounter\"\xcf\x01\n\x14SubscriptionResponse\x12h\n\x16signed_sequenced_event\x18\x01 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v0.SignedContentR\x14signedSequencedEvent\x12M\n\rtrace_context\x18\x02 \x01(\x0b\x32(.com.digitalasset.canton.v0.TraceContextR\x0ctraceContext\"\x8c\x02\n\x1dVersionedSubscriptionResponse\x12\x34\n\x16signed_sequenced_event\x18\x01 \x01(\x0cR\x14signedSequencedEvent\x12M\n\rtrace_context\x18\x02 \x01(\x0b\x32(.com.digitalasset.canton.v0.TraceContextR\x0ctraceContext\x12\x66\n\rtraffic_state\x18\x03 \x01(\x0b\x32\x41.com.digitalasset.canton.domain.api.v0.SequencedEventTrafficStateR\x0ctrafficState\"f\n\x12\x41\x63knowledgeRequest\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x38\n\ttimestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\"\x8a\x01\n\x1aSequencedEventTrafficState\x12\x36\n\x17\x65xtra_traffic_remainder\x18\x02 \x01(\x04R\x15\x65xtraTrafficRemainder\x12\x34\n\x16\x65xtra_traffic_consumed\x18\x03 \x01(\x04R\x14\x65xtraTrafficConsumed2\x9b\x0c\n\x10SequencerService\x12}\n\tSendAsync\x12\x36.com.digitalasset.canton.protocol.v0.SubmissionRequest\x1a\x38.com.digitalasset.canton.domain.api.v0.SendAsyncResponse\x12\x85\x01\n\x0fSendAsyncSigned\x12\x32.com.digitalasset.canton.protocol.v0.SignedContent\x1a>.com.digitalasset.canton.domain.api.v0.SendAsyncSignedResponse\x12\x8c\x01\n\x18SendAsyncUnauthenticated\x12\x36.com.digitalasset.canton.protocol.v0.SubmissionRequest\x1a\x38.com.digitalasset.canton.domain.api.v0.SendAsyncResponse\x12\x96\x01\n\x12SendAsyncVersioned\x12@.com.digitalasset.canton.domain.api.v0.SendAsyncVersionedRequest\x1a>.com.digitalasset.canton.domain.api.v0.SendAsyncSignedResponse\x12\xae\x01\n!SendAsyncUnauthenticatedVersioned\x12O.com.digitalasset.canton.domain.api.v0.SendAsyncUnauthenticatedVersionedRequest\x1a\x38.com.digitalasset.canton.domain.api.v0.SendAsyncResponse\x12\x86\x01\n\tSubscribe\x12:.com.digitalasset.canton.domain.api.v0.SubscriptionRequest\x1a;.com.digitalasset.canton.domain.api.v0.SubscriptionResponse0\x01\x12\x98\x01\n\x12SubscribeVersioned\x12:.com.digitalasset.canton.domain.api.v0.SubscriptionRequest\x1a\x44.com.digitalasset.canton.domain.api.v0.VersionedSubscriptionResponse0\x01\x12\x95\x01\n\x18SubscribeUnauthenticated\x12:.com.digitalasset.canton.domain.api.v0.SubscriptionRequest\x1a;.com.digitalasset.canton.domain.api.v0.SubscriptionResponse0\x01\x12\xa7\x01\n!SubscribeUnauthenticatedVersioned\x12:.com.digitalasset.canton.domain.api.v0.SubscriptionRequest\x1a\x44.com.digitalasset.canton.domain.api.v0.VersionedSubscriptionResponse0\x01\x12`\n\x0b\x41\x63knowledge\x12\x39.com.digitalasset.canton.domain.api.v0.AcknowledgeRequest\x1a\x16.google.protobuf.Empty\x12_\n\x11\x41\x63knowledgeSigned\x12\x32.com.digitalasset.canton.protocol.v0.SignedContent\x1a\x16.google.protobuf.EmptyBVZTgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.api.v0.sequencer_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZTgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/api/v0'
  _globals['_SENDASYNCVERSIONEDREQUEST']._serialized_start=268
  _globals['_SENDASYNCVERSIONEDREQUEST']._serialized_end=355
  _globals['_SENDASYNCUNAUTHENTICATEDVERSIONEDREQUEST']._serialized_start=357
  _globals['_SENDASYNCUNAUTHENTICATEDVERSIONEDREQUEST']._serialized_end=446
  _globals['_SENDASYNCRESPONSE']._serialized_start=449
  _globals['_SENDASYNCRESPONSE']._serialized_end=859
  _globals['_SENDASYNCRESPONSE_ERROR']._serialized_start=557
  _globals['_SENDASYNCRESPONSE_ERROR']._serialized_end=859
  _globals['_SENDASYNCSIGNEDRESPONSE']._serialized_start=862
  _globals['_SENDASYNCSIGNEDRESPONSE']._serialized_end=1342
  _globals['_SENDASYNCSIGNEDRESPONSE_ERROR']._serialized_start=982
  _globals['_SENDASYNCSIGNEDRESPONSE_ERROR']._serialized_end=1342
  _globals['_SUBSCRIPTIONREQUEST']._serialized_start=1344
  _globals['_SUBSCRIPTIONREQUEST']._serialized_end=1415
  _globals['_SUBSCRIPTIONRESPONSE']._serialized_start=1418
  _globals['_SUBSCRIPTIONRESPONSE']._serialized_end=1625
  _globals['_VERSIONEDSUBSCRIPTIONRESPONSE']._serialized_start=1628
  _globals['_VERSIONEDSUBSCRIPTIONRESPONSE']._serialized_end=1896
  _globals['_ACKNOWLEDGEREQUEST']._serialized_start=1898
  _globals['_ACKNOWLEDGEREQUEST']._serialized_end=2000
  _globals['_SEQUENCEDEVENTTRAFFICSTATE']._serialized_start=2003
  _globals['_SEQUENCEDEVENTTRAFFICSTATE']._serialized_end=2141
  _globals['_SEQUENCERSERVICE']._serialized_start=2144
  _globals['_SEQUENCERSERVICE']._serialized_end=3707
# @@protoc_insertion_point(module_scope)
