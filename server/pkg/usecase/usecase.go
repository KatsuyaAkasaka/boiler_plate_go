package usecase

import "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"

type Usecases struct {
	User   *UserUsecase
	System *SystemUsecase
}

func NewUsecase(r *repository.Repositories) *Usecases {
	return &Usecases{
		User:   newUserUsecase(r.User),
		System: newSystemUsecase(r.System),
	}
}
