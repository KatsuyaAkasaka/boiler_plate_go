package datastore

import (
	"fmt"
	"time"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBRepo struct {
	Conn *gorm.DB
}

var count = 0

func Seed(db *gorm.DB) {
}

func InitDB() *DBRepo {
	dbConf := config.GetConf().DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
	newLogger := logger.Default.LogMode(logger.LogLevel(dbConf.LogMode))
	var dbCon *gorm.DB
	var er error
	for i := 0; i < 10; i++ {
		dbCon, er = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if er != nil {
			log.Warnf("[RDS] primary Not ready. Retry connecting... host: %s", dsn)
			time.Sleep(time.Second)
			continue
		}
		break
	}
	if er != nil {
		log.Error(er)
		panic(er)
	}
	log.Info("Successfully connected to RDS")
	return &DBRepo{
		Conn: dbCon,
	}
}
