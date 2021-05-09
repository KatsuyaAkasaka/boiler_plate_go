package usecase

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/mail"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/middleware"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/input"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase/output"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func newUserUsecase(
	ur repository.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu UserUsecase) CreateUser(user *input.UserReq) (*output.UserWithTokenRes, e.Err) {
	entityUser, err := user.GetUserEntity()
	if err != nil {
		return nil, err
	}
	entityUser.ApplyUUID()
	// ユーザの作成とstripeのcustomerIDの作成
	afterUser, err := uu.userRepo.Create(entityUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token, err := middleware.GenerateUUIDJWT(&entityUser.UUID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateUserWithTokenRes(afterUser, token)
}

func (uu UserUsecase) GetUserByUUID(id *entity.UUID) (*output.MeRes, e.Err) {
	user, err := uu.userRepo.FindByUUID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateMeRes(user)
}

func (uu UserUsecase) GetUserByUserID(uuid *entity.UUID, id *entity.UserID) (*output.OthersRes, e.Err) {
	user, err := uu.userRepo.FindByUserID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateOthersRes(user)
}

func (uu UserUsecase) UpdateUser(user *input.UserReq) (*output.MeRes, e.Err) {
	entityUser, err := user.GetUserEntity()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	afterUser, err := uu.userRepo.Update(entityUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateMeRes(afterUser)
}

func (uu UserUsecase) ChangeEmailNotify(id *entity.UUID) (*output.EmailNotifyRes, e.Err) {
	user, err := uu.userRepo.ToggleEmailNotify(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return output.CreateEmailNotifyRes(user)
}

func (uu UserUsecase) DeleteUser(id *entity.UUID) e.Err {
	err := uu.userRepo.Delete(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (uu UserUsecase) SignInUser(req *input.SignInReq) (*output.UserWithTokenRes, e.Err) {
	entityEmail, err := req.GetUserEmail()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	entityUser, err := uu.userRepo.SignIn(entityEmail)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token, err := middleware.GenerateEmailJWT(entityEmail)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	mail.SendSignInAuthEmail(entityEmail, token)
	return output.CreateUserWithTokenRes(entityUser, token)
}

func (uu UserUsecase) SignUpUser(req *input.SignUpReq) e.Err {
	entityEmail, err := req.GetUserEmail()
	if err != nil {
		log.Error(err)
		return err
	}
	token, err := middleware.GenerateEmailJWT(entityEmail)
	if err != nil {
		log.Error(err)
		return err
	}
	mail.SendSignUpAuthEmail(entityEmail, token)
	return nil
}
