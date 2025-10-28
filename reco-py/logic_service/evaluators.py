from __future__ import annotations
import ast
import math
from typing import Dict, Callable, Any

from logic_service.errors import LogicError


class SafeExpressionEvaluator(ast.NodeVisitor):
    """
    Safely evaluate math expressions with a small whitelist:
      - + - * / // % **, unary +/- , parentheses
      - numeric constants
      - variables (from provided dict)
      - functions: abs, min, max, round, sqrt, pow
    """

    _allowed_funcs: Dict[str, Callable[..., float]] = {
        "abs": abs,
        "min": min,
        "max": max,
        "round": round,
        "sqrt": math.sqrt,
        "pow": pow,
    }

    _allowed_binops: Dict[type, Callable[[float, float], float]] = {
        ast.Add: lambda a, b: a + b,
        ast.Sub: lambda a, b: a - b,
        ast.Mult: lambda a, b: a * b,
        ast.Div: lambda a, b: a / b,
        ast.FloorDiv: lambda a, b: a // b,
        ast.Mod: lambda a, b: a % b,
        ast.Pow: lambda a, b: a ** b,
    }

    _allowed_unops: Dict[type, Callable[[float], float]] = {
        ast.UAdd: lambda a: +a,
        ast.USub: lambda a: -a,
    }

    _allowed_cmp = {
        ast.Eq: lambda a, b: 1.0 if a == b else 0.0,
        ast.NotEq: lambda a, b: 1.0 if a != b else 0.0,
        ast.Lt: lambda a, b: 1.0 if a < b else 0.0,
        ast.LtE: lambda a, b: 1.0 if a <= b else 0.0,
        ast.Gt: lambda a, b: 1.0 if a > b else 0.0,
        ast.GtE: lambda a, b: 1.0 if a >= b else 0.0,
    }

    def __init__(self, variables: Dict[str, float] | None = None):
        self.variables = variables or {}

    # Public API
    def evaluate(self, expr: str) -> float:
        try:
            node = ast.parse(expr, mode="eval")
            return float(self.visit(node.body))
        except LogicError:
            raise
        except ZeroDivisionError:
            raise
        except Exception as e:
            raise LogicError(f"invalid expression: {e}")

    # Visitors
    def visit_BinOp(self, node: ast.BinOp) -> float:
        op_type = type(node.op)
        if op_type not in self._allowed_binops:
            raise LogicError(f"operator {op_type.__name__} is not allowed")
        left = float(self.visit(node.left))
        right = float(self.visit(node.right))
        return float(self._allowed_binops[op_type](left, right))

    def visit_UnaryOp(self, node: ast.UnaryOp) -> float:
        op_type = type(node.op)
        if op_type not in self._allowed_unops:
            raise LogicError(f"unary operator {op_type.__name__} is not allowed")
        operand = float(self.visit(node.operand))
        return float(self._allowed_unops[op_type](operand))

    def visit_Name(self, node: ast.Name) -> float:
        if node.id not in self.variables:
            raise LogicError(f"unknown variable '{node.id}'")
        return float(self.variables[node.id])

    def visit_Constant(self, node: ast.Constant) -> float:
        if isinstance(node.value, (int, float)):
            return float(node.value)
        raise LogicError("only numeric constants are allowed")

    def visit_Call(self, node: ast.Call) -> float:
        if not isinstance(node.func, ast.Name):
            raise LogicError("only simple function names are allowed")
        fn_name = node.func.id
        if fn_name not in self._allowed_funcs:
            raise LogicError(f"function '{fn_name}' is not allowed")
        if node.keywords:
            raise LogicError("keyword arguments are not allowed")
        args = [float(self.visit(arg)) for arg in node.args]
        return float(self._allowed_funcs[fn_name](*args))

    # Disallow everything else
    def generic_visit(self, node: ast.AST) -> Any:
        raise LogicError(f"illegal expression node: {type(node).__name__}")

    def visit_Compare(self, node: ast.Compare) -> float:
        # Support chained comparisons: a < b < c
        left = float(self.visit(node.left))
        assert len(node.ops) == len(node.comparators)
        current = left
        for op, comp in zip(node.ops, node.comparators):
            right = float(self.visit(comp))
            op_type = type(op)
            if op_type not in self._allowed_cmp:
                raise LogicError(f"comparison {op_type.__name__} not allowed")
            ok = self._allowed_cmp[op_type](current, right)
            if ok == 0.0:
                return 0.0
            current = right
        return 1.0
