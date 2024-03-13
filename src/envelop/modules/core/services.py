from typing import final

from google.protobuf.empty_pb2 import Empty
from grpc import ServicerContext
from typing_extensions import AsyncIterator

from envelop.services.process import AbstractProcessService
from envelop.services.proto.process_pb2 import Command, Log
from envelop.services.proto.system_pb2 import Event
from envelop.services.system import AbstractSystemService


@final
class ProcessService(AbstractProcessService):
    async def WriteCommand(
        self, request: Command, context: ServicerContext
    ) -> Empty: ...

    async def StreamLogs(
        self, request: Empty, context: ServicerContext
    ) -> AsyncIterator[Log]: ...


@final
class SystemService(AbstractSystemService):
    async def StreamEvents(
        self, request: Empty, context: ServicerContext
    ) -> AsyncIterator[Event]: ...
