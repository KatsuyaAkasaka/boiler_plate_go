package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBInfo struct {
	UserName string
	Password string
	Host     string
	Port     string
	Name     string
}

func InitDB(info *DBInfo) *gorm.DB {
	dbcon, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", info.UserName, info.Password, info.Host, info.Port, info.Name))
	if err != nil {
		log.Fatalln(err)
	}

	return dbcon
}
