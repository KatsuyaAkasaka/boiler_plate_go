package repository

import (
	"mime/multipart"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
)

type SystemRepository interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, e.Err)
}
