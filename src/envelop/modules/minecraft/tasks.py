import asyncio

from envelop.modules.minecraft.protocol.ping import ping
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

    async def _wait_port(self, port: int, delay: float = 5) -> None:
        done = False
        while not done:
            try:
                _, writer = await asyncio.open_connection("127.0.0.1", port)
                writer.close()
                await writer.wait_closed()
                done = True
            except ConnectionError:
                await asyncio.sleep(delay)

    async def run(self) -> None:
        await self._wait_port(25565)

        try:
            while True:
                response = await ping("127.0.0.1", 25565)
                if response is not None:
                    await PlayersState(
                        player_count=response.players.total[0],
                        max_players=response.players.total[1],
                        players=response.players.names,
                    ).write(self._ctx)
                await asyncio.sleep(30)
        except ConnectionError:
            pass
