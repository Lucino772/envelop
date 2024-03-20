# ruff: noqa: N802, ARG002
from __future__ import annotations

import datetime as dt
import uuid
from typing import TYPE_CHECKING

from google.protobuf.empty_pb2 import Empty
from google.protobuf.struct_pb2 import Struct
from google.protobuf.timestamp_pb2 import Timestamp
from typing_extensions import final

from envelop.services.process import AbstractProcessService
from envelop.services.proto.process_pb2 import Command, Log
from envelop.services.proto.system_pb2 import Event
from envelop.services.system import AbstractSystemService

if TYPE_CHECKING:
    from collections.abc import AsyncIterator

    from grpc import ServicerContext

    from envelop.types import Context


@final
class ProcessService(AbstractProcessService):
    def __init__(self, ctx: Context):
        self._ctx = ctx

    async def WriteCommand(self, request: Command, context: ServicerContext) -> Empty:
        await self._ctx.write_stdin(request.value)
        return Empty()

    async def StreamLogs(
        self,
        request: Empty,
        context: ServicerContext,
    ) -> AsyncIterator[Log]:
        async for line in self._ctx.iter_logs():
            timestamp = Timestamp()
            timestamp.FromDatetime(dt.datetime.now(tz=dt.timezone.utc))
            yield Log(id=uuid.uuid4().hex, timestamp=timestamp, value=line)


@final
class SystemService(AbstractSystemService):
    def __init__(self, ctx: Context):
        self._ctx = ctx

    async def StreamEvents(
        self,
        request: Empty,
        context: ServicerContext,
    ) -> AsyncIterator[Event]:
        async for event in self._ctx.iter_events():
            data = Struct()
            data.update(event.get_data())
            yield Event(id=event.get_uid(), name=event.get_name(), data=data)
