# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import vault_service_pb2 as com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2

GRPC_GENERATED_VERSION = '1.67.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in com/digitalasset/canton/crypto/admin/v30/vault_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class VaultServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListMyKeys = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/ListMyKeys',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysResponse.FromString,
                _registered_method=True)
        self.GenerateSigningKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/GenerateSigningKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyResponse.FromString,
                _registered_method=True)
        self.GenerateEncryptionKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/GenerateEncryptionKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyResponse.FromString,
                _registered_method=True)
        self.RegisterKmsEncryptionKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/RegisterKmsEncryptionKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyResponse.FromString,
                _registered_method=True)
        self.RegisterKmsSigningKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/RegisterKmsSigningKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyResponse.FromString,
                _registered_method=True)
        self.ImportPublicKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/ImportPublicKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyResponse.FromString,
                _registered_method=True)
        self.ListPublicKeys = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/ListPublicKeys',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysResponse.FromString,
                _registered_method=True)
        self.RotateWrapperKey = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/RotateWrapperKey',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyResponse.FromString,
                _registered_method=True)
        self.GetWrapperKeyId = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/GetWrapperKeyId',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdResponse.FromString,
                _registered_method=True)
        self.ExportKeyPair = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/ExportKeyPair',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairResponse.FromString,
                _registered_method=True)
        self.ImportKeyPair = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/ImportKeyPair',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairResponse.FromString,
                _registered_method=True)
        self.DeleteKeyPair = channel.unary_unary(
                '/com.digitalasset.canton.crypto.admin.v30.VaultService/DeleteKeyPair',
                request_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairRequest.SerializeToString,
                response_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairResponse.FromString,
                _registered_method=True)


class VaultServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListMyKeys(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateSigningKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateEncryptionKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RegisterKmsEncryptionKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RegisterKmsSigningKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ImportPublicKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListPublicKeys(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RotateWrapperKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetWrapperKeyId(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExportKeyPair(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ImportKeyPair(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteKeyPair(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_VaultServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListMyKeys': grpc.unary_unary_rpc_method_handler(
                    servicer.ListMyKeys,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysResponse.SerializeToString,
            ),
            'GenerateSigningKey': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateSigningKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyResponse.SerializeToString,
            ),
            'GenerateEncryptionKey': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateEncryptionKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyResponse.SerializeToString,
            ),
            'RegisterKmsEncryptionKey': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterKmsEncryptionKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyResponse.SerializeToString,
            ),
            'RegisterKmsSigningKey': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterKmsSigningKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyResponse.SerializeToString,
            ),
            'ImportPublicKey': grpc.unary_unary_rpc_method_handler(
                    servicer.ImportPublicKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyResponse.SerializeToString,
            ),
            'ListPublicKeys': grpc.unary_unary_rpc_method_handler(
                    servicer.ListPublicKeys,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysResponse.SerializeToString,
            ),
            'RotateWrapperKey': grpc.unary_unary_rpc_method_handler(
                    servicer.RotateWrapperKey,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyResponse.SerializeToString,
            ),
            'GetWrapperKeyId': grpc.unary_unary_rpc_method_handler(
                    servicer.GetWrapperKeyId,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdResponse.SerializeToString,
            ),
            'ExportKeyPair': grpc.unary_unary_rpc_method_handler(
                    servicer.ExportKeyPair,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairResponse.SerializeToString,
            ),
            'ImportKeyPair': grpc.unary_unary_rpc_method_handler(
                    servicer.ImportKeyPair,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairResponse.SerializeToString,
            ),
            'DeleteKeyPair': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteKeyPair,
                    request_deserializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairRequest.FromString,
                    response_serializer=com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'com.digitalasset.canton.crypto.admin.v30.VaultService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('com.digitalasset.canton.crypto.admin.v30.VaultService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class VaultService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListMyKeys(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/ListMyKeys',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListMyKeysResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GenerateSigningKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/GenerateSigningKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateSigningKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GenerateEncryptionKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/GenerateEncryptionKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GenerateEncryptionKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def RegisterKmsEncryptionKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/RegisterKmsEncryptionKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsEncryptionKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def RegisterKmsSigningKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/RegisterKmsSigningKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RegisterKmsSigningKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ImportPublicKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/ImportPublicKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportPublicKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ListPublicKeys(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/ListPublicKeys',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ListPublicKeysResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def RotateWrapperKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/RotateWrapperKey',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.RotateWrapperKeyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetWrapperKeyId(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/GetWrapperKeyId',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.GetWrapperKeyIdResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ExportKeyPair(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/ExportKeyPair',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ExportKeyPairResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ImportKeyPair(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/ImportKeyPair',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.ImportKeyPairResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DeleteKeyPair(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/com.digitalasset.canton.crypto.admin.v30.VaultService/DeleteKeyPair',
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairRequest.SerializeToString,
            com_dot_digitalasset_dot_canton_dot_crypto_dot_admin_dot_v30_dot_vault__service__pb2.DeleteKeyPairResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
