# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/domain/admin/v0/domain_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2
from ....protocol.v1 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_sequencing__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n<com/digitalasset/canton/domain/admin/v0/domain_service.proto\x12\'com.digitalasset.canton.domain.admin.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a\x34\x63om/digitalasset/canton/protocol/v1/sequencing.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x83\x02\n\x13GetDomainParameters\x1a\t\n\x07Request\x1a\xe0\x01\n\x08Response\x12\x62\n\rparameters_v0\x18\x01 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v0.StaticDomainParametersH\x00R\x0cparametersV0\x12\x62\n\rparameters_v1\x18\x02 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v1.StaticDomainParametersH\x00R\x0cparametersV1B\x0c\n\nparameters\"\x84\x01\n\x1bServiceAgreementAcceptances\x12\x65\n\x0b\x61\x63\x63\x65ptances\x18\x01 \x03(\x0b\x32\x43.com.digitalasset.canton.domain.admin.v0.ServiceAgreementAcceptanceR\x0b\x61\x63\x63\x65ptances\"\xec\x01\n\x1aServiceAgreementAcceptance\x12!\n\x0c\x61greement_id\x18\x01 \x01(\tR\x0b\x61greementId\x12%\n\x0eparticipant_id\x18\x02 \x01(\tR\rparticipantId\x12J\n\tsignature\x18\x03 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\tsignature\x12\x38\n\ttimestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp2\xaa\x03\n\rDomainService\x12\x7f\n\x1fListServiceAgreementAcceptances\x12\x16.google.protobuf.Empty\x1a\x44.com.digitalasset.canton.domain.admin.v0.ServiceAgreementAcceptances\x12j\n\x13GetDomainParameters\x12\x16.google.protobuf.Empty\x1a;.com.digitalasset.canton.protocol.v0.StaticDomainParameters\x12\xab\x01\n\x1cGetDomainParametersVersioned\x12\x44.com.digitalasset.canton.domain.admin.v0.GetDomainParameters.Request\x1a\x45.com.digitalasset.canton.domain.admin.v0.GetDomainParameters.ResponseBXZVgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/admin/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.admin.v0.domain_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZVgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/domain/admin/v0'
  _globals['_GETDOMAINPARAMETERS']._serialized_start=324
  _globals['_GETDOMAINPARAMETERS']._serialized_end=583
  _globals['_GETDOMAINPARAMETERS_REQUEST']._serialized_start=347
  _globals['_GETDOMAINPARAMETERS_REQUEST']._serialized_end=356
  _globals['_GETDOMAINPARAMETERS_RESPONSE']._serialized_start=359
  _globals['_GETDOMAINPARAMETERS_RESPONSE']._serialized_end=583
  _globals['_SERVICEAGREEMENTACCEPTANCES']._serialized_start=586
  _globals['_SERVICEAGREEMENTACCEPTANCES']._serialized_end=718
  _globals['_SERVICEAGREEMENTACCEPTANCE']._serialized_start=721
  _globals['_SERVICEAGREEMENTACCEPTANCE']._serialized_end=957
  _globals['_DOMAINSERVICE']._serialized_start=960
  _globals['_DOMAINSERVICE']._serialized_end=1386
# @@protoc_insertion_point(module_scope)
