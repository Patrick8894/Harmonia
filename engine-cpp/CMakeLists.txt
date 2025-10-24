cmake_minimum_required(VERSION 3.15)
project(engine_cpp)

set(CMAKE_CXX_STANDARD 17)

# Find Thrift headers and library manually
find_path(THRIFT_INCLUDE_DIR thrift/Thrift.h)
find_library(THRIFT_LIB thrift)

include_directories(${THRIFT_INCLUDE_DIR} gen-cpp)

add_executable(engine main.cpp
    gen-cpp/EngineService.cpp
    gen-cpp/engine_types.cpp
)

target_link_libraries(engine ${THRIFT_LIB})
