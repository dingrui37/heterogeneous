from concurrent import futures
import time
import logging

import grpc

import calculate_pb2
import calculate_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class Calculator(calculate_pb2_grpc.CalculateServicer):
    def Add(self, request, context):
        logging.debug("Receive %s %s" % (request.a, request.b))
        response = calculate_pb2.AddResponse()
        response.result = request.a + request.b
        response.server_type = calculate_pb2.AddResponse.ServerType.PYTHON
        server_id = calculate_pb2.AddResponse.ServerId()
        server_id.id = 2
        response.server_id.CopyFrom(server_id)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    calculate_pb2_grpc.add_CalculateServicer_to_server(Calculator(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(name)s:%(levelname)s:%(message)s", 
                        datefmt="%d-%M-%Y %H:%M:%S", level=logging.DEBUG)
    serve()