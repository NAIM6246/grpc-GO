package conn

import (
	"fmt"
	"sync"

	"github.com/naim6246/grpc-GO/product/config"
	"github.com/naim6246/grpc-GO/product/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var (
	connDBOnce sync.Once
	dbInstance *DB
)

func ConnectDB(config *config.DBConfig) *DB {
	connDBOnce.Do(func() {
		err := connectDB(config)
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("Database Connected successfully")
	return dbInstance
}

func connectDB(config *config.DBConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dbInstance = &DB{conn}
	return nil
}

func (db *DB) Migration() {
	db.AutoMigrate(&models.Product{})
}
