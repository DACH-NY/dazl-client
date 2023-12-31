# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/domain/admin/v1/sequencer_initialization_snapshot.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from ..v0 import sequencer_administration_service_pb2 as com_dot_digitalasset_dot_canton_dot_domain_dot_admin_dot_v0_dot_sequencer__administration__service__pb2
from ..v0 import sequencer_initialization_snapshot_pb2 as com_dot_digitalasset_dot_canton_dot_domain_dot_admin_dot_v0_dot_sequencer__initialization__snapshot__pb2
from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nOcom/digitalasset/canton/domain/admin/v1/sequencer_initialization_snapshot.proto\x12\'com.digitalasset.canton.domain.admin.v1\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1aNcom/digitalasset/canton/domain/admin/v0/sequencer_administration_service.proto\x1aOcom/digitalasset/canton/domain/admin/v0/sequencer_initialization_snapshot.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xfb\r\n\x11SequencerSnapshot\x12\x45\n\x10latest_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0flatestTimestamp\x12z\n\x14head_member_counters\x18\x02 \x03(\x0b\x32H.com.digitalasset.canton.domain.admin.v1.SequencerSnapshot.MemberCounterR\x12headMemberCounters\x12W\n\x06status\x18\x03 \x01(\x0b\x32?.com.digitalasset.canton.domain.admin.v0.SequencerPruningStatusR\x06status\x12\x63\n\nadditional\x18\x04 \x01(\x0b\x32\x43.com.digitalasset.canton.domain.admin.v0.ImplementationSpecificInfoR\nadditional\x12\x8a\x01\n\x16in_flight_aggregations\x18\x05 \x03(\x0b\x32T.com.digitalasset.canton.domain.admin.v1.SequencerSnapshot.InFlightAggregationWithIdR\x14inFlightAggregations\x12}\n\x11traffic_snapshots\x18\x06 \x03(\x0b\x32P.com.digitalasset.canton.domain.admin.v1.SequencerSnapshot.MemberTrafficSnapshotR\x10trafficSnapshots\x1aT\n\rMemberCounter\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12+\n\x11sequencer_counter\x18\x02 \x01(\x03R\x10sequencerCounter\x1a\xee\x02\n\x19InFlightAggregationWithId\x12%\n\x0e\x61ggregation_id\x18\x01 \x01(\x0cR\raggregationId\x12_\n\x10\x61ggregation_rule\x18\x02 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.AggregationRuleR\x0f\x61ggregationRule\x12J\n\x13max_sequencing_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x11maxSequencingTime\x12}\n\x12\x61ggregated_senders\x18\x04 \x03(\x0b\x32N.com.digitalasset.canton.domain.admin.v1.SequencerSnapshot.AggregationBySenderR\x11\x61ggregatedSenders\x1a\x85\x02\n\x13\x41ggregationBySender\x12\x16\n\x06sender\x18\x01 \x01(\tR\x06sender\x12M\n\x14sequencing_timestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x13sequencingTimestamp\x12\x86\x01\n\x16signatures_by_envelope\x18\x03 \x03(\x0b\x32P.com.digitalasset.canton.domain.admin.v1.SequencerSnapshot.SignaturesForEnvelopeR\x14signaturesByEnvelope\x1a\x65\n\x15SignaturesForEnvelope\x12L\n\nsignatures\x18\x03 \x03(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\nsignatures\x1a\xa2\x02\n\x15MemberTrafficSnapshot\x12\x16\n\x06member\x18\x01 \x01(\tR\x06member\x12\x36\n\x17\x65xtra_traffic_remainder\x18\x02 \x01(\x04R\x15\x65xtraTrafficRemainder\x12\x34\n\x16\x65xtra_traffic_consumed\x18\x03 \x01(\x04R\x14\x65xtraTrafficConsumed\x12\x34\n\x16\x62\x61se_traffic_remainder\x18\x04 \x01(\x04R\x14\x62\x61seTrafficRemainder\x12M\n\x14sequencing_timestamp\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x13sequencingTimestampb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.admin.v1.sequencer_initialization_snapshot_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_SEQUENCERSNAPSHOT']._serialized_start=421
  _globals['_SEQUENCERSNAPSHOT']._serialized_end=2208
  _globals['_SEQUENCERSNAPSHOT_MEMBERCOUNTER']._serialized_start=1095
  _globals['_SEQUENCERSNAPSHOT_MEMBERCOUNTER']._serialized_end=1179
  _globals['_SEQUENCERSNAPSHOT_INFLIGHTAGGREGATIONWITHID']._serialized_start=1182
  _globals['_SEQUENCERSNAPSHOT_INFLIGHTAGGREGATIONWITHID']._serialized_end=1548
  _globals['_SEQUENCERSNAPSHOT_AGGREGATIONBYSENDER']._serialized_start=1551
  _globals['_SEQUENCERSNAPSHOT_AGGREGATIONBYSENDER']._serialized_end=1812
  _globals['_SEQUENCERSNAPSHOT_SIGNATURESFORENVELOPE']._serialized_start=1814
  _globals['_SEQUENCERSNAPSHOT_SIGNATURESFORENVELOPE']._serialized_end=1915
  _globals['_SEQUENCERSNAPSHOT_MEMBERTRAFFICSNAPSHOT']._serialized_start=1918
  _globals['_SEQUENCERSNAPSHOT_MEMBERTRAFFICSNAPSHOT']._serialized_end=2208
# @@protoc_insertion_point(module_scope)
