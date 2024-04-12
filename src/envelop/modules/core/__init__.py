# ruff: noqa: ARG001
from envelop.modules.core.services import PlayersService, ProcessService, SystemService
from envelop.types import Builder, Context


def register(builder: Builder, context: Context, config: dict) -> Builder:
    return (
        builder.add_service(ProcessService(context))
        .add_service(SystemService(context))
        .add_service(PlayersService(context))
    )
