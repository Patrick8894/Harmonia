#include "src/lib/MatrixOps.h"

namespace mat {

void matmul(const engine::Matrix& A, const engine::Matrix& B, engine::Matrix& C) {
    const int m = A.rows, k = A.cols, n = B.cols;

    for (int i = 0; i < m; ++i) {
        for (int p = 0; p < k; ++p) {
            const double a_ip = A.data[idxRM(i, p, A.cols)];
            if (a_ip == 0.0) continue;
            const std::size_t bRow = idxRM(p, 0, B.cols);
            const std::size_t cRow = idxRM(i, 0, C.cols);
            for (int j = 0; j < n; ++j) {
                C.data[cRow + j] += a_ip * B.data[bRow + j];
            }
        }
    }
}

} // namespace mat
