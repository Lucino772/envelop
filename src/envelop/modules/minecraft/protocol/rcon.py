import asyncio
import enum
import secrets
import struct


class PacketTypes(enum.IntEnum):
    LOGIN = 3
    COMMAND = 2
    RESPONSE = 0


class RconError(Exception):
    pass


class RconAuthenticationError(RconError):
    def __init__(self) -> None:
        super().__init__("Rcon Authentication Failed")


class RconCommandError(RconError):
    def __init__(self, command: str) -> None:
        msg = f"Failed to execute `{command}`"
        super().__init__(msg)


async def _write_packet(writer: asyncio.StreamWriter, _type: int, payload: str) -> int:
    packet_id = secrets.randbelow(2**31)
    packet = struct.pack(
        f"<iii{len(payload) + 2}s",
        10 + len(payload),
        packet_id,
        _type,
        payload.encode("ascii") + b"\00",
    )
    writer.write(packet)
    await writer.drain()
    return packet_id


async def _read_packet(reader: asyncio.StreamReader):
    size_bytes = await reader.readexactly(4)
    size = int.from_bytes(size_bytes, "little")
    data = await reader.readexactly(size)
    packet_id, _type = struct.unpack("<ii", data[:8])
    return packet_id, _type, data[8:], size - 8


async def rcon_send(host: str, port: int, password: str, command: str) -> str:
    try:
        reader, writer = await asyncio.open_connection(host, port)

        # Login
        packet_id = await _write_packet(writer, PacketTypes.LOGIN, password)
        r_packet_id, r_type, _, _ = await _read_packet(reader)
        if r_packet_id != packet_id or r_type != PacketTypes.COMMAND:
            raise RconAuthenticationError

        # Send Command
        packet_id = await _write_packet(writer, PacketTypes.COMMAND, command)
        r_packet_id, r_type, payload, _ = await _read_packet(reader)
        if r_packet_id != packet_id or r_type != PacketTypes.RESPONSE:
            raise RconCommandError(command)

        return payload.strip(b"\0").decode("ascii")
    finally:
        writer.close()
        await writer.wait_closed()
