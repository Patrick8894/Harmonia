from __future__ import annotations
from typing import Iterable, List
import math

from logic_service.evaluators import SafeExpressionEvaluator
from logic_service.errors import LogicError


def _is_finite(x: float) -> bool:
    return not (math.isnan(x) or math.isinf(x))


def transform_map(data: Iterable[float], expr: str, var_name: str = "x") -> List[float]:
    out: List[float] = []
    for el in data:
        ev = SafeExpressionEvaluator({var_name: float(el)})
        val = ev.evaluate(expr)
        if not _is_finite(val):
            raise LogicError("non-finite value produced (NaN/Inf)")
        out.append(float(val))
    return out


def transform_filter(data: Iterable[float], expr: str, var_name: str = "x") -> List[float]:
    out: List[float] = []
    for el in data:
        ev = SafeExpressionEvaluator({var_name: float(el)})
        keep = ev.evaluate(expr)
        if not _is_finite(keep):
            raise LogicError("non-finite predicate produced (NaN/Inf)")
        if keep != 0.0:  # non-zero is truthy
            out.append(float(el))
    return out


def transform_sum(data: Iterable[float], expr: str, var_name: str = "x") -> float:
    total = 0.0
    for el in data:
        ev = SafeExpressionEvaluator({var_name: float(el)})
        val = ev.evaluate(expr)
        if not _is_finite(val):
            raise LogicError("non-finite value produced in SUM (NaN/Inf)")
        total += float(val)
    return total
