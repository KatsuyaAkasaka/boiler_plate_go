package repository

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
)

type UserRepository interface {
	Create(model *entity.User) (*entity.User, e.Err)
	FindByUserID(id entity.UserID) (*entity.User, e.Err)
	Update(model *entity.User) (*entity.User, e.Err)
	Delete(id entity.UserID) e.Err
}
