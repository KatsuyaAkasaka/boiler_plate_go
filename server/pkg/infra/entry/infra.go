package entry

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	"github.com/jinzhu/gorm"
)

func NewRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		User: NewUserRepository(db),
	}
}
