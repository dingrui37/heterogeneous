# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import calculate_pb2 as calculate__pb2


class CalculateStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Add = channel.unary_unary(
        '/proto.Calculate/Add',
        request_serializer=calculate__pb2.AddRequest.SerializeToString,
        response_deserializer=calculate__pb2.AddResponse.FromString,
        )


class CalculateServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Add(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CalculateServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Add': grpc.unary_unary_rpc_method_handler(
          servicer.Add,
          request_deserializer=calculate__pb2.AddRequest.FromString,
          response_serializer=calculate__pb2.AddResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'proto.Calculate', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))