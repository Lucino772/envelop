import abc
from typing import ClassVar, Self

from pydantic import BaseModel, Field

from envelop.types import Context


class _BaseState(BaseModel, abc.ABC):
    name: ClassVar[str]

    async def write(self, ctx: Context) -> None:
        await ctx.write_store(self.name, self.model_dump())

    @classmethod
    async def read(cls, ctx: Context) -> Self:
        data = await ctx.read_store(cls.name)
        return cls.model_construct(**data)


class ProcessStatus(_BaseState):
    name = "/process/status"

    description: str = Field(default="Unknown")
