package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/leetpy/cactus/handler"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    nil,
	})
}
