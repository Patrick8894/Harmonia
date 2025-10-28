import argparse
import grpc
import logic_pb2, logic_pb2_grpc

def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--addr", default="localhost:9002", help="server address host:port")
    ap.add_argument("--goal", default="Add a Python gRPC Transform service and containerize it",
                    help="short goal text")
    ap.add_argument("--hint", action="append", default=[],
                    help="keyword hints (repeat flag), e.g. --hint tests --hint logging")
    ap.add_argument("--max_steps", type=int, default=8, help="cap number of steps")
    args = ap.parse_args()

    with grpc.insecure_channel(args.addr) as ch:
        stub = logic_pb2_grpc.LogicServiceStub(ch)
        resp = stub.PlanTasks(logic_pb2.PlanRequest(goal=args.goal, hints=args.hint, max_steps=args.max_steps))
        if resp.error:
            print("Error:", resp.error)
            return

        print("Notes:", resp.notes)
        for t in resp.tasks:
            deps = ",".join(t.depends_on) if t.depends_on else "-"
            print(f"- {t.id} [P{t.priority}] {t.title} ({t.estimate_min}m) deps={deps}")
            if t.detail:
                print(f"  {t.detail}")

if __name__ == "__main__":
    main()
