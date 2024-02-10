# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/topology/admin/v1/topology_manager_write_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ....protocol.v2 import topology_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v2_dot_topology__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nNcom/digitalasset/canton/topology/admin/v1/topology_manager_write_service.proto\x12)com.digitalasset.canton.topology.admin.v1\x1a\x32\x63om/digitalasset/canton/protocol/v2/topology.proto\"\xe3\x03\n\x10\x41uthorizeRequest\x12\x62\n\x08proposal\x18\x01 \x01(\x0b\x32\x44.com.digitalasset.canton.topology.admin.v1.AuthorizeRequest.ProposalH\x00R\x08proposal\x12+\n\x10transaction_hash\x18\x02 \x01(\tH\x00R\x0ftransactionHash\x12\x30\n\x14must_fully_authorize\x18\x03 \x01(\x08R\x12mustFullyAuthorize\x12!\n\x0c\x66orce_change\x18\x04 \x01(\x08R\x0b\x66orceChange\x12\x1b\n\tsigned_by\x18\x05 \x03(\tR\x08signedBy\x1a\xc3\x01\n\x08Proposal\x12N\n\x06\x63hange\x18\x01 \x01(\x0e\x32\x36.com.digitalasset.canton.protocol.v2.TopologyChangeOpXR\x06\x63hange\x12\x16\n\x06serial\x18\x02 \x01(\rR\x06serial\x12O\n\x07mapping\x18\x03 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v2.TopologyMappingXR\x07mappingB\x06\n\x04type\"v\n\x11\x41uthorizeResponse\x12\x61\n\x0btransaction\x18\x01 \x01(\x0b\x32?.com.digitalasset.canton.protocol.v2.SignedTopologyTransactionXR\x0btransaction\"\xa0\x01\n\x16\x41\x64\x64TransactionsRequest\x12\x63\n\x0ctransactions\x18\x01 \x03(\x0b\x32?.com.digitalasset.canton.protocol.v2.SignedTopologyTransactionXR\x0ctransactions\x12!\n\x0c\x66orce_change\x18\x02 \x01(\x08R\x0b\x66orceChange\"\x19\n\x17\x41\x64\x64TransactionsResponse2\xc2\x02\n\x1cTopologyManagerWriteServiceX\x12\x86\x01\n\tAuthorize\x12;.com.digitalasset.canton.topology.admin.v1.AuthorizeRequest\x1a<.com.digitalasset.canton.topology.admin.v1.AuthorizeResponse\x12\x98\x01\n\x0f\x41\x64\x64Transactions\x12\x41.com.digitalasset.canton.topology.admin.v1.AddTransactionsRequest\x1a\x42.com.digitalasset.canton.topology.admin.v1.AddTransactionsResponseBZZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.topology.admin.v1.topology_manager_write_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZXgithub.com/digital-asset/dazl-client/v7/go/api/com/digitalasset/canton/topology/admin/v1'
  _globals['_AUTHORIZEREQUEST']._serialized_start=178
  _globals['_AUTHORIZEREQUEST']._serialized_end=661
  _globals['_AUTHORIZEREQUEST_PROPOSAL']._serialized_start=458
  _globals['_AUTHORIZEREQUEST_PROPOSAL']._serialized_end=653
  _globals['_AUTHORIZERESPONSE']._serialized_start=663
  _globals['_AUTHORIZERESPONSE']._serialized_end=781
  _globals['_ADDTRANSACTIONSREQUEST']._serialized_start=784
  _globals['_ADDTRANSACTIONSREQUEST']._serialized_end=944
  _globals['_ADDTRANSACTIONSRESPONSE']._serialized_start=946
  _globals['_ADDTRANSACTIONSRESPONSE']._serialized_end=971
  _globals['_TOPOLOGYMANAGERWRITESERVICEX']._serialized_start=974
  _globals['_TOPOLOGYMANAGERWRITESERVICEX']._serialized_end=1296
# @@protoc_insertion_point(module_scope)
