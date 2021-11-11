package model

import (
	"fmt"
	"sync"

	"github.com/leetpy/cactus/pkg/auth"
	"github.com/leetpy/cactus/pkg/constvar"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "user"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.DB.Create(&u).Error
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// // Validate the fields.
// func (u *UserModel) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(u)
// }

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.DB.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count int64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.DB.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.DB.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}
