# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/synchronization.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from . import causality_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_causality__pb2
from . import participant_transaction_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_participant__transaction__pb2
from . import participant_transfer_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_participant__transfer__pb2
from . import topology_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_topology__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n9com/digitalasset/canton/protocol/v0/synchronization.proto\x12#com.digitalasset.canton.protocol.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x33\x63om/digitalasset/canton/protocol/v0/causality.proto\x1a\x41\x63om/digitalasset/canton/protocol/v0/participant_transaction.proto\x1a>com/digitalasset/canton/protocol/v0/participant_transfer.proto\x1a\x32\x63om/digitalasset/canton/protocol/v0/topology.proto\"\x84\x03\n\x15SignedProtocolMessage\x12J\n\tsignature\x18\x01 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\tsignature\x12-\n\x11mediator_response\x18\x02 \x01(\x0cH\x00R\x10mediatorResponse\x12/\n\x12transaction_result\x18\x03 \x01(\x0cH\x00R\x11transactionResult\x12K\n!malformed_mediator_request_result\x18\x04 \x01(\x0cH\x00R\x1emalformedMediatorRequestResult\x12)\n\x0ftransfer_result\x18\x05 \x01(\x0cH\x00R\x0etransferResult\x12\'\n\x0e\x61\x63s_commitment\x18\x06 \x01(\x0cH\x00R\racsCommitmentB\x1e\n\x1csome_signed_protocol_message\"\x94\n\n\x0f\x45nvelopeContent\x12\x61\n\x10informee_message\x18\x01 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.InformeeMessageH\x00R\x0finformeeMessage\x12\x63\n\x0esigned_message\x18\x02 \x01(\x0b\x32:.com.digitalasset.canton.protocol.v0.SignedProtocolMessageH\x00R\rsignedMessage\x12q\n\x16\x65ncrypted_view_message\x18\x03 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.EncryptedViewMessageH\x00R\x14\x65ncryptedViewMessage\x12\x96\x01\n#domain_topology_transaction_message\x18\x05 \x01(\x0b\x32\x45.com.digitalasset.canton.protocol.v0.DomainTopologyTransactionMessageH\x00R domainTopologyTransactionMessage\x12\x84\x01\n\x1dtransfer_out_mediator_message\x18\x06 \x01(\x0b\x32?.com.digitalasset.canton.protocol.v0.TransferOutMediatorMessageH\x00R\x1atransferOutMediatorMessage\x12\x81\x01\n\x1ctransfer_in_mediator_message\x18\x07 \x01(\x0b\x32>.com.digitalasset.canton.protocol.v0.TransferInMediatorMessageH\x00R\x19transferInMediatorMessage\x12\x62\n\x11root_hash_message\x18\x08 \x01(\x0b\x32\x34.com.digitalasset.canton.protocol.v0.RootHashMessageH\x00R\x0frootHashMessage\x12\x9c\x01\n%register_topology_transaction_request\x18\t \x01(\x0b\x32G.com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionRequestH\x00R\"registerTopologyTransactionRequest\x12\x9f\x01\n&register_topology_transaction_response\x18\n \x01(\x0b\x32H.com.digitalasset.canton.protocol.v0.RegisterTopologyTransactionResponseH\x00R#registerTopologyTransactionResponse\x12\x64\n\x11\x63\x61usality_message\x18\x0b \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.CausalityMessageH\x00R\x10\x63\x61usalityMessageB\x17\n\x15some_envelope_contentBTZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.synchronization_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZRgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/protocol/v0'
  _globals['_SIGNEDPROTOCOLMESSAGE']._serialized_start=383
  _globals['_SIGNEDPROTOCOLMESSAGE']._serialized_end=771
  _globals['_ENVELOPECONTENT']._serialized_start=774
  _globals['_ENVELOPECONTENT']._serialized_end=2074
# @@protoc_insertion_point(module_scope)
