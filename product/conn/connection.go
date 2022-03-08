package conn

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/naim6246/grpc-GO/product/config"
	"github.com/naim6246/grpc-GO/product/models"

	//sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	*gorm.DB
}

const dbType = "sqlite3"

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
	connectionString := fmt.Sprintf("%s.db", config.DBName)
	conn, err := gorm.Open(dbType, connectionString)
	if err != nil {
		return err
	}
	dbInstance = &DB{conn}
	return nil
}

func (db *DB) Migration() {
	db.AutoMigrate(&models.Product{})
}
