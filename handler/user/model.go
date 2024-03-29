package user

import (
	"github.com/xiaozefeng/goserver/vo/uservo"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	UserList   []*uservo.UserInfo `json:"userList"`
}
