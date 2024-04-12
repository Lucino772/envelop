import asyncio
import re

from envelop.modules.minecraft.protocol.query import ServerStats, query_stats
from envelop.states import PlayersState, ProcessStatus
from envelop.types import Context


class MinecraftStatusTask:
    _REGEX_STARTING = (
        r"\[Server thread\/INFO\]\: Starting Minecraft server on \*:(?P<port>[0-9]+)"
    )
    _REGEX_PREPARING = r"\[(.*?)\]: Preparing spawn area: (?P<progress>[0-9]+)%"
    _REGEX_READY = r"\[Server thread\/INFO\]\: Done \((.*?)\)! For help, type \"help\""
    _REGEX_STOPPING = r"\[Server thread\/INFO\]\: Stopping server"

    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx

    async def run(self) -> None:
        async for log in self._ctx.iter_logs():
            if re.search(self._REGEX_STARTING, log) is not None:
                await ProcessStatus(description="Starting").write(self._ctx)
            elif (_match := re.search(self._REGEX_PREPARING, log)) is not None:
                progress = _match.group("progress")
                await ProcessStatus(description=f"Preparing ({progress}%)").write(
                    self._ctx
                )
            elif re.search(self._REGEX_READY, log) is not None:
                await ProcessStatus(description="Ready").write(self._ctx)
            elif re.search(self._REGEX_STOPPING, log) is not None:
                await ProcessStatus(description="Stopping").write(self._ctx)


class PlayersTask:
    _REGEX_READY = r"\[Server thread\/INFO\]\: Done \((.*?)\)! For help, type \"help\""
    _REGEX_QUERY = (
        r"\[Query Listener #1\/INFO\]\: Query running on (.*?):(?P<port>[0-9]+)"
    )

    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx
        self._server_ready = asyncio.Event()
        self._query_ready = asyncio.Event()

    async def _process_logs(self):
        async for log in self._ctx.iter_logs():
            if (
                re.search(self._REGEX_READY, log) is not None
                and not self._server_ready.is_set()
            ):
                self._server_ready.set()
            elif (
                re.search(self._REGEX_QUERY, log) is not None
                and not self._query_ready.is_set()
            ):
                self._query_ready.set()

    async def _wait_query_ready(self, timeout: float = 5) -> bool:
        task = asyncio.create_task(self._process_logs())
        await self._server_ready.wait()

        try:
            await asyncio.wait_for(self._query_ready.wait(), timeout)
        except asyncio.TimeoutError:
            return False
        else:
            return True
        finally:
            task.cancel()

    def __aiter__(self):
        return self

    async def __anext__(self) -> ServerStats | None:
        process_task = asyncio.create_task(self._ctx.wait_process(), name="process")
        query_task = asyncio.create_task(query_stats("127.0.0.1", 25565), name="query")

        done, pending = await asyncio.wait(
            [process_task, query_task], return_when=asyncio.FIRST_COMPLETED
        )

        for task in pending:
            task.cancel()

        completed_task = done.pop()
        if completed_task.get_name() == "process":
            raise StopAsyncIteration

        return completed_task.result()  # type: ignore

    async def run(self) -> None:
        query_ready = await self._wait_query_ready()
        if query_ready is False:
            return

        try:
            async for response in self:
                if response is not None:
                    await PlayersState(
                        player_count=response.players[0],
                        max_players=response.players[1],
                        players=response.player_list,
                    ).write(self._ctx)
                await asyncio.sleep(5)
        except ConnectionError:
            pass
