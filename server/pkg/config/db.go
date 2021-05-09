package config

import "os"

type DBInfo struct {
	UserName string
	Password string
	Host     string
	Name     string
	Port     string
	LogMode  int
}

func parseDBConf(dbConf map[string]interface{}) *DBInfo {
	username := os.Getenv("DB_USER")
	if username == "" {
		username = dbConf["user"].(string)
	}
	password := os.Getenv("DB_PASS")
	if password == "" {
		password = dbConf["password"].(string)
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = dbConf["host"].(string)
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = dbConf["port"].(string)
	}
	return &DBInfo{
		UserName: username,
		Password: password,
		Port:     port,
		Host:     host,
		Name:     dbConf["name"].(string),
		LogMode:  dbConf["log_mode"].(int),
	}
}
