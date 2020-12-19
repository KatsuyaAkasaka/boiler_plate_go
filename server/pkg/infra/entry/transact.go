package entry

import (
	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/error"
	"github.com/jinzhu/gorm"
)

func transact(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, e.Err)) (data interface{}, err e.Err) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, e.SQL.FailedToBeginTx
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			db = tx.Commit()
			if db.Error != nil {
				err = e.SQL.FailedToCommitTx
			}
		}
	}()
	data, err = txFunc(tx)
	return
}
