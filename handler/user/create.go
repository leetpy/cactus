package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/leetpy/cactus/handler"
	"github.com/leetpy/cactus/model"
	"github.com/leetpy/cactus/pkg/errno"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	// if err := u.Validate(); err != nil {
	// 	SendResponse(c, errno.ErrValidation, nil)
	// 	return
	// }

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
