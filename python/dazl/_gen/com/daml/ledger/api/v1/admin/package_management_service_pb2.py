# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/admin/package_management_service.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='com/daml/ledger/api/v1/admin/package_management_service.proto',
  package='com.daml.ledger.api.v1.admin',
  syntax='proto3',
  serialized_options=b'\n\034com.daml.ledger.api.v1.adminB\"PackageManagementServiceOuterClass\252\002\034Com.Daml.Ledger.Api.V1.Admin',
  serialized_pb=b'\n=com/daml/ledger/api/v1/admin/package_management_service.proto\x12\x1c\x63om.daml.ledger.api.v1.admin\x1a\x1fgoogle/protobuf/timestamp.proto\"\x1a\n\x18ListKnownPackagesRequest\"b\n\x19ListKnownPackagesResponse\x12\x45\n\x0fpackage_details\x18\x01 \x03(\x0b\x32,.com.daml.ledger.api.v1.admin.PackageDetails\"\x87\x01\n\x0ePackageDetails\x12\x12\n\npackage_id\x18\x01 \x01(\t\x12\x14\n\x0cpackage_size\x18\x02 \x01(\x04\x12/\n\x0bknown_since\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1a\n\x12source_description\x18\x04 \x01(\t\"?\n\x14UploadDarFileRequest\x12\x10\n\x08\x64\x61r_file\x18\x01 \x01(\x0c\x12\x15\n\rsubmission_id\x18\x02 \x01(\t\"\x17\n\x15UploadDarFileResponse2\x9b\x02\n\x18PackageManagementService\x12\x84\x01\n\x11ListKnownPackages\x12\x36.com.daml.ledger.api.v1.admin.ListKnownPackagesRequest\x1a\x37.com.daml.ledger.api.v1.admin.ListKnownPackagesResponse\x12x\n\rUploadDarFile\x12\x32.com.daml.ledger.api.v1.admin.UploadDarFileRequest\x1a\x33.com.daml.ledger.api.v1.admin.UploadDarFileResponseBa\n\x1c\x63om.daml.ledger.api.v1.adminB\"PackageManagementServiceOuterClass\xaa\x02\x1c\x43om.Daml.Ledger.Api.V1.Adminb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_LISTKNOWNPACKAGESREQUEST = _descriptor.Descriptor(
  name='ListKnownPackagesRequest',
  full_name='com.daml.ledger.api.v1.admin.ListKnownPackagesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=128,
  serialized_end=154,
)


_LISTKNOWNPACKAGESRESPONSE = _descriptor.Descriptor(
  name='ListKnownPackagesResponse',
  full_name='com.daml.ledger.api.v1.admin.ListKnownPackagesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='package_details', full_name='com.daml.ledger.api.v1.admin.ListKnownPackagesResponse.package_details', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=156,
  serialized_end=254,
)


_PACKAGEDETAILS = _descriptor.Descriptor(
  name='PackageDetails',
  full_name='com.daml.ledger.api.v1.admin.PackageDetails',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='package_id', full_name='com.daml.ledger.api.v1.admin.PackageDetails.package_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='package_size', full_name='com.daml.ledger.api.v1.admin.PackageDetails.package_size', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='known_since', full_name='com.daml.ledger.api.v1.admin.PackageDetails.known_since', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source_description', full_name='com.daml.ledger.api.v1.admin.PackageDetails.source_description', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=257,
  serialized_end=392,
)


_UPLOADDARFILEREQUEST = _descriptor.Descriptor(
  name='UploadDarFileRequest',
  full_name='com.daml.ledger.api.v1.admin.UploadDarFileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dar_file', full_name='com.daml.ledger.api.v1.admin.UploadDarFileRequest.dar_file', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='submission_id', full_name='com.daml.ledger.api.v1.admin.UploadDarFileRequest.submission_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=394,
  serialized_end=457,
)


_UPLOADDARFILERESPONSE = _descriptor.Descriptor(
  name='UploadDarFileResponse',
  full_name='com.daml.ledger.api.v1.admin.UploadDarFileResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=459,
  serialized_end=482,
)

_LISTKNOWNPACKAGESRESPONSE.fields_by_name['package_details'].message_type = _PACKAGEDETAILS
_PACKAGEDETAILS.fields_by_name['known_since'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['ListKnownPackagesRequest'] = _LISTKNOWNPACKAGESREQUEST
DESCRIPTOR.message_types_by_name['ListKnownPackagesResponse'] = _LISTKNOWNPACKAGESRESPONSE
DESCRIPTOR.message_types_by_name['PackageDetails'] = _PACKAGEDETAILS
DESCRIPTOR.message_types_by_name['UploadDarFileRequest'] = _UPLOADDARFILEREQUEST
DESCRIPTOR.message_types_by_name['UploadDarFileResponse'] = _UPLOADDARFILERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListKnownPackagesRequest = _reflection.GeneratedProtocolMessageType('ListKnownPackagesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTKNOWNPACKAGESREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.admin.package_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.ListKnownPackagesRequest)
  })
_sym_db.RegisterMessage(ListKnownPackagesRequest)

ListKnownPackagesResponse = _reflection.GeneratedProtocolMessageType('ListKnownPackagesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTKNOWNPACKAGESRESPONSE,
  '__module__' : 'com.daml.ledger.api.v1.admin.package_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.ListKnownPackagesResponse)
  })
_sym_db.RegisterMessage(ListKnownPackagesResponse)

PackageDetails = _reflection.GeneratedProtocolMessageType('PackageDetails', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGEDETAILS,
  '__module__' : 'com.daml.ledger.api.v1.admin.package_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.PackageDetails)
  })
_sym_db.RegisterMessage(PackageDetails)

UploadDarFileRequest = _reflection.GeneratedProtocolMessageType('UploadDarFileRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADDARFILEREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.admin.package_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.UploadDarFileRequest)
  })
_sym_db.RegisterMessage(UploadDarFileRequest)

UploadDarFileResponse = _reflection.GeneratedProtocolMessageType('UploadDarFileResponse', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADDARFILERESPONSE,
  '__module__' : 'com.daml.ledger.api.v1.admin.package_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.admin.UploadDarFileResponse)
  })
_sym_db.RegisterMessage(UploadDarFileResponse)


DESCRIPTOR._options = None

_PACKAGEMANAGEMENTSERVICE = _descriptor.ServiceDescriptor(
  name='PackageManagementService',
  full_name='com.daml.ledger.api.v1.admin.PackageManagementService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=485,
  serialized_end=768,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListKnownPackages',
    full_name='com.daml.ledger.api.v1.admin.PackageManagementService.ListKnownPackages',
    index=0,
    containing_service=None,
    input_type=_LISTKNOWNPACKAGESREQUEST,
    output_type=_LISTKNOWNPACKAGESRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UploadDarFile',
    full_name='com.daml.ledger.api.v1.admin.PackageManagementService.UploadDarFile',
    index=1,
    containing_service=None,
    input_type=_UPLOADDARFILEREQUEST,
    output_type=_UPLOADDARFILERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PACKAGEMANAGEMENTSERVICE)

DESCRIPTOR.services_by_name['PackageManagementService'] = _PACKAGEMANAGEMENTSERVICE

# @@protoc_insertion_point(module_scope)
