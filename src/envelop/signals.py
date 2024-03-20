# ruff: noqa: ARG001
from __future__ import annotations

import asyncio
import platform
import signal
from contextlib import contextmanager
from typing import TYPE_CHECKING

from envelop.queue import AsyncIterableQueue

if TYPE_CHECKING:
    from collections.abc import Callable, Generator, Iterable
    from types import FrameType


# The following signals handlers are inspired by the python package `trio`:
# https://github.com/python-trio/trio/blob/e94e92a3ed21ac8f8cc15940b4f4347e9ced9465/src/trio/_signals.py#L69


class DefaultSignalsHandler:
    def __init__(self, signals: Iterable[int], callback: Callable[[], None]) -> None:
        if platform.system() != "Windows":
            self._handler = NativeSignalsHandler(signals, callback)
        else:
            self._handler = WindowsSignalsHandler(signals, callback)

    async def __aenter__(self):
        await self._handler.__aenter__()

    async def __aexit__(self, *args, **kwargs):
        await self._handler.__aexit__(*args, **kwargs)


class NativeSignalsHandler:
    def __init__(self, signals: Iterable[int], callback: Callable[[], None]):
        self._signals = signals
        self._callback = callback

    async def __aenter__(self):
        _loop = asyncio.get_running_loop()
        for sig in self._signals:
            _loop.remove_signal_handler(sig)
            _loop.add_signal_handler(sig, self._callback)

    async def __aexit__(self, *args, **kwargs):
        _loop = asyncio.get_running_loop()
        for sig in self._signals:
            _loop.remove_signal_handler(sig)


class WindowsSignalsHandler:
    def __init__(self, signals: Iterable[int], callback: Callable[[], None]):
        self._signals = signals
        self._callback = callback
        self._task: asyncio.Task | None = None

    @contextmanager
    def _signal_handler(
        self,
        handler: Callable[[int, FrameType | None], object]
        | int
        | signal.Handlers
        | None,
    ) -> Generator[None, None, None]:
        original_handlers = {}
        try:
            for signum in set(self._signals):
                original_handlers[signum] = signal.signal(signum, handler)
            yield
        finally:
            for signum, original_handler in original_handlers.items():
                signal.signal(signum, original_handler)

    async def _handle_signals(self):
        queue = AsyncIterableQueue[int]()
        loop = asyncio.get_running_loop()

        def _handler(signum: int, frame: FrameType | None):
            loop.call_soon_threadsafe(queue.put_nowait, signum)

        with self._signal_handler(_handler):
            async for sig in queue:
                if sig in self._signals:
                    self._callback()
                    break

        queue.dispose()

    async def __aenter__(self):
        self._task = asyncio.create_task(self._handle_signals())

    async def __aexit__(self, *args, **kwargs):
        if self._task is not None:
            self._task.cancel()
