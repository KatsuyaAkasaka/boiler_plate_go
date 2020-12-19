package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID          string `gorm:"column:user_id;unique;not null"`
	NickName        string `gorm:"column:nick_name;not null"`
	ProfileImageURI string `gorm:"column:profile_image_uri"`
	Email           string `gorm:"column:email;unique;not null"`
	Description     string `gorm:"column:description"`
	SocialLink      string `gorm:"column:social_link"`
	Gender          int32  `gorm:"column:gender"`
	IdentifyStatus  int32  `gorm:"column:identify_status"`
	CustomerID      string `gorm:"column:customer_id"`
}

type Users []User

type UserID string

func GetUserID(s string) UserID {
	return UserID(s)
}

func (u UserID) GetUserIDStr() string {
	return string(u)
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
