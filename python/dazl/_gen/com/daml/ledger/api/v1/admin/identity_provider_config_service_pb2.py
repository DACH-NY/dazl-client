# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: com/daml/ledger/api/v1/admin/identity_provider_config_service.proto
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
    'com/daml/ledger/api/v1/admin/identity_provider_config_service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nCcom/daml/ledger/api/v1/admin/identity_provider_config_service.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\x1a google/protobuf/field_mask.proto\"\xc0\x01\n\x16IdentityProviderConfig\x12\x30\n\x14identity_provider_id\x18\x01 \x01(\tR\x12identityProviderId\x12%\n\x0eis_deactivated\x18\x02 \x01(\x08R\risDeactivated\x12\x16\n\x06issuer\x18\x03 \x01(\tR\x06issuer\x12\x19\n\x08jwks_url\x18\x04 \x01(\tR\x07jwksUrl\x12\x1a\n\x08\x61udience\x18\x05 \x01(\tR\x08\x61udience\"\x95\x01\n#CreateIdentityProviderConfigRequest\x12n\n\x18identity_provider_config\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x16identityProviderConfig\"\x96\x01\n$CreateIdentityProviderConfigResponse\x12n\n\x18identity_provider_config\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x16identityProviderConfig\"T\n GetIdentityProviderConfigRequest\x12\x30\n\x14identity_provider_id\x18\x01 \x01(\tR\x12identityProviderId\"\x93\x01\n!GetIdentityProviderConfigResponse\x12n\n\x18identity_provider_config\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x16identityProviderConfig\"$\n\"ListIdentityProviderConfigsRequest\"\x97\x01\n#ListIdentityProviderConfigsResponse\x12p\n\x19identity_provider_configs\x18\x01 \x03(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x17identityProviderConfigs\"\xd2\x01\n#UpdateIdentityProviderConfigRequest\x12n\n\x18identity_provider_config\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x16identityProviderConfig\x12;\n\x0bupdate_mask\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\nupdateMask\"\x96\x01\n$UpdateIdentityProviderConfigResponse\x12n\n\x18identity_provider_config\x18\x01 \x01(\x0b\x32\x34.com.daml.ledger.api.v1.admin.IdentityProviderConfigR\x16identityProviderConfig\"W\n#DeleteIdentityProviderConfigRequest\x12\x30\n\x14identity_provider_id\x18\x01 \x01(\tR\x12identityProviderId\"&\n$DeleteIdentityProviderConfigResponse2\xdb\x06\n\x1dIdentityProviderConfigService\x12\xa5\x01\n\x1c\x43reateIdentityProviderConfig\x12\x41.com.daml.ledger.api.v1.admin.CreateIdentityProviderConfigRequest\x1a\x42.com.daml.ledger.api.v1.admin.CreateIdentityProviderConfigResponse\x12\x9c\x01\n\x19GetIdentityProviderConfig\x12>.com.daml.ledger.api.v1.admin.GetIdentityProviderConfigRequest\x1a?.com.daml.ledger.api.v1.admin.GetIdentityProviderConfigResponse\x12\xa5\x01\n\x1cUpdateIdentityProviderConfig\x12\x41.com.daml.ledger.api.v1.admin.UpdateIdentityProviderConfigRequest\x1a\x42.com.daml.ledger.api.v1.admin.UpdateIdentityProviderConfigResponse\x12\xa2\x01\n\x1bListIdentityProviderConfigs\x12@.com.daml.ledger.api.v1.admin.ListIdentityProviderConfigsRequest\x1a\x41.com.daml.ledger.api.v1.admin.ListIdentityProviderConfigsResponse\x12\xa5\x01\n\x1c\x44\x65leteIdentityProviderConfig\x12\x41.com.daml.ledger.api.v1.admin.DeleteIdentityProviderConfigRequest\x1a\x42.com.daml.ledger.api.v1.admin.DeleteIdentityProviderConfigResponseB\xb3\x01\n\x1c\x63om.daml.ledger.api.v1.adminB\'IdentityProviderConfigServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'com.daml.ledger.api.v1.admin.identity_provider_config_service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\034com.daml.ledger.api.v1.adminB\'IdentityProviderConfigServiceOuterClassZKgithub.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v1/admin\252\002\034Com.Daml.Ledger.Api.V1.Admin'
  _globals['_IDENTITYPROVIDERCONFIG']._serialized_start=136
  _globals['_IDENTITYPROVIDERCONFIG']._serialized_end=328
  _globals['_CREATEIDENTITYPROVIDERCONFIGREQUEST']._serialized_start=331
  _globals['_CREATEIDENTITYPROVIDERCONFIGREQUEST']._serialized_end=480
  _globals['_CREATEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_start=483
  _globals['_CREATEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_end=633
  _globals['_GETIDENTITYPROVIDERCONFIGREQUEST']._serialized_start=635
  _globals['_GETIDENTITYPROVIDERCONFIGREQUEST']._serialized_end=719
  _globals['_GETIDENTITYPROVIDERCONFIGRESPONSE']._serialized_start=722
  _globals['_GETIDENTITYPROVIDERCONFIGRESPONSE']._serialized_end=869
  _globals['_LISTIDENTITYPROVIDERCONFIGSREQUEST']._serialized_start=871
  _globals['_LISTIDENTITYPROVIDERCONFIGSREQUEST']._serialized_end=907
  _globals['_LISTIDENTITYPROVIDERCONFIGSRESPONSE']._serialized_start=910
  _globals['_LISTIDENTITYPROVIDERCONFIGSRESPONSE']._serialized_end=1061
  _globals['_UPDATEIDENTITYPROVIDERCONFIGREQUEST']._serialized_start=1064
  _globals['_UPDATEIDENTITYPROVIDERCONFIGREQUEST']._serialized_end=1274
  _globals['_UPDATEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_start=1277
  _globals['_UPDATEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_end=1427
  _globals['_DELETEIDENTITYPROVIDERCONFIGREQUEST']._serialized_start=1429
  _globals['_DELETEIDENTITYPROVIDERCONFIGREQUEST']._serialized_end=1516
  _globals['_DELETEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_start=1518
  _globals['_DELETEIDENTITYPROVIDERCONFIGRESPONSE']._serialized_end=1556
  _globals['_IDENTITYPROVIDERCONFIGSERVICE']._serialized_start=1559
  _globals['_IDENTITYPROVIDERCONFIGSERVICE']._serialized_end=2418
# @@protoc_insertion_point(module_scope)
