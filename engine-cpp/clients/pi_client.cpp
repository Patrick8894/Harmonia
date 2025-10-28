#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <chrono>
#include <iostream>
#include "gen-cpp/EngineService.h"

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

int main(int argc, char** argv) {
    long long samples = (argc > 1) ? std::stoll(argv[1]) : 1000000;

    std::shared_ptr<TTransport> socket   = std::make_shared<TSocket>("127.0.0.1", 9101);
    std::shared_ptr<TTransport> transport= std::make_shared<TBufferedTransport>(socket);
    std::shared_ptr<TProtocol> protocol  = std::make_shared<TBinaryProtocol>(transport);
    engine::EngineServiceClient client(protocol);

    try {
        transport->open();
        engine::PiRequest req;
        req.samples = samples;
        engine::PiReply rep;

        auto t0 = std::chrono::high_resolution_clock::now();
        client.EstimatePi(rep, req);
        auto t1 = std::chrono::high_resolution_clock::now();
        double ms = std::chrono::duration<double, std::milli>(t1 - t0).count();

        std::cout << "π ≈ " << rep.pi << "  (inside=" << rep.inside
                  << ", total=" << rep.total
                  << ", seed=" << rep.seed
                  << ", time=" << ms << "ms)" << std::endl;
        transport->close();
    } catch (const TException& tx) {
        std::cerr << "Thrift exception: " << tx.what() << std::endl;
    }
}
