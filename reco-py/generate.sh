#!/usr/bin/env bash
set -euo pipefail
python -m grpc_tools.protoc \
  -I proto \
  --python_out=gen \
  --grpc_python_out=gen \
  proto/logic.proto
echo "Generated Python gRPC code in ./gen"
