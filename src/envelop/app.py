from __future__ import annotations

import asyncio
import shlex
from typing import TYPE_CHECKING, final

import grpc

from envelop.process import ProcessBuilder
from envelop.queue import Producer

if TYPE_CHECKING:
    from collections.abc import AsyncIterator, Mapping

    from envelop.types import Event, Module, Process, Runnable, Servicer


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
    def __init__(self):
        self._events: Producer[Event] = Producer()
        self._process: Process | None = None
        self._tasks: list[asyncio.Task] = []

    def iter_logs(self) -> AsyncIterator[str]:
        if self._process is None:
            raise RuntimeError

        return aiter(self._process)

    async def write_stdin(self, command: str) -> None:
        if self._process is None:
            raise RuntimeError
        await self._process.write(command)

    def iter_events(self) -> AsyncIterator[Event]:
        return aiter(self._events)

    async def emit_event(self, event: Event) -> None:
        await self._events.put(event)

    async def run(
        self, server: grpc.aio.Server, process: Process, tasks: list[Runnable]
    ) -> None:
        tasks.append(self._events)

        try:
            await server.start()
            for task in [*tasks]:
                self._tasks.append(asyncio.create_task(task.run()))
            self._process = process
            await self._process.run()
        finally:
            await server.stop(10)
            for task in self._tasks:
                task.cancel()


@final
class AppBuilder:
    def __init__(self) -> None:
        self._services: list[Servicer] = []
        self._tasks: list[Runnable] = []

    def add_service(self, service: Servicer) -> AppBuilder:
        self._services.append(service)
        return self

    def add_task(self, task: Runnable) -> AppBuilder:
        self._tasks.append(task)
        return self

    def build(self, config: dict, registry: Mapping[str, Module]) -> Application:
        context = AppContext()

        # Create process
        command = shlex.split(config["process"]["command"])
        env = config["process"].get("env", {})
        graceful = config["process"]["graceful"]

        builder = (
            ProcessBuilder(command[0])
            .args(command[1:])
            .envs(env)
            .stdin(asyncio.subprocess.PIPE)
            .stdout(asyncio.subprocess.PIPE)
            .stderr(asyncio.subprocess.STDOUT)
        )
        timeout = graceful.get("timeout", 30)
        if "signal" in graceful:
            builder = builder.graceful(int(graceful["signal"]), timeout)
        else:
            builder = builder.graceful(graceful["cmd"], timeout)
        process = builder.build()

        # Register each module
        for module_settings in config.get("modules", []):
            module_name = module_settings["uses"]
            module_config = module_settings.get("with", {})

            module = registry[module_name]
            module.register(self, context, module_config)

        # Create server
        server = grpc.aio.server()
        server.add_insecure_port("0.0.0.0:8791")
        for service in self._services:
            service.add_rpc_handlers(server)

        return Application(context, server, process, self._tasks)
