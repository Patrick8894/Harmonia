package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	// gRPC
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// thrift
	"github.com/apache/thrift/lib/go/thrift"

	// Swagger UI
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Patrick8894/harmonia/api-gw/docs"

	eng "github.com/Patrick8894/harmonia/api-gw/gen/engine"
	lg "github.com/Patrick8894/harmonia/api-gw/gen/logic/v1"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

var logicAddr = getEnv("LOGIC_ADDR", "localhost:9002")
var engineAddr = getEnv("ENGINE_ADDR", "localhost:9101")

// @title           Harmonia API
// @version         0.1.1
// @description     REST gateway for the Harmonia project. Orchestrates Python (gRPC) and C++ (Thrift) services.
// @BasePath        /api

// @tag.name root
// @tag.description Root endpoints

// @tag.name logic
// @tag.description Python gRPC LogicService

// @tag.name engine
// @tag.description C++ Thrift EngineService

// @tag.name health
// @tag.description Liveness & readiness

// Hello godoc
// @Summary      Hello endpoint
// @Description  Basic greeting from Harmonia API Gateway
// @Tags         root
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Harmonia API Gateway!"})
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Liveness/readiness probe endpoint
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /healthz [get]
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// HelloLogicRPC godoc
// @Summary      Call Python LogicService Hello RPC
// @Description  Triggers the Hello RPC on the Python gRPC LogicService
// @Tags         logic
// @Produce      json
// @Param        name  query   string  false  "Name to greet"  default(World)
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /logic/hello [get]
func helloLogicRPCHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		logicAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to connect to logic service: " + err.Error()})
		return
	}
	defer conn.Close()

	client := lg.NewLogicServiceClient(conn)
	resp, err := client.Hello(ctx, &lg.HelloRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}

// HelloEngineRPC godoc
// @Summary      Call C++ EngineService Hello RPC
// @Description  Triggers the Hello RPC on the C++ Thrift EngineService
// @Tags         engine
// @Produce      json
// @Param        name  query   string  false  "Name to greet"  default(World)
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /engine/hello [get]
func helloEngineRPCHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// Buffered transport to match your C++ server
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)

	cfg := &thrift.TConfiguration{
		ConnectTimeout: 1 * time.Second, // timeout for establishing the TCP connection
		SocketTimeout:  2 * time.Second, // timeout for read/write operations
	}

	sock := thrift.NewTSocketConf(engineAddr, cfg)
	if sock == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create socket"})
		return
	}

	transport, err := transportFactory.GetTransport(sock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to wrap transport: " + err.Error()})
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open transport: " + err.Error()})
		return
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	tclient := thrift.NewTStandardClient(iprot, oprot)
	client := eng.NewEngineServiceClient(tclient)

	fmt.Println("[API-GW] Calling EngineService Hello RPC with name=", name, " transport=buffered")

	resp, err := client.Hello(ctx, &eng.HelloRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil) // recommended for safety in dev

	// Base API group: all future endpoints start with /api
	api := r.Group("/api")
	{
		api.GET("/hello", helloHandler)
		api.GET("/healthz", healthHandler)
		api.GET("/logic/hello", helloLogicRPCHandler)
		api.GET("/engine/hello", helloEngineRPCHandler)
	}

	// Swagger endpoint (still available at /swagger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the HTTP server
	r.Run(":8080")
}
