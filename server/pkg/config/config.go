package config

import (
	"os"
	"path/filepath"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/spf13/viper"
)

var confMap *ConfMap

func GetConf() *ConfMap {
	return confMap
}

type ConfMap struct {
	Gateway *GatewayInfo
	DB      *DBInfo
	Stripe  *StripeInfo
	AWS     *AWSInfo
	Redis   *RedisInfo
}

func InitConf(env string) *ConfMap {
	rootPath, _ := os.Getwd()
	viper.AddConfigPath(filepath.Join(rootPath, "pkg", "config", "yaml"))
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
	confMap = &ConfMap{
		Gateway: parseGatewayConf(viper.GetStringMap("gateway")),
		DB:      parseDBConf(viper.GetStringMap("db")),
		AWS:     parseAWSConf(viper.GetStringMap("aws")),
		Stripe:  parseStripeConf(viper.GetStringMap("stripe")),
	}
	return confMap
}
