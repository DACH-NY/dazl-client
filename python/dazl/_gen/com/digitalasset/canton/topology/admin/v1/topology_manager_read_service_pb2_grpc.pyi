# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file

import builtins as _builtins, typing as _typing

import grpc as _grpc
from grpc import aio as _grpc_aio

from .topology_manager_read_service_pb2 import ListAllRequest, ListAllResponse, ListAuthorityOfRequest, ListAuthorityOfResult, ListAvailableStoresRequest, ListAvailableStoresResult, ListDomainParametersStateRequest, ListDomainParametersStateResult, ListDomainTrustCertificateRequest, ListDomainTrustCertificateResult, ListIdentifierDelegationRequest, ListIdentifierDelegationResult, ListMediatorDomainStateRequest, ListMediatorDomainStateResult, ListNamespaceDelegationRequest, ListNamespaceDelegationResult, ListOwnerToKeyMappingRequest, ListOwnerToKeyMappingResult, ListParticipantDomainPermissionRequest, ListParticipantDomainPermissionResult, ListPartyHostingLimitsRequest, ListPartyHostingLimitsResult, ListPartyToParticipantRequest, ListPartyToParticipantResult, ListPurgeTopologyTransactionXRequest, ListPurgeTopologyTransactionXResult, ListSequencerDomainStateRequest, ListSequencerDomainStateResult, ListTrafficStateRequest, ListTrafficStateResult, ListUnionspaceDefinitionRequest, ListUnionspaceDefinitionResult, ListVettedPackagesRequest, ListVettedPackagesResult

__all__ = [
    "TopologyManagerReadServiceXStub",
]


# noinspection PyPep8Naming,DuplicatedCode
class TopologyManagerReadServiceXStub:
    @classmethod  # type: ignore
    @_typing.overload
    def __new__(cls, channel: _grpc.Channel) -> _TopologyManagerReadServiceXBlockingStub: ...  # type: ignore
    @classmethod  # type: ignore
    @_typing.overload
    def __new__(cls, channel: _grpc_aio.Channel) -> _TopologyManagerReadServiceXAsyncStub: ...  # type: ignore
    def ListNamespaceDelegation(self, __1: ListNamespaceDelegationRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListNamespaceDelegationResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListNamespaceDelegationResult]]: ...
    def ListUnionspaceDefinition(self, __1: ListUnionspaceDefinitionRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListUnionspaceDefinitionResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListUnionspaceDefinitionResult]]: ...
    def ListIdentifierDelegation(self, __1: ListIdentifierDelegationRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListIdentifierDelegationResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListIdentifierDelegationResult]]: ...
    def ListOwnerToKeyMapping(self, __1: ListOwnerToKeyMappingRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListOwnerToKeyMappingResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListOwnerToKeyMappingResult]]: ...
    def ListDomainTrustCertificate(self, __1: ListDomainTrustCertificateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListDomainTrustCertificateResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListDomainTrustCertificateResult]]: ...
    def ListParticipantDomainPermission(self, __1: ListParticipantDomainPermissionRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListParticipantDomainPermissionResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListParticipantDomainPermissionResult]]: ...
    def ListPartyHostingLimits(self, __1: ListPartyHostingLimitsRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListPartyHostingLimitsResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListPartyHostingLimitsResult]]: ...
    def ListVettedPackages(self, __1: ListVettedPackagesRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListVettedPackagesResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListVettedPackagesResult]]: ...
    def ListPartyToParticipant(self, __1: ListPartyToParticipantRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListPartyToParticipantResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListPartyToParticipantResult]]: ...
    def ListAuthorityOf(self, __1: ListAuthorityOfRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListAuthorityOfResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListAuthorityOfResult]]: ...
    def ListDomainParametersState(self, __1: ListDomainParametersStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListDomainParametersStateResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListDomainParametersStateResult]]: ...
    def ListMediatorDomainState(self, __1: ListMediatorDomainStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListMediatorDomainStateResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListMediatorDomainStateResult]]: ...
    def ListSequencerDomainState(self, __1: ListSequencerDomainStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListSequencerDomainStateResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListSequencerDomainStateResult]]: ...
    def ListPurgeTopologyTransactionX(self, __1: ListPurgeTopologyTransactionXRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListPurgeTopologyTransactionXResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListPurgeTopologyTransactionXResult]]: ...
    def ListAvailableStores(self, __1: ListAvailableStoresRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListAvailableStoresResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListAvailableStoresResult]]: ...
    def ListAll(self, __1: ListAllRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListAllResponse, _grpc_aio.UnaryUnaryCall[_typing.Any, ListAllResponse]]: ...
    def ListTrafficState(self, __1: ListTrafficStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _typing.Union[ListTrafficStateResult, _grpc_aio.UnaryUnaryCall[_typing.Any, ListTrafficStateResult]]: ...

# noinspection PyPep8Naming,DuplicatedCode
class _TopologyManagerReadServiceXBlockingStub(TopologyManagerReadServiceXStub):
    def ListNamespaceDelegation(self, __1: ListNamespaceDelegationRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListNamespaceDelegationResult: ...
    def ListUnionspaceDefinition(self, __1: ListUnionspaceDefinitionRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListUnionspaceDefinitionResult: ...
    def ListIdentifierDelegation(self, __1: ListIdentifierDelegationRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListIdentifierDelegationResult: ...
    def ListOwnerToKeyMapping(self, __1: ListOwnerToKeyMappingRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListOwnerToKeyMappingResult: ...
    def ListDomainTrustCertificate(self, __1: ListDomainTrustCertificateRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListDomainTrustCertificateResult: ...
    def ListParticipantDomainPermission(self, __1: ListParticipantDomainPermissionRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListParticipantDomainPermissionResult: ...
    def ListPartyHostingLimits(self, __1: ListPartyHostingLimitsRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListPartyHostingLimitsResult: ...
    def ListVettedPackages(self, __1: ListVettedPackagesRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListVettedPackagesResult: ...
    def ListPartyToParticipant(self, __1: ListPartyToParticipantRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListPartyToParticipantResult: ...
    def ListAuthorityOf(self, __1: ListAuthorityOfRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListAuthorityOfResult: ...
    def ListDomainParametersState(self, __1: ListDomainParametersStateRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListDomainParametersStateResult: ...
    def ListMediatorDomainState(self, __1: ListMediatorDomainStateRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListMediatorDomainStateResult: ...
    def ListSequencerDomainState(self, __1: ListSequencerDomainStateRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListSequencerDomainStateResult: ...
    def ListPurgeTopologyTransactionX(self, __1: ListPurgeTopologyTransactionXRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListPurgeTopologyTransactionXResult: ...
    def ListAvailableStores(self, __1: ListAvailableStoresRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListAvailableStoresResult: ...
    def ListAll(self, __1: ListAllRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListAllResponse: ...
    def ListTrafficState(self, __1: ListTrafficStateRequest, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_typing.Tuple[_typing.Tuple[str, _typing.Union[str, bytes]], ...]] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> ListTrafficStateResult: ...

# noinspection PyPep8Naming,DuplicatedCode
class _TopologyManagerReadServiceXAsyncStub(TopologyManagerReadServiceXStub):
    def ListNamespaceDelegation(self, __1: ListNamespaceDelegationRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListNamespaceDelegationResult]: ...  # type: ignore
    def ListUnionspaceDefinition(self, __1: ListUnionspaceDefinitionRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListUnionspaceDefinitionResult]: ...  # type: ignore
    def ListIdentifierDelegation(self, __1: ListIdentifierDelegationRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListIdentifierDelegationResult]: ...  # type: ignore
    def ListOwnerToKeyMapping(self, __1: ListOwnerToKeyMappingRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListOwnerToKeyMappingResult]: ...  # type: ignore
    def ListDomainTrustCertificate(self, __1: ListDomainTrustCertificateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListDomainTrustCertificateResult]: ...  # type: ignore
    def ListParticipantDomainPermission(self, __1: ListParticipantDomainPermissionRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListParticipantDomainPermissionResult]: ...  # type: ignore
    def ListPartyHostingLimits(self, __1: ListPartyHostingLimitsRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListPartyHostingLimitsResult]: ...  # type: ignore
    def ListVettedPackages(self, __1: ListVettedPackagesRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListVettedPackagesResult]: ...  # type: ignore
    def ListPartyToParticipant(self, __1: ListPartyToParticipantRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListPartyToParticipantResult]: ...  # type: ignore
    def ListAuthorityOf(self, __1: ListAuthorityOfRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListAuthorityOfResult]: ...  # type: ignore
    def ListDomainParametersState(self, __1: ListDomainParametersStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListDomainParametersStateResult]: ...  # type: ignore
    def ListMediatorDomainState(self, __1: ListMediatorDomainStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListMediatorDomainStateResult]: ...  # type: ignore
    def ListSequencerDomainState(self, __1: ListSequencerDomainStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListSequencerDomainStateResult]: ...  # type: ignore
    def ListPurgeTopologyTransactionX(self, __1: ListPurgeTopologyTransactionXRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListPurgeTopologyTransactionXResult]: ...  # type: ignore
    def ListAvailableStores(self, __1: ListAvailableStoresRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListAvailableStoresResult]: ...  # type: ignore
    def ListAll(self, __1: ListAllRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListAllResponse]: ...  # type: ignore
    def ListTrafficState(self, __1: ListTrafficStateRequest, *, timeout: _typing.Optional[float] = ..., metadata: _typing.Optional[_grpc_aio.Metadata] = ..., credentials: _typing.Optional[_grpc.CallCredentials] = ..., wait_for_ready: _typing.Optional[bool] = ..., compression: _typing.Optional[_grpc.Compression] = ...) -> _grpc_aio.UnaryUnaryCall[_typing.Any, ListTrafficStateResult]: ...  # type: ignore
