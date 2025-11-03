package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/Patrick8894/harmonia/api-gw/docs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Patrick8894/harmonia/api-gw/internal/auth"
	"github.com/Patrick8894/harmonia/api-gw/internal/cache"
	"github.com/Patrick8894/harmonia/api-gw/internal/config"
	"github.com/Patrick8894/harmonia/api-gw/internal/engine"
	"github.com/Patrick8894/harmonia/api-gw/internal/health"
	"github.com/Patrick8894/harmonia/api-gw/internal/hello"
	"github.com/Patrick8894/harmonia/api-gw/internal/httpserver"
	"github.com/Patrick8894/harmonia/api-gw/internal/logic"
	"github.com/redis/go-redis/v9"
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

	// --- DB (users)
	db, err := sql.Open("mysql", cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Migrations + dev seed data
	if err := auth.RunMigrations(ctx, db); err != nil {
		log.Fatal(err)
	}
	if err := auth.SeedDevData(ctx, db); err != nil {
		log.Fatal(err)
	}

	// --- Sessions backend selection
	var sessStore auth.SessionStore
	switch cfg.SessionBackend {
	case "redis":
		rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
		if err := rdb.Ping(context.Background()).Err(); err != nil {
			log.Fatal(err)
		}
		sessStore = auth.NewRedisStore(rdb, time.Duration(cfg.CookieMaxAge)*time.Second, "sess:")
	default:
		sessStore = auth.NewMemoryStore(time.Duration(cfg.CookieMaxAge) * time.Second)
	}

	// --- Cache backend selection
	var resultCache cache.Store
	switch cfg.CacheBackend {
	case "redis":
		rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
		if err := rdb.Ping(context.Background()).Err(); err != nil {
			log.Fatal(err)
		}
		resultCache = cache.NewRedisStore(rdb, "rc:") // result cache namespace
	default:
		resultCache = cache.NewMemoryStore()
	}

	// Services (pass cache + TTL)
	userRepo := auth.NewUserRepo(db)
	authCtrl := auth.NewController(userRepo, sessStore, cfg.CookieName, cfg.CookieDomain, cfg.CookieSecure, cfg.CookieMaxAge)

	cacheTTL := time.Duration(cfg.CacheTTLSeconds) * time.Second

	engineClient := engine.NewClient(cfg.EngineAddr)
	engineSvc := engine.NewService(engineClient, resultCache, cacheTTL)

	logicClient := logic.NewClient(cfg.LogicAddr)
	logicSvc := logic.NewService(logicClient, resultCache, cacheTTL)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Register routes; pass sessStore to middleware inside httpserver.RegisterRoutes
	httpserver.RegisterRoutes(r, cfg, engineSvc, logicSvc, health.New(), hello.New(), authCtrl, sessStore)

	log.Println("Harmonia API listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
