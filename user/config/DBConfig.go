package config

import (
	"os"
	"sync"
)

type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func NewDBConfig() *DBConfig {
	var loadDBOnce sync.Once
	loadDBOnce.Do(mapDBConfig)
	return dbConfig
}

var dbConfig *DBConfig

func mapDBConfig() {
	dbConfig = &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     3306,
		DBName:   "UserDatabase",
		User:     "root",
		Password: "",
	}
}
