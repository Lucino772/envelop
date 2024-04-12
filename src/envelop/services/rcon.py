# ruff: noqa: N802
from __future__ import annotations

from abc import ABCMeta, abstractmethod
from typing import TYPE_CHECKING

from envelop.services.proto.rcon_pb2_grpc import add_RconServicer_to_server

if TYPE_CHECKING:
    from collections.abc import Coroutine

    import grpc

    from envelop.services.proto.rcon_pb2 import RconCommand, RconResponse
    from envelop.types import SupportsGenericRpcHandlers


class AbstractRconService(metaclass=ABCMeta):
    @abstractmethod
    def SendCommand(
        self, request: RconCommand, context: grpc.ServicerContext
    ) -> Coroutine[None, None, RconResponse]:
        raise NotImplementedError

    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        add_RconServicer_to_server(self, server)
