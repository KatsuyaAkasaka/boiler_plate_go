package entry

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/charge"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/datastore"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/infra/dao"
	"gorm.io/gorm"
)

type userRepository struct {
	db         *datastore.DBRepo
	chargeRepo *charge.Repo
	userDao    *dao.User
}

func NewUserRepository(db *datastore.DBRepo, chargeRepo *charge.Repo) repository.UserRepository {
	return userRepository{
		db:         db,
		chargeRepo: chargeRepo,
		userDao:    dao.NewUserDao(),
	}
}

func (r userRepository) Create(model *entity.User) (*entity.User, e.Err) {
	// stripeのカスタマーID作成
	data, err := transact(r.db.Conn, func(tx *gorm.DB) (interface{}, e.Err) {
		customerID, err := r.chargeRepo.CreateCustomer(&model.UUID, &model.Email)
		if err != nil {
			return nil, err
		}
		model.CustomerID = customerID
		afterUser, err := r.userDao.Create(tx, model)
		if err != nil {
			return nil, err
		}
		return afterUser, err
	})
	if err != nil {
		return nil, err
	}
	return data.(*entity.User), nil
}

func (r userRepository) FindByUUID(id *entity.UUID) (*entity.User, e.Err) {
	return r.userDao.FetchByUUID(r.db.Conn, id)
}

func (r userRepository) FindByUserID(id *entity.UserID) (*entity.User, e.Err) {
	return r.userDao.FetchByUserID(r.db.Conn, id)
}

func (r userRepository) FindByEmail(email *entity.Email) (*entity.User, e.Err) {
	return r.userDao.FetchByEmail(r.db.Conn, email)
}

func (r userRepository) Update(model *entity.User) (*entity.User, e.Err) {
	data, err := transact(r.db.Conn, func(tx *gorm.DB) (interface{}, e.Err) {
		err := r.userDao.Update(tx, model)
		if err != nil {
			return nil, err
		}
		user, err := r.userDao.FetchByUserID(tx, &model.UserID)
		if err != nil {
			return nil, err
		}
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	return data.(*entity.User), nil
}

func (r userRepository) ToggleEmailNotify(id *entity.UUID) (*entity.User, e.Err) {
	data, err := transact(r.db.Conn, func(tx *gorm.DB) (interface{}, e.Err) {
		err := r.userDao.ToggleEmailNotify(tx, id)
		if err != nil {
			return nil, err
		}
		user, err := r.userDao.FetchByUUID(tx, id)
		if err != nil {
			return nil, err
		}
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	return data.(*entity.User), nil
}

func (r userRepository) Delete(uuid *entity.UUID) e.Err {
	_, err := transact(r.db.Conn, func(tx *gorm.DB) (interface{}, e.Err) {
		_, err := r.userDao.FetchByUUID(tx, uuid)
		if err != nil {
			return nil, err
		}
		return nil, r.userDao.Remove(tx, uuid)
	})
	return err
}

func (r userRepository) SignIn(email *entity.Email) (*entity.User, e.Err) {
	user, err := transact(r.db.Conn, func(tx *gorm.DB) (interface{}, e.Err) {
		user, err := r.userDao.FetchByEmail(tx, email)
		if err != nil {
			return nil, err
		}
		return user, nil
	})
	return user.(*entity.User), err
}
