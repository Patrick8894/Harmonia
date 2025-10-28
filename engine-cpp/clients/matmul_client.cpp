#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <iostream>
#include <vector>
#include "gen-cpp/EngineService.h"

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

static engine::Matrix makeMatrix(int rows, int cols, const std::vector<double>& data) {
    engine::Matrix M;
    M.rows = rows;
    M.cols = cols;
    M.data = data;
    return M;
}

static void printMatrix(const engine::Matrix& M) {
    std::cout << M.rows << "x" << M.cols << ":\n";
    for (int r = 0; r < M.rows; ++r) {
        for (int c = 0; c < M.cols; ++c) {
            std::cout << M.data[r * M.cols + c] << (c + 1 == M.cols ? '\n' : ' ');
        }
    }
}

int main() {
    std::shared_ptr<TTransport> socket   = std::make_shared<TSocket>("127.0.0.1", 9101);
    std::shared_ptr<TTransport> transport= std::make_shared<TBufferedTransport>(socket);
    std::shared_ptr<TProtocol> protocol  = std::make_shared<TBinaryProtocol>(transport);
    engine::EngineServiceClient client(protocol);

    try {
        transport->open();

        engine::Matrix A = makeMatrix(2, 3, {1,2,3,4,5,6});
        engine::Matrix B = makeMatrix(3, 2, {7,8,9,10,11,12});
        engine::MatMulRequest req; req.a = A; req.b = B;
        engine::MatReply rep;
        client.MatMul(rep, req);

        std::cout << "C = A x B =\n";
        printMatrix(rep.c);

        transport->close();
    } catch (const TException& tx) {
        std::cerr << "Thrift exception: " << tx.what() << std::endl;
    }
}
