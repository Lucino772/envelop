from __future__ import annotations

import uuid
from dataclasses import dataclass, field

from typing_extensions import Any, ClassVar, Mapping, TypedDict


@dataclass
class _BaseEvent:
    name: ClassVar[str]
    uid: str = field(default_factory=lambda: str(uuid.uuid4()), init=False)
    data: Mapping[str, Any] = field(init=False, default_factory=dict)

    def get_uid(self) -> str:
        return self.uid

    def get_name(self) -> str:
        return self.name

    def get_data(self) -> Mapping[str, Any]:
        return self.data


@dataclass
class ProcessLog(_BaseEvent):
    LogData = TypedDict("LogData", {"log": str})
    name = "/process/log"

    data: LogData = field()
