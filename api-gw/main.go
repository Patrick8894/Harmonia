package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// Swagger UI
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// IMPORTANT: replace module path below with your actual module path if different
	_ "github.com/Patrick8894/harmonia/api-gw/docs"
)

// @title           Harmonia API
// @version         0.1.1
// @description     REST gateway for the Harmonia project. Orchestrates Python (gRPC) and C++ (Thrift) services.
// @BasePath        /api

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

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil) // recommended for safety in dev

	// Base API group: all future endpoints start with /api
	api := r.Group("/api")
	{
		api.GET("/hello", helloHandler)
		api.GET("/healthz", healthHandler)
	}

	// Swagger endpoint (still available at /swagger)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the HTTP server
	r.Run(":8080")
}
