# Copyright (c) 2017-2021 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# fmt: off

from typing import (
    Any as __Any,
    Dict as __Map,
    List as __List,
    Literal as __Literal,
    NoReturn as __Void,
    Optional as __Optional,
    Sequence as __Sequence,
    Tuple as __Tuple,
    overload,
)

from google.protobuf.descriptor import FieldDescriptor as __FieldDescriptor
from google.protobuf.message import Message as __Message

from .commands_pb2 import Commands
from .trace_context_pb2 import TraceContext

__all__ = [
    "SubmitAndWaitRequest",
    "SubmitAndWaitForTransactionIdResponse",
    "SubmitAndWaitForTransactionResponse",
    "SubmitAndWaitForTransactionTreeResponse",
]

from .transaction_pb2 import Transaction, TransactionTree

class SubmitAndWaitRequest(__Message):
    commands: Commands
    trace_context: TraceContext
    def __init__(self, *, commands: __Optional[Commands] = ..., trace_context: __Optional[TraceContext] = ...): ...
    def __eq__(self, other_msg: __Optional[__Any]) -> bool: ...
    def __str__(self) -> str: ...
    def __unicode__(self) -> str: ...
    def MergeFrom(self, other_msg: SubmitAndWaitRequest) -> None: ...
    def Clear(self) -> None: ...
    def SetInParent(self) -> None: ...
    def IsInitialized(self) -> bool: ...
    def MergeFromString(self, serialized: bytes) -> int: ...
    def SerializeToString(self, *, deterministic: bool = ...) -> bytes: ...
    def SerializePartialToString(self, *, deterministic: bool = ...) -> bytes: ...
    def ListFields(self) -> __Sequence[__Tuple[__FieldDescriptor, __Any]]: ...
    def HasField(self, field_name: __Literal["commands", "trace_context"]) -> bool: ...
    def ClearField(self, field_name: __Literal["commands", "trace_context"]) -> None: ...
    def WhichOneof(self, oneof_group: __Void) -> __Void: ...
    def HasExtension(self, extension_handle: __Any) -> bool: ...
    def ClearExtension(self, extension_handle: __Any) -> None: ...
    def UnknownFields(self) -> __Any: ...
    def DiscardUnknownFields(self) -> None: ...
    def ByteSize(self) -> int: ...
    def _SetListener(self, message_listener: __Any) -> None: ...

class SubmitAndWaitForTransactionIdResponse(__Message):
    transaction_id: str
    def __init__(self, *, transaction_id: __Optional[str] = ...): ...
    def __eq__(self, other_msg: __Optional[__Any]) -> bool: ...
    def __str__(self) -> str: ...
    def __unicode__(self) -> str: ...
    def MergeFrom(self, other_msg: SubmitAndWaitRequest) -> None: ...
    def Clear(self) -> None: ...
    def SetInParent(self) -> None: ...
    def IsInitialized(self) -> bool: ...
    def MergeFromString(self, serialized: bytes) -> int: ...
    def SerializeToString(self, *, deterministic: bool = ...) -> bytes: ...
    def SerializePartialToString(self, *, deterministic: bool = ...) -> bytes: ...
    def ListFields(self) -> __Sequence[__Tuple[__FieldDescriptor, __Any]]: ...
    def HasField(self, field_name: __Literal["transaction_id"]) -> bool: ...
    def ClearField(self, field_name: __Literal["transaction_id"]) -> None: ...
    def WhichOneof(self, oneof_group: __Void) -> __Void: ...
    def HasExtension(self, extension_handle: __Any) -> bool: ...
    def ClearExtension(self, extension_handle: __Any) -> None: ...
    def UnknownFields(self) -> __Any: ...
    def DiscardUnknownFields(self) -> None: ...
    def ByteSize(self) -> int: ...
    def _SetListener(self, message_listener: __Any) -> None: ...

class SubmitAndWaitForTransactionResponse(__Message):
    transaction: Transaction
    def __init__(self, *, transaction_id: __Optional[Transaction] = ...): ...
    def __eq__(self, other_msg: __Optional[__Any]) -> bool: ...
    def __str__(self) -> str: ...
    def __unicode__(self) -> str: ...
    def MergeFrom(self, other_msg: SubmitAndWaitRequest) -> None: ...
    def Clear(self) -> None: ...
    def SetInParent(self) -> None: ...
    def IsInitialized(self) -> bool: ...
    def MergeFromString(self, serialized: bytes) -> int: ...
    def SerializeToString(self, *, deterministic: bool = ...) -> bytes: ...
    def SerializePartialToString(self, *, deterministic: bool = ...) -> bytes: ...
    def ListFields(self) -> __Sequence[__Tuple[__FieldDescriptor, __Any]]: ...
    def HasField(self, field_name: __Literal["transaction"]) -> bool: ...
    def ClearField(self, field_name: __Literal["transaction"]) -> None: ...
    def WhichOneof(self, oneof_group: __Void) -> __Void: ...
    def HasExtension(self, extension_handle: __Any) -> bool: ...
    def ClearExtension(self, extension_handle: __Any) -> None: ...
    def UnknownFields(self) -> __Any: ...
    def DiscardUnknownFields(self) -> None: ...
    def ByteSize(self) -> int: ...
    def _SetListener(self, message_listener: __Any) -> None: ...

class SubmitAndWaitForTransactionTreeResponse(__Message):
    transaction: TransactionTree
    def __init__(self, *, transaction: __Optional[TransactionTree] = ...): ...
    def __eq__(self, other_msg: __Optional[__Any]) -> bool: ...
    def __str__(self) -> str: ...
    def __unicode__(self) -> str: ...
    def MergeFrom(self, other_msg: SubmitAndWaitRequest) -> None: ...
    def Clear(self) -> None: ...
    def SetInParent(self) -> None: ...
    def IsInitialized(self) -> bool: ...
    def MergeFromString(self, serialized: bytes) -> int: ...
    def SerializeToString(self, *, deterministic: bool = ...) -> bytes: ...
    def SerializePartialToString(self, *, deterministic: bool = ...) -> bytes: ...
    def ListFields(self) -> __Sequence[__Tuple[__FieldDescriptor, __Any]]: ...
    def HasField(self, field_name: __Literal["transaction"]) -> bool: ...
    def ClearField(self, field_name: __Literal["transaction"]) -> None: ...
    def WhichOneof(self, oneof_group: __Void) -> __Void: ...
    def HasExtension(self, extension_handle: __Any) -> bool: ...
    def ClearExtension(self, extension_handle: __Any) -> None: ...
    def UnknownFields(self) -> __Any: ...
    def DiscardUnknownFields(self) -> None: ...
    def ByteSize(self) -> int: ...
    def _SetListener(self, message_listener: __Any) -> None: ...