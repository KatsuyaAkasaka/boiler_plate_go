package entry

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/repository"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/infra/dao"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db      *gorm.DB
	userDao *dao.User
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userRepository{
		db:      db,
		userDao: dao.NewUserDao(),
	}
}

func (r userRepository) Create(model *entity.User) (*entity.User, e.Err) {
	data, err := transact(r.db, func(tx *gorm.DB) (interface{}, e.Err) {
		afterUser, err := r.userDao.Create(tx, model)
		return afterUser, err
	})
	if err != nil {
		return nil, err
	}
	return data.(*entity.User), nil
}

func (r userRepository) FindByUserID(id entity.UserID) (*entity.User, e.Err) {
	var user entity.User
	dbc := r.db.Where("user_id = ?", id.GetUserIDStr()).First(&user)
	if dbc.Error != nil {
		return nil, e.User.CheckDBError(dbc.Error)
	}
	return &user, nil
}

func (r userRepository) Update(model *entity.User) (*entity.User, e.Err) {
	data, err := transact(r.db, func(tx *gorm.DB) (interface{}, e.Err) {
		afterUser, err := r.userDao.Update(tx, model)
		return afterUser, err

	})
	if err != nil {
		return nil, err
	}
	return data.(*entity.User), nil
}

func (r userRepository) Delete(id entity.UserID) e.Err {
	_, err := transact(r.db, func(tx *gorm.DB) (interface{}, e.Err) {
		err := r.userDao.Delete(tx, &id)
		return nil, err
	})
	return err
}
