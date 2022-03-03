package main

import (
	"github.com/naim6246/grpc-GO/user/config"
	"github.com/naim6246/grpc-GO/user/conn"
	"github.com/naim6246/grpc-GO/user/handlers"
	"github.com/naim6246/grpc-GO/user/models"
	"github.com/naim6246/grpc-GO/user/repositories"
	"github.com/naim6246/grpc-GO/user/services"
)

var dbInstance *conn.DB

func init() {
	dbConfig := config.NewDBConfig()
	dbInstance = conn.ConnectDB(dbConfig)
	dbInstance.Migration()
}

func main() {
	userRepository := repositories.NewUserRepository(dbInstance)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	go userService.StartGrpcUserService()
	models.Wg.Add(1)
	go userHandler.Handler()
	models.Wg.Add(1)
	models.Wg.Wait()
}
