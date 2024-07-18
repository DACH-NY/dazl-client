# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/sequencing.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from ...v0 import trace_context_pb2 as com_dot_digitalasset_dot_canton_dot_v0_dot_trace__context__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4com/digitalasset/canton/protocol/v0/sequencing.proto\x12#com.digitalasset.canton.protocol.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a.com/digitalasset/canton/v0/trace_context.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xf5\x02\n\x0eSequencedEvent\x12\x18\n\x07\x63ounter\x18\x01 \x01(\x03R\x07\x63ounter\x12\x38\n\ttimestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12\x1b\n\tdomain_id\x18\x03 \x01(\tR\x08\x64omainId\x12;\n\nmessage_id\x18\x04 \x01(\x0b\x32\x1c.google.protobuf.StringValueR\tmessageId\x12J\n\x05\x62\x61tch\x18\x05 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.CompressedBatchR\x05\x62\x61tch\x12i\n\x14\x64\x65liver_error_reason\x18\x06 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.DeliverErrorReasonR\x12\x64\x65liverErrorReason\"\xb5\x02\n\x1dPossiblyIgnoredSequencedEvent\x12\x18\n\x07\x63ounter\x18\x01 \x01(\x03R\x07\x63ounter\x12\x38\n\ttimestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12M\n\rtrace_context\x18\x03 \x01(\x0b\x32(.com.digitalasset.canton.v0.TraceContextR\x0ctraceContext\x12\x1d\n\nis_ignored\x18\x04 \x01(\x08R\tisIgnored\x12R\n\nunderlying\x18\x05 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v0.SignedContentR\nunderlying\"\x81\x01\n\x0eRecipientsTree\x12\x1e\n\nrecipients\x18\x01 \x03(\tR\nrecipients\x12O\n\x08\x63hildren\x18\x02 \x03(\x0b\x32\x33.com.digitalasset.canton.protocol.v0.RecipientsTreeR\x08\x63hildren\"j\n\nRecipients\x12\\\n\x0frecipients_tree\x18\x01 \x03(\x0b\x32\x33.com.digitalasset.canton.protocol.v0.RecipientsTreeR\x0erecipientsTree\"u\n\x08\x45nvelope\x12\x18\n\x07\x63ontent\x18\x01 \x01(\x0cR\x07\x63ontent\x12O\n\nrecipients\x18\x02 \x01(\x0b\x32/.com.digitalasset.canton.protocol.v0.RecipientsR\nrecipients\"\xd2\x01\n\x0f\x43ompressedBatch\x12g\n\talgorithm\x18\x01 \x01(\x0e\x32I.com.digitalasset.canton.protocol.v0.CompressedBatch.CompressionAlgorithmR\talgorithm\x12)\n\x10\x63ompressed_batch\x18\x02 \x01(\x0cR\x0f\x63ompressedBatch\"+\n\x14\x43ompressionAlgorithm\x12\t\n\x05None_\x10\x00\x12\x08\n\x04Gzip\x10\x01\"T\n\x05\x42\x61tch\x12K\n\tenvelopes\x18\x01 \x03(\x0b\x32-.com.digitalasset.canton.protocol.v0.EnvelopeR\tenvelopes\"\xe9\x01\n\rSignedContent\x12\x35\n\x07\x63ontent\x18\x01 \x01(\x0b\x32\x1b.google.protobuf.BytesValueR\x07\x63ontent\x12L\n\nsignatures\x18\x02 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\nsignatures\x12S\n\x18timestamp_of_signing_key\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x15timestampOfSigningKey\"l\n\x12\x44\x65liverErrorReason\x12%\n\rbatch_invalid\x18\x01 \x01(\tH\x00R\x0c\x62\x61tchInvalid\x12%\n\rbatch_refused\x18\x02 \x01(\tH\x00R\x0c\x62\x61tchRefusedB\x08\n\x06reason\"A\n\x10ServiceAgreement\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n\nlegal_text\x18\x02 \x01(\tR\tlegalText\"\x8b\x07\n\x16StaticDomainParameters\x12R\n\x17reconciliation_interval\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationR\x16reconciliationInterval\x12\x37\n\x18max_rate_per_participant\x18\x02 \x01(\rR\x15maxRatePerParticipant\x12\x37\n\x18max_inbound_message_size\x18\x03 \x01(\rR\x15maxInboundMessageSize\x12\x30\n\x14unique_contract_keys\x18\x04 \x01(\x08R\x12uniqueContractKeys\x12t\n\x1crequired_signing_key_schemes\x18\x05 \x03(\x0e\x32\x33.com.digitalasset.canton.crypto.v0.SigningKeySchemeR\x19requiredSigningKeySchemes\x12}\n\x1frequired_encryption_key_schemes\x18\x06 \x03(\x0e\x32\x36.com.digitalasset.canton.crypto.v0.EncryptionKeySchemeR\x1crequiredEncryptionKeySchemes\x12z\n\x1erequired_symmetric_key_schemes\x18\x07 \x03(\x0e\x32\x35.com.digitalasset.canton.crypto.v0.SymmetricKeySchemeR\x1brequiredSymmetricKeySchemes\x12j\n\x18required_hash_algorithms\x18\x08 \x03(\x0e\x32\x30.com.digitalasset.canton.crypto.v0.HashAlgorithmR\x16requiredHashAlgorithms\x12q\n\x1brequired_crypto_key_formats\x18\t \x03(\x0e\x32\x32.com.digitalasset.canton.crypto.v0.CryptoKeyFormatR\x18requiredCryptoKeyFormats\x12)\n\x10protocol_version\x18\n \x01(\tR\x0fprotocolVersion\"\xde\x03\n\x17\x44ynamicDomainParameters\x12[\n\x1cparticipant_response_timeout\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationR\x1aparticipantResponseTimeout\x12U\n\x19mediator_reaction_timeout\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x17mediatorReactionTimeout\x12[\n\x1ctransfer_exclusivity_timeout\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x1atransferExclusivityTimeout\x12M\n\x15topology_change_delay\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x13topologyChangeDelay\x12\x63\n!ledger_time_record_time_tolerance\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x1dledgerTimeRecordTimeTolerance\"\xcd\x03\n\tHandshake\x1a\x9b\x01\n\x07Request\x12\x38\n\x18\x63lient_protocol_versions\x18\x01 \x03(\tR\x16\x63lientProtocolVersions\x12V\n\x18minimum_protocol_version\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValueR\x16minimumProtocolVersion\x1a\xf3\x01\n\x08Response\x12\x36\n\x17server_protocol_version\x18\x01 \x01(\tR\x15serverProtocolVersion\x12R\n\x07success\x18\x02 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.Handshake.SuccessH\x00R\x07success\x12R\n\x07\x66\x61ilure\x18\x03 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.Handshake.FailureH\x00R\x07\x66\x61ilureB\x07\n\x05value\x1a\t\n\x07Success\x1a!\n\x07\x46\x61ilure\x12\x16\n\x06reason\x18\x01 \x01(\tR\x06reason\"\xd6\x02\n\x11SubmissionRequest\x12\x16\n\x06sender\x18\x01 \x01(\tR\x06sender\x12\x1d\n\nmessage_id\x18\x02 \x01(\tR\tmessageId\x12\x1d\n\nis_request\x18\x03 \x01(\x08R\tisRequest\x12J\n\x05\x62\x61tch\x18\x04 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.CompressedBatchR\x05\x62\x61tch\x12J\n\x13max_sequencing_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x11maxSequencingTime\x12S\n\x18timestamp_of_signing_key\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x15timestampOfSigningKeyBTZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.sequencing_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0'
  _globals['_SEQUENCEDEVENT']._serialized_start=287
  _globals['_SEQUENCEDEVENT']._serialized_end=660
  _globals['_POSSIBLYIGNOREDSEQUENCEDEVENT']._serialized_start=663
  _globals['_POSSIBLYIGNOREDSEQUENCEDEVENT']._serialized_end=972
  _globals['_RECIPIENTSTREE']._serialized_start=975
  _globals['_RECIPIENTSTREE']._serialized_end=1104
  _globals['_RECIPIENTS']._serialized_start=1106
  _globals['_RECIPIENTS']._serialized_end=1212
  _globals['_ENVELOPE']._serialized_start=1214
  _globals['_ENVELOPE']._serialized_end=1331
  _globals['_COMPRESSEDBATCH']._serialized_start=1334
  _globals['_COMPRESSEDBATCH']._serialized_end=1544
  _globals['_COMPRESSEDBATCH_COMPRESSIONALGORITHM']._serialized_start=1501
  _globals['_COMPRESSEDBATCH_COMPRESSIONALGORITHM']._serialized_end=1544
  _globals['_BATCH']._serialized_start=1546
  _globals['_BATCH']._serialized_end=1630
  _globals['_SIGNEDCONTENT']._serialized_start=1633
  _globals['_SIGNEDCONTENT']._serialized_end=1866
  _globals['_DELIVERERRORREASON']._serialized_start=1868
  _globals['_DELIVERERRORREASON']._serialized_end=1976
  _globals['_SERVICEAGREEMENT']._serialized_start=1978
  _globals['_SERVICEAGREEMENT']._serialized_end=2043
  _globals['_STATICDOMAINPARAMETERS']._serialized_start=2046
  _globals['_STATICDOMAINPARAMETERS']._serialized_end=2953
  _globals['_DYNAMICDOMAINPARAMETERS']._serialized_start=2956
  _globals['_DYNAMICDOMAINPARAMETERS']._serialized_end=3434
  _globals['_HANDSHAKE']._serialized_start=3437
  _globals['_HANDSHAKE']._serialized_end=3898
  _globals['_HANDSHAKE_REQUEST']._serialized_start=3451
  _globals['_HANDSHAKE_REQUEST']._serialized_end=3606
  _globals['_HANDSHAKE_RESPONSE']._serialized_start=3609
  _globals['_HANDSHAKE_RESPONSE']._serialized_end=3852
  _globals['_HANDSHAKE_SUCCESS']._serialized_start=3854
  _globals['_HANDSHAKE_SUCCESS']._serialized_end=3863
  _globals['_HANDSHAKE_FAILURE']._serialized_start=3865
  _globals['_HANDSHAKE_FAILURE']._serialized_end=3898
  _globals['_SUBMISSIONREQUEST']._serialized_start=3901
  _globals['_SUBMISSIONREQUEST']._serialized_end=4243
# @@protoc_insertion_point(module_scope)
