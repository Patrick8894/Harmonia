package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello godoc
// @Summary      Hello endpoint
// @Description  Basic greeting from Harmonia API Gateway
// @Tags         root
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /hello [get]
func (c *Controller) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello from Harmonia API Gateway!"})
}
