package entry

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/charge"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/cloud"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/datastore"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
)

func NewRepository(db *datastore.DBRepo, chargeRepo *charge.Repo, cloudRepo *cloud.Repo) *repository.Repositories {
	return &repository.Repositories{
		User:   NewUserRepository(db, chargeRepo),
		System: NewSystemRepository(db, cloudRepo),
	}
}
