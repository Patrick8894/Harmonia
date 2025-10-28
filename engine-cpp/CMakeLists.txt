cmake_minimum_required(VERSION 3.15)
project(engine_server CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Thrift (assumes installed system-wide)
find_package(Threads REQUIRED)

# You may need to hint Thrift libs/headers depending on your system:
# include_directories(/usr/local/include)
# link_directories(/usr/local/lib)

include_directories(${CMAKE_SOURCE_DIR})
include_directories(${CMAKE_SOURCE_DIR}/gen-cpp)

add_library(engine_gen
    gen-cpp/EngineService.cpp
    gen-cpp/engine_types.cpp
)

add_library(engine_lib
    src/lib/MonteCarlo.cpp
    src/lib/MatrixOps.cpp
    src/lib/Stats.cpp
)
target_link_libraries(engine_lib PUBLIC engine_gen)

add_executable(engine_server
    src/server/main.cpp
    src/server/EngineServiceHandler.cpp
)
target_link_libraries(engine_server
    engine_lib
    engine_gen
    thrift
    Threads::Threads
)

# Optional clients
add_executable(pi_client     clients/pi_client.cpp)
target_link_libraries(pi_client engine_gen thrift Threads::Threads)

add_executable(matmul_client clients/matmul_client.cpp)
target_link_libraries(matmul_client engine_gen thrift Threads::Threads)

add_executable(stats_client  clients/stats_client.cpp)
target_link_libraries(stats_client engine_gen thrift Threads::Threads)
