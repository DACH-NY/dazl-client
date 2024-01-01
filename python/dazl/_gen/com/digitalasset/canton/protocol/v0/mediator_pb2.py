# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/mediator.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import common_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_common__pb2
from . import mediator_response_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_mediator__response__pb2
from . import merkle_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_merkle__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2com/digitalasset/canton/protocol/v0/mediator.proto\x12#com.digitalasset.canton.protocol.v0\x1a\x30\x63om/digitalasset/canton/protocol/v0/common.proto\x1a;com/digitalasset/canton/protocol/v0/mediator_response.proto\x1a\x30\x63om/digitalasset/canton/protocol/v0/merkle.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xfd\x01\n\x18TransactionResultMessage\x12\x39\n\nrequest_id\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\trequestId\x12\x46\n\x07verdict\x18\x02 \x01(\x0b\x32,.com.digitalasset.canton.protocol.v0.VerdictR\x07verdict\x12^\n\x11notification_tree\x18\x05 \x01(\x0b\x32\x31.com.digitalasset.canton.protocol.v0.InformeeTreeR\x10notificationTree\"\xc8\x02\n\x07Verdict\x12\x32\n\x07\x61pprove\x18\x01 \x01(\x0b\x32\x16.google.protobuf.EmptyH\x00R\x07\x61pprove\x12\x62\n\x10validator_reject\x18\x02 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.RejectionReasonsH\x00R\x0fvalidatorReject\x12\x61\n\x0fmediator_reject\x18\x03 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.MediatorRejectionH\x00R\x0emediatorReject\x12\x32\n\x07timeout\x18\x04 \x01(\x0b\x32\x16.google.protobuf.EmptyH\x00R\x07timeoutB\x0e\n\x0csome_verdict\"b\n\x10RejectionReasons\x12N\n\x07reasons\x18\x01 \x03(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.RejectionReasonR\x07reasons\"u\n\x0fRejectionReason\x12\x18\n\x07parties\x18\x01 \x03(\tR\x07parties\x12H\n\x06reject\x18\x02 \x01(\x0b\x32\x30.com.digitalasset.canton.protocol.v0.LocalRejectR\x06reject\"\xe7\x02\n\x11MediatorRejection\x12O\n\x04\x63ode\x18\x01 \x01(\x0e\x32;.com.digitalasset.canton.protocol.v0.MediatorRejection.CodeR\x04\x63ode\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\"\xe8\x01\n\x04\x43ode\x12\x0f\n\x0bMissingCode\x10\x00\x12)\n%InformeesNotHostedOnActiveParticipant\x10\x01\x12\x1e\n\x1aNotEnoughConfirmingParties\x10\x02\x12&\n\"ViewThresholdBelowMinimumThreshold\x10\x03\x12\x1a\n\x16InvalidRootHashMessage\x10\x04\x12\x0b\n\x07Timeout\x10\x05\x12\x19\n\x15WrongDeclaredMediator\x10\x06\x12\x18\n\x14NonUniqueRequestUuid\x10\x07\"[\n\x0cInformeeTree\x12K\n\x04tree\x18\x01 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.GenTransactionTreeR\x04tree\"\x9a\x02\n\x1eMalformedMediatorRequestResult\x12\x39\n\nrequest_id\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\trequestId\x12\x1b\n\tdomain_id\x18\x02 \x01(\tR\x08\x64omainId\x12J\n\tview_type\x18\x03 \x01(\x0e\x32-.com.digitalasset.canton.protocol.v0.ViewTypeR\x08viewType\x12T\n\trejection\x18\x04 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.MediatorRejectionR\trejection\"\x89\x02\n\x0eTransferResult\x12\x39\n\nrequest_id\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\trequestId\x12%\n\rorigin_domain\x18\x02 \x01(\tH\x00R\x0coriginDomain\x12%\n\rtarget_domain\x18\x03 \x01(\tH\x00R\x0ctargetDomain\x12\x1c\n\tinformees\x18\x04 \x03(\tR\tinformees\x12\x46\n\x07verdict\x18\x05 \x01(\x0b\x32,.com.digitalasset.canton.protocol.v0.VerdictR\x07verdictB\x08\n\x06\x64omainb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.mediator_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_TRANSACTIONRESULTMESSAGE']._serialized_start=315
  _globals['_TRANSACTIONRESULTMESSAGE']._serialized_end=568
  _globals['_VERDICT']._serialized_start=571
  _globals['_VERDICT']._serialized_end=899
  _globals['_REJECTIONREASONS']._serialized_start=901
  _globals['_REJECTIONREASONS']._serialized_end=999
  _globals['_REJECTIONREASON']._serialized_start=1001
  _globals['_REJECTIONREASON']._serialized_end=1118
  _globals['_MEDIATORREJECTION']._serialized_start=1121
  _globals['_MEDIATORREJECTION']._serialized_end=1480
  _globals['_MEDIATORREJECTION_CODE']._serialized_start=1248
  _globals['_MEDIATORREJECTION_CODE']._serialized_end=1480
  _globals['_INFORMEETREE']._serialized_start=1482
  _globals['_INFORMEETREE']._serialized_end=1573
  _globals['_MALFORMEDMEDIATORREQUESTRESULT']._serialized_start=1576
  _globals['_MALFORMEDMEDIATORREQUESTRESULT']._serialized_end=1858
  _globals['_TRANSFERRESULT']._serialized_start=1861
  _globals['_TRANSFERRESULT']._serialized_end=2126
# @@protoc_insertion_point(module_scope)