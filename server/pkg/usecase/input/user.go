package input

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
)

type UserReq struct {
	UUID            string `json:"uuid"`
	UserID          string `json:"user_id"`
	NickName        string `json:"nick_name"`
	ProfileImageURI string `json:"profile_image_uri"`
	Email           string `json:"email"`
	Description     string `json:"description"`
	SocialLink      string `json:"social_link"`
	Gender          int8   `json:"gender"`
	IsOfficial      bool   `json:"is_official"`
	SendMailStatus  int8   `json:"send_mail_status"`
}

func (r UserReq) Valid() bool {
	return r.UserID != "" && r.NickName != "" && r.Email != ""
}

func (r UserReq) GetUserEntity() (*entity.User, e.Err) {
	if !r.Valid() {
		return nil, e.System.BadRequest
	}
	user := &entity.User{}
	if err := entity.Decode(r, user); err != nil {
		return nil, e.System.BadRequest
	}
	return user, nil
}

func (r *UserReq) BindUUID(id *entity.UUID) {
	r.UUID = id.ToStr()
}

func (r *UserReq) BindEmail(email *entity.Email) {
	r.Email = email.ToStr()
}

type UserIdReq string

func (r UserIdReq) GetUserId() (*entity.UserID, e.Err) {
	return entity.GetUserID(string(r))
}

type SignInReq struct {
	Email string `json:"email"`
}

type SignUpReq struct {
	Email string `json:"email"`
}

func (r *SignInReq) GetUserEmail() (*entity.Email, e.Err) {
	return entity.GetUserEmail(r.Email)
}

func (r *SignUpReq) GetUserEmail() (*entity.Email, e.Err) {
	return entity.GetUserEmail(r.Email)
}

type EmailNotifyReq struct {
	Status int8 `json:"send_mail_status"`
}
