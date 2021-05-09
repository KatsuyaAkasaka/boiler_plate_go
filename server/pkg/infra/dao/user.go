package dao

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"gorm.io/gorm"
)

type User struct{}

func NewUserDao() *User {
	return &User{}
}

func (d *User) FetchByUUID(tx *gorm.DB, id *entity.UUID) (*entity.User, e.Err) {
	user := entity.User{}
	if err := tx.Where("uuid = ?", id.ToStr()).First(&user).Error; err != nil {
		return nil, e.User.CheckDBError(err, true)
	}
	return &user, nil
}

func (d *User) FetchByUserID(tx *gorm.DB, id *entity.UserID) (*entity.User, e.Err) {
	var user entity.User
	if err := tx.Where("user_id = ?", id.ToStr()).First(&user).Error; err != nil {
		return nil, e.User.CheckDBError(err, true)
	}
	return &user, nil
}

func (d *User) FetchByEmail(tx *gorm.DB, email *entity.Email) (*entity.User, e.Err) {
	var user entity.User
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, e.User.CheckDBError(err, true)
	}
	return &user, nil
}

func (d *User) ToggleEmailNotify(tx *gorm.DB, id *entity.UUID) e.Err {
	if err := tx.Exec("UPDATE users SET send_mail_status = 1 - send_mail_status WHERE uuid = ? and deleted_at is null", id.ToStr()).Error; err != nil {
		return e.User.CheckDBError(err, true)
	}
	return nil
}

func (d *User) Create(tx *gorm.DB, model *entity.User) (*entity.User, e.Err) {
	if err := tx.Create(model).Error; err != nil {
		return nil, e.User.CheckDBError(err, false)
	}
	return model, nil
}

func (d *User) Remove(tx *gorm.DB, uuid *entity.UUID) e.Err {
	if err := tx.Where("uuid = ?", uuid.ToStr()).Delete(&entity.User{}).Error; err != nil {
		return e.User.CheckDBError(err, false)
	}
	return nil
}

func (d *User) Update(tx *gorm.DB, model *entity.User) e.Err {
	err := tx.Updates(&model).Error
	return e.User.CheckDBError(err, true)
}
