#pragma once
#include "gen-cpp/EngineService.h"
#include <cstddef>

namespace mat {

inline bool valid_shape(const engine::Matrix& m) {
    long long need = static_cast<long long>(m.rows) * static_cast<long long>(m.cols);
    return m.rows >= 0 && m.cols >= 0 &&
           need == static_cast<long long>(m.data.size());
}

inline std::size_t idxRM(int r, int c, int cols) {
    return static_cast<std::size_t>(r) * static_cast<std::size_t>(cols) + static_cast<std::size_t>(c);
}

// C = A x B, assumes shapes already validated
void matmul(const engine::Matrix& A, const engine::Matrix& B, engine::Matrix& C);

} // namespace mat
