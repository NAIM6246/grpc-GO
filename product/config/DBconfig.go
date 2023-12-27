package config

import (
	"os"
	"sync"
)

type DBConfig struct {
	Port     int
	Host     string
	DBName   string
	User     string
	Password string
}

var (
	loadDBConfig sync.Once
	dbConfig     *DBConfig
)

func NewDBConfig() *DBConfig {
	loadDBConfig.Do(mapDBConfig)
	return dbConfig
}

func mapDBConfig() {
	dbConfig = &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     3306,
		DBName:   "ProductDatabase",
		User:     "root",
		Password: "",
	}
}
