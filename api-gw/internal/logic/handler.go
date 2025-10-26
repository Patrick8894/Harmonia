package logic

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HelloLogicRPC godoc
// @Summary      Call Python LogicService Hello RPC
// @Description  Triggers the Hello RPC on the Python gRPC LogicService
// @Tags         logic
// @Produce      json
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /logic/hello [get]
func Register(rg *gin.RouterGroup, svc *Service) {
	g := rg.Group("/logic")

	g.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "World")
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		msg, err := svc.Hello(ctx, name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "RPC failed: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": msg})
	})
}
