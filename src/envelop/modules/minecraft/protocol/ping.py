import asyncio
import enum

from envelop.modules.minecraft.protocol.slp import (
    AbstractProtocol,
    ProtocolFE,
    ProtocolFE01,
    ProtocolNetty,
    SLPResponse,
)


class PingVersion(enum.IntFlag):
    """
    :param int V1_7: Use the ping protocol for version 1.7 and higher
    :param int V1_6: Use the ping protocol for version 1.6
    :param int V1_4: Use the ping protocol for version 1.4 and 1.5
    :param int V1_3: Use the ping protocol for version 1.3 and lower
    :param int V_ALL: Use all the ping protocol
    """

    V1_7 = 1
    V1_6 = 2
    V1_4 = 4
    V1_3 = 8
    V_ALL = 15


async def ping(
    host: str,
    port: int,
    timeout: int = 3,
    flags: int = PingVersion.V_ALL,
) -> SLPResponse | None:
    protocols: list[AbstractProtocol] = []
    if PingVersion.V1_7 & flags:
        protocols.append(ProtocolNetty(host, port))
    if PingVersion.V1_6 & flags:
        protocols.append(ProtocolFE01(host, port))
    if PingVersion.V1_4 & flags:
        protocols.append(ProtocolFE01())
    if PingVersion.V1_3 & flags:
        protocols.append(ProtocolFE())

    response = None
    while len(protocols) > 0 and response is None:
        protocol = protocols.pop(0)
        reader, writer = await asyncio.open_connection(host, port)
        try:
            response = await asyncio.wait_for(_ping(reader, writer, protocol), timeout)
        except asyncio.TimeoutError:
            response = None

    return response


async def _ping(
    reader: asyncio.StreamReader,
    writer: asyncio.StreamWriter,
    protocol: AbstractProtocol,
) -> SLPResponse:
    protocol.init()

    while not protocol.is_done():
        data = protocol.get_data_to_send()
        if len(data) > 0:
            writer.write(data)
            await writer.drain()

        data = await reader.read(1024)
        protocol.feed_data(data)

    return protocol.get_response()
