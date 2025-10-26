package health

import "github.com/gin-gonic/gin"

type Controller struct{}

func New() *Controller { return &Controller{} }

// Register wires endpoints for /healthz.
func Register(rg *gin.RouterGroup, c *Controller) {
	rg.GET("/healthz", c.Health)
}
