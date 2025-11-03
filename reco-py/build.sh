#!/usr/bin/env bash
set -e

# Simple build helper for the Python gRPC Logic Service
# Usage:
#   ./build.sh              → generate gRPC stubs
#   ./build.sh run          → generate stubs and run main.py
#   ./build.sh clean        → remove generated files

PROTO_DIR="../proto"
GEN_DIR="."
PROTO_FILE="$PROTO_DIR/logic.proto"
VENV_DIR=".venv"
CLIENT_DIR="./clients"

echo "🚀 [1/3] Checking virtual environment..."
if [ ! -d "$VENV_DIR" ]; then
  echo "⚙️  Creating venv..."
  python3 -m venv $VENV_DIR
fi
source $VENV_DIR/bin/activate

echo "📦 [2/3] Installing dependencies..."
pip install -q --upgrade pip
pip install -q grpcio grpcio-tools

if [ "$1" == "clean" ]; then
  echo "🧹 Cleaning generated stubs..."
  rm -f logic_pb2.py logic_pb2_grpc.py
  deactivate
  exit 0
fi

echo "🔧 [3/3] Generating gRPC stubs from $PROTO_FILE..."
python -m grpc_tools.protoc -I $PROTO_DIR --python_out=$GEN_DIR --grpc_python_out=$GEN_DIR $PROTO_FILE
cp logic_pb2.py logic_pb2_grpc.py $CLIENT_DIR/
echo "✅ gRPC stubs generated successfully."

if [ "$1" == "run" ]; then
  echo "🚀 Starting Python LogicService..."
  echo "------------------------------------"
  python main.py
else
  echo "✅ Done. Run './build.sh run' to start the service."
fi
