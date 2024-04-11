from typing import IO, Literal


class NotEnoughBytesError(Exception):
    def __init__(self, expected: int, current: int) -> None:
        msg = f"Expected {expected} bytes, got {current} bytes"
        super().__init__(msg)


def read_exactly(buffer: IO[bytes], n: int) -> bytes:
    data = buffer.read(n)
    if len(data) != n:
        raise NotEnoughBytesError(n, len(data))
    return data


def read_null_terminated_string(buffer: IO[bytes], encoding: str = "utf-8") -> str:
    res = b""
    char = read_exactly(buffer, 1)
    while char != b"\0":
        res += char
        char = read_exactly(buffer, 1)

    return res.decode(encoding)


def write_varint(buffer: IO[bytes], value: int) -> None:
    buff = []
    while True:
        byte = value & 0x7F
        value >>= 7
        if (value == 0 and byte & 0x40 == 0) or (value == -1 and byte & 0x40 != 0):
            buff.append(byte)
            break
        buff.append(0x80 | byte)

    buffer.write(bytearray(buff))


def read_varint(buffer: IO[bytes], size: int = 5) -> int:
    buff = bytearray()
    while True:
        byte = ord(read_exactly(buffer, 1))
        buff.append(byte)
        if (byte & 0x80) == 0:
            break

    if len(buff) > size:
        msg = "VarInt is too big"
        raise RuntimeError(msg)

    val = 0
    for i, byte in enumerate(buff):
        val = val + ((byte & 0x7F) << (i * 7))
    if byte & 0x40 != 0:
        val |= -(1 << (i * 7) + 7)
    return val


def write_string(buffer: IO[bytes], value: str, encoding: str = "utf-8"):
    encoded = value.encode(encoding)
    write_varint(buffer, len(encoded))
    buffer.write(encoded)


def read_string(buffer: IO[bytes], encoding: str = "utf-8") -> str:
    length = read_varint(buffer)
    data = read_exactly(buffer, length)
    return data.decode(encoding)


def read_int(
    buffer: IO[bytes],
    size: int,
    byteorder: Literal["little", "big"] = "big",
    *,
    signed: bool = False,
) -> int:
    data = read_exactly(buffer, size)
    return int.from_bytes(data, byteorder, signed=signed)
