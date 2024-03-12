from __future__ import annotations

import asyncio
import contextlib
import signal
from typing import AsyncIterator, Iterable, List, Optional

from envelop.consumer import AsyncConsumer
from envelop.types import ProcessStopper, ProcessWriter


class Process:
    @classmethod
    async def create(
        cls, command: Iterable[str], stopper: ProcessStopper, env: Optional[dict] = None
    ) -> Process:
        subproc = await asyncio.create_subprocess_exec(
            *command,
            stdin=asyncio.subprocess.PIPE,
            stdout=asyncio.subprocess.PIPE,
            stderr=asyncio.subprocess.STDOUT,
            env=env,
        )

        return Process(subproc, stopper)

    def __init__(
        self,
        process: asyncio.subprocess.Process,
        stopper: ProcessStopper,
    ):
        self.__process = process
        self.__stopper = stopper
        self.__consumers: List[AsyncConsumer[str]] = []

    def __aiter__(self) -> AsyncIterator[str]:
        consumer = AsyncConsumer()
        self.__consumers.append(consumer)
        return consumer

    async def write(self, value: str, encoding: str = "utf-8") -> None:
        if self.__process.stdin is None:
            return

        self.__process.stdin.write(f"{value}\n".encode(encoding))
        await self.__process.stdin.drain()

    def _setup_interrupts(self, stop_flag: asyncio.Event):
        _loop = asyncio.get_running_loop()
        for sig in {signal.SIGINT, signal.SIGTERM}:
            _loop.remove_signal_handler(sig)
            _loop.add_signal_handler(sig, lambda: stop_flag.set())

    async def _produce_logs(self):
        if self.__process.stdout is None:
            return

        async for line in self.__process.stdout:
            decoded = line.decode("utf-8").strip("\n")
            await asyncio.gather(
                *[consumer.put(decoded) for consumer in self.__consumers]
            )

        for consumer in self.__consumers:
            consumer.dispose()

    async def wait(self) -> None:
        stop_flag = asyncio.Event()
        produce_logs_task = asyncio.create_task(self._produce_logs(), name="logs")
        self._setup_interrupts(stop_flag)
        stop_flag_task = asyncio.create_task(stop_flag.wait(), name="stop")
        done, pending = await asyncio.wait(
            [stop_flag_task, produce_logs_task], return_when=asyncio.FIRST_COMPLETED
        )

        with contextlib.suppress(Exception):
            completed_task = done.pop()
            if (
                completed_task.get_name() == "stop"
                and self.__process.returncode is None
            ):
                await self.__stopper.stop(self, self.__process)

        for task in pending:
            task.cancel()

        if self.__process.returncode is None:
            self.__process.kill()


class CommandProcessStopper:
    def __init__(self, command: str, timeout: int = 30, encoding: str = "utf-8"):
        self.__command = command
        self.__timeout = timeout
        self.__encoding = encoding

    async def stop(
        self, writter: ProcessWriter, process: asyncio.subprocess.Process
    ) -> None:
        with contextlib.suppress(Exception):
            await writter.write(self.__command, encoding=self.__encoding)
            await asyncio.wait_for(process.wait(), self.__timeout)


class SignalProcessStopper:
    def __init__(self, signal: int, timeout: int = 30):
        self.__signal = signal
        self.__timeout = timeout

    async def stop(
        self, writter: ProcessWriter, process: asyncio.subprocess.Process
    ) -> None:
        with contextlib.suppress(Exception):
            process.send_signal(self.__signal)
            await asyncio.wait_for(process.wait(), self.__timeout)
