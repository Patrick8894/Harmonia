#include "src/lib/MonteCarlo.h"
#include <random>

namespace mc {

PiResult estimate_pi(long long samples, long long seed) {
    std::mt19937_64 rng(static_cast<uint64_t>(seed));
    std::uniform_real_distribution<double> dist(-1.0, 1.0);

    long long inside = 0;
    for (long long i = 0; i < samples; ++i) {
        double x = dist(rng), y = dist(rng);
        if (x*x + y*y <= 1.0) ++inside;
    }
    double pi = 4.0 * static_cast<double>(inside) / static_cast<double>(samples);
    return PiResult{pi, inside, samples, seed};
}

} // namespace mc
