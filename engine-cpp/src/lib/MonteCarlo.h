// MonteCarlo.h
#pragma once
#include <cstdint>

namespace mc {

struct PiResult {
    double pi{0.0};
    long long inside{0};
    long long total{0};
    long long seed{0};
};

PiResult estimate_pi(long long samples, long long seed);

} // namespace mc
