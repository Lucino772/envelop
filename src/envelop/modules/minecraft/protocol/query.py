from __future__ import annotations

import asyncio
import contextlib
import io
import struct
import time
from contextlib import contextmanager
from typing import IO, Any, NamedTuple

from envelop.modules.minecraft.protocol.types import (
    NotEnoughBytesError,
    read_exactly,
    read_null_terminated_string,
)


async def query_stats(host: str, port: int, timeout: int = 3) -> ServerStats | None:
    try:
        return await asyncio.wait_for(_query_stats(host, port), timeout)
    except asyncio.TimeoutError:
        return None


class QueryClientProtocol(asyncio.DatagramProtocol):
    def __init__(self, future: asyncio.Future[ServerStats | None]):
        self._future = future
        self._protocol = QueryProtocol()
        self.transport = None

    def _send_data(self, transport: asyncio.DatagramTransport):
        data_to_send = self._protocol.get_data_to_send()
        if len(data_to_send):
            transport.sendto(data_to_send)

    def connection_made(self, transport: asyncio.DatagramTransport) -> None:
        self.transport = transport
        self._protocol.init()
        self._send_data(self.transport)

    def datagram_received(self, data: bytes, _: tuple[str | Any, int]) -> None:
        if self.transport is None:
            return

        self._protocol.feed_data(data)
        self._send_data(self.transport)
        if self._protocol.is_done():
            self._future.set_result(self._protocol.get_stats())
            self.transport.close()

    def error_received(self, exc: Exception) -> None:
        self._future.set_exception(exc)

    def connection_lost(self, exc: Exception | None) -> None:
        if self._future.done():
            return

        if exc is not None:
            self._future.set_exception(exc)
        else:
            self._future.set_result(None)


async def _query_stats(host: str, port: int) -> ServerStats | None:
    loop = asyncio.get_running_loop()

    fut: asyncio.Future[ServerStats | None] = loop.create_future()
    transport, _ = await loop.create_datagram_endpoint(
        lambda: QueryClientProtocol(fut), remote_addr=(host, port)
    )

    try:
        return await fut
    finally:
        transport.close()


class ServerStats(NamedTuple):
    motd: str
    game_type: str
    game_id: str
    version: str
    map_name: str
    host: tuple[str, int]
    plugins: list[str]
    players: tuple[int, int]
    player_list: list[str]


class QueryProtocol:
    _PKT_TYPE_HANDSHAKE = 9
    _PKT_TYPE_STAT = 0

    def __init__(self) -> None:
        self._wbuffer = bytearray()
        self._rbuffer = bytearray()
        self._session_id = int(time.time()) & 0x0F0F0F0F
        self._token = None
        self._stats = None
        self._done = False

    def is_done(self) -> bool:
        return self._done

    def get_data_to_send(self) -> bytes:
        data = self._wbuffer[:]
        self._wbuffer = bytearray()
        return data

    def get_stats(self) -> ServerStats:
        if self._stats is None:
            raise ValueError
        return self._stats

    def _write_packet(self, __type: int, __sess_id: int, data: bytes = b"") -> None:
        with self._writer() as buffer:
            data = struct.pack(f">Hbi{len(data)}s", 0xFEFD, __type, __sess_id, data)
            buffer.write(data)

    def _parse_packet_header(self, buffer: IO[bytes]) -> tuple[int, int]:
        _type, sess_id = struct.unpack(">bi", read_exactly(buffer, 5))
        return _type, sess_id

    def _parse_stats(self, buffer: IO[bytes]) -> ServerStats:
        _ = read_exactly(buffer, 11)

        # Read server info
        info = {}
        for _ in range(10):
            key = read_null_terminated_string(buffer)
            value = read_null_terminated_string(buffer)
            info[key] = value

        motd = str(info.pop("hostname"))
        game_type = str(info.pop("gametype"))
        game_id = str(info.pop("game_id"))
        version = str(info.pop("version"))
        map_name = str(info.pop("map"))
        host = (str(info.pop("hostip")), int(info.pop("hostport")))
        # TODO: Parse plugins
        plugins: list[str] = []
        players = (
            int(info.pop("numplayers")),
            int(info.pop("maxplayers")),
        )

        _ = read_exactly(buffer, 11)

        # Read players
        player_list = []
        player_name = read_null_terminated_string(buffer)
        while len(player_name) != 0:
            player_list.append(player_name)
            player_name = read_null_terminated_string(buffer)

        return ServerStats(
            motd=motd,
            game_type=game_type,
            game_id=game_id,
            version=version,
            map_name=map_name,
            host=host,
            plugins=plugins,
            players=players,
            player_list=player_list,
        )

    def _parse_packet(self) -> None:
        with self._reader() as buffer:
            _type, sess_id = self._parse_packet_header(buffer)
            if sess_id != self._session_id:
                self._done = True
            elif _type == self._PKT_TYPE_HANDSHAKE:
                self._token = int(read_null_terminated_string(buffer))
                self._write_packet(
                    self._PKT_TYPE_STAT,
                    self._session_id,
                    struct.pack(">iI", self._token, 0xFFFFFF01),
                )
            elif _type == self._PKT_TYPE_STAT:
                self._stats = self._parse_stats(buffer)
                self._done = True

    def init(self) -> None:
        self._write_packet(self._PKT_TYPE_HANDSHAKE, self._session_id)

    def feed_data(self, data: bytes) -> None:
        self._rbuffer += data
        with contextlib.suppress(NotEnoughBytesError):
            self._parse_packet()

    @contextmanager
    def _writer(self):
        buffer = io.BytesIO()
        try:
            yield buffer
        finally:
            self._wbuffer += buffer.getvalue()
            buffer.close()

    @contextmanager
    def _reader(self):
        buffer = io.BytesIO(self._rbuffer)
        try:
            start_index = buffer.tell()
            yield buffer
            bytes_read = buffer.tell() - start_index
            self._rbuffer = self._rbuffer[bytes_read:]
        finally:
            buffer.close()
