import asyncio
import os

from envelop.app import AppBuilder
from envelop.config import Config
from envelop.modules import core as envelop_core_module
from envelop.types import Module

registry: dict[str, Module] = {"envelop.core": envelop_core_module}


async def _run():
    config = Config.from_file(os.path.join(os.getcwd(), "envelop.yaml"))
    app = AppBuilder().build(config, registry)
    await app.run()


def cli():
    asyncio.run(_run())
