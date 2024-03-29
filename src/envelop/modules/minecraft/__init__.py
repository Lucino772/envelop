# ruff: noqa: ARG001
from envelop.modules.minecraft.tasks import MinecraftStatusTask
from envelop.types import Builder, Context


def register(builder: Builder, context: Context, config: dict) -> Builder:
    return builder.add_task(MinecraftStatusTask(context))
