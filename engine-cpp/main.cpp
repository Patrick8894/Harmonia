#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <iostream>

#include "gen-cpp/EngineService.h"

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;
using namespace ::engine;

class EngineServiceHandler : public EngineServiceIf {
public:
    EngineServiceHandler() {}

    void Hello(HelloReply& _return, const HelloRequest& req) override {
        
        std::cout << "[Engine] Received HelloRequest from name=" << req.name << std::endl;
        
        _return.message = "Hello " + req.name + " from C++ Engine!";
    }
};

int main() {
    int port = 9101;
    ::std::shared_ptr<EngineServiceHandler> handler(new EngineServiceHandler());
    ::std::shared_ptr<TProcessor> processor(new EngineServiceProcessor(handler));
    ::std::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
    ::std::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
    ::std::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

    TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
    std::cout << "✅ C++ Thrift EngineService running on port " << port << std::endl;
    server.serve();
    return 0;
}
