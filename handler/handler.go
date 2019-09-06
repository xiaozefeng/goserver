package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/goserver/pkg/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//always return http.StatusOk
	c.JSON(
		http.StatusOK,
		Response{
			Code:    code,
			Message: message,
			Data:    data,
		},
	)
}
