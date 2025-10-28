import argparse
import json
import grpc
import logic_pb2, logic_pb2_grpc

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--addr", default="localhost:9002", help="server address host:port")
    ap.add_argument("--expr", default="2 + 3 * x - sqrt(y)", help="expression to evaluate")
    ap.add_argument("--vars", default='{"x":4, "y":9}', help='JSON map of variables, e.g. \'{"x":4,"y":9}\'')
    args = ap.parse_args()

    try:
        variables = json.loads(args.vars) if args.vars else {}
        if not isinstance(variables, dict):
            raise ValueError("variables must be a JSON object")
    except Exception as e:
        print("Invalid --vars:", e)
        return

    with grpc.insecure_channel(args.addr) as ch:
        stub = logic_pb2_grpc.LogicServiceStub(ch)
        resp = stub.Evaluate(logic_pb2.EvalRequest(expression=args.expr, variables=variables))
        if resp.error:
            print("Error:", resp.error)
        else:
            print("Result:", resp.result)

if __name__ == "__main__":
    main()
