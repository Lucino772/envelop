from __future__ import annotations

from typing import TYPE_CHECKING

from typing_extensions import Any

if TYPE_CHECKING:
    from collections.abc import Mapping, MutableMapping


class MemoryStore:
    def __init__(self) -> None:
        self._data: MutableMapping[str, Mapping[str, Any]] = {}

    async def write(self, key: str, data: Mapping[str, Any]) -> None:
        self._data[key] = data

    async def read(self, key: str) -> Mapping[str, Any]:
        return self._data[key]
