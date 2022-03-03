package config

import "sync"

type DBConfig struct {
	Server   string
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
		Server:   "",
		Port:     0,
		DBName:   "UserDatabase",
		User:     "",
		Password: "",
	}
}