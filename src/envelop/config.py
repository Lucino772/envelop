from __future__ import annotations

from typing import Any

import yaml
from pydantic import BaseModel, Field, model_validator
from pydantic_settings import BaseSettings


class MutuallyExclusiveError(Exception):
    def __init__(self, *args) -> None:
        names = ",".join([f"'{val!s}'" for val in args])
        super().__init__(f"The following attributes are mutually exclusive: {names}")


class _ProcessGracefulConfig(BaseModel):
    timeout: int = Field(default=30)
    signal: int | None = None
    cmd: str | None = None

    @model_validator(mode="before")
    @classmethod
    def mutually_exclusive(cls, data: Any):
        if (
            isinstance(data, dict)
            and data.get("signal", None) is not None
            and data.get("cmd", None) is not None
        ):
            raise MutuallyExclusiveError("signal", "cmd")  # noqa: EM101
        return data


class _ProcessConfig(BaseModel):
    command: str
    env: dict[str, Any] = Field(default_factory=dict)
    graceful: _ProcessGracefulConfig


class _ModuleConfig(BaseModel):
    uses: str
    config: dict[str, Any] = Field(alias="with", default_factory=dict)


class Config(BaseSettings):
    process: _ProcessConfig
    modules: list[_ModuleConfig]

    @classmethod
    def from_file(cls, filepath: str) -> Config:
        with open(filepath, "rb") as fp:
            data = yaml.safe_load(fp)
        return cls(**data)
