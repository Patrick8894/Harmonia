import grpc
from concurrent import futures
import time

import logic_pb2, logic_pb2_grpc

class LogicService(logic_pb2_grpc.LogicServiceServicer):
    def Hello(self, request, context):
        """
        Implements the Hello RPC defined in logic.proto.
        """
        print(f"[Reco] Received Hello request for name={request.name}")
        message = f"Hello, {request.name} from Python LogicService!"
        return logic_pb2.HelloReply(message=message)

def serve():
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