#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <iostream>
#include <vector>
#include "gen-cpp/EngineService.h"

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

int main(int argc, char** argv) {
    bool sample = true;
    if (argc > 1 && std::string(argv[1]) == "pop") sample = false;

    std::vector<double> data = {1.0, 2.0, 2.0, 3.0, 9.0};

    std::shared_ptr<TTransport> socket   = std::make_shared<TSocket>("127.0.0.1", 9101);
    std::shared_ptr<TTransport> transport= std::make_shared<TBufferedTransport>(socket);
    std::shared_ptr<TProtocol> protocol  = std::make_shared<TBinaryProtocol>(transport);
    engine::EngineServiceClient client(protocol);

    try {
        transport->open();

        engine::VectorStatsRequest req;
        req.data = data;
        req.sample = sample;
        engine::VectorStatsReply rep;

        client.ComputeStats(rep, req);

        std::cout << (sample ? "Sample" : "Population") << " stats:\n";
        std::cout << "count: " << rep.count << "\n"
                  << "sum:   " << rep.sum   << "\n"
                  << "mean:  " << rep.mean  << "\n"
                  << "var:   " << rep.variance << "\n"
                  << "stddev:" << rep.stddev   << "\n"
                  << "min:   " << rep.min   << "\n"
                  << "max:   " << rep.max   << "\n";

        transport->close();
    } catch (const TException& tx) {
        std::cerr << "Thrift exception: " << tx.what() << std::endl;
    }
}
