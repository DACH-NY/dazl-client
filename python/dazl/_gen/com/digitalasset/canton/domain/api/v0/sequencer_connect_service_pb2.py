# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/domain/api/v0/sequencer_connect_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import service_agreement_pb2 as com_dot_digitalasset_dot_canton_dot_domain_dot_api_dot_v0_dot_service__agreement__pb2
from ....protocol.v0 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_sequencing__pb2
from ....protocol.v1 import sequencing_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v1_dot_sequencing__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nEcom/digitalasset/canton/domain/api/v0/sequencer_connect_service.proto\x12%com.digitalasset.canton.domain.api.v0\x1a=com/digitalasset/canton/domain/api/v0/service_agreement.proto\x1a\x34\x63om/digitalasset/canton/protocol/v0/sequencing.proto\x1a\x34\x63om/digitalasset/canton/protocol/v1/sequencing.proto\"\xcf\x05\n\x10SequencerConnect\x1a\x64\n\x0bGetDomainId\x1a\t\n\x07Request\x1aJ\n\x08Response\x12\x1b\n\tdomain_id\x18\x01 \x01(\tR\x08\x64omainId\x12!\n\x0csequencer_id\x18\x02 \x01(\tR\x0bsequencerId\x1a\x83\x02\n\x13GetDomainParameters\x1a\t\n\x07Request\x1a\xe0\x01\n\x08Response\x12\x62\n\rparameters_v0\x18\x01 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v0.StaticDomainParametersH\x00R\x0cparametersV0\x12\x62\n\rparameters_v1\x18\x02 \x01(\x0b\x32;.com.digitalasset.canton.protocol.v1.StaticDomainParametersH\x00R\x0cparametersV1B\x0c\n\nparameters\x1a\xce\x02\n\x0cVerifyActive\x1a\t\n\x07Request\x1a\xe7\x01\n\x08Response\x12h\n\x07success\x18\x01 \x01(\x0b\x32L.com.digitalasset.canton.domain.api.v0.SequencerConnect.VerifyActive.SuccessH\x00R\x07success\x12h\n\x07\x66\x61ilure\x18\x02 \x01(\x0b\x32L.com.digitalasset.canton.domain.api.v0.SequencerConnect.VerifyActive.FailureH\x00R\x07\x66\x61ilureB\x07\n\x05value\x1a&\n\x07Success\x12\x1b\n\tis_active\x18\x01 \x01(\x08R\x08isActive\x1a!\n\x07\x46\x61ilure\x12\x16\n\x06reason\x18\x01 \x01(\tR\x06reason2\xd2\x06\n\x17SequencerConnectService\x12|\n\tHandshake\x12\x36.com.digitalasset.canton.protocol.v0.Handshake.Request\x1a\x37.com.digitalasset.canton.protocol.v0.Handshake.Response\x12\xa8\x01\n\x0bGetDomainId\x12K.com.digitalasset.canton.domain.api.v0.SequencerConnect.GetDomainId.Request\x1aL.com.digitalasset.canton.domain.api.v0.SequencerConnect.GetDomainId.Response\x12\xc0\x01\n\x13GetDomainParameters\x12S.com.digitalasset.canton.domain.api.v0.SequencerConnect.GetDomainParameters.Request\x1aT.com.digitalasset.canton.domain.api.v0.SequencerConnect.GetDomainParameters.Response\x12\xab\x01\n\x0cVerifyActive\x12L.com.digitalasset.canton.domain.api.v0.SequencerConnect.VerifyActive.Request\x1aM.com.digitalasset.canton.domain.api.v0.SequencerConnect.VerifyActive.Response\x12\x9c\x01\n\x13GetServiceAgreement\x12\x41.com.digitalasset.canton.domain.api.v0.GetServiceAgreementRequest\x1a\x42.com.digitalasset.canton.domain.api.v0.GetServiceAgreementResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.domain.api.v0.sequencer_connect_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_SEQUENCERCONNECT']._serialized_start=284
  _globals['_SEQUENCERCONNECT']._serialized_end=1003
  _globals['_SEQUENCERCONNECT_GETDOMAINID']._serialized_start=304
  _globals['_SEQUENCERCONNECT_GETDOMAINID']._serialized_end=404
  _globals['_SEQUENCERCONNECT_GETDOMAINID_REQUEST']._serialized_start=319
  _globals['_SEQUENCERCONNECT_GETDOMAINID_REQUEST']._serialized_end=328
  _globals['_SEQUENCERCONNECT_GETDOMAINID_RESPONSE']._serialized_start=330
  _globals['_SEQUENCERCONNECT_GETDOMAINID_RESPONSE']._serialized_end=404
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS']._serialized_start=407
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS']._serialized_end=666
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS_REQUEST']._serialized_start=319
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS_REQUEST']._serialized_end=328
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS_RESPONSE']._serialized_start=442
  _globals['_SEQUENCERCONNECT_GETDOMAINPARAMETERS_RESPONSE']._serialized_end=666
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE']._serialized_start=669
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE']._serialized_end=1003
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_REQUEST']._serialized_start=319
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_REQUEST']._serialized_end=328
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_RESPONSE']._serialized_start=697
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_RESPONSE']._serialized_end=928
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_SUCCESS']._serialized_start=930
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_SUCCESS']._serialized_end=968
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_FAILURE']._serialized_start=970
  _globals['_SEQUENCERCONNECT_VERIFYACTIVE_FAILURE']._serialized_end=1003
  _globals['_SEQUENCERCONNECTSERVICE']._serialized_start=1006
  _globals['_SEQUENCERCONNECTSERVICE']._serialized_end=1856
# @@protoc_insertion_point(module_scope)
