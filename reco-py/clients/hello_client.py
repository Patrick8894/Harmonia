import argparse
import grpc
import logic_pb2, logic_pb2_grpc

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--addr", default="localhost:9002", help="server address host:port")
    ap.add_argument("--name", default="Patrick", help="name for Hello")
    args = ap.parse_args()

    with grpc.insecure_channel(args.addr) as ch:
        stub = logic_pb2_grpc.LogicServiceStub(ch)
        resp = stub.Hello(logic_pb2.HelloRequest(name=args.name))
        print(resp.message)

if __name__ == "__main__":
    main()
