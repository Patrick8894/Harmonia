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

    void Compute(ComputeResult& _return, const ComputeTask& task) override {
        std::cout << "[Engine] Received task_id=" << task.task_id
                  << ", payload=" << task.payload_json << std::endl;
        _return.task_id = task.task_id;
        _return.code = 0;
        _return.result_json = R"({"result":"ok","from":"C++ Engine"})";
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
