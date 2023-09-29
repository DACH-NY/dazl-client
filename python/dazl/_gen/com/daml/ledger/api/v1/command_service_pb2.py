# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/command_service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import commands_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_commands__pb2
from . import transaction_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_transaction__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,com/daml/ledger/api/v1/command_service.proto\x12\x16\x63om.daml.ledger.api.v1\x1a%com/daml/ledger/api/v1/commands.proto\x1a(com/daml/ledger/api/v1/transaction.proto\x1a\x1bgoogle/protobuf/empty.proto\"T\n\x14SubmitAndWaitRequest\x12<\n\x08\x63ommands\x18\x01 \x01(\x0b\x32 .com.daml.ledger.api.v1.CommandsR\x08\x63ommands\"{\n%SubmitAndWaitForTransactionIdResponse\x12%\n\x0etransaction_id\x18\x01 \x01(\tR\rtransactionId\x12+\n\x11\x63ompletion_offset\x18\x02 \x01(\tR\x10\x63ompletionOffset\"\x99\x01\n#SubmitAndWaitForTransactionResponse\x12\x45\n\x0btransaction\x18\x01 \x01(\x0b\x32#.com.daml.ledger.api.v1.TransactionR\x0btransaction\x12+\n\x11\x63ompletion_offset\x18\x02 \x01(\tR\x10\x63ompletionOffset\"\xa1\x01\n\'SubmitAndWaitForTransactionTreeResponse\x12I\n\x0btransaction\x18\x01 \x01(\x0b\x32\'.com.daml.ledger.api.v1.TransactionTreeR\x0btransaction\x12+\n\x11\x63ompletion_offset\x18\x02 \x01(\tR\x10\x63ompletionOffset2\x94\x04\n\x0e\x43ommandService\x12U\n\rSubmitAndWait\x12,.com.daml.ledger.api.v1.SubmitAndWaitRequest\x1a\x16.google.protobuf.Empty\x12\x8c\x01\n\x1dSubmitAndWaitForTransactionId\x12,.com.daml.ledger.api.v1.SubmitAndWaitRequest\x1a=.com.daml.ledger.api.v1.SubmitAndWaitForTransactionIdResponse\x12\x88\x01\n\x1bSubmitAndWaitForTransaction\x12,.com.daml.ledger.api.v1.SubmitAndWaitRequest\x1a;.com.daml.ledger.api.v1.SubmitAndWaitForTransactionResponse\x12\x90\x01\n\x1fSubmitAndWaitForTransactionTree\x12,.com.daml.ledger.api.v1.SubmitAndWaitRequest\x1a?.com.daml.ledger.api.v1.SubmitAndWaitForTransactionTreeResponseB\x92\x01\n\x16\x63om.daml.ledger.api.v1B\x18\x43ommandServiceOuterClassZEgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1\xaa\x02\x16\x43om.Daml.Ledger.Api.V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.command_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026com.daml.ledger.api.v1B\030CommandServiceOuterClassZEgithub.com/digital-asset/dazl-client/v7/go/api/com/daml/ledger/api/v1\252\002\026Com.Daml.Ledger.Api.V1'
  _globals['_SUBMITANDWAITREQUEST']._serialized_start=182
  _globals['_SUBMITANDWAITREQUEST']._serialized_end=266
  _globals['_SUBMITANDWAITFORTRANSACTIONIDRESPONSE']._serialized_start=268
  _globals['_SUBMITANDWAITFORTRANSACTIONIDRESPONSE']._serialized_end=391
  _globals['_SUBMITANDWAITFORTRANSACTIONRESPONSE']._serialized_start=394
  _globals['_SUBMITANDWAITFORTRANSACTIONRESPONSE']._serialized_end=547
  _globals['_SUBMITANDWAITFORTRANSACTIONTREERESPONSE']._serialized_start=550
  _globals['_SUBMITANDWAITFORTRANSACTIONTREERESPONSE']._serialized_end=711
  _globals['_COMMANDSERVICE']._serialized_start=714
  _globals['_COMMANDSERVICE']._serialized_end=1246
# @@protoc_insertion_point(module_scope)
