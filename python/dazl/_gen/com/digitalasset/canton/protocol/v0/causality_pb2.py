# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/causality.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import participant_transfer_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_participant__transfer__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n3com/digitalasset/canton/protocol/v0/causality.proto\x12#com.digitalasset.canton.protocol.v0\x1a>com/digitalasset/canton/protocol/v0/participant_transfer.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd6\x01\n\x10\x43\x61usalityMessage\x12(\n\x10target_domain_id\x18\x01 \x01(\tR\x0etargetDomainId\x12P\n\x0btransfer_id\x18\x02 \x01(\x0b\x32/.com.digitalasset.canton.protocol.v0.TransferIdR\ntransferId\x12\x46\n\x05\x63lock\x18\x03 \x01(\x0b\x32\x30.com.digitalasset.canton.protocol.v0.VectorClockR\x05\x63lock\"\xb2\x02\n\x0bVectorClock\x12(\n\x10origin_domain_id\x18\x01 \x01(\tR\x0eoriginDomainId\x12\x35\n\x08local_ts\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07localTs\x12\x19\n\x08party_id\x18\x04 \x01(\tR\x07partyId\x12Q\n\x05\x63lock\x18\x05 \x03(\x0b\x32;.com.digitalasset.canton.protocol.v0.VectorClock.ClockEntryR\x05\x63lock\x1aT\n\nClockEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x30\n\x05value\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x05value:\x02\x38\x01\"\xf3\x03\n\x0f\x43\x61usalityUpdate\x12\x32\n\x14informeeStakeholders\x18\x01 \x03(\tR\x14informeeStakeholders\x12*\n\x02ts\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02ts\x12\x1b\n\tdomain_id\x18\x03 \x01(\tR\x08\x64omainId\x12\'\n\x0frequest_counter\x18\x04 \x01(\x03R\x0erequestCounter\x12\x66\n\x11transactionUpdate\x18\x05 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.TransactionUpdateH\x00R\x11transactionUpdate\x12\x66\n\x11transferOutUpdate\x18\x06 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.TransferOutUpdateH\x00R\x11transferOutUpdate\x12\x63\n\x10transferInUpdate\x18\x07 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.TransferInUpdateH\x00R\x10transferInUpdateB\x05\n\x03tag\"\x13\n\x11TransactionUpdate\"e\n\x11TransferOutUpdate\x12P\n\x0btransfer_id\x18\x01 \x01(\x0b\x32/.com.digitalasset.canton.protocol.v0.TransferIdR\ntransferId\"d\n\x10TransferInUpdate\x12P\n\x0btransfer_id\x18\x01 \x01(\x0b\x32/.com.digitalasset.canton.protocol.v0.TransferIdR\ntransferIdBTZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.causality_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0'
  _VECTORCLOCK_CLOCKENTRY._options = None
  _VECTORCLOCK_CLOCKENTRY._serialized_options = b'8\001'
  _globals['_CAUSALITYMESSAGE']._serialized_start=190
  _globals['_CAUSALITYMESSAGE']._serialized_end=404
  _globals['_VECTORCLOCK']._serialized_start=407
  _globals['_VECTORCLOCK']._serialized_end=713
  _globals['_VECTORCLOCK_CLOCKENTRY']._serialized_start=629
  _globals['_VECTORCLOCK_CLOCKENTRY']._serialized_end=713
  _globals['_CAUSALITYUPDATE']._serialized_start=716
  _globals['_CAUSALITYUPDATE']._serialized_end=1215
  _globals['_TRANSACTIONUPDATE']._serialized_start=1217
  _globals['_TRANSACTIONUPDATE']._serialized_end=1236
  _globals['_TRANSFEROUTUPDATE']._serialized_start=1238
  _globals['_TRANSFEROUTUPDATE']._serialized_end=1339
  _globals['_TRANSFERINUPDATE']._serialized_start=1341
  _globals['_TRANSFERINUPDATE']._serialized_end=1441
# @@protoc_insertion_point(module_scope)
