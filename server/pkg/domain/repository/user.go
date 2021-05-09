package repository

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
)

type UserRepository interface {
	Create(model *entity.User) (*entity.User, e.Err)
	FindByUUID(id *entity.UUID) (*entity.User, e.Err)
	FindByEmail(email *entity.Email) (*entity.User, e.Err)
	FindByUserID(id *entity.UserID) (*entity.User, e.Err)
	Update(model *entity.User) (*entity.User, e.Err)
	ToggleEmailNotify(id *entity.UUID) (*entity.User, e.Err)
	Delete(id *entity.UUID) e.Err
	SignIn(email *entity.Email) (*entity.User, e.Err)
}
