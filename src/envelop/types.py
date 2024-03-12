import asyncio
from typing import Coroutine, Protocol


class ProcessWriter(Protocol):
    def write(self, value: str, encoding: str = ...) -> Coroutine[None, None, None]: ...


class ProcessStopper(Protocol):
    def stop(
        self, writter: ProcessWriter, process: asyncio.subprocess.Process
    ) -> Coroutine[None, None, None]: ...
