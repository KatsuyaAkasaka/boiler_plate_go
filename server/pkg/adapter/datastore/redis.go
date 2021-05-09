package datastore

import (
	"time"

	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/go-redis/redis"
)

type RedisRepo struct {
	Primary  *redis.Client
	ReadOnly *redis.Client
	Keys     Keys
}
type Keys struct {
	MovieHotRanking string
}

func InitRedis() *RedisRepo {
	redisConf := config.GetConf().Redis
	primary := redis.NewClient(&redis.Options{
		Addr: redisConf.Primary,
		DB:   0,
	})
	readOnly := redis.NewClient(&redis.Options{
		Addr: redisConf.ReadOnly,
		DB:   0,
	})
	_, err := primary.Ping().Result()
	if err != nil {
		log.Errorf("[redis] primary connection failed. host: %s", redisConf.Primary)
	} else {
		log.Infof("[redis] primary successfuly conneced. host: %s", redisConf.Primary)
	}
	_, err = readOnly.Ping().Result()
	if err != nil {
		log.Errorf("[redis] readOnly connection failed. host: %s", redisConf.ReadOnly)
	} else {
		log.Infof("[redis] primary successfuly conneced. host: %s", redisConf.ReadOnly)
	}
	return &RedisRepo{
		Primary:  primary,
		ReadOnly: readOnly,
		Keys: Keys{
			MovieHotRanking: "movie:trend",
		},
	}
}

func CalcMovieHotExpireTime() time.Duration {
	t := time.Now()
	ut := t.Unix()
	_, offset := t.Zone()
	day := time.Unix(((ut+int64(offset))/86400)*86400-int64(offset), 0).In(t.Location())

	tomorrow0AM := day.AddDate(0, 0, 1)
	return tomorrow0AM.Sub(t)
}
