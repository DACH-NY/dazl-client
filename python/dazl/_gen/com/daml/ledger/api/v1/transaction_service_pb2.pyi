# Copyright (c) 2017-2024 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off
# isort: skip_file

import builtins as _builtins, sys, typing as _typing

from google.protobuf.internal.containers import RepeatedCompositeFieldContainer, RepeatedScalarFieldContainer
from google.protobuf.message import Message as _Message

from .ledger_offset_pb2 import LedgerOffset
from .transaction_filter_pb2 import TransactionFilter
from .transaction_pb2 import Transaction, TransactionTree

if sys.version_info >= (3, 8):
    from typing import Literal as _L
else:
    from typing_extensions import Literal as _L

__all__ = [
    "GetTransactionsRequest",
    "GetTransactionsResponse",
    "GetTransactionTreesResponse",
    "GetTransactionByEventIdRequest",
    "GetTransactionByIdRequest",
    "GetTransactionResponse",
    "GetFlatTransactionResponse",
    "GetLedgerEndRequest",
    "GetLedgerEndResponse",
]


class GetTransactionsRequest(_Message):
    ledger_id: _builtins.str
    @property
    def begin(self) -> LedgerOffset: ...
    @property
    def end(self) -> LedgerOffset: ...
    @property
    def filter(self) -> TransactionFilter: ...
    verbose: _builtins.bool
    def __init__(self, *, ledger_id: _typing.Optional[_builtins.str] = ..., begin: _typing.Optional[LedgerOffset] = ..., end: _typing.Optional[LedgerOffset] = ..., filter: _typing.Optional[TransactionFilter] = ..., verbose: _typing.Optional[_builtins.bool] = ...): ...
    def HasField(self, field_name: _L["ledger_id", "begin", "end", "filter", "verbose"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["ledger_id", "begin", "end", "filter", "verbose"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetTransactionsResponse(_Message):
    @property
    def transactions(self) -> RepeatedCompositeFieldContainer[Transaction]: ...
    def __init__(self, *, transactions: _typing.Optional[_typing.Iterable[Transaction]] = ...): ...
    def HasField(self, field_name: _L["transactions"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["transactions"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetTransactionTreesResponse(_Message):
    @property
    def transactions(self) -> RepeatedCompositeFieldContainer[TransactionTree]: ...
    def __init__(self, *, transactions: _typing.Optional[_typing.Iterable[TransactionTree]] = ...): ...
    def HasField(self, field_name: _L["transactions"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["transactions"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetTransactionByEventIdRequest(_Message):
    ledger_id: _builtins.str
    event_id: _builtins.str
    @property
    def requesting_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    def __init__(self, *, ledger_id: _typing.Optional[_builtins.str] = ..., event_id: _typing.Optional[_builtins.str] = ..., requesting_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ...): ...
    def HasField(self, field_name: _L["ledger_id", "event_id", "requesting_parties"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["ledger_id", "event_id", "requesting_parties"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetTransactionByIdRequest(_Message):
    ledger_id: _builtins.str
    transaction_id: _builtins.str
    @property
    def requesting_parties(self) -> RepeatedScalarFieldContainer[_builtins.str]: ...
    def __init__(self, *, ledger_id: _typing.Optional[_builtins.str] = ..., transaction_id: _typing.Optional[_builtins.str] = ..., requesting_parties: _typing.Optional[_typing.Iterable[_builtins.str]] = ...): ...
    def HasField(self, field_name: _L["ledger_id", "transaction_id", "requesting_parties"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["ledger_id", "transaction_id", "requesting_parties"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetTransactionResponse(_Message):
    @property
    def transaction(self) -> TransactionTree: ...
    def __init__(self, *, transaction: _typing.Optional[TransactionTree] = ...): ...
    def HasField(self, field_name: _L["transaction"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["transaction"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetFlatTransactionResponse(_Message):
    @property
    def transaction(self) -> Transaction: ...
    def __init__(self, *, transaction: _typing.Optional[Transaction] = ...): ...
    def HasField(self, field_name: _L["transaction"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["transaction"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetLedgerEndRequest(_Message):
    ledger_id: _builtins.str
    def __init__(self, *, ledger_id: _typing.Optional[_builtins.str] = ...): ...
    def HasField(self, field_name: _L["ledger_id"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["ledger_id"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...

class GetLedgerEndResponse(_Message):
    @property
    def offset(self) -> LedgerOffset: ...
    def __init__(self, *, offset: _typing.Optional[LedgerOffset] = ...): ...
    def HasField(self, field_name: _L["offset"]) -> _builtins.bool: ...
    def ClearField(self, field_name: _L["offset"]) -> None: ...
    def WhichOneof(self, oneof_group: _typing.NoReturn) -> _typing.NoReturn: ...
