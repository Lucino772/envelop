import asyncio
from typing import Any, Coroutine, Mapping, Protocol


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
