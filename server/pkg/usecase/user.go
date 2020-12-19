package usecase

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/input"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func newUserUsecase(ur repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu UserUsecase) CreateUser(user *input.UserReq) (*entity.User, e.Err) {
	afterUser, err := uu.userRepo.Create(user.GetUser())
	if err != nil {
		return nil, err
	}
	return afterUser, nil
}

func (uu UserUsecase) GetUser(id string) (*entity.User, e.Err) {
	afterUser, err := uu.userRepo.FindByUserID(entity.GetUserID(id))
	if err != nil {
		return nil, err
	}
	return afterUser, nil
}

func (uu UserUsecase) UpdateUser(user *input.UserReq) (*entity.User, e.Err) {
	afterUser, err := uu.userRepo.Update(user.GetUser())
	if err != nil {
		return nil, err
	}
	return afterUser, nil
}

func (uu UserUsecase) DeleteUser(id string) e.Err {
	err := uu.userRepo.Delete(entity.GetUserID(id))
	return err
}
