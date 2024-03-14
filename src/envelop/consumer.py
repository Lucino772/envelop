import asyncio
from collections.abc import AsyncIterator
from typing import Generic, TypeVar

T = TypeVar("T")


class AsyncConsumer(Generic[T]):
    def __init__(self):
        self.__queue = asyncio.Queue[T]()
        self.__disposed = asyncio.Event()

    def __aiter__(self) -> AsyncIterator[T]:
        return self

    async def __anext__(self) -> T:
        disposed_task = asyncio.create_task(self.__disposed.wait(), name="dispose")
        queue_task = asyncio.create_task(self.__queue.get(), name="queue")

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
        await self.__queue.put(item)

    def dispose(self) -> None:
        if not self.__disposed.is_set():
            self.__disposed.set()
