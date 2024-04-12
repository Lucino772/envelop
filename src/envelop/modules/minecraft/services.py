# ruff: noqa: N802, ARG002
import os
from configparser import ConfigParser
from typing import final

from grpc import ServicerContext

from envelop.modules.minecraft.protocol.rcon import rcon_send
from envelop.services.proto.rcon_pb2 import RconCommand, RconResponse
from envelop.services.rcon import AbstractRconService
from envelop.types import Context


@final
class MinecraftRconService(AbstractRconService):
    def __init__(self, ctx: Context) -> None:
        self._ctx = ctx

    def _get_settings(self) -> tuple[bool, int, str]:
        settings_path = os.path.join(os.getcwd(), "server.properties")

        with open(settings_path) as fp:
            content = "[settings]\n" + fp.read()

        config = ConfigParser()
        config.read_string(content)

        settings = config["settings"]
        rcon_port = settings.getint("rcon.port")
        rcon_password = settings.get("rcon.password")
        rcon_enabled = settings.getboolean("enable-rcon")
        return rcon_enabled, rcon_port, rcon_password

    async def SendCommand(
        self, request: RconCommand, context: ServicerContext
    ) -> RconResponse:
        enabled, port, password = self._get_settings()
        if not enabled:
            return RconResponse(value="error")  # TODO: Return error

        resp = await rcon_send("localhost", port, password, request.value)
        return RconResponse(value=resp)
