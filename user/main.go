package main

import (
	"github.com/naim6246/grpc-GO/user/handlers"
	"github.com/naim6246/grpc-GO/user/models"
)

func main() {
	userHandler := handlers.NewUserHandler()
	go userHandler.StartGrpcUserService()
	models.Wg.Add(1)
	go userHandler.StartUsersApi()
	models.Wg.Add(1)
	models.Wg.Wait()
}
