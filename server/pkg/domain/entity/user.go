package entity

import (
	"strings"
	"time"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
)

type User struct {
	UUID            UUID   `gorm:"column:uuid;primaryKey;size:191"`
	UserID          UserID `gorm:"column:user_id;not null;unique"`
	NickName        string `gorm:"column:nick_name;not null"`
	ProfileImageURI string `gorm:"column:profile_image_uri"`
	Email           Email  `gorm:"column:email;unique;not null"`
	Description     string `gorm:"column:description"`
	SocialLink      string `gorm:"column:social_link"`
	Gender          int8   `gorm:"column:gender"`
	IsOfficial      bool   `gorm:"column:is_official"`
	SendMailStatus  int8   `gorm:"column:send_mail_status"`
	CustomerID      string `gorm:"column:customer_id;unique"`
	Model
}

const (
	userIDPrefix = "user"
)

func (u User) Valid() bool {
	return u.Email != "" && u.UserID.ToStr() != "" && u.NickName != "" && u.UUID.ToStr() != ""
}

func (u User) IsEmpty() bool {
	return u.UUID.ToStr() == ""
}

type UUID string

func GetUUID(s string) (*UUID, e.Err) {
	if !IsUUID(s) {
		return nil, e.User.InvalidParameter
	}
	uuid := UUID(s)
	return &uuid, nil
}

func (id UUID) ToStr() string {
	return string(id)
}

func IsUUID(id string) bool {
	slice := strings.Split(id, idSeparater)
	return slice[0] == userIDPrefix
}

func (u *User) ApplyUUID() {
	elem := []string{
		u.UserID.ToStr(),
		time.Now().String(),
	}
	uuid, _ := GetUUID(userIDPrefix + idSeparater + createShaID(elem))
	u.UUID = *uuid
	return
}

func (user *User) IsAlreadyHasCustomerID() bool {
	return user.CustomerID != ""
}

func (u User) EmailSendEnabled() bool {
	return u.SendMailStatus == 1
}

type Users []User

type UserID string

func GetUserID(s string) (*UserID, e.Err) {
	if s == "" {
		return nil, e.User.InvalidParameter
	}
	ID := UserID(s)
	return &ID, nil
}

func (u UserID) ToStr() string {
	return string(u)
}

type Email string

func GetUserEmail(s string) (*Email, e.Err) {
	if s == "" {
		return nil, e.User.InvalidParameter
	}
	mail := Email(s)
	return &mail, nil
}

func (e Email) ToStr() string {
	return string(e)
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

func ParseUserIDs(ids []string) (*[]UserID, e.Err) {
	var users []UserID
	var err e.Err = nil
	for n := range ids {
		id := ids[n]
		userID, er := GetUserID(id)
		if er != nil {
			err = er
			break
		}
		users = append(users, *userID)
	}
	return &users, err
}

func (us Users) ToSlice() []User {
	var users []User
	for _, v := range us {
		users = append(users, v)
	}
	return users
}

func ToUsers(users *[]User) *Users {
	var res Users = *users
	return &res
}
