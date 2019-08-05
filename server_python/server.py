from concurrent import futures
import time
import argparse
import logging
import socket
import grpc

import calculate_pb2
import calculate_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

parser = argparse.ArgumentParser()
parser.add_argument("-p", "--port", help="Server listen port")
args = parser.parse_args()

class Calculator(calculate_pb2_grpc.CalculateServicer):
    def Add(self, request, context):
        logging.debug("Receive %s %s" % (request.a, request.b))
        response = calculate_pb2.AddResponse()
        response.result = request.a + request.b
        response.server_type = calculate_pb2.AddResponse.ServerType.PYTHON
        server_id = calculate_pb2.AddResponse.ServerId()
        server_id.id = socket.gethostname()
        response.server_id.CopyFrom(server_id)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    calculate_pb2_grpc.add_CalculateServicer_to_server(Calculator(), server)
    server.add_insecure_port('[::]:' + args.port)
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