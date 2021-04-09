# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off

from typing import (
    AsyncIterable as __AsyncStream,
    Awaitable as __Awaitable,
    Iterable as __Stream,
    Union as __Union,
    overload,
)

from grpc import Channel as __Channel
from grpc.aio import Channel as __AsyncChannel

from .transaction_service_pb2 import (
    GetFlatTransactionResponse,
    GetLedgerEndRequest,
    GetLedgerEndResponse,
    GetTransactionByEventIdRequest,
    GetTransactionByIdRequest,
    GetTransactionsRequest,
    GetTransactionsResponse,
    GetTransactionTreesResponse,
)

__all__ = [
    "TransactionServiceStub"
]

# noinspection PyPep8Naming
class TransactionServiceStub:
    @classmethod
    @overload
    def __new__(cls, channel: __Channel) -> _TransactionServiceStub: ...
    @classmethod
    @overload
    def __new__(cls, channel: __AsyncChannel) -> _TransactionServiceStub_Async: ...
    def GetTransactions(self, __1: GetTransactionsRequest) -> __Union[__Stream[GetTransactionsResponse], __AsyncStream[GetTransactionsResponse]]: ...
    def GetTransactionTrees(self, __1: GetTransactionsRequest) -> __Union[__Stream[GetTransactionTreesResponse], __AsyncStream[GetTransactionTreesResponse]]: ...
    def GetTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> __Union[GetTransactionsResponse, __Awaitable[GetTransactionsResponse]]: ...
    def GetTransactionById(self, __1: GetTransactionByIdRequest) -> __Union[GetTransactionsResponse, __Awaitable[GetTransactionsResponse]]: ...
    def GetFlatTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> __Union[GetFlatTransactionResponse, __Awaitable[GetFlatTransactionResponse]]: ...
    def GetFlatTransactionById(self, __1: GetTransactionByIdRequest) -> __Union[GetFlatTransactionResponse, __Awaitable[GetFlatTransactionResponse]]: ...
    def GetLedgerEnd(self, __1: GetLedgerEndRequest) -> __Union[GetLedgerEndResponse, __Awaitable[GetLedgerEndResponse]]: ...

class _TransactionServiceStub(TransactionServiceStub):
    def GetTransactions(self, __1: GetTransactionsRequest) -> __Stream[GetTransactionsResponse]: ...
    def GetTransactionTrees(self, __1: GetTransactionsRequest) -> __Stream[GetTransactionTreesResponse]: ...
    def GetTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> GetTransactionsResponse: ...
    def GetTransactionById(self, __1: GetTransactionByIdRequest) -> GetTransactionsResponse: ...
    def GetFlatTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> GetFlatTransactionResponse: ...
    def GetFlatTransactionById(self, __1: GetTransactionByIdRequest) -> GetFlatTransactionResponse: ...
    def GetLedgerEnd(self, __1: GetLedgerEndRequest) -> GetLedgerEndResponse: ...

class _TransactionServiceStub_Async(TransactionServiceStub):
    def GetTransactions(self, __1: GetTransactionsRequest) -> __AsyncStream[GetTransactionsResponse]: ...
    def GetTransactionTrees(self, __1: GetTransactionsRequest) -> __AsyncStream[GetTransactionTreesResponse]: ...
    def GetTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> __Awaitable[GetTransactionsResponse]: ...
    def GetTransactionById(self, __1: GetTransactionByIdRequest) -> __Awaitable[GetTransactionsResponse]: ...
    def GetFlatTransactionByEventId(self, __1: GetTransactionByEventIdRequest) -> __Awaitable[GetFlatTransactionResponse]: ...
    def GetFlatTransactionById(self, __1: GetTransactionByIdRequest) -> __Awaitable[GetFlatTransactionResponse]: ...
    def GetLedgerEnd(self, __1: GetLedgerEndRequest) -> __Awaitable[GetLedgerEndResponse]: ...