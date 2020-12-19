package dao

import (
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/jinzhu/gorm"
)

type User struct{}

func NewUserDao() *User {
	return &User{}
}

func (d *User) Create(tx *gorm.DB, model *entity.User) (*entity.User, e.Err) {
	if err := tx.Create(model).Error; err != nil {
		return nil, e.User.CheckDBError(err)
	}
	return model, nil
}

func (d *User) Delete(tx *gorm.DB, id *entity.UserID) e.Err {
	if err := tx.Delete(&entity.User{}, "user_id = ?", id.GetUserIDStr()).Error; err != nil {
		return e.User.CheckDBError(err)
	}
	return nil
}

func (d *User) Update(tx *gorm.DB, model *entity.User) (*entity.User, e.Err) {
	if err := tx.Model(&entity.User{}).Where("user_id = ?", model.UserID).Update(&model).Error; err != nil {
		return nil, e.User.CheckDBError(err)
	}
	return model, nil
}
