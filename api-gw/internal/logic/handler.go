// internal/logic/handler.go
package logic

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller { return &Controller{svc: svc} }

// Wire routes to named methods.
func Register(rg *gin.RouterGroup, ctrl *Controller) {
	g := rg.Group("/logic")
	g.GET("/hello", ctrl.Hello)
	g.POST("/eval", ctrl.Evaluate)
	g.POST("/transform", ctrl.Transform)
	g.POST("/plan", ctrl.Plan)
}

// HelloLogicRPC godoc
// @Summary      Call Python LogicService Hello RPC
// @Description  Triggers the Hello RPC on the Python gRPC LogicService
// @Tags         logic
// @Produce      json
// @Param        name  query  string  false  "Name to greet"  default(World)
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /logic/hello [get]
func (c *Controller) Hello(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "World")
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	msg, err := c.svc.Hello(reqCtx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": msg})
}

// Evaluate godoc
// @Summary      Evaluate expression
// @Description  Evaluate a numeric expression with optional variables via LogicService.Evaluate
// @Tags         logic
// @Accept       json
// @Produce      json
// @Param        payload  body  EvalDTO  true  "Eval input"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /logic/eval [post]
func (c *Controller) Evaluate(ctx *gin.Context) {
	var req EvalDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second)
	defer cancel()

	resp, fromCache, err := c.svc.Evaluate(reqCtx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": resp.GetResult(),
		"error":  resp.GetError(),
		"cached": fromCache,
	})
}

// Transform godoc
// @Summary      Transform dataset
// @Description  Apply MAP/FILTER/SUM with an optional expression/var on numeric data via LogicService.Transform
// @Tags         logic
// @Accept       json
// @Produce      json
// @Param        payload  body  TransformDTO  true  "Transform input"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /logic/transform [post]
func (c *Controller) Transform(ctx *gin.Context) {
	var req TransformDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()

	resp, fromCache, err := c.svc.Transform(reqCtx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":   resp.GetData(),
		"result": resp.GetResult(),
		"error":  resp.GetError(),
		"cached": fromCache,
	})
}

// Plan godoc
// @Summary      Create a task plan
// @Description  Generate a step plan from a goal (+ optional hints) via LogicService.PlanTasks
// @Tags         logic
// @Accept       json
// @Produce      json
// @Param        payload  body  PlanDTO  true  "Plan input"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /logic/plan [post]
func (c *Controller) Plan(ctx *gin.Context) {
	var req PlanDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 8*time.Second)
	defer cancel()

	resp, cached, err := c.svc.PlanTasks(reqCtx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"tasks":  resp.GetTasks(),
		"notes":  resp.GetNotes(),
		"error":  resp.GetError(),
		"cached": cached,
	})
}
