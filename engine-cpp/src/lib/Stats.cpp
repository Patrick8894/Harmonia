#include "src/lib/Stats.h"
#include <vector>

namespace stats {

template <typename It>
Summary compute(It first, It last, bool sample) {
    Summary s;
    if (first == last) return s;

    double mean = 0.0, M2 = 0.0, sum = 0.0;
    double mn = std::numeric_limits<double>::infinity();
    double mx = -std::numeric_limits<double>::infinity();
    long long k = 0;

    for (It it = first; it != last; ++it) {
        const double x = static_cast<double>(*it);
        ++k;
        sum += x;
        if (x < mn) mn = x;
        if (x > mx) mx = x;

        double delta = x - mean;
        mean += delta / static_cast<double>(k);
        double delta2 = x - mean;
        M2 += delta * delta2;
    }

    s.count = k;
    s.sum = sum;
    s.mean = mean;
    if (sample) {
        s.variance = (k >= 2) ? (M2 / static_cast<double>(k - 1)) : 0.0;
    } else {
        s.variance = (k >= 1) ? (M2 / static_cast<double>(k)) : 0.0;
    }
    s.stddev = std::sqrt(s.variance);
    s.min = mn;
    s.max = mx;
    return s;
}

// Explicit instantiation for common iterator types (vector<double>)
template Summary compute<std::vector<double>::const_iterator>(
    std::vector<double>::const_iterator, std::vector<double>::const_iterator, bool);

} // namespace stats
