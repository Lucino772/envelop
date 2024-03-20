from __future__ import annotations

import asyncio
import shlex
from typing import TYPE_CHECKING

import grpc
import structlog
from typing_extensions import Any, Self, final

from envelop.events import ProcessLog
from envelop.process import AppProcess
from envelop.queue import Producer
from envelop.store import MemoryStore

if TYPE_CHECKING:
    from collections.abc import AsyncIterator, Mapping

    from envelop.config import Config
    from envelop.types import Context, Event, Module, Process, Runnable, Servicer, Store

logger: structlog.stdlib.BoundLogger = structlog.get_logger()


class _ForwardLogTasks:
    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx

    async def run(self):
        async for log in self._ctx.iter_logs():
            await self._ctx.emit_event(ProcessLog(log=log))


class Application:
    def __init__(
        self,
        context: AppContext,
        server: grpc.aio.Server,
        process: Process,
        tasks: list[Runnable],
    ) -> None:
        self._context = context
        self._server = server
        self._process = process
        self._tasks = tasks

    async def run(self) -> None:
        await self._context.run(self._server, self._process, self._tasks)


class AppContext:
    def __init__(
        self, event_producer: Producer[Event], log_producer: Producer[str], store: Store
    ):
        self._events: Producer[Event] = event_producer
        self._logs: Producer[str] = log_producer
        self._process: Process | None = None
        self._tasks: list[asyncio.Task] = []
        self._store: Store = store

    def iter_logs(self) -> AsyncIterator[str]:
        return aiter(self._logs)

    async def write_stdin(self, command: str) -> None:
        if self._process is None:
            raise RuntimeError
        await self._process.write(command)

    def iter_events(self) -> AsyncIterator[Event]:
        return aiter(self._events)

    async def emit_event(self, event: Event) -> None:
        log = logger.bind()
        await self._events.put(event)
        log.debug(
            "app.events:%s", event.get_name(), id=event.get_uid(), data=event.get_data()
        )

    async def write_store(self, key: str, data: Mapping[str, Any]) -> None:
        await self._store.write(key, data)

    async def read_store(self, key: str) -> Mapping[str, Any]:
        return await self._store.read(key)

    async def run(
        self, server: grpc.aio.Server, process: Process, tasks: list[Runnable]
    ) -> None:
        log = logger.bind()
        tasks.append(self._events)
        tasks.append(self._logs)

        try:
            await server.start()
            log.debug("app.server.started")
            for task in [*tasks]:
                self._tasks.append(asyncio.create_task(task.run()))
            log.debug("app.tasks.started")
            self._process = process
            await self._process.run()
        finally:
            await server.stop(10)
            log.debug("app.server.stopped")
            for task in self._tasks:
                task.cancel()
            log.debug("app.tasks.stopped")


@final
class AppBuilder:
    def __init__(self) -> None:
        self._services: list[Servicer] = []
        self._tasks: list[Runnable] = []

    def add_service(self, service: Servicer) -> Self:
        self._services.append(service)
        return self

    def add_task(self, task: Runnable) -> Self:
        self._tasks.append(task)
        return self

    def build(self, config: Config, registry: Mapping[str, Module]) -> Application:
        log_producer: Producer[str] = Producer()
        context = AppContext(
            event_producer=Producer(), log_producer=log_producer, store=MemoryStore()
        )

        # Add forward log task
        self.add_task(_ForwardLogTasks(context))

        # Create process
        command = shlex.split(config.process.command)
        process = (
            AppProcess(command[0], log_producer)
            .args(command[1:])
            .envs(config.process.env)
        )
        timeout = config.process.graceful.timeout
        if config.process.graceful.signal is not None:
            process = process.graceful(config.process.graceful.signal, timeout)
        elif config.process.graceful.cmd is not None:
            process = process.graceful(config.process.graceful.cmd, timeout)

        # Register each module
        for mod in config.modules:
            module = registry[mod.uses]
            module.register(self, context, mod.config)

        # Create server
        server = grpc.aio.server()
        server.add_insecure_port("0.0.0.0:8791")
        for service in self._services:
            service.add_rpc_handlers(server)

        return Application(context, server, process, self._tasks)
