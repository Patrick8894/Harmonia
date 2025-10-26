#!/usr/bin/env bash
set -e

# Simple build helper for the C++ Thrift Engine
# Usage:
#   ./build.sh          → build only
#   ./build.sh run      → build + run
#   ./build.sh clean    → remove build and generated files

THRIFT_FILE="../thrift/engine.thrift"
BUILD_DIR="build"
GEN_DIR="gen-cpp"
EXECUTABLE="$BUILD_DIR/engine"

if [ "$1" == "clean" ]; then
  echo "🧹 Cleaning build artifacts and generated files..."
  rm -rf "$BUILD_DIR" "$GEN_DIR"
  echo "✅ Clean complete."
  exit 0
fi

echo "🚀 [1/4] Generating Thrift stubs..."
thrift -r --gen cpp "$THRIFT_FILE"

echo "🏗️  [2/4] Preparing build directory..."
mkdir -p "$BUILD_DIR"
cd "$BUILD_DIR"

echo "🔧 [3/4] Running CMake + make..."
cmake .. > /dev/null
make -j"$(nproc)"

cd ..

if [ "$1" == "run" ]; then
  echo "✅ [4/4] Starting C++ Thrift EngineService..."
  echo "---------------------------------------------"
  "$EXECUTABLE"
else
  echo "✅ Build complete. Run './build.sh run' to start the server."
fi
