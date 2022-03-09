package main

import (
	"github.com/naim6246/grpc-GO/product/config"
	"github.com/naim6246/grpc-GO/product/conn"
	"github.com/naim6246/grpc-GO/product/handlers"
	"github.com/naim6246/grpc-GO/product/models"
	"github.com/naim6246/grpc-GO/product/repositories"
	"github.com/naim6246/grpc-GO/product/services"
)

var dbInstance *conn.DB

func init() {
	config := config.NewDBConfig()
	dbInstance = conn.ConnectDB(config)
	dbInstance.Migration()
}

func main() {
	productRepository := repositories.NewProductRepository(dbInstance)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	go productService.StartProductGrpcServer()
	models.Wg.Add(1)
	go productHandler.Handler()
	models.Wg.Add(1)
	models.Wg.Wait()
}
