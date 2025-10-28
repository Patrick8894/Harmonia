from typing import Dict, List
import math

import logic_pb2
import logic_pb2_grpc
from logic_service.evaluators import SafeExpressionEvaluator
from logic_service.errors import LogicError
from logic_service.transforms import transform_map, transform_filter, transform_sum
from logic_service.planners import plan_tasks


class LogicService(logic_pb2_grpc.LogicServiceServicer):
    def Hello(self, request, context):
        name = request.name or "there"
        print(f"[Logic] Hello request name={name!r}")
        return logic_pb2.HelloReply(message=f"Hello, {name} from Python LogicService!")

    def Evaluate(self, request, context):
        expr: str = (request.expression or "").strip()
        vars_dict: Dict[str, float] = dict(request.variables)
        print(f"[Logic] Evaluate expr={expr!r} vars={vars_dict}")

        if not expr:
            return logic_pb2.EvalReply(error="expression is empty")

        try:
            result = SafeExpressionEvaluator(vars_dict).evaluate(expr)
            if math.isnan(result) or math.isinf(result):
                return logic_pb2.EvalReply(error="result is not finite (NaN/Inf)")
            return logic_pb2.EvalReply(result=float(result))
        except LogicError as le:
            return logic_pb2.EvalReply(error=str(le))
        except ZeroDivisionError:
            return logic_pb2.EvalReply(error="division by zero")
        except Exception as e:
            return logic_pb2.EvalReply(error=f"evaluation error: {e}")

    def Transform(self, request, context):
        data: List[float] = list(request.data)
        expr: str = (request.expr or "").strip()
        var_name: str = (request.var_name or "x").strip() or "x"
        op = request.op

        print(f"[Logic] Transform op={op} expr={expr!r} var={var_name!r} data_len={len(data)}")

        if not expr:
            return logic_pb2.TransformReply(error="expr is empty")

        try:
            if op == logic_pb2.MAP:
                out = transform_map(data, expr, var_name)
                return logic_pb2.TransformReply(data=out)
            elif op == logic_pb2.FILTER:
                out = transform_filter(data, expr, var_name)
                return logic_pb2.TransformReply(data=out)
            elif op == logic_pb2.SUM:
                s = transform_sum(data, expr, var_name)
                return logic_pb2.TransformReply(result=s)
            else:
                return logic_pb2.TransformReply(error="unsupported op")
        except LogicError as le:
            return logic_pb2.TransformReply(error=str(le))
        except ZeroDivisionError:
            return logic_pb2.TransformReply(error="division by zero")
        except Exception as e:
            return logic_pb2.TransformReply(error=f"transform error: {e}")

    def PlanTasks(self, request, context):
        goal = (request.goal or "").strip()
        hints = list(request.hints)
        max_steps = int(request.max_steps) if request.max_steps else None

        print(f"[Logic] PlanTasks goal={goal!r} hints={hints} max_steps={max_steps}")

        if not goal:
            return logic_pb2.PlanReply(error="goal is empty")

        try:
            tasks, notes = plan_tasks(goal, hints, max_steps)
            return logic_pb2.PlanReply(
                tasks=[
                    logic_pb2.Task(
                        id=t["id"],
                        title=t["title"],
                        detail=t["detail"],
                        priority=t["priority"],
                        estimate_min=t["estimate_min"],
                        depends_on=t["depends_on"],
                    )
                    for t in tasks
                ],
                notes=notes,
            )
        except LogicError as le:
            return logic_pb2.PlanReply(error=str(le))
        except Exception as e:
            return logic_pb2.PlanReply(error=f"planner error: {e}")
