import asyncio
from collections.abc import AsyncIterator

from envelop.consumer import AsyncConsumer
from envelop.types import Event


class EventEmitter:
    def __init__(self):
        self.__consumers: list[AsyncConsumer[Event]] = []

    async def emit(self, event: Event) -> None:
        await asyncio.gather(*[consumer.put(event) for consumer in self.__consumers])

    def __aiter__(self) -> AsyncIterator[Event]:
        consumer = AsyncConsumer()
        self.__consumers.append(consumer)
        return consumer

    def shutdown(self) -> None:
        for consumer in self.__consumers:
            consumer.dispose()
