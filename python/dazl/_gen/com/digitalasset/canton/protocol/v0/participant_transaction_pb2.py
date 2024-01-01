# Copyright (c) 2017-2023 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/canton/protocol/v0/participant_transaction.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from ...crypto.v0 import crypto_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_v0_dot_crypto__pb2
from . import common_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_common__pb2
from . import merkle_pb2 as com_dot_digitalasset_dot_canton_dot_protocol_dot_v0_dot_merkle__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nAcom/digitalasset/canton/protocol/v0/participant_transaction.proto\x12#com.digitalasset.canton.protocol.v0\x1a.com/digitalasset/canton/crypto/v0/crypto.proto\x1a\x30\x63om/digitalasset/canton/protocol/v0/common.proto\x1a\x30\x63om/digitalasset/canton/protocol/v0/merkle.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x91\x03\n\x14\x45ncryptedViewMessage\x12\x1b\n\tview_tree\x18\x01 \x01(\x0cR\x08viewTree\x12t\n\x1fsubmitter_participant_signature\x18\x02 \x01(\x0b\x32,.com.digitalasset.canton.crypto.v0.SignatureR\x1dsubmitterParticipantSignature\x12\x1b\n\tview_hash\x18\x03 \x01(\x0cR\x08viewHash\x12`\n\nrandomness\x18\x04 \x03(\x0b\x32@.com.digitalasset.canton.protocol.v0.ParticipantRandomnessLookupR\nrandomness\x12\x1b\n\tdomain_id\x18\x05 \x01(\tR\x08\x64omainId\x12J\n\tview_type\x18\x06 \x01(\x0e\x32-.com.digitalasset.canton.protocol.v0.ViewTypeR\x08viewType\"g\n\x18LightTransactionViewTree\x12K\n\x04tree\x18\x01 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.GenTransactionTreeR\x04tree\"\xe4\x02\n\x11SubmitterMetadata\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12\x15\n\x06\x61\x63t_as\x18\x02 \x03(\tR\x05\x61\x63tAs\x12%\n\x0e\x61pplication_id\x18\x03 \x01(\tR\rapplicationId\x12\x1d\n\ncommand_id\x18\x04 \x01(\tR\tcommandId\x12\x33\n\x15submitter_participant\x18\x05 \x01(\tR\x14submitterParticipant\x12#\n\rsubmission_id\x18\x06 \x01(\tR\x0csubmissionId\x12[\n\x0c\x64\x65\x64up_period\x18\x07 \x01(\x0b\x32\x38.com.digitalasset.canton.protocol.v0.DeduplicationPeriodR\x0b\x64\x65\x64upPeriod\"r\n\x13\x44\x65\x64uplicationPeriod\x12\x37\n\x08\x64uration\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00R\x08\x64uration\x12\x18\n\x06offset\x18\x02 \x01(\x0cH\x00R\x06offsetB\x08\n\x06period\"\xd0\x01\n\x0e\x43ommonMetadata\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12/\n\x13\x63onfirmation_policy\x18\x02 \x01(\x0cR\x12\x63onfirmationPolicy\x12\x1b\n\tdomain_id\x18\x03 \x01(\tR\x08\x64omainId\x12\x12\n\x04uuid\x18\x04 \x01(\tR\x04uuid\x12\x1f\n\x0bmediator_id\x18\x05 \x01(\tR\nmediatorId\"\xf5\x01\n\x13ParticipantMetadata\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12;\n\x0bledger_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nledgerTime\x12\x43\n\x0fsubmission_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0esubmissionTime\x12\x1f\n\x0bworkflow_id\x18\x04 \x01(\tR\nworkflowId\"\xa0\x02\n\x08ViewNode\x12\\\n\x10view_common_data\x18\x01 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v0.BlindableNodeR\x0eviewCommonData\x12\x66\n\x15view_participant_data\x18\x02 \x01(\x0b\x32\x32.com.digitalasset.canton.protocol.v0.BlindableNodeR\x13viewParticipantData\x12N\n\x08subviews\x18\x03 \x03(\x0b\x32\x32.com.digitalasset.canton.protocol.v0.BlindableNodeR\x08subviews\"\xb8\x01\n\x0eViewCommonData\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12K\n\tinformees\x18\x02 \x03(\x0b\x32-.com.digitalasset.canton.protocol.v0.InformeeR\tinformees\x12\x1c\n\tthreshold\x18\x03 \x01(\x05R\tthreshold\"8\n\x08Informee\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x12\x16\n\x06weight\x18\x02 \x01(\x05R\x06weight\"\xd8\n\n\x13ViewParticipantData\x12;\n\x04salt\x18\x01 \x01(\x0b\x32\'.com.digitalasset.canton.crypto.v0.SaltR\x04salt\x12g\n\x0b\x63ore_inputs\x18\x02 \x03(\x0b\x32\x46.com.digitalasset.canton.protocol.v0.ViewParticipantData.InputContractR\ncoreInputs\x12k\n\x0c\x63reated_core\x18\x03 \x03(\x0b\x32H.com.digitalasset.canton.protocol.v0.ViewParticipantData.CreatedContractR\x0b\x63reatedCore\x12K\n#created_in_subview_archived_in_core\x18\x04 \x03(\tR\x1e\x63reatedInSubviewArchivedInCore\x12i\n\rresolved_keys\x18\x05 \x03(\x0b\x32\x44.com.digitalasset.canton.protocol.v0.ViewParticipantData.ResolvedKeyR\x0cresolvedKeys\x12\x65\n\x12\x61\x63tion_description\x18\x06 \x01(\x0b\x32\x36.com.digitalasset.canton.protocol.v0.ActionDescriptionR\x11\x61\x63tionDescription\x12s\n\x10rollback_context\x18\x07 \x01(\x0b\x32H.com.digitalasset.canton.protocol.v0.ViewParticipantData.RollbackContextR\x0frollbackContext\x1a\xb3\x01\n\x0f\x43reatedContract\x12U\n\x08\x63ontract\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.SerializableContractR\x08\x63ontract\x12(\n\x10\x63onsumed_in_core\x18\x02 \x01(\x08R\x0e\x63onsumedInCore\x12\x1f\n\x0brolled_back\x18\x03 \x01(\x08R\nrolledBack\x1a\x82\x01\n\rInputContract\x12U\n\x08\x63ontract\x18\x01 \x01(\x0b\x32\x39.com.digitalasset.canton.protocol.v0.SerializableContractR\x08\x63ontract\x12\x1a\n\x08\x63onsumed\x18\x02 \x01(\x08R\x08\x63onsumed\x1a\xd8\x01\n\x0bResolvedKey\x12@\n\x03key\x18\x01 \x01(\x0b\x32..com.digitalasset.canton.protocol.v0.GlobalKeyR\x03key\x12!\n\x0b\x63ontract_id\x18\x02 \x01(\tH\x00R\ncontractId\x12V\n\x04\x66ree\x18\x03 \x01(\x0b\x32@.com.digitalasset.canton.protocol.v0.ViewParticipantData.FreeKeyH\x00R\x04\x66reeB\x0c\n\nresolution\x1a+\n\x07\x46reeKey\x12 \n\x0bmaintainers\x18\x01 \x03(\tR\x0bmaintainers\x1aW\n\x0fRollbackContext\x12%\n\x0erollback_scope\x18\x01 \x03(\x05R\rrollbackScope\x12\x1d\n\nnext_child\x18\x02 \x01(\x05R\tnextChild\"\xc6\x08\n\x11\x41\x63tionDescription\x12h\n\x06\x63reate\x18\x01 \x01(\x0b\x32N.com.digitalasset.canton.protocol.v0.ActionDescription.CreateActionDescriptionH\x00R\x06\x63reate\x12n\n\x08\x65xercise\x18\x02 \x01(\x0b\x32P.com.digitalasset.canton.protocol.v0.ActionDescription.ExerciseActionDescriptionH\x00R\x08\x65xercise\x12\x65\n\x05\x66\x65tch\x18\x03 \x01(\x0b\x32M.com.digitalasset.canton.protocol.v0.ActionDescription.FetchActionDescriptionH\x00R\x05\x66\x65tch\x12y\n\rlookup_by_key\x18\x04 \x01(\x0b\x32S.com.digitalasset.canton.protocol.v0.ActionDescription.LookupByKeyActionDescriptionH\x00R\x0blookupByKey\x1aq\n\x17\x43reateActionDescription\x12\x1f\n\x0b\x63ontract_id\x18\x01 \x01(\tR\ncontractId\x12\x1b\n\tnode_seed\x18\x02 \x01(\x0cR\x08nodeSeed\x12\x18\n\x07version\x18\x03 \x01(\tR\x07version\x1a\x80\x02\n\x19\x45xerciseActionDescription\x12*\n\x11input_contract_id\x18\x01 \x01(\tR\x0finputContractId\x12\x16\n\x06\x63hoice\x18\x02 \x01(\tR\x06\x63hoice\x12!\n\x0c\x63hosen_value\x18\x03 \x01(\x0cR\x0b\x63hosenValue\x12\x16\n\x06\x61\x63tors\x18\x04 \x03(\tR\x06\x61\x63tors\x12\x15\n\x06\x62y_key\x18\x05 \x01(\x08R\x05\x62yKey\x12\x1b\n\tnode_seed\x18\x06 \x01(\x0cR\x08nodeSeed\x12\x18\n\x07version\x18\x07 \x01(\tR\x07version\x12\x16\n\x06\x66\x61iled\x18\x08 \x01(\x08R\x06\x66\x61iled\x1a\x8d\x01\n\x16\x46\x65tchActionDescription\x12*\n\x11input_contract_id\x18\x01 \x01(\tR\x0finputContractId\x12\x16\n\x06\x61\x63tors\x18\x02 \x03(\tR\x06\x61\x63tors\x12\x15\n\x06\x62y_key\x18\x03 \x01(\x08R\x05\x62yKey\x12\x18\n\x07version\x18\x04 \x01(\tR\x07version\x1a`\n\x1cLookupByKeyActionDescription\x12@\n\x03key\x18\x01 \x01(\x0b\x32..com.digitalasset.canton.protocol.v0.GlobalKeyR\x03keyB\r\n\x0b\x64\x65scription\"_\n\x1bParticipantRandomnessLookup\x12 \n\x0bparticipant\x18\x01 \x01(\tR\x0bparticipant\x12\x1e\n\nrandomness\x18\x02 \x01(\x0cR\nrandomness\"v\n\x0fInformeeMessage\x12\x63\n\x12\x66ull_informee_tree\x18\x01 \x01(\x0b\x32\x35.com.digitalasset.canton.protocol.v0.FullInformeeTreeR\x10\x66ullInformeeTree\"_\n\x10\x46ullInformeeTree\x12K\n\x04tree\x18\x01 \x01(\x0b\x32\x37.com.digitalasset.canton.protocol.v0.GenTransactionTreeR\x04tree\"\xb1\x01\n\x0fRootHashMessage\x12\x1b\n\troot_hash\x18\x01 \x01(\x0cR\x08rootHash\x12\x1b\n\tdomain_id\x18\x02 \x01(\tR\x08\x64omainId\x12J\n\tview_type\x18\x03 \x01(\x0e\x32-.com.digitalasset.canton.protocol.v0.ViewTypeR\x08viewType\x12\x18\n\x07payload\x18\x04 \x01(\x0cR\x07payloadb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.digitalasset.canton.protocol.v0.participant_transaction_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_ENCRYPTEDVIEWMESSAGE']._serialized_start=320
  _globals['_ENCRYPTEDVIEWMESSAGE']._serialized_end=721
  _globals['_LIGHTTRANSACTIONVIEWTREE']._serialized_start=723
  _globals['_LIGHTTRANSACTIONVIEWTREE']._serialized_end=826
  _globals['_SUBMITTERMETADATA']._serialized_start=829
  _globals['_SUBMITTERMETADATA']._serialized_end=1185
  _globals['_DEDUPLICATIONPERIOD']._serialized_start=1187
  _globals['_DEDUPLICATIONPERIOD']._serialized_end=1301
  _globals['_COMMONMETADATA']._serialized_start=1304
  _globals['_COMMONMETADATA']._serialized_end=1512
  _globals['_PARTICIPANTMETADATA']._serialized_start=1515
  _globals['_PARTICIPANTMETADATA']._serialized_end=1760
  _globals['_VIEWNODE']._serialized_start=1763
  _globals['_VIEWNODE']._serialized_end=2051
  _globals['_VIEWCOMMONDATA']._serialized_start=2054
  _globals['_VIEWCOMMONDATA']._serialized_end=2238
  _globals['_INFORMEE']._serialized_start=2240
  _globals['_INFORMEE']._serialized_end=2296
  _globals['_VIEWPARTICIPANTDATA']._serialized_start=2299
  _globals['_VIEWPARTICIPANTDATA']._serialized_end=3667
  _globals['_VIEWPARTICIPANTDATA_CREATEDCONTRACT']._serialized_start=3002
  _globals['_VIEWPARTICIPANTDATA_CREATEDCONTRACT']._serialized_end=3181
  _globals['_VIEWPARTICIPANTDATA_INPUTCONTRACT']._serialized_start=3184
  _globals['_VIEWPARTICIPANTDATA_INPUTCONTRACT']._serialized_end=3314
  _globals['_VIEWPARTICIPANTDATA_RESOLVEDKEY']._serialized_start=3317
  _globals['_VIEWPARTICIPANTDATA_RESOLVEDKEY']._serialized_end=3533
  _globals['_VIEWPARTICIPANTDATA_FREEKEY']._serialized_start=3535
  _globals['_VIEWPARTICIPANTDATA_FREEKEY']._serialized_end=3578
  _globals['_VIEWPARTICIPANTDATA_ROLLBACKCONTEXT']._serialized_start=3580
  _globals['_VIEWPARTICIPANTDATA_ROLLBACKCONTEXT']._serialized_end=3667
  _globals['_ACTIONDESCRIPTION']._serialized_start=3670
  _globals['_ACTIONDESCRIPTION']._serialized_end=4764
  _globals['_ACTIONDESCRIPTION_CREATEACTIONDESCRIPTION']._serialized_start=4135
  _globals['_ACTIONDESCRIPTION_CREATEACTIONDESCRIPTION']._serialized_end=4248
  _globals['_ACTIONDESCRIPTION_EXERCISEACTIONDESCRIPTION']._serialized_start=4251
  _globals['_ACTIONDESCRIPTION_EXERCISEACTIONDESCRIPTION']._serialized_end=4507
  _globals['_ACTIONDESCRIPTION_FETCHACTIONDESCRIPTION']._serialized_start=4510
  _globals['_ACTIONDESCRIPTION_FETCHACTIONDESCRIPTION']._serialized_end=4651
  _globals['_ACTIONDESCRIPTION_LOOKUPBYKEYACTIONDESCRIPTION']._serialized_start=4653
  _globals['_ACTIONDESCRIPTION_LOOKUPBYKEYACTIONDESCRIPTION']._serialized_end=4749
  _globals['_PARTICIPANTRANDOMNESSLOOKUP']._serialized_start=4766
  _globals['_PARTICIPANTRANDOMNESSLOOKUP']._serialized_end=4861
  _globals['_INFORMEEMESSAGE']._serialized_start=4863
  _globals['_INFORMEEMESSAGE']._serialized_end=4981
  _globals['_FULLINFORMEETREE']._serialized_start=4983
  _globals['_FULLINFORMEETREE']._serialized_end=5078
  _globals['_ROOTHASHMESSAGE']._serialized_start=5081
  _globals['_ROOTHASHMESSAGE']._serialized_end=5258
# @@protoc_insertion_point(module_scope)