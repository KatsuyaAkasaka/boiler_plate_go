package output

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
)

type UserSummaryRes struct {
	UserID          string `json:"user_id"`
	NickName        string `json:"nick_name"`
	ProfileImageURI string `json:"profile_image_uri"`
	IsOfficial      bool   `json:"is_official"`
}

type UserRes struct {
	UserID          string `json:"user_id"`
	UUID            string `json:"uuid"`
	NickName        string `json:"nick_name"`
	ProfileImageURI string `json:"profile_image_uri"`
	Email           string `json:"email"`
	Description     string `json:"description"`
	SocialLink      string `json:"social_link"`
	Gender          int32  `json:"gender"`
	IsOfficial      bool   `json:"is_official"`
	SendMailStatus  int32  `json:"send_mail_status"`
	CustomerID      string `json:"customer_id"`
	entity.Model
}

type UserWithTokenRes struct {
	User  UserRes `json:"user"`
	Token string  `json:"token"`
}

type MeRes struct {
	User UserRes `json:"user"`
}

type OthersRes struct {
	User UserRes `json:"user"`
}

type EmailNotifyRes struct {
	SendMailStatus int8 `json:"send_mail_status"`
}

func CreateUserSummaryRes(
	user *entity.User,
) (*UserSummaryRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	var u UserSummaryRes
	if err := entity.Decode(user, &u); err != nil {
		return nil, e.User.InvalidParameter
	}
	return &u, nil
}

func CreateUserRes(
	user *entity.User,
) (*UserRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	var u UserRes
	if err := entity.Decode(user, &u); err != nil {
		return nil, e.User.InvalidParameter
	}
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	u.DeletedAt = user.DeletedAt
	return &u, nil
}

func CreateUserWithTokenRes(
	user *entity.User,
	token string,
) (*UserWithTokenRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	userRes, err := CreateUserRes(user)
	if err != nil {
		return nil, err
	}
	tokenRes, err := CreateTokenRes(token)
	if err != nil {
		return nil, err
	}
	return &UserWithTokenRes{
		User:  *userRes,
		Token: tokenRes.Token,
	}, nil
}

func CreateMeRes(
	user *entity.User,
) (*MeRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	userDetail, err := CreateUserRes(user)
	if err != nil {
		return nil, err
	}
	return &MeRes{
		User: *userDetail,
	}, nil
}

func CreateOthersRes(
	user *entity.User,
) (*OthersRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	userDetail, err := CreateUserRes(user)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &OthersRes{
		User: *userDetail,
	}, nil
}

func CreateEmailNotifyRes(
	user *entity.User,
) (*EmailNotifyRes, e.Err) {
	if !user.Valid() {
		return nil, e.User.InvalidParameter
	}

	return &EmailNotifyRes{
		SendMailStatus: user.SendMailStatus,
	}, nil
}
