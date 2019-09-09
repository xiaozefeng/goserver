package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/goserver/handler"
	"github.com/xiaozefeng/goserver/pkg/errno"
	"github.com/xiaozefeng/goserver/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse the json web token
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
