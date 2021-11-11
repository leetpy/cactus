package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/leetpy/cactus/handler"
	"github.com/leetpy/cactus/model"
	"github.com/leetpy/cactus/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
