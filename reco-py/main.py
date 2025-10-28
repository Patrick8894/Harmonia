import time
from concurrent import futures
import grpc

import logic_pb2_grpc
from logic_service.service import LogicService
# (optional) from logic_service.logging_setup import configure_logging

def serve() -> None:
    # configure_logging()  # if you want structured logs
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logic_pb2_grpc.add_LogicServiceServicer_to_server(LogicService(), server)
    server.add_insecure_port("[::]:9002")
    server.start()
    print("✅ Python gRPC LogicService running on port 9002")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)
        print("🛑 Server stopped")

if __name__ == "__main__":
    serve()
