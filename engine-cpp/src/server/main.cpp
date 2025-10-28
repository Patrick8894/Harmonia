#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <iostream>

#include "gen-cpp/EngineService.h"
#include "src/server/EngineServiceHandler.h"

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;
using namespace apache::thrift::server;

int main() {
    const int port = 9101;

    auto handler   = std::make_shared<engine::EngineServiceHandler>();
    auto processor = std::make_shared<engine::EngineServiceProcessor>(handler);
    auto serverTransport  = std::make_shared<TServerSocket>(port);
    auto transportFactory = std::make_shared<TBufferedTransportFactory>();
    auto protocolFactory  = std::make_shared<TBinaryProtocolFactory>();

    TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
    std::cout << "✅ C++ Thrift EngineService running on port " << port << std::endl;
    server.serve();
    return 0;
}
