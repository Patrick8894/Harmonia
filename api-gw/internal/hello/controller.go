package hello

import "github.com/gin-gonic/gin"

type Controller struct{}

func New() *Controller { return &Controller{} }

func Register(rg *gin.RouterGroup, c *Controller) {
	rg.GET("/hello", c.Hello)
}
