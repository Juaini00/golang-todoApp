package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_app/pkg/utils"
)

// Test handler returns a simple message.
// @Summary Test endpoint
// @Tags test
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, utils.BuildResponse(http.StatusOK, "Hello World!", nil))
}
