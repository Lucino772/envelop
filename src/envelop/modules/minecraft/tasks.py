import asyncio

from envelop.modules.minecraft.protocol.query import query_stats
from envelop.states import PlayersState, ProcessStatus
from envelop.types import Context


class MinecraftStatusTask:
    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx

    async def run(self) -> None:
        async for log in self._ctx.iter_logs():
            state = await ProcessStatus.read(self._ctx)
            if "done" in log.lower() and state.description != "Ready":
                await ProcessStatus(description="Ready").write(self._ctx)
            elif "stopping" in log.lower() and state.description != "Stopping":
                await ProcessStatus(description="Stopping").write(self._ctx)
            elif state.description == "Unknown":
                await ProcessStatus(description="Starting").write(self._ctx)


class PlayersTask:
    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx
        self._server_ready = asyncio.Event()
        self._query_ready = asyncio.Event()

    async def _process_logs(self):
        async for log in self._ctx.iter_logs():
            if "done" in log.lower() and not self._server_ready.is_set():
                self._server_ready.set()
            elif "query running on" in log.lower() and not self._query_ready.is_set():
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

    async def run(self) -> None:
        query_ready = await self._wait_query_ready()
        if query_ready is False:
            return

        try:
            while True:
                response = await query_stats("127.0.0.1", 25565)
                if response is not None:
                    await PlayersState(
                        player_count=response.players[0],
                        max_players=response.players[1],
                        players=response.player_list,
                    ).write(self._ctx)
                await asyncio.sleep(30)
        except ConnectionError:
            pass
