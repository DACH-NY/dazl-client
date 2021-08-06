# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off

import typing as _typing

from grpc import Channel as __Channel
from grpc.aio import Channel as __AsyncChannel

from .participant_pruning_service_pb2 import PruneRequest, PruneResponse

__all__ = [
    "ParticipantPruningServiceStub",
]

class ParticipantPruningServiceStub:
    @classmethod
    @_typing.overload
    def __new__(cls, channel: __Channel) -> _ParticipantPruningServiceStub: ...
    @classmethod
    @_typing.overload
    def __new__(cls, channel: __AsyncChannel) -> _ParticipantPruningServiceStub_Async: ...
    def Prune(self, __1: PruneRequest) -> _typing.Union[PruneResponse, _typing.Awaitable[PruneResponse]]: ...

class _ParticipantPruningServiceStub(ParticipantPruningServiceStub):
    def Prune(self, __1: PruneRequest) -> PruneResponse: ...

class _ParticipantPruningServiceStub_Async(ParticipantPruningServiceStub):
    def Prune(self, __1: PruneRequest) -> _typing.Awaitable[PruneResponse]: ...
