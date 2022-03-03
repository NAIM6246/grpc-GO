package main

import (
	"github.com/naim6246/grpc-GO/user/config"
	"github.com/naim6246/grpc-GO/user/conn"
	"github.com/naim6246/grpc-GO/user/handlers"
	"github.com/naim6246/grpc-GO/user/models"
	"github.com/naim6246/grpc-GO/user/services"
)

func init() {
	dbConfig := config.NewDBConfig()
	dbInstance := conn.ConnectDB(dbConfig)
	dbInstance.Migration()
}

func main() {

	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	go userService.StartGrpcUserService()
	models.Wg.Add(1)
	go userHandler.Handler()
	models.Wg.Add(1)
	models.Wg.Wait()
}
