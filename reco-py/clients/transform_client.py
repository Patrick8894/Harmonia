import argparse
import json
import grpc
import logic_pb2, logic_pb2_grpc

OP_NAME_TO_ENUM = {
    "MAP": logic_pb2.MAP,
    "FILTER": logic_pb2.FILTER,
    "SUM": logic_pb2.SUM,
}

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--addr", default="localhost:9002", help="server address host:port")
    ap.add_argument("--data", default="[1,2,3,4,5,6]", help="JSON array of numbers")
    ap.add_argument("--expr", default="x % 2 == 0", help="per-element expression")
    ap.add_argument("--var",  default="x", help="per-element variable name")
    ap.add_argument("--op",   default="FILTER", choices=OP_NAME_TO_ENUM.keys(), help="operation")
    args = ap.parse_args()

    try:
        data = json.loads(args.data) if args.data else []
        if not isinstance(data, list):
            raise ValueError("data must be a JSON array")
        data = [float(x) for x in data]
    except Exception as e:
        print("Invalid --data:", e)
        return

    with grpc.insecure_channel(args.addr) as ch:
        stub = logic_pb2_grpc.LogicServiceStub(ch)
        resp = stub.Transform(logic_pb2.TransformRequest(
            data=data,
            expr=args.expr,
            var_name=args.var,
            op=OP_NAME_TO_ENUM[args.op],
        ))
        if resp.error:
            print("Error:", resp.error)
        else:
            if args.op == "SUM":
                print("Sum:", resp.result)
            else:
                print("Data:", list(resp.data))

if __name__ == "__main__":
    main()
