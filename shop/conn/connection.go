package conn

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/naim6246/grpc-GO/shop/config"
	"github.com/naim6246/grpc-GO/shop/models"

	//sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	connectionString := fmt.Sprintf("%s.db", config.DBName)
	conn, err := gorm.Open(dbType, connectionString)
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
