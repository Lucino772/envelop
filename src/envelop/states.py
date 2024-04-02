from __future__ import annotations

import abc
from typing import TYPE_CHECKING, ClassVar

from pydantic import BaseModel, Field
from typing_extensions import Self

from envelop.events import StateUpdate

if TYPE_CHECKING:
    from envelop.types import Context


class _BaseState(BaseModel, abc.ABC):
    name: ClassVar[str]

    async def write(self, ctx: Context) -> None:
        data = self.model_dump()
        await ctx.write_store(self.name, data)
        await ctx.emit_event(StateUpdate(state=self.name, data=data))

    @classmethod
    async def read(cls, ctx: Context) -> Self:
        data = await ctx.read_store(cls.name, {})
        return cls.model_validate(data, from_attributes=False)


class ProcessStatus(_BaseState):
    name = "/process/status"

    description: str = Field(default="Unknown")


class PlayersState(_BaseState):
    name = "/player/list"

    player_count: int
    max_players: int
    players: list[str]
