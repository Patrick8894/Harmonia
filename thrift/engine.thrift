namespace cpp engine
namespace go engine

struct HelloRequest { 1: string name }
struct HelloReply  { 1: string message }

struct PiRequest { 1: i64 samples }
struct PiReply {
  1: double pi
  2: i64 inside
  3: i64 total
  4: i64 seed
}

struct Matrix {
  1: i32 rows
  2: i32 cols
  3: list<double> data
}
struct MatMulRequest { 1: Matrix a, 2: Matrix b }
struct MatReply { 1: Matrix c }

struct VectorStatsRequest {
  1: list<double> data
  2: bool sample = true
}
struct VectorStatsReply {
  1: i64 count
  2: double sum
  3: double mean
  4: double variance
  5: double stddev
  6: double min
  7: double max
}

service EngineService {
  HelloReply Hello(1: HelloRequest req)
  PiReply    EstimatePi(1: PiRequest req)
  MatReply   MatMul(1: MatMulRequest req)
  VectorStatsReply ComputeStats(1: VectorStatsRequest req)
}
