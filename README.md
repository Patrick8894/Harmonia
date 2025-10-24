# Harmonia

**Harmonia** is a hybrid compute dashboard designed to demonstrate cross-language RPC integration and modern full-stack architecture. It harmonizes Go, Python, and C++ services into a unified system that coordinates complex tasks, high-performance computations, and real-time visualization.

## Overview
- **Frontend:** React / Next.js dashboard for submitting and monitoring compute tasks.
- **Gateway:** Go (Gin) REST API as the orchestrator, managing authentication, routing, and aggregation.
- **Logic Service:** Python (gRPC) microservice handling complex logical and data-processing tasks.
- **Compute Engine:** C++ (Thrift) service performing high-speed numerical computations.
- **Databases:** Oracle for persistent storage (auth, records) and Redis for caching task results.

## Architecture
```
React / Next.js  ── REST ──▶  Go API Gateway (Gin)
                               │
                               ├── gRPC ─▶ Python Logic Service
                               └── Thrift ─▶ C++ Compute Engine
                               │
                               ▼
                        Oracle + Redis
```

## Key Features
- Cross-language orchestration via gRPC and Thrift
- RESTful interface with JWT-based authentication
- Real-time task visualization and monitoring
- Docker Compose setup, extendable to Kubernetes
- Modular architecture emphasizing clarity and scalability

## Tech Stack
- **Frontend:** React, Next.js
- **Backend:** Go (Gin, Viper, Zap), Python (gRPC), C++ (Thrift)
- **Database:** Oracle, Redis
- **Deployment:** Docker Compose / Kubernetes-ready