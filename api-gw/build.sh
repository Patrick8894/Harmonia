#!/usr/bin/env bash
set -e

# Dev-only rebuild helper:
# - Clean + regenerate Go gRPC stubs from ../proto/*.proto -> ./gen/logic/v1
# - Rebuild Swagger docs into ./docs (imported via _ "github.com/Patrick8894/harmonia/api-gw/docs")
# - Generate Go Thrift stubs from ../thrift/engine.thrift -> ./gen/engine
#
# Usage:
#   ./build.sh

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR="${SCRIPT_DIR}/../proto"
OUT_DIR="${SCRIPT_DIR}/gen/logic/v1"
SWAGGER_OUT="${SCRIPT_DIR}/docs"
GENERAL_INFO="${SCRIPT_DIR}/main.go"   # entry scanned by swag

THRIFT_IDL="${SCRIPT_DIR}/../thrift/engine.thrift"
THRIFT_OUT_BASE="${SCRIPT_DIR}/gen"    # thrift will create ${THRIFT_OUT_BASE}/engine/

# --- minimal tool checks ---
need() { command -v "$1" >/dev/null 2>&1 || { echo "❌ '$1' not found in PATH"; exit 1; }; }
need protoc
need protoc-gen-go
need protoc-gen-go-grpc
need swag

# --- Protobufs ---
echo "🧹 Cleaning ${OUT_DIR} ..."
rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

# Collect protos (non-recursive; change -maxdepth for recursive)
mapfile -t PROTOS < <(find "${PROTO_DIR}" -maxdepth 1 -type f -name '*.proto' | sort)
if [ ${#PROTOS[@]} -eq 0 ]; then
  echo "ℹ️  No .proto files found in ${PROTO_DIR}"
else
  echo "🔧 Generating Go stubs → ${OUT_DIR}"
  protoc -I "${PROTO_DIR}" \
    --go_out="${OUT_DIR}" --go_opt=paths=source_relative \
    --go-grpc_out="${OUT_DIR}" --go-grpc_opt=paths=source_relative \
    "${PROTOS[@]}"
fi

# --- Thrift (Go) ---
if [ -f "${THRIFT_IDL}" ]; then
  if command -v thrift >/dev/null 2>&1; then
    echo "🧹 Cleaning ${THRIFT_OUT_BASE}/engine ..."
    rm -rf "${THRIFT_OUT_BASE}/engine"
    mkdir -p "${THRIFT_OUT_BASE}"

    echo "🔧 Generating Thrift Go stubs → ${THRIFT_OUT_BASE}/engine"
    # package_prefix ensures imports like: github.com/Patrick8894/harmonia/api-gw/gen/engine
    thrift -r \
      --gen go:package_prefix=github.com/Patrick8894/harmonia/api-gw/gen/ \
      -out "${THRIFT_OUT_BASE}" \
      "${THRIFT_IDL}"
  else
    echo "⚠️  'thrift' compiler not found; skipping Thrift Go generation"
  fi
else
  echo "ℹ️  No Thrift IDL found at ${THRIFT_IDL}; skipping Thrift Go generation"
fi

# --- Swagger (swag) ---
echo "🧽 Refreshing Swagger docs in ${SWAGGER_OUT} ..."
rm -rf "${SWAGGER_OUT}"

swag init \
  --parseDependency \
  --parseInternal \
  --dir "${SCRIPT_DIR}" \
  --generalInfo "main.go" \
  --output "${SWAGGER_OUT}"

echo "✅ Done."
