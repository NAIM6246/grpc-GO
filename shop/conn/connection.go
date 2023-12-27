package conn

import (
	"fmt"
	"log"
	"sync"

	"github.com/naim6246/grpc-GO/shop/config"
	"github.com/naim6246/grpc-GO/shop/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbType = "sqlite3"

var (
	connDBOnce sync.Once
	dbInstance *DB
)

type DB struct {
	*gorm.DB
}

func connectDB(config *config.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
		return err
	}
	fmt.Println("Database connected successfully.")
	dbInstance = &DB{conn}
	return nil
}

func ConnectDB(config *config.DBConfig) *DB {
	connDBOnce.Do(func() {
		err := connectDB(config)
		if err != nil {
			panic("failed to connect DB: " + err.Error())
		}
	})
	return dbInstance
}

func (db *DB) Migration() {
	db.AutoMigrate(&models.Shop{})
}
