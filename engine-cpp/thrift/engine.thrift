namespace cpp engine
namespace go engine

struct ComputeTask {
  1: string task_id
  2: string payload_json
}

struct ComputeResult {
  1: string task_id
  2: i32 code
  3: string result_json
}

service EngineService {
  ComputeResult Compute(1: ComputeTask task)
}
