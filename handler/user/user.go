package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/xiaozefeng/goserver/handler"
	"github.com/xiaozefeng/goserver/pkg/errno"
)

// Create creates a new user account
func Create(c *gin.Context) {
	var r CreateRequest

	var err error
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).
			Add("this is a add message")
		handler.SendResponse(c, err, nil)
		return
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
		handler.SendResponse(c, err, nil)
		return
	}

	resp := CreateResponse{Username: r.Username}
	handler.SendResponse(c, err, resp)
}
