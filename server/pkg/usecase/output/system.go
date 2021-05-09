package output

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
)

type UploadImageRes struct {
	URL string `json:"url"`
}

func CreateUploadImageRes(path string) (*UploadImageRes, e.Err) {
	// 画像アップロードのレスポンスにassets.wantty.appのCDN形式に変換させる
	return &UploadImageRes{
		URL: entity.BuildImagePathForCDN(path),
	}, nil
}
