import grpc
from concurrent import futures
import time

# Import generated classes
import logic_pb2, logic_pb2_grpc

# Implement the service defined in logic.proto
class LogicService(logic_pb2_grpc.LogicServiceServicer):
    def Recommend(self, request, context):
        """
        Example RPC method that returns a list of recommended items.
        """
        print(f"[Reco] Received request for user_id={request.user_id}, limit={request.limit}")
        # Dummy response data
        recommendations = [f"item_{i}" for i in range(1, request.limit + 1)]
        return logic_pb2.RecommendResponse(item_ids=recommendations)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logic_pb2_grpc.add_LogicServiceServicer_to_server(LogicService(), server)

    # Run on all interfaces (for Docker / Go gateway access)
    server.add_insecure_port("[::]:9002")
    server.start()
    print("✅ Python gRPC LogicService running on port 9002")
    try:
        while True:
            time.sleep(86400)  # keep alive
    except KeyboardInterrupt:
        server.stop(0)
        print("🛑 Server stopped")

if __name__ == "__main__":
    serve()
