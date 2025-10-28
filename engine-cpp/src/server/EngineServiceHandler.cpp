#include "src/server/EngineServiceHandler.h"
#include <iostream>
#include <limits>
#include <chrono>

#include "src/lib/MonteCarlo.h"
#include "src/lib/MatrixOps.h"
#include "src/lib/Stats.h"

namespace engine {

EngineServiceHandler::EngineServiceHandler() = default;

void EngineServiceHandler::Hello(HelloReply& _return, const HelloRequest& req) {
    std::cout << "[Engine] Hello name=" << req.name << std::endl;
    _return.message = "Hello " + req.name + " from C++ Engine!";
}

void EngineServiceHandler::EstimatePi(PiReply& _return, const PiRequest& req) {
    const int64_t n = req.samples;
    if (n <= 0) {
        _return = PiReply(); // zeros
        std::cerr << "[Engine] EstimatePi: samples must be > 0\n";
        return;
    }
    const int64_t seed = static_cast<int64_t>(
        std::chrono::steady_clock::now().time_since_epoch().count()
    );
    auto res = mc::estimate_pi(n, seed); // returns {pi, inside, total, seed}

    _return.pi = res.pi;
    _return.inside = res.inside;
    _return.total = res.total;
    _return.seed = res.seed;

    std::cout << "[Engine] EstimatePi(samples=" << n << ") -> " << res.pi << "\n";
}

void EngineServiceHandler::MatMul(MatReply& _return, const MatMulRequest& req) {
    const Matrix& A = req.a;
    const Matrix& B = req.b;

    if (!mat::valid_shape(A) || !mat::valid_shape(B) || A.cols != B.rows) {
        std::cerr << "[Engine] MatMul: invalid shapes: (" << A.rows << "x" << A.cols
                  << ") * (" << B.rows << "x" << B.cols << ")\n";
        _return.c = Matrix(); // empty
        return;
    }

    Matrix C;
    C.rows = A.rows;
    C.cols = B.cols;
    C.data.assign(static_cast<size_t>(C.rows) * C.cols, 0.0);

    mat::matmul(A, B, C); // fills C.data
    _return.c = std::move(C);

    std::cout << "[Engine] MatMul: (" << A.rows << "x" << A.cols << ") * ("
              << B.rows << "x" << B.cols << ") -> (" << _return.c.rows
              << "x" << _return.c.cols << ")\n";
}

void EngineServiceHandler::ComputeStats(VectorStatsReply& _return, const VectorStatsRequest& req) {
    const auto& v = req.data;
    bool sample = req.sample;

    stats::Summary s = stats::compute(v.begin(), v.end(), sample);
    _return.count = s.count;
    _return.sum   = s.sum;
    _return.mean  = s.mean;
    _return.variance = s.variance;
    _return.stddev   = s.stddev;
    _return.min   = s.min;
    _return.max   = s.max;

    std::cout << "[Engine] ComputeStats: n=" << s.count
              << " mean=" << s.mean << " std=" << s.stddev << "\n";
}

} // namespace engine
