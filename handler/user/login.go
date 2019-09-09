package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/goserver/handler"
	"github.com/xiaozefeng/goserver/model"
	"github.com/xiaozefeng/goserver/pkg/auth"
	"github.com/xiaozefeng/goserver/pkg/errno"
	"github.com/xiaozefeng/goserver/pkg/token"
)

// login generates the authentication token
// if the password was matched with the specified account.
func Login(c *gin.Context) {
	var u model.UserInfo
	// bind
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// get account by username from db
	dbUser, err := model.GetUserByUsername(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// compare the password
	if err := auth.Compare(dbUser.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// generate token
	t, err := token.Sign(c, token.Context{
		Id:       dbUser.Id,
		Username: u.Username,
	}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, gin.H{"token": t})
}
