package httpserver

import (
	"github.com/gin-gonic/gin"

	"github.com/Patrick8894/harmonia/api-gw/internal/auth"
	"github.com/Patrick8894/harmonia/api-gw/internal/config"
	"github.com/Patrick8894/harmonia/api-gw/internal/engine"
	"github.com/Patrick8894/harmonia/api-gw/internal/health"
	"github.com/Patrick8894/harmonia/api-gw/internal/hello"
	"github.com/Patrick8894/harmonia/api-gw/internal/logic"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(
	r *gin.Engine,
	cfg config.Config,
	engSvc *engine.Service,
	lgSvc *logic.Service,
	healthCtrl *health.Controller,
	helloCtrl *hello.Controller,
	authCtrl *auth.Controller,
	authStore *auth.Store,
) {
	// Global auth middleware to parse cookie (non-fatal)
	r.Use(auth.Middleware(cfg.CookieName, authStore))

	api := r.Group("/api")

	// Auth
	authCtrl.Register(api)

	// Health + Hello
	hello.Register(api, helloCtrl)
	health.Register(api, healthCtrl)

	// Protected feature groups
	engineParent := api.Group("", auth.RequireAuth(cfg.CookieName, authStore))
	logicParent := api.Group("", auth.RequireAuth(cfg.CookieName, authStore))

	// Features
	engine.Register(engineParent, engSvc)
	logic.Register(logicParent, lgSvc)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
