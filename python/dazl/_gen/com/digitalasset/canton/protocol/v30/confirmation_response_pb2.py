# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v30/confirmation_response.proto
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
    'com/digitalasset/canton/protocol/v30/confirmation_response.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n@com/digitalasset/canton/protocol/v30/confirmation_response.proto\x12$com.digitalasset.canton.protocol.v30\x1a\x17google/rpc/status.proto\"\x9d\x02\n\x0cLocalVerdict\x12R\n\x04\x63ode\x18\x01 \x01(\x0e\x32>.com.digitalasset.canton.protocol.v30.LocalVerdict.VerdictCodeR\x04\x63ode\x12*\n\x06reason\x18\x02 \x01(\x0b\x32\x12.google.rpc.StatusR\x06reason\"\x8c\x01\n\x0bVerdictCode\x12\x1c\n\x18VERDICT_CODE_UNSPECIFIED\x10\x00\x12\x1e\n\x1aVERDICT_CODE_LOCAL_APPROVE\x10\x01\x12\x1d\n\x19VERDICT_CODE_LOCAL_REJECT\x10\x02\x12 \n\x1cVERDICT_CODE_LOCAL_MALFORMED\x10\x03\"\xe8\x02\n\x14\x43onfirmationResponse\x12\x1d\n\nrequest_id\x18\x01 \x01(\x03R\trequestId\x12\x16\n\x06sender\x18\x02 \x01(\tR\x06sender\x12W\n\rlocal_verdict\x18\x03 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v30.LocalVerdictR\x0clocalVerdict\x12\x1b\n\troot_hash\x18\x04 \x01(\x0cR\x08rootHash\x12-\n\x12\x63onfirming_parties\x18\x05 \x03(\tR\x11\x63onfirmingParties\x12\x1b\n\tdomain_id\x18\x06 \x01(\tR\x08\x64omainId\x12W\n\rview_position\x18\x07 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v30.ViewPositionR\x0cviewPosition\"`\n\x0cViewPosition\x12P\n\x08position\x18\x01 \x03(\x0b\x32\x34.com.digitalasset.canton.protocol.v30.MerkleSeqIndexR\x08position\"+\n\x0eMerkleSeqIndex\x12\x19\n\x08is_right\x18\x01 \x03(\x08R\x07isRightBUZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v30.confirmation_response_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZSgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v30'
  _globals['_LOCALVERDICT']._serialized_start=132
  _globals['_LOCALVERDICT']._serialized_end=417
  _globals['_LOCALVERDICT_VERDICTCODE']._serialized_start=277
  _globals['_LOCALVERDICT_VERDICTCODE']._serialized_end=417
  _globals['_CONFIRMATIONRESPONSE']._serialized_start=420
  _globals['_CONFIRMATIONRESPONSE']._serialized_end=780
  _globals['_VIEWPOSITION']._serialized_start=782
  _globals['_VIEWPOSITION']._serialized_end=878
  _globals['_MERKLESEQINDEX']._serialized_start=880
  _globals['_MERKLESEQINDEX']._serialized_end=923
# @@protoc_insertion_point(module_scope)
