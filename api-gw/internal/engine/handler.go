package engine

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

func Register(rg *gin.RouterGroup, ctrl *Controller) {
	g := rg.Group("/engine")
	g.GET("/hello", ctrl.Hello)
	g.POST("/pi", ctrl.Pi)
	g.POST("/matmul", ctrl.MatMul)
	g.POST("/stats", ctrl.Stats)
}

// HelloEngineRPC godoc
// @Summary      Call C++ EngineService Hello RPC
// @Description  Triggers the Hello RPC on the C++ Thrift EngineService
// @Tags         engine
// @Produce      json
// @Param        name  query  string  false  "Name to greet"  default(World)
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /engine/hello [get]
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

// Pi godoc
// @Summary      Estimate π via Monte Carlo
// @Description  Calls EngineService.EstimatePi with given sample size
// @Tags         engine
// @Accept       json
// @Produce      json
// @Param        payload  body  PiDTO  true  "Pi input"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /engine/pi [post]
func (c *Controller) Pi(ctx *gin.Context) {
	var req PiDTO
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Samples <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 4*time.Second)
	defer cancel()

	resp, err := c.svc.EstimatePi(reqCtx, req.Samples)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"pi":     resp.GetPi(),
		"inside": resp.GetInside(),
		"total":  resp.GetTotal(),
		"seed":   resp.GetSeed(),
	})
}

// MatMul godoc
// @Summary      Matrix multiply
// @Description  Calls EngineService.MatMul with two matrices A and B
// @Tags         engine
// @Accept       json
// @Produce      json
// @Param        payload  body  MatMulDTO  true  "A and B matrices"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /engine/matmul [post]
func (c *Controller) MatMul(ctx *gin.Context) {
	var req MatMulDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	// basic validation before RPC
	if req.A.Cols != req.B.Rows {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dimension mismatch: A.cols must equal B.rows"})
		return
	}
	if int64(req.A.Rows)*int64(req.A.Cols) != int64(len(req.A.Data)) ||
		int64(req.B.Rows)*int64(req.B.Cols) != int64(len(req.B.Data)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "data length must equal rows*cols for A and B"})
		return
	}

	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 6*time.Second)
	defer cancel()

	resp, err := c.svc.MatMul(reqCtx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	C := resp.GetC()
	ctx.JSON(http.StatusOK, gin.H{
		"c": gin.H{
			"rows": C.GetRows(),
			"cols": C.GetCols(),
			"data": C.GetData(),
		},
	})
}

// Stats godoc
// @Summary      Compute vector statistics
// @Description  Calls EngineService.ComputeStats on a dataset (sample variance by default)
// @Tags         engine
// @Accept       json
// @Produce      json
// @Param        payload  body  StatsDTO  true  "Stats input"
// @Success      200      {object}  map[string]any
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /engine/stats [post]
func (c *Controller) Stats(ctx *gin.Context) {
	var req StatsDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	sample := true
	if req.Sample != nil {
		sample = *req.Sample
	}
	reqCtx, cancel := context.WithTimeout(ctx.Request.Context(), 3*time.Second)
	defer cancel()

	resp, err := c.svc.ComputeStats(reqCtx, StatsDTO{Data: req.Data, Sample: &sample})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"count":    resp.GetCount(),
		"sum":      resp.GetSum(),
		"mean":     resp.GetMean(),
		"variance": resp.GetVariance(),
		"stddev":   resp.GetStddev(),
		"min":      resp.GetMin(),
		"max":      resp.GetMax(),
	})
}
