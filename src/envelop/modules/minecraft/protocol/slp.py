import contextlib
import io
import time
from abc import ABCMeta, abstractmethod
from contextlib import contextmanager
from typing import NamedTuple

from envelop.modules.minecraft.protocol.packets import (
    read_fe01_response,
    read_fe_response,
    read_netty_ping_response,
    read_netty_status_response,
    write_fe01_request,
    write_fe_request,
    write_netty_handshake,
    write_netty_ping_request,
    write_netty_status_request,
)
from envelop.modules.minecraft.protocol.types import (
    NotEnoughBytesError,
    read_exactly,
    read_varint,
)


class SLPResponsePlayers(NamedTuple):
    total: tuple[int, int]
    names: list[str]


class SLPResponse(NamedTuple):
    protocol_version: int
    version: str
    motd: str
    players: SLPResponsePlayers
    ping: float


class AbstractProtocol(metaclass=ABCMeta):
    def __init__(self) -> None:
        self._rbuffer = bytearray()
        self._wbuffer = bytearray()
        self._done = False

    def is_done(self) -> bool:
        return self._done

    def get_data_to_send(self) -> bytes:
        data = self._wbuffer[:]
        self._wbuffer = bytearray()
        return data

    @abstractmethod
    def init(self) -> None:
        raise NotImplementedError

    @abstractmethod
    def get_response(self) -> SLPResponse:
        raise NotImplementedError

    @abstractmethod
    def _parse_packet(self) -> None:
        raise NotImplementedError

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


class ProtocolNetty(AbstractProtocol):
    def __init__(self, hostname: str, port: int) -> None:
        super().__init__()
        self._hostname = hostname
        self._port = port
        self._response = None
        self._ping_start = None

    def init(self) -> None:
        with self._writer() as buffer:
            write_netty_handshake(buffer, self._hostname, self._port)
            write_netty_status_request(buffer)

    def _parse_packet(self) -> None:
        with self._reader() as buffer:
            packet_len = read_varint(buffer)
            packet_data = read_exactly(buffer, packet_len)

        with io.BytesIO(packet_data) as buffer:
            packet_id = read_varint(buffer)
            if packet_id == 0x00:
                self._response = read_netty_status_response(buffer)
                with self._writer() as buff:
                    write_netty_ping_request(buff, int(time.time() * 1000))
            elif packet_id == 0x01:
                self._ping_start = read_netty_ping_response(buffer)
                self._done = True

    def get_response(self) -> SLPResponse:
        if self._ping_start is None or self._response is None:
            raise ValueError

        players = SLPResponsePlayers((-1, -1), [])
        if "players" in self._response:
            players = SLPResponsePlayers(
                (
                    int(self._response["players"]["online"]),
                    int(self._response["players"]["max"]),
                ),
                self._response["players"].get("sample", []),
            )

        return SLPResponse(
            protocol_version=self._response.get("version", {}).get(
                "protocol", "unknown"
            ),
            version=self._response.get("version", {}).get("name", "unknown"),
            motd=self._response.get("description", None),
            players=players,
            ping=(time.time() * 1000) - self._ping_start,
        )


class ProtocolFE01(AbstractProtocol):
    def __init__(self, hostname: str | None = None, port: int | None = None) -> None:
        super().__init__()
        self._hostname = hostname
        self._port = port
        self._response = None
        self._ping_start = None

    def init(self) -> None:
        with self._writer() as buffer:
            write_fe01_request(buffer, self._hostname, self._port)
        self._ping_start = time.time() * 1000

    def _parse_packet(self) -> None:
        with self._reader() as buffer:
            self._response = read_fe01_response(buffer)
            self._done = True

    def get_response(self) -> SLPResponse:
        if self._ping_start is None or self._response is None:
            raise ValueError

        return SLPResponse(
            protocol_version=self._response[0],
            version=self._response[1],
            motd=self._response[2],
            players=SLPResponsePlayers(
                total=(self._response[3], self._response[4]), names=[]
            ),
            ping=(time.time() * 1000) - self._ping_start,
        )


class ProtocolFE(AbstractProtocol):
    def __init__(self) -> None:
        super().__init__()
        self._response = None
        self._ping_start = None

    def init(self) -> None:
        with self._writer() as buffer:
            write_fe_request(buffer)
        self._ping_start = time.time() * 1000

    def _parse_packet(self) -> None:
        with self._reader() as buffer:
            self._response = read_fe_response(buffer)
            self._done = True

    def get_response(self) -> SLPResponse:
        if self._ping_start is None or self._response is None:
            raise ValueError

        return SLPResponse(
            protocol_version=-1,
            version="",
            motd=self._response[0],
            players=SLPResponsePlayers((self._response[1], self._response[2]), []),
            ping=(time.time() * 1000) - self._ping_start,
        )
