from __future__ import annotations

from abc import ABCMeta, abstractmethod

import grpc
from google.protobuf.empty_pb2 import Empty
from typing_extensions import AsyncIterator

from envelop.services.proto.system_pb2 import Event
from envelop.services.proto.system_pb2_grpc import add_SystemServicer_to_server
from envelop.types import SupportsGenericRpcHandlers


class AbstractSystemService(metaclass=ABCMeta):
    @abstractmethod
    def StreamEvents(
        self, request: Empty, context: grpc.ServicerContext
    ) -> AsyncIterator[Event]:
        raise NotImplementedError

    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        add_SystemServicer_to_server(self, server)
