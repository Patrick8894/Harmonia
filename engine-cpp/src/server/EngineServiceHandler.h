#pragma once
#include <memory>
#include "gen-cpp/EngineService.h"

namespace engine {

class EngineServiceHandler : public EngineServiceIf {
public:
    EngineServiceHandler();

    // RPCs
    void Hello(HelloReply& _return, const HelloRequest& req) override;
    void EstimatePi(PiReply& _return, const PiRequest& req) override;
    void MatMul(MatReply& _return, const MatMulRequest& req) override;
    void ComputeStats(VectorStatsReply& _return, const VectorStatsRequest& req) override;

private:
    // if you later need shared state (rng pool, thread pool, caches), add here
};

} // namespace engine
