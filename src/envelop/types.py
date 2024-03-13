import asyncio
from typing import Any, Coroutine, Mapping, Protocol, Sequence

import grpc


class ProcessWriter(Protocol):
    def write(self, value: str, encoding: str = ...) -> Coroutine[None, None, None]: ...


class ProcessStopper(Protocol):
    def stop(
        self, writter: ProcessWriter, process: asyncio.subprocess.Process
    ) -> Coroutine[None, None, None]: ...


class Event(Protocol):
    def get_uid(self) -> str: ...

    def get_name(self) -> str: ...

    def get_data(self) -> Mapping[str, Any]: ...


class SupportsGenericRpcHandlers(Protocol):
    def add_generic_rpc_handlers(
        self, generic_rpc_handlers: Sequence[grpc.GenericRpcHandler]
    ) -> None: ...


class Servicer(Protocol):
    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None: ...


class Builder(Protocol):
    def add_service(self, service: Servicer) -> None: ...


class Module(Protocol):
    def register(self, builder: Builder) -> Builder: ...
