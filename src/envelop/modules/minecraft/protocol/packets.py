import io
import json
import struct
from typing import IO

from envelop.modules.minecraft.protocol.types import (
    read_exactly,
    read_int,
    read_string,
    write_string,
    write_varint,
)


# Post-Netty
def write_netty_handshake(buffer: IO[bytes], hostname: str, port: int) -> None:
    with io.BytesIO() as buff:
        write_varint(buff, 0x00)
        write_varint(buff, -1)
        write_string(buff, hostname)
        buff.write(struct.pack(">H", port))
        write_varint(buff, 0x01)
        packet_data = buff.getvalue()

    write_varint(buffer, len(packet_data))
    buffer.write(packet_data)


def write_netty_status_request(buffer: IO[bytes]) -> None:
    with io.BytesIO() as buff:
        write_varint(buff, 0x00)
        packet_data = buff.getvalue()

    write_varint(buffer, len(packet_data))
    buffer.write(packet_data)


def write_netty_ping_request(buffer: IO[bytes], payload: int) -> None:
    with io.BytesIO() as buff:
        write_varint(buff, 0x01)
        buff.write(struct.pack(">q", payload))
        packet_data = buff.getvalue()

    write_varint(buffer, len(packet_data))
    buffer.write(packet_data)


def read_netty_status_response(buffer: IO[bytes]) -> dict:
    data = read_string(buffer)
    return json.loads(data)


def read_netty_ping_response(buffer: IO[bytes]) -> int:
    return read_int(buffer, 8)


# Pre-Netty
def write_fe01_request(
    buffer: IO[bytes], hostname: str | None = None, port: int | None = None
) -> int:
    if hostname is not None and len(hostname) > 0 and port is not None:
        encoded_hostname = hostname.encode("utf-16be")
        packet = struct.pack(
            f">BBBh22shBh{len(encoded_hostname)}si",
            0xFE,  # Packet ID
            0x01,  # Packet Payload (always 1)
            0xFA,  # Packet identifier for a plugin message
            11,  # Length of following string (always 11)
            "MC|PingHost".encode("utf-16be"),  # MC|Pingost
            7 + len(encoded_hostname),  # Length of the rest of the data
            0x4F,  # Protocol version (80)
            len(hostname),
            encoded_hostname,
            port,  # Port
        )
    else:
        packet = struct.pack(">BB", 0xFE, 0x01)

    return buffer.write(packet)


def read_fe01_response(buffer: IO[bytes]):
    # TODO: Check packet_id and response
    _ = read_exactly(buffer, 9)  # Skip 9 first bytes
    response = buffer.read().decode("utf-16-be").split("\x00")
    proto_version, server_version, motd, player_cnt, max_players = response
    return int(proto_version), server_version, motd, int(player_cnt), int(max_players)


def write_fe_request(buffer: IO[bytes]) -> int:
    return buffer.write(struct.pack(">B", 0xFE))


def read_fe_response(buffer: IO[bytes]):
    # TODO: Check packet_id and response
    _ = read_exactly(buffer, 3)  # Skip 3 first bytes
    response = buffer.read().decode("utf-16-be").split("\xa7")
    motd, player_cnt, max_players = response
    return motd, int(player_cnt), int(max_players)
