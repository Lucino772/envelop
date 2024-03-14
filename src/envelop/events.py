import abc
import uuid
from collections.abc import Mapping
from typing import Any, ClassVar

from pydantic import BaseModel, Field


class _BaseEvent(BaseModel, abc.ABC):
    name: ClassVar[str]
    uid: str = Field(default_factory=lambda: uuid.uuid4().hex, init=False)

    def get_name(self) -> str:
        return self.name

    def get_uid(self) -> str:
        return self.uid

    def get_data(self) -> Mapping[str, Any]:
        return self.model_dump()


class ProcessLog(_BaseEvent):
    name = "/process/log"

    log: str
