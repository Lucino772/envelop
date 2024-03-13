from __future__ import annotations

from abc import ABCMeta, abstractmethod

import grpc
from google.protobuf.empty_pb2 import Empty
from typing_extensions import AsyncIterator

from envelop.services.proto.process_pb2 import Command, Log
from envelop.services.proto.process_pb2_grpc import add_ProcessServicer_to_server
from envelop.types import SupportsGenericRpcHandlers


class AbstractProcessService(metaclass=ABCMeta):
    @abstractmethod
    def WriteCommand(self, request: Command, context: grpc.ServicerContext) -> Empty:
        raise NotImplementedError

    @abstractmethod
    def StreamLogs(
        self, request: Empty, context: grpc.ServicerContext
    ) -> AsyncIterator[Log]:
        raise NotImplementedError

    def add_rpc_handlers(self, server: SupportsGenericRpcHandlers) -> None:
        add_ProcessServicer_to_server(self, server)
