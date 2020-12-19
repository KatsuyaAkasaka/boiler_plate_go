package input

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
)

type UserReq struct {
	UserID          string `json:"user_id"`
	NickName        string `json:"nick_name"`
	ProfileImageURI string `json:"profile_image_uri"`
	Email           string `json:"email"`
	Description     string `json:"description"`
	SocialLink      string `json:"social_link"`
	Gender          int32  `json:"gender"`
	IdentifyStatus  int32  `json:"identify_status"`
	CustomerID      int32  `json:"customer_id"`
}

func (r *UserReq) GetUser() *entity.User {
	return &entity.User{
		UserID:          r.UserID,
		NickName:        r.NickName,
		ProfileImageURI: r.ProfileImageURI,
		Email:           r.Email,
		Description:     r.Description,
		SocialLink:      r.SocialLink,
		Gender:          r.Gender,
		IdentifyStatus:  r.IdentifyStatus,
	}
}

func (r *UserReq) BindUserID(id string) {
	r.UserID = id
}
