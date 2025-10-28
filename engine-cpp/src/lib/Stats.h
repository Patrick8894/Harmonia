#pragma once
#include <cstdint>
#include <iterator>
#include <limits>
#include <cmath>

namespace stats {

struct Summary {
    long long count{0};
    double sum{0.0};
    double mean{0.0};
    double variance{0.0};
    double stddev{0.0};
    double min{std::numeric_limits<double>::quiet_NaN()};
    double max{std::numeric_limits<double>::quiet_NaN()};
};

template <typename It>
Summary compute(It first, It last, bool sample);

} // namespace stats
