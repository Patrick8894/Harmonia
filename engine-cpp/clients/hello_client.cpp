#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <iostream>
#include "gen-cpp/EngineService.h"

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

int main(int argc, char** argv) {
    const char* name = (argc > 1 ? argv[1] : "Patrick");

    std::shared_ptr<TTransport> socket   = std::make_shared<TSocket>("127.0.0.1", 9101);
    std::shared_ptr<TTransport> transport= std::make_shared<TBufferedTransport>(socket);
    std::shared_ptr<TProtocol> protocol  = std::make_shared<TBinaryProtocol>(transport);
    engine::EngineServiceClient client(protocol);

    try {
        transport->open();
        engine::HelloRequest req;
        req.name = name;
        engine::HelloReply rep;
        client.Hello(rep, req);
        std::cout << rep.message << std::endl;
        transport->close();
    } catch (const TException& tx) {
        std::cerr << "Thrift exception: " << tx.what() << std::endl;
    }
}
