package main

import (
	"os"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/database"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/handler"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/infra/entry"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
)

func main() {
	env := os.Getenv("ENV")
	log.NewLogger(env)
	conf := config.GetConf(env)

	apiConf := conf.Api
	dbConf := conf.DB
	// stripeConf := conf.Stripe

	dbInfo := database.DBInfo{
		UserName: dbConf["user"].(string),
		Password: dbConf["password"].(string),
		Host:     dbConf["host"].(string),
		Port:     dbConf["port"].(string),
		Name:     dbConf["name"].(string),
	}
	db := database.InitDB(&dbInfo)
	defer db.Close()

	if env == "local" {
		db.LogMode(true)
	}

	repos := entry.NewRepository(db)
	uc := usecase.NewUsecase(repos)
	port := apiConf["port"].(string)
	ver := apiConf["version"].(string)
	handler.Start(uc, port, ver)
}
