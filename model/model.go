package model

import "time"

type BaseModel struct {
	Id          uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdateTime  time.Time `gorm:"column:updated_time" json:"updatedTime"`
	DeletedTime time.Time `gorm:"column:deleted_time" json:"deletedTime"`
}

type UserInfo struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

