package usecase

import (
	"mime/multipart"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/output"
)

type SystemUsecase struct {
	systemRepo repository.SystemRepository
}

func newSystemUsecase(sr repository.SystemRepository) *SystemUsecase {
	return &SystemUsecase{
		systemRepo: sr,
	}
}

func (su SystemUsecase) UploadImage(fileHeader *multipart.FileHeader) (*output.UploadImageRes, e.Err) {
	path, err := su.systemRepo.UploadImage(fileHeader)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateUploadImageRes(path)
}
