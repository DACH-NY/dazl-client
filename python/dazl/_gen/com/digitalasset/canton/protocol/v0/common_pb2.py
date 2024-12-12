# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/digitalasset/canton/protocol/v0/common.proto
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
    'com/digitalasset/canton/protocol/v0/common.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0com/digitalasset/canton/protocol/v0/common.proto\x12#com.digitalasset.canton.protocol.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x82\x04\n\x14SerializableContract\x12\x1f\n\x0b\x63ontract_id\x18\x01 \x01(\tR\ncontractId\x12\x32\n\x15raw_contract_instance\x18\x02 \x01(\x0cR\x13rawContractInstance\x12^\n\x08metadata\x18\x03 \x01(\x0b\x32\x42.com.digitalasset.canton.protocol.v0.SerializableContract.MetadataR\x08metadata\x12H\n\x12ledger_create_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x10ledgerCreateTime\x1a\xea\x01\n\x08Metadata\x12<\n\x1anon_maintainer_signatories\x18\x01 \x03(\tR\x18nonMaintainerSignatories\x12<\n\x1anon_signatory_stakeholders\x18\x02 \x03(\tR\x18nonSignatoryStakeholders\x12@\n\x03key\x18\x03 \x01(\x0b\x32..com.digitalasset.canton.protocol.v0.GlobalKeyR\x03key\x12 \n\x0bmaintainers\x18\x04 \x03(\tR\x0bmaintainers\">\n\tGlobalKey\x12\x1f\n\x0btemplate_id\x18\x01 \x01(\x0cR\ntemplateId\x12\x10\n\x03key\x18\x02 \x01(\x0cR\x03key\"f\n\x16\x44riverContractMetadata\x12L\n\rcontract_salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x0c\x63ontractSalt*i\n\x08ViewType\x12\x13\n\x0fMissingViewType\x10\x00\x12\x17\n\x13TransactionViewType\x10\x01\x12\x17\n\x13TransferOutViewType\x10\x02\x12\x16\n\x12TransferInViewType\x10\x03\x42TZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.common_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/protocol/v0'
  _globals['_VIEWTYPE']._serialized_start=855
  _globals['_VIEWTYPE']._serialized_end=960
  _globals['_SERIALIZABLECONTRACT']._serialized_start=171
  _globals['_SERIALIZABLECONTRACT']._serialized_end=685
  _globals['_SERIALIZABLECONTRACT_METADATA']._serialized_start=451
  _globals['_SERIALIZABLECONTRACT_METADATA']._serialized_end=685
  _globals['_GLOBALKEY']._serialized_start=687
  _globals['_GLOBALKEY']._serialized_end=749
  _globals['_DRIVERCONTRACTMETADATA']._serialized_start=751
  _globals['_DRIVERCONTRACTMETADATA']._serialized_end=853
# @@protoc_insertion_point(module_scope)
