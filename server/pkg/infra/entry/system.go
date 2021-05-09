package entry

import (
	"mime/multipart"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/cloud"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/datastore"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
)

type systemRepository struct {
	db         *datastore.DBRepo
	uploadRepo *cloud.Repo
}

func NewSystemRepository(db *datastore.DBRepo, cr *cloud.Repo) repository.SystemRepository {
	return systemRepository{
		db:         db,
		uploadRepo: cr,
	}
}

func (r systemRepository) UploadImage(fileHeader *multipart.FileHeader) (string, e.Err) {
	return r.uploadRepo.UploadImageToS3(fileHeader)
}
