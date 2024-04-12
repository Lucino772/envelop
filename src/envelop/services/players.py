# ruff: noqa: N802
from __future__ import annotations

from abc import ABCMeta, abstractmethod
from typing import TYPE_CHECKING

from envelop.services.proto.players_pb2_grpc import add_PlayersServicer_to_server

if TYPE_CHECKING:
    from collections.abc import AsyncIterator, Coroutine

    import grpc
    from google.protobuf.empty_pb2 import Empty

    from envelop.services.proto.players_pb2 import PlayerList
    from envelop.types import SupportsGenericRpcHandlers


class AbstractPlayersService(metaclass=ABCMeta):
    @abstractmethod
    def ListPlayers(
        self, request: Empty, context: grpc.ServicerContext
    ) -> Coroutine[None, None, PlayerList]:
        raise NotImplementedError

    @abstractmethod
    def StreamPlayers(
        self, request: Empty, context: grpc.ServicerContext
    ) -> AsyncIterator[PlayerList]:
        raise NotImplementedError

    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        add_PlayersServicer_to_server(self, server)
