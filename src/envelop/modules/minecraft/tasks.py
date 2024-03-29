from envelop.states import ProcessStatus
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
