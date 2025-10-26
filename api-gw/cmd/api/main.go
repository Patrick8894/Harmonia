package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/Patrick8894/harmonia/api-gw/docs"

	"github.com/Patrick8894/harmonia/api-gw/internal/auth"
	"github.com/Patrick8894/harmonia/api-gw/internal/config"
	"github.com/Patrick8894/harmonia/api-gw/internal/engine"
	"github.com/Patrick8894/harmonia/api-gw/internal/health"
	"github.com/Patrick8894/harmonia/api-gw/internal/hello"
	"github.com/Patrick8894/harmonia/api-gw/internal/httpserver"
	"github.com/Patrick8894/harmonia/api-gw/internal/logic"
)

// @title           Harmonia API
// @version         0.1.1
// @description     REST gateway for the Harmonia project. Orchestrates Python (gRPC) and C++ (Thrift) services.
// @BasePath        /api

// @tag.name root
// @tag.description Root endpoints

//@tag.name auth
// @tag.description Authentication endpoints

// @tag.name logic
// @tag.description Python gRPC LogicService

// @tag.name engine
// @tag.description C++ Thrift EngineService

// @tag.name health
// @tag.description Liveness & readiness

func main() {
	cfg := config.Load()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// RPC clients & services
	engineClient := engine.NewClient(cfg.EngineAddr)
	engineSvc := engine.NewService(engineClient)

	logicClient := logic.NewClient(cfg.LogicAddr)
	logicSvc := logic.NewService(logicClient)

	// Auth store & controller
	authStore := auth.NewStore(time.Duration(cfg.CookieMaxAge)*time.Second, cfg.SessionSecret)
	authCtrl := auth.NewController(authStore, cfg.CookieName, cfg.CookieDomain, cfg.CookieSecure, cfg.CookieMaxAge)

	// Register routes
	httpserver.RegisterRoutes(r, cfg, engineSvc, logicSvc, health.New(), hello.New(), authCtrl, authStore)

	log.Println("Harmonia API listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
