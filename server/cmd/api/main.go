package main

import (
	"os"

	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/charge"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/cloud"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/datastore"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/handler"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/infra/entry"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/middleware"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/usecase"
)

func main() {
	env := os.Getenv("ENV")
	log.NewLogger(env)
	conf := config.InitConf(env)

	// stripeConf := conf.Stripe

	dbRepo := datastore.InitDB()
	dbConn, _ := dbRepo.Conn.DB()
	defer dbConn.Close()

	// redisRepo := datastore.InitRedis()
	scRepo := charge.InitStripe()
	upRepo := cloud.InitUploader()

	repos := entry.NewRepository(dbRepo, scRepo, upRepo)
	uc := usecase.NewUsecase(repos)
	middles := middleware.NewMiddleware(repos)
	handler.Start(uc, middles, conf.Gateway)
	// handler.Start(uc, port, ver)
}
