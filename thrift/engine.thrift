namespace cpp engine
namespace go engine

struct HelloRequest {
  1: string name
}

struct HelloReply {
  1: string message
}

service EngineService {
  HelloReply Hello(1: HelloRequest req)
}
