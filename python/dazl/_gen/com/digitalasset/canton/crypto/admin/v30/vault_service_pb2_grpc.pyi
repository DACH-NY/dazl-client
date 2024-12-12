# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file

import builtins as _builtins, typing as _typing

import grpc as _grpc
from grpc import aio as _grpc_aio

from .vault_service_pb2 import DeleteKeyPairRequest, DeleteKeyPairResponse, ExportKeyPairRequest, ExportKeyPairResponse, GenerateEncryptionKeyRequest, GenerateEncryptionKeyResponse, GenerateSigningKeyRequest, GenerateSigningKeyResponse, GetWrapperKeyIdRequest, GetWrapperKeyIdResponse, ImportKeyPairRequest, ImportKeyPairResponse, ImportPublicKeyRequest, ImportPublicKeyResponse, ListMyKeysRequest, ListMyKeysResponse, ListPublicKeysRequest, ListPublicKeysResponse, RegisterKmsEncryptionKeyRequest, RegisterKmsEncryptionKeyResponse, RegisterKmsSigningKeyRequest, RegisterKmsSigningKeyResponse, RotateWrapperKeyRequest, RotateWrapperKeyResponse

__all__ = [
    "VaultServiceStub",
]


# noinspection PyPep8Naming,DuplicatedCode
class VaultServiceStub:
    @classmethod  # type: ignore
    @_typing.overload
    def __new__(cls, channel: _grpc.Channel) -> _VaultServiceBlockingStub: ...  # type: ignore
    @classmethod  # type: ignore
    @_typing.overload
    def __new__(cls, channel: _grpc_aio.Channel) -> _VaultServiceAsyncStub: ...  # type: ignore
    def ListMyKeys(self, __1: ListMyKeysRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListMyKeysResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ListMyKeysResponse]]: ...
    def GenerateSigningKey(self, __1: GenerateSigningKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[GenerateSigningKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, GenerateSigningKeyResponse]]: ...
    def GenerateEncryptionKey(self, __1: GenerateEncryptionKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[GenerateEncryptionKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, GenerateEncryptionKeyResponse]]: ...
    def RegisterKmsEncryptionKey(self, __1: RegisterKmsEncryptionKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[RegisterKmsEncryptionKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, RegisterKmsEncryptionKeyResponse]]: ...
    def RegisterKmsSigningKey(self, __1: RegisterKmsSigningKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[RegisterKmsSigningKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, RegisterKmsSigningKeyResponse]]: ...
    def ImportPublicKey(self, __1: ImportPublicKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ImportPublicKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ImportPublicKeyResponse]]: ...
    def ListPublicKeys(self, __1: ListPublicKeysRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListPublicKeysResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ListPublicKeysResponse]]: ...
    def RotateWrapperKey(self, __1: RotateWrapperKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[RotateWrapperKeyResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, RotateWrapperKeyResponse]]: ...
    def GetWrapperKeyId(self, __1: GetWrapperKeyIdRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[GetWrapperKeyIdResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, GetWrapperKeyIdResponse]]: ...
    def ExportKeyPair(self, __1: ExportKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ExportKeyPairResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ExportKeyPairResponse]]: ...
    def ImportKeyPair(self, __1: ImportKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ImportKeyPairResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ImportKeyPairResponse]]: ...
    def DeleteKeyPair(self, __1: DeleteKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[DeleteKeyPairResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, DeleteKeyPairResponse]]: ...

# noinspection PyPep8Naming,DuplicatedCode
class _VaultServiceBlockingStub(VaultServiceStub):
    def ListMyKeys(self, __1: ListMyKeysRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListMyKeysResponse: ...
    def GenerateSigningKey(self, __1: GenerateSigningKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> GenerateSigningKeyResponse: ...
    def GenerateEncryptionKey(self, __1: GenerateEncryptionKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> GenerateEncryptionKeyResponse: ...
    def RegisterKmsEncryptionKey(self, __1: RegisterKmsEncryptionKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> RegisterKmsEncryptionKeyResponse: ...
    def RegisterKmsSigningKey(self, __1: RegisterKmsSigningKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> RegisterKmsSigningKeyResponse: ...
    def ImportPublicKey(self, __1: ImportPublicKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ImportPublicKeyResponse: ...
    def ListPublicKeys(self, __1: ListPublicKeysRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListPublicKeysResponse: ...
    def RotateWrapperKey(self, __1: RotateWrapperKeyRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> RotateWrapperKeyResponse: ...
    def GetWrapperKeyId(self, __1: GetWrapperKeyIdRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> GetWrapperKeyIdResponse: ...
    def ExportKeyPair(self, __1: ExportKeyPairRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ExportKeyPairResponse: ...
    def ImportKeyPair(self, __1: ImportKeyPairRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ImportKeyPairResponse: ...
    def DeleteKeyPair(self, __1: DeleteKeyPairRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> DeleteKeyPairResponse: ...

# noinspection PyPep8Naming,DuplicatedCode
class _VaultServiceAsyncStub(VaultServiceStub):
    def ListMyKeys(self, __1: ListMyKeysRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListMyKeysResponse]: ...  # type: ignore
    def GenerateSigningKey(self, __1: GenerateSigningKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, GenerateSigningKeyResponse]: ...  # type: ignore
    def GenerateEncryptionKey(self, __1: GenerateEncryptionKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, GenerateEncryptionKeyResponse]: ...  # type: ignore
    def RegisterKmsEncryptionKey(self, __1: RegisterKmsEncryptionKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, RegisterKmsEncryptionKeyResponse]: ...  # type: ignore
    def RegisterKmsSigningKey(self, __1: RegisterKmsSigningKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, RegisterKmsSigningKeyResponse]: ...  # type: ignore
    def ImportPublicKey(self, __1: ImportPublicKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ImportPublicKeyResponse]: ...  # type: ignore
    def ListPublicKeys(self, __1: ListPublicKeysRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListPublicKeysResponse]: ...  # type: ignore
    def RotateWrapperKey(self, __1: RotateWrapperKeyRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, RotateWrapperKeyResponse]: ...  # type: ignore
    def GetWrapperKeyId(self, __1: GetWrapperKeyIdRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, GetWrapperKeyIdResponse]: ...  # type: ignore
    def ExportKeyPair(self, __1: ExportKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ExportKeyPairResponse]: ...  # type: ignore
    def ImportKeyPair(self, __1: ImportKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ImportKeyPairResponse]: ...  # type: ignore
    def DeleteKeyPair(self, __1: DeleteKeyPairRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, DeleteKeyPairResponse]: ...  # type: ignore