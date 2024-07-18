# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/admin/user_management_service.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import object_meta_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_object__meta__pb2
from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n:com/daml/ledger/api/v1/admin/user_management_service.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\x1a.com/daml/ledger/api/v1/admin/object_meta.proto\x1a google/protobuf/field_mask.proto\"\xda\x01\n\x04User\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12#\n\rprimary_party\x18\x02 \x01(\tR\x0cprimaryParty\x12%\n\x0eis_deactivated\x18\x03 \x01(\x08R\risDeactivated\x12\x44\n\x08metadata\x18\x04 \x01(\x0b\x32(.com.daml.ledger.api.v1.admin.ObjectMetaR\x08metadata\x12\x30\n\x14identity_provider_id\x18\x05 \x01(\tR\x12identityProviderId\"\xfa\x03\n\x05Right\x12\x63\n\x11participant_admin\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.Right.ParticipantAdminH\x00R\x10participantAdmin\x12L\n\ncan_act_as\x18\x02 \x01(\x0b\x32,.com.daml.ledger.api.v1.admin.Right.CanActAsH\x00R\x08\x63\x61nActAs\x12O\n\x0b\x63\x61n_read_as\x18\x03 \x01(\x0b\x32-.com.daml.ledger.api.v1.admin.Right.CanReadAsH\x00R\tcanReadAs\x12s\n\x17identity_provider_admin\x18\x04 \x01(\x0b\x32\x39.com.daml.ledger.api.v1.admin.Right.IdentityProviderAdminH\x00R\x15identityProviderAdmin\x1a\x12\n\x10ParticipantAdmin\x1a \n\x08\x43\x61nActAs\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x1a!\n\tCanReadAs\x12\x14\n\x05party\x18\x01 \x01(\tR\x05party\x1a\x17\n\x15IdentityProviderAdminB\x06\n\x04kind\"\x88\x01\n\x11\x43reateUserRequest\x12\x36\n\x04user\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x04user\x12;\n\x06rights\x18\x02 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x06rights\"L\n\x12\x43reateUserResponse\x12\x36\n\x04user\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x04user\"[\n\x0eGetUserRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12\x30\n\x14identity_provider_id\x18\x02 \x01(\tR\x12identityProviderId\"I\n\x0fGetUserResponse\x12\x36\n\x04user\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x04user\"\x88\x01\n\x11UpdateUserRequest\x12\x36\n\x04user\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x04user\x12;\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nupdateMask\"L\n\x12UpdateUserResponse\x12\x36\n\x04user\x18\x01 \x01(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x04user\"^\n\x11\x44\x65leteUserRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12\x30\n\x14identity_provider_id\x18\x02 \x01(\tR\x12identityProviderId\"\x14\n\x12\x44\x65leteUserResponse\"\x80\x01\n\x10ListUsersRequest\x12\x1d\n\npage_token\x18\x02 \x01(\tR\tpageToken\x12\x1b\n\tpage_size\x18\x03 \x01(\x05R\x08pageSize\x12\x30\n\x14identity_provider_id\x18\x04 \x01(\tR\x12identityProviderId\"u\n\x11ListUsersResponse\x12\x38\n\x05users\x18\x01 \x03(\x0b\x32\".com.daml.ledger.api.v1.admin.UserR\x05users\x12&\n\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xa0\x01\n\x16GrantUserRightsRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12;\n\x06rights\x18\x02 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x06rights\x12\x30\n\x14identity_provider_id\x18\x03 \x01(\tR\x12identityProviderId\"p\n\x17GrantUserRightsResponse\x12U\n\x14newly_granted_rights\x18\x01 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x12newlyGrantedRights\"\xa1\x01\n\x17RevokeUserRightsRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12;\n\x06rights\x18\x02 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x06rights\x12\x30\n\x14identity_provider_id\x18\x03 \x01(\tR\x12identityProviderId\"q\n\x18RevokeUserRightsResponse\x12U\n\x14newly_revoked_rights\x18\x01 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x12newlyRevokedRights\"b\n\x15ListUserRightsRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12\x30\n\x14identity_provider_id\x18\x02 \x01(\tR\x12identityProviderId\"U\n\x16ListUserRightsResponse\x12;\n\x06rights\x18\x01 \x03(\x0b\x32#.com.daml.ledger.api.v1.admin.RightR\x06rights\"\xba\x01\n!UpdateUserIdentityProviderRequest\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12=\n\x1bsource_identity_provider_id\x18\x02 \x01(\tR\x18sourceIdentityProviderId\x12=\n\x1btarget_identity_provider_id\x18\x03 \x01(\tR\x18targetIdentityProviderId\"$\n\"UpdateUserIdentityProviderResponse2\xe5\x08\n\x15UserManagementService\x12o\n\nCreateUser\x12/.com.daml.ledger.api.v1.admin.CreateUserRequest\x1a\x30.com.daml.ledger.api.v1.admin.CreateUserResponse\x12\x66\n\x07GetUser\x12,.com.daml.ledger.api.v1.admin.GetUserRequest\x1a-.com.daml.ledger.api.v1.admin.GetUserResponse\x12o\n\nUpdateUser\x12/.com.daml.ledger.api.v1.admin.UpdateUserRequest\x1a\x30.com.daml.ledger.api.v1.admin.UpdateUserResponse\x12o\n\nDeleteUser\x12/.com.daml.ledger.api.v1.admin.DeleteUserRequest\x1a\x30.com.daml.ledger.api.v1.admin.DeleteUserResponse\x12l\n\tListUsers\x12..com.daml.ledger.api.v1.admin.ListUsersRequest\x1a/.com.daml.ledger.api.v1.admin.ListUsersResponse\x12~\n\x0fGrantUserRights\x12\x34.com.daml.ledger.api.v1.admin.GrantUserRightsRequest\x1a\x35.com.daml.ledger.api.v1.admin.GrantUserRightsResponse\x12\x81\x01\n\x10RevokeUserRights\x12\x35.com.daml.ledger.api.v1.admin.RevokeUserRightsRequest\x1a\x36.com.daml.ledger.api.v1.admin.RevokeUserRightsResponse\x12{\n\x0eListUserRights\x12\x33.com.daml.ledger.api.v1.admin.ListUserRightsRequest\x1a\x34.com.daml.ledger.api.v1.admin.ListUserRightsResponse\x12\xa1\x01\n\x1cUpdateUserIdentityProviderId\x12?.com.daml.ledger.api.v1.admin.UpdateUserIdentityProviderRequest\x1a@.com.daml.ledger.api.v1.admin.UpdateUserIdentityProviderResponseB\xab\x01\n\x1c\x63om.daml.ledger.api.v1.adminB\x1fUserManagementServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.admin.user_management_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\034com.daml.ledger.api.v1.adminB\037UserManagementServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\252\002\034Com.Daml.Ledger.Api.V1.Admin'
  _globals['_USER']._serialized_start=175
  _globals['_USER']._serialized_end=393
  _globals['_RIGHT']._serialized_start=396
  _globals['_RIGHT']._serialized_end=902
  _globals['_RIGHT_PARTICIPANTADMIN']._serialized_start=782
  _globals['_RIGHT_PARTICIPANTADMIN']._serialized_end=800
  _globals['_RIGHT_CANACTAS']._serialized_start=802
  _globals['_RIGHT_CANACTAS']._serialized_end=834
  _globals['_RIGHT_CANREADAS']._serialized_start=836
  _globals['_RIGHT_CANREADAS']._serialized_end=869
  _globals['_RIGHT_IDENTITYPROVIDERADMIN']._serialized_start=871
  _globals['_RIGHT_IDENTITYPROVIDERADMIN']._serialized_end=894
  _globals['_CREATEUSERREQUEST']._serialized_start=905
  _globals['_CREATEUSERREQUEST']._serialized_end=1041
  _globals['_CREATEUSERRESPONSE']._serialized_start=1043
  _globals['_CREATEUSERRESPONSE']._serialized_end=1119
  _globals['_GETUSERREQUEST']._serialized_start=1121
  _globals['_GETUSERREQUEST']._serialized_end=1212
  _globals['_GETUSERRESPONSE']._serialized_start=1214
  _globals['_GETUSERRESPONSE']._serialized_end=1287
  _globals['_UPDATEUSERREQUEST']._serialized_start=1290
  _globals['_UPDATEUSERREQUEST']._serialized_end=1426
  _globals['_UPDATEUSERRESPONSE']._serialized_start=1428
  _globals['_UPDATEUSERRESPONSE']._serialized_end=1504
  _globals['_DELETEUSERREQUEST']._serialized_start=1506
  _globals['_DELETEUSERREQUEST']._serialized_end=1600
  _globals['_DELETEUSERRESPONSE']._serialized_start=1602
  _globals['_DELETEUSERRESPONSE']._serialized_end=1622
  _globals['_LISTUSERSREQUEST']._serialized_start=1625
  _globals['_LISTUSERSREQUEST']._serialized_end=1753
  _globals['_LISTUSERSRESPONSE']._serialized_start=1755
  _globals['_LISTUSERSRESPONSE']._serialized_end=1872
  _globals['_GRANTUSERRIGHTSREQUEST']._serialized_start=1875
  _globals['_GRANTUSERRIGHTSREQUEST']._serialized_end=2035
  _globals['_GRANTUSERRIGHTSRESPONSE']._serialized_start=2037
  _globals['_GRANTUSERRIGHTSRESPONSE']._serialized_end=2149
  _globals['_REVOKEUSERRIGHTSREQUEST']._serialized_start=2152
  _globals['_REVOKEUSERRIGHTSREQUEST']._serialized_end=2313
  _globals['_REVOKEUSERRIGHTSRESPONSE']._serialized_start=2315
  _globals['_REVOKEUSERRIGHTSRESPONSE']._serialized_end=2428
  _globals['_LISTUSERRIGHTSREQUEST']._serialized_start=2430
  _globals['_LISTUSERRIGHTSREQUEST']._serialized_end=2528
  _globals['_LISTUSERRIGHTSRESPONSE']._serialized_start=2530
  _globals['_LISTUSERRIGHTSRESPONSE']._serialized_end=2615
  _globals['_UPDATEUSERIDENTITYPROVIDERREQUEST']._serialized_start=2618
  _globals['_UPDATEUSERIDENTITYPROVIDERREQUEST']._serialized_end=2804
  _globals['_UPDATEUSERIDENTITYPROVIDERRESPONSE']._serialized_start=2806
  _globals['_UPDATEUSERIDENTITYPROVIDERRESPONSE']._serialized_end=2842
  _globals['_USERMANAGEMENTSERVICE']._serialized_start=2845
  _globals['_USERMANAGEMENTSERVICE']._serialized_end=3970
# @@protoc_insertion_point(module_scope)
