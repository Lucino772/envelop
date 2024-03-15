from collections.abc import Mapping, MutableMapping
from typing import Any


class MemoryStore:
    def __init__(self) -> None:
        self._data: MutableMapping[str, Mapping[str, Any]] = {}

    async def write(self, key: str, data: Mapping[str, Any]) -> None:
        self._data[key] = data

    async def read(self, key: str) -> Mapping[str, Any]:
        return self._data[key]
