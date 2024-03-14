from __future__ import annotations

import asyncio
import contextlib
import os
import signal
from typing import (
    IO,
    Any,
    AsyncIterator,
    Iterable,
    List,
    Mapping,
    MutableMapping,
    final,
)

from envelop.consumer import AsyncConsumer
from envelop.types import Process


@final
class ProcessBuilder:
    def __init__(self, program: str) -> None:
        self._program: str = program
        self._cwd: str | None = None
        self._args: List[str] = []
        self._env_vars: MutableMapping[str, Any] = {**os.environ}
        self._stdin: int | IO[Any] | None = None
        self._stdout: int | IO[Any] | None = None
        self._stderr: int | IO[Any] | None = None
        self._graceful_command: str | None = None
        self._graceful_signal: int | None = None
        self._graceful_timeout: int = 30

    def cwd(self, directory: str) -> ProcessBuilder:
        self._cwd = directory
        return self

    def args(self, args: Iterable[str]) -> ProcessBuilder:
        self._args.extend(args)
        return self

    def envs(self, vars: Mapping[str, Any]) -> ProcessBuilder:
        self._env_vars.update(vars)
        return self

    def stdin(self, stdin: int | IO[Any]) -> ProcessBuilder:
        self._stdin = stdin
        return self

    def stdout(self, stdout: int | IO[Any]) -> ProcessBuilder:
        self._stdout = stdout
        return self

    def stderr(self, stderr: int | IO[Any]) -> ProcessBuilder:
        self._stderr = stderr
        return self

    def graceful(self, command_or_signal: str | int, timeout: int) -> ProcessBuilder:
        self._graceful_timeout = timeout
        if isinstance(command_or_signal, str):
            self._graceful_command = command_or_signal
        else:
            self._graceful_signal = command_or_signal
        return self

    def build(self) -> Process:
        cmd_args = [self._program] + self._args
        return AppProcess(
            command=cmd_args,
            env=self._env_vars,
            stdin=self._stdin,
            stdout=self._stdout,
            stderr=self._stderr,
            graceful_timeout=self._graceful_timeout,
            graceful_command=self._graceful_command,
            graceful_signal=self._graceful_signal,
        )


class AppProcess:
    def __init__(
        self,
        command: Iterable[str],
        env: Mapping[str, Any],
        stdin: int | IO[Any] | None,
        stdout: int | IO[Any] | None,
        stderr: int | IO[Any] | None,
        graceful_timeout: int,
        graceful_command: str | None,
        graceful_signal: int | None,
    ) -> None:
        self._command = command
        self._env = env
        self._stdin = stdin
        self._stdout = stdout
        self._stderr = stderr

        self._process: asyncio.subprocess.Process | None = None
        self._consumers: List[AsyncConsumer[str]] = []
        self._graceful_timeout: int = graceful_timeout
        self._graceful_command: str | None = graceful_command
        self._graceful_signal: int | None = graceful_signal

    def __aiter__(self) -> AsyncIterator[str]:
        consumer = AsyncConsumer()
        self._consumers.append(consumer)
        return consumer

    async def write(self, value: str) -> None:
        if self._process is None or self._process.stdin is None:
            return  # TODO: Return error

        self._process.stdin.write(f"{value}\n".encode("utf-8"))
        await self._process.stdin.drain()

    async def run(self) -> None:
        self._process = await asyncio.create_subprocess_exec(
            *self._command,
            stdin=self._stdin,
            stdout=self._stdout,
            stderr=self._stderr,
            env=self._env,
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
        for sig in {signal.SIGINT, signal.SIGTERM}:
            _loop.remove_signal_handler(sig)
            _loop.add_signal_handler(sig, lambda: stop_flag.set())

    async def _produce_logs(self, process: asyncio.subprocess.Process):
        if process.stdout is None:
            return

        async for line in process.stdout:
            decoded = line.decode("utf-8").strip("\n")
            await asyncio.gather(
                *[consumer.put(decoded) for consumer in self._consumers]
            )

        for consumer in self._consumers:
            consumer.dispose()

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
