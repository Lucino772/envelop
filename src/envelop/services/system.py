# ruff: noqa: N802
from __future__ import annotations

from abc import ABCMeta, abstractmethod
from typing import TYPE_CHECKING

from envelop.services.proto.system_pb2_grpc import add_SystemServicer_to_server

if TYPE_CHECKING:
    from collections.abc import AsyncIterator

    import grpc
    from google.protobuf.empty_pb2 import Empty

    from envelop.services.proto.system_pb2 import Event
    from envelop.types import SupportsGenericRpcHandlers


class AbstractSystemService(metaclass=ABCMeta):
    @abstractmethod
    def StreamEvents(
        self, request: Empty, context: grpc.ServicerContext
    ) -> AsyncIterator[Event]:
        raise NotImplementedError

    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        add_SystemServicer_to_server(self, server)
