import grpc
import logic_pb2, logic_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:9002')
    stub = logic_pb2_grpc.LogicServiceStub(channel)
    response = stub.Hello(logic_pb2.HelloRequest(name="TestUser"))
    print("Server replied:", response.message)

if __name__ == "__main__":
    run()