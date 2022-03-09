package config

import "sync"

type DBConfig struct {
	Port     int
	Server   string
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
		Port:     0,
		Server:   "",
		DBName:   "ProductDatabase",
		User:     "",
		Password: "",
	}
}
