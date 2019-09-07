package uservo

type UserInfo struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedTime string  `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}
