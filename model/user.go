package model

import (
	"fmt"
	"github.com/xiaozefeng/goserver/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username; not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password; not null" binding:"required" validate:"min=5,max=128"`
}

func (c UserModel) TableName() string {
	return "tb_user"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func DeleteUserById(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

func GetUserByUsername(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ? ", username).First(&u)
	return u, d.Error
}

func GetUserByBy(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ? ", id).First(&u)
	return u, d.Error
}

func (u *UserModel) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}

func (u *UserModel) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = 10
	}
	where := fmt.Sprintf("username like '%%%s%%'", username)
	var total uint64
	users := make([]*UserModel, 0)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&total).Error; err != nil {
		return users, 0, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, 0, err
	}
	return users, total, nil
}
