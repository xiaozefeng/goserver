package model

import "gopkg.in/go-playground/validator.v9"

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

func (u *UserModel) Compare(pwd string) bool {
	return false
}

func (u *UserModel) Encrypt() error {
	return nil
}

func (u *UserModel) Validate() error {
	v := validator.New()
	return v.Struct(u)
}
