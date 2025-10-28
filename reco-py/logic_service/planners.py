from __future__ import annotations
from typing import List, Dict, Tuple
import re

from logic_service.errors import LogicError

_BASE_STEPS: List[Tuple[str, str, int, int]] = [
    ("Clarify scope", "Write 1–2 sentences of the goal + success criteria.", 1, 10),
    ("Design surface", "Sketch API/proto & inputs/outputs; decide return schema.", 1, 15),
    ("Implement MVP", "Code minimal path; keep pure logic isolated in its module.", 1, 40),
    ("Add tests", "Unit tests for happy-path + 1–2 edge cases.", 2, 25),
    ("Wire endpoint", "Expose via gRPC method; integrate with service layer.", 2, 20),
    ("Docs & examples", "README snippet + simple client sample.", 3, 10),
]

# Keyword -> extra step templates
_KEYWORD_TEMPLATES: Dict[str, List[Tuple[str, str, int, int, str]]] = {
    "grpc": [
        ("Update proto", "Add RPC/messages; regenerate stubs.", 1, 10, "Clarify scope"),
        ("Server hook", "Register handler in server bootstrap.", 2, 10, "Implement MVP"),
    ],
    "docker": [
        ("Containerize", "Add Dockerfile + dev compose target.", 2, 20, "Implement MVP"),
    ],
    "k8s": [
        ("K8s manifest", "Deployment/Service; set resource requests.", 3, 25, "Containerize"),
    ],
    "tests": [
        ("More tests", "Edge cases: empty input, invalid params, timeouts.", 2, 20, "Add tests"),
    ],
    "thrift": [
        ("Cross-RPC note", "Document how this composes with Thrift services.", 3, 10, "Docs & examples"),
    ],
    "logging": [
        ("Observability", "Add structured logs around request/response (no PII).", 2, 10, "Implement MVP"),
    ],
}

def _pick_keywords(text: str, hints: List[str]) -> List[str]:
    text = text.lower()
    found = set()
    for k in _KEYWORD_TEMPLATES:
        if k in text:
            found.add(k)
    for h in hints or []:
        h = h.strip().lower()
        if h in _KEYWORD_TEMPLATES:
            found.add(h)
    return list(found)

def _bound(n: int, lo: int, hi: int) -> int:
    return max(lo, min(hi, n))

def plan_tasks(goal: str, hints: List[str] | None, max_steps: int | None) -> Tuple[List[dict], str]:
    if not goal or not goal.strip():
        raise LogicError("goal is empty")
    max_steps = _bound(max_steps or 8, 3, 20)

    # Derive keywords
    kws = _pick_keywords(goal, hints or [])

    # Build base tasks
    tasks: List[dict] = []
    id_map: Dict[str, str] = {}  # title -> id

    def add_task(title: str, detail: str, priority: int, est: int, depends_title: str | None = None):
        tid = f"T{len(tasks)+1}"
        depends_on = [id_map[depends_title]] if depends_title and depends_title in id_map else []
        task = {
            "id": tid, "title": title, "detail": detail,
            "priority": int(priority), "estimate_min": int(est),
            "depends_on": depends_on,
        }
        tasks.append(task)
        id_map[title] = tid

    for t in _BASE_STEPS:
        add_task(*t)

    # Keyword-specific extras
    for kw in kws:
        for (title, detail, prio, est, dep) in _KEYWORD_TEMPLATES[kw]:
            add_task(title, detail, prio, est, dep)

    # Lightweight trimming heuristic:
    # keep all priority-1, then 2, then 3 until max_steps
    tasks.sort(key=lambda x: (x["priority"], x["title"]))
    tasks = tasks[:max_steps]

    # Fix up dependencies after trimming (drop missing deps)
    kept_ids = {t["id"] for t in tasks}
    for t in tasks:
        t["depends_on"] = [d for d in t["depends_on"] if d in kept_ids]

    notes = "Keywords detected: " + (", ".join(sorted(kws)) if kws else "none")
    return tasks, notes
