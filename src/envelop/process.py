from __future__ import annotations

import asyncio
import contextlib
import os
import signal
from typing import TYPE_CHECKING, Any

if TYPE_CHECKING:
    from collections.abc import Iterable, Mapping, MutableMapping

    from envelop.queue import Producer


class AppProcess:
    def __init__(
        self,
        command: Iterable[str],
        producer: Producer[str],
        init_env: Mapping[str, Any] = os.environ,
    ) -> None:
        self._producer = producer
        self._command = command
        self._env_vars: MutableMapping[str, Any] = {**init_env}
        self._cwd: str = os.getcwd()
        self._args: list[str] = []

        self._process: asyncio.subprocess.Process | None = None
        self._graceful_timeout: int = 30
        self._graceful_command: str | None = None
        self._graceful_signal: int | None = None

    def cwd(self, directory: str) -> AppProcess:
        self._cwd = directory
        return self

    def args(self, args: Iterable[str]) -> AppProcess:
        self._args.extend(args)
        return self

    def envs(self, env_vars: Mapping[str, Any]) -> AppProcess:
        self._env_vars.update(env_vars)
        return self

    def graceful(self, command_or_signal: str | int, timeout: int) -> AppProcess:
        self._graceful_timeout = timeout
        if isinstance(command_or_signal, str):
            self._graceful_command = command_or_signal
        else:
            self._graceful_signal = command_or_signal
        return self

    async def write(self, value: str) -> None:
        if self._process is None or self._process.stdin is None:
            return  # TODO: Return error

        self._process.stdin.write(f"{value}\n".encode())
        await self._process.stdin.drain()

    async def run(self) -> None:
        self._process = await asyncio.create_subprocess_exec(
            *self._command,
            stdin=asyncio.subprocess.PIPE,
            stdout=asyncio.subprocess.PIPE,
            stderr=asyncio.subprocess.STDOUT,
            env=self._env_vars,
            cwd=self._cwd,
        )

        stop_flag = asyncio.Event()
        produce_logs_task = asyncio.create_task(self._produce_logs(self._process))
        self._setup_interrupts(stop_flag)

        stop_flag_task = asyncio.create_task(stop_flag.wait(), name="stop")
        proc_wait_task = asyncio.create_task(self._process.wait(), name="wait")

        done, pending = await asyncio.wait(
            [stop_flag_task, proc_wait_task], return_when=asyncio.FIRST_COMPLETED
        )

        for task in pending:
            task.cancel()

        with contextlib.suppress(Exception):
            completed_task = done.pop()
            if completed_task.get_name() == "stop" and self._process.returncode is None:
                await self._stop(self._process)

        produce_logs_task.cancel()

    def _setup_interrupts(self, stop_flag: asyncio.Event):
        _loop = asyncio.get_running_loop()
        for sig in (signal.SIGINT, signal.SIGTERM):
            _loop.remove_signal_handler(sig)
            _loop.add_signal_handler(sig, lambda: stop_flag.set())

    async def _produce_logs(self, process: asyncio.subprocess.Process):
        if process.stdout is None:
            return

        async for line in process.stdout:
            decoded = line.decode("utf-8").strip("\n")
            await self._producer.put(decoded)

    async def _stop(self, process: asyncio.subprocess.Process) -> None:
        if self._graceful_command is not None:
            await self._stop_with_command(
                process, self._graceful_command, self._graceful_timeout
            )
        elif self._graceful_signal is not None:
            await self._stop_with_signal(
                process, self._graceful_signal, self._graceful_timeout
            )

        if process.returncode is None:
            process.kill()

    async def _stop_with_command(
        self, process: asyncio.subprocess.Process, command: str, timeout: int
    ):
        with contextlib.suppress(Exception):
            await self.write(command)
            await asyncio.wait_for(process.wait(), timeout)

    async def _stop_with_signal(
        self, process: asyncio.subprocess.Process, stop_signal: int, timeout: int
    ):
        with contextlib.suppress(Exception):
            process.send_signal(stop_signal)
            await asyncio.wait_for(process.wait(), timeout)
