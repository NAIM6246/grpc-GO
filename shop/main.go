package main

import (
	"github.com/naim6246/grpc-GO/shop/config"
	"github.com/naim6246/grpc-GO/shop/conn"
	"github.com/naim6246/grpc-GO/shop/handlers"
	"github.com/naim6246/grpc-GO/shop/repositories"
	"github.com/naim6246/grpc-GO/shop/services"
)

var dbInstance *conn.DB

func init() {
	dbConfig := config.NewDBConfig()
	dbInstance = conn.ConnectDB(dbConfig)
}

func main() {
	shopRepository := repositories.NewShopRepository(dbInstance)
	shopService := services.NewShopService(shopRepository)
	shopHandler := handlers.NewShopHandler(shopService)

	go shopService.StartGrpcShopServer()
	handlers.Wg.Add(1)

	go shopHandler.Handler()
	handlers.Wg.Add(1)
	handlers.Wg.Wait()
}
