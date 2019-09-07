package service

import (
	"github.com/xiaozefeng/goserver/model"
	"github.com/xiaozefeng/goserver/vo/uservo"
)

func ListUser(username string, offset,limit int ) ([]*uservo.UserInfo ,uint64, error) {
	results:= make([]*uservo.UserInfo, 0)
	models, total, err := model.ListUser(username, offset, limit)
	for _,u := range models {
		results = append(results, &uservo.UserInfo{
			Username:u.Username,
			Password:u.Password,
			CreatedTime:u.CreatedTime.Format("2006-01-02 15:04:05"),
			UpdatedTime:u.UpdateTime.Format("2006-01-02 15:04:05"),
			Id:u.Id,
		})
	}
	return results, total, err
}
