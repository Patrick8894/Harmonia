# Harmonia

**Harmonia** is a hybrid compute dashboard demonstrating **cross-language RPC integration** and **modern full-stack architecture**.  
It harmonizes Go, Python, and C++ services into a unified system capable of coordinating logical planning, numerical computation, and real-time visualization.

---

## 🧩 Overview
- **Frontend:** React / Next.js dashboard (Material UI) for login, registration, home, and dashboard pages.
- **Gateway:** Go (Gin) REST API orchestrator managing authentication (via cookies), routing, aggregation, and pluggable caching for RPC results.
- **Logic Service:** Python (gRPC) microservice handling expression evaluation, task planning, and data transformation.
- **Compute Engine:** C++ (Thrift) microservice performing high-speed numerical, matrix, and statistical computations.
- **Databases:** MySQL for persistent user/auth data, and Redis or in-memory map for session management.

---

## 🏗 Architecture
```
React / Next.js (Material UI)
        │
        ▼
   REST API (Go + Gin)
        │
        ├── gRPC ─▶ Python Logic Service
        └── Thrift ─▶ C++ Compute Engine
        │
        ▼
 MySQL ──┬─ User/Auth Data
         ├─ Redis / In‑Memory Session Store
         └─ Redis / In‑Memory **RPC Result Cache**
```

---

## 🚀 Key Features
- **Cross-language orchestration:** Go ↔ Python (gRPC) ↔ C++ (Thrift)
- **RESTful interface with Cookie-based authentication**
- **Pluggable backends**
  - Sessions: **Redis** or **in-memory**
  - **RPC result cache:** **Redis** or **in-memory** (JSON-serialized values, SHA‑256 request keys, TTL)
- **Swagger UI** for interactive API documentation and testing
- **Air live reload** for hot-reloading during backend development
- **Docker Compose** setup, extendable to Kubernetes

---

## 📁 Project Structure
```
harmonia/
 ├── api-gw/           # Go API Gateway
 │   ├── cmd/          # Entrypoint (main.go)
 │   ├── db/           # MySQL / Redis init and connection logic
 │   ├── docs/         # Auto-generated Swagger documentation
 │   ├── gen/          # gRPC and Thrift stubs
 │   ├── internal/     # Application modules
 │   │   ├── auth/     # Cookie-based auth + session management
 │   │   ├── cache/    # 🔹 Pluggable cache (memory/redis) for RPC results
 │   │   ├── logic/    # gRPC client for Python LogicService
 │   │   ├── engine/   # Thrift client for C++ EngineService
 │   │   ├── hello/    # Sample hello endpoints
 │   │   ├── health/   # Health check endpoints
 │   │   └── httpserver/ # Gin router & route registration
 │   ├── build.sh      # Proto/Thrift/Swagger generation
 │   ├── Dockerfile.dev
 │   └── tmp/
 │
 ├── engine-cpp/       # C++ Compute Engine
 │   ├── src/          # Thrift service implementation
 │   ├── clients/      # Example Thrift test clients
 │   ├── build.sh      # Build and run script
 │   ├── CMakeLists.txt
 │   ├── Dockerfile.dev
 │   └── tmp/
 │
 ├── reco-py/          # Python Logic Service (gRPC)
 │   ├── logic_service/ # Core logic and RPC handlers
 │   ├── clients/       # Example gRPC clients
 │   ├── build.sh       # Build/run helper
 │   ├── generate.sh    # Generate gRPC stubs
 │   ├── main.py        # Entrypoint
 │   └── requirements.txt
 │
 ├── frontend/         # Next.js + Material UI frontend
 │   ├── app/          # Next.js app router
 │   ├── pages/        # Page components (login, register, dashboard)
 │   ├── components/   # Shared UI components
 │   ├── contexts/     # React contexts for auth/session
 │   ├── lib/          # API utilities and fetch hooks
 │   ├── public/       # Static assets
 │   ├── package.json
 │   ├── Dockerfile.dev
 │   └── next.config.ts
 │
 ├── proto/            # gRPC IDL
 │   └── logic.proto
 │
 ├── thrift/           # Thrift IDL
 │   └── engine.thrift
 │
 ├── docker-compose.yml
 └── README.md
```

---

## ⚙️ Deployment & Development

### 🧠 API Gateway (Go)
```bash
cd api-gw
./build.sh   # regenerate gRPC/Thrift/Swagger
air          # run in hot-reload mode
```
- Accessible at: `http://localhost:8080`
- Swagger docs: `http://localhost:8080/swagger/index.html`

### 🧮 Compute Engine (C++)
```bash
cd engine-cpp
./build.sh run
```
- Runs the Thrift server on port `9101`.

### 🧠 Logic Service (Python)
```bash
cd reco-py
./build.sh run
```
- Starts the Python gRPC server on port `9002`.

### 💻 Frontend (Next.js)
```bash
cd frontend
npm install
npm run dev
```
- Access at: `http://localhost:3000`

### 🐳 Docker Compose (Full stack)
```bash
docker-compose up --build
```
All services (gateway, logic, engine, frontend) will start together.

---

## 🧩 Services Summary
| Service | Language | Protocol | Port | Description |
|----------|-----------|-----------|-------|--------------|
| API Gateway | Go | REST / JSON | 8080 | Routes requests, manages cookies/sessions, caches RPC results |
| Logic Service | Python | gRPC | 9002 | Evaluates expressions, transforms data, plans tasks |
| Compute Engine | C++ | Thrift | 9101 | Performs numerical and matrix computations |
| Frontend | Next.js | HTTP | 3000 | User interface for login and compute dashboard |

---

## 🧭 Tech Stack
| Layer | Tech |
|-------|------|
| Frontend | React, Next.js, Material UI |
| API Gateway | Go (Gin), Swagger, Air |
| RPC | gRPC (Python), Thrift (C++) |
| Database | MySQL |
| Session Store | Redis / In-memory |
| RPC Result Cache | Redis / In-memory |
| Deployment | Docker Compose, Kubernetes-ready |

---

## 🏁 Status
Harmonia is actively evolving to demonstrate **cross-language orchestration**, **hybrid compute pipelines**, and **real-world service integration** between Go, Python, and C++ systems — with a **pluggable cache** for fast, idempotent RPC responses.
