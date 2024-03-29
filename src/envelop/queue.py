from __future__ import annotations

import asyncio
from typing import TYPE_CHECKING, Generic

from typing_extensions import TypeVar

if TYPE_CHECKING:
    from collections.abc import AsyncIterator

T = TypeVar("T")


class AsyncIterableQueue(Generic[T]):
    def __init__(self):
        self._queue = asyncio.Queue[T]()
        self._disposed = asyncio.Event()

    def __aiter__(self) -> AsyncIterator[T]:
        return self

    async def __anext__(self) -> T:
        disposed_task = asyncio.create_task(self._disposed.wait(), name="dispose")
        queue_task = asyncio.create_task(self._queue.get(), name="queue")

        done, pending = await asyncio.wait(
            [disposed_task, queue_task], return_when=asyncio.FIRST_COMPLETED
        )

        for task in pending:
            task.cancel()

        completed_task = done.pop()
        if completed_task.get_name() == "dispose":
            raise StopAsyncIteration

        return completed_task.result()  # type: ignore

    async def put(self, item: T) -> None:
        await self._queue.put(item)

    def put_nowait(self, item: T) -> None:
        self._queue.put_nowait(item)

    def dispose(self) -> None:
        if not self._disposed.is_set():
            self._disposed.set()


class Producer(Generic[T]):
    def __init__(self) -> None:
        self._queue = AsyncIterableQueue[T]()
        self._consumers: list[AsyncIterableQueue[T]] = []

    async def put(self, item: T) -> None:
        await self._queue.put(item)

    def __aiter__(self) -> AsyncIterator[T]:
        consumer = AsyncIterableQueue[T]()
        self._consumers.append(consumer)
        return consumer

    async def run(self) -> None:
        try:
            async for item in self._queue:
                await asyncio.gather(
                    *[consumer.put(item) for consumer in self._consumers]
                )
        finally:
            for consumer in self._consumers:
                consumer.dispose()

    def dispose(self) -> None:
        self._queue.dispose()
