package config

import (
	"os"
	"path/filepath"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/logger"
	"github.com/spf13/viper"
)

type ConfMap struct {
	Api    map[string]interface{}
	DB     map[string]interface{}
	Stripe map[string]interface{}
}

func GetConf(env string) *ConfMap {
	rootPath := os.Getenv("ROOT_PATH")
	viper.AddConfigPath(filepath.Join(rootPath, "server", "pkg", "config"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("common")
	err := viper.ReadInConfig()
	if err != nil {
		log.GetLogger(env).Panic(err)
	}
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	err = viper.MergeInConfig()
	if err != nil {
		log.GetLogger(env).Panic(err)
	}
	conf := ConfMap{
		Api:    viper.GetStringMap("api"),
		DB:     viper.GetStringMap("db"),
		Stripe: viper.GetStringMap("stripe"),
	}
	return &conf
}
