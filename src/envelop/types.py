from __future__ import annotations

from typing import TYPE_CHECKING

from typing_extensions import Any, Protocol, Self

if TYPE_CHECKING:
    from collections.abc import AsyncIterator, Coroutine, Mapping, Sequence

    import grpc


class Event(Protocol):
    def get_uid(self) -> str:
        ...

    def get_name(self) -> str:
        ...

    def get_data(self) -> Mapping[str, Any]:
        ...


class Store(Protocol):
    def write(self, key: str, data: Mapping[str, Any]) -> Coroutine[None, None, None]:
        ...

    def read(
        self, key: str, default: Mapping[str, Any]
    ) -> Coroutine[None, None, Mapping[str, Any]]:
        ...


class SupportsGenericRpcHandlers(Protocol):
    def add_generic_rpc_handlers(
        self, generic_rpc_handlers: Sequence[grpc.GenericRpcHandler]
    ) -> None:
        ...


class Servicer(Protocol):
    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        ...


class Runnable(Protocol):
    def run(self) -> Coroutine[None, None, None]:
        ...


class Context(Protocol):
    def wait_process(self) -> Coroutine[None, None, None]:
        ...

    def iter_logs(self) -> AsyncIterator[str]:
        ...

    def write_stdin(self, command: str) -> Coroutine[None, None, None]:
        ...

    def iter_events(self) -> AsyncIterator[Event]:
        ...

    def emit_event(self, event: Event) -> Coroutine[None, None, None]:
        ...

    def write_store(
        self, key: str, data: Mapping[str, Any]
    ) -> Coroutine[None, None, None]:
        ...

    def read_store(
        self, key: str, default: Mapping[str, Any]
    ) -> Coroutine[None, None, Mapping[str, Any]]:
        ...


class Builder(Protocol):
    def add_service(self, service: Servicer) -> Self:
        ...

    def add_task(self, task: Runnable) -> Self:
        ...


class Module(Protocol):
    def register(self, builder: Builder, context: Context, config: dict) -> Builder:
        ...


class Process(Protocol):
    def write(self, value: str) -> Coroutine[Any, Any, None]:
        ...

    def run(self) -> Coroutine[Any, Any, None]:
        ...
