package output

import e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"

type TokenRes struct {
	Token string `json:"token"`
}

func CreateTokenRes(token string) (*TokenRes, e.Err) {
	return &TokenRes{
		Token: token,
	}, nil
}
