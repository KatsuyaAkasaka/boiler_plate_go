package config

import "os"

type RedisInfo struct {
	Primary  string
	ReadOnly string
}

func parseRedisConf(redisConf map[string]interface{}) *RedisInfo {
	primary := os.Getenv("REDIS_PRIMARY")
	if primary == "" {
		primary = redisConf["primary"].(string)
	}
	readOnly := os.Getenv("REDIS_READONLY")
	if readOnly == "" {
		readOnly = redisConf["read_only"].(string)
	}
	return &RedisInfo{
		Primary:  primary,
		ReadOnly: readOnly,
	}
}
