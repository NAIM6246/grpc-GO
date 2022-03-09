package mapper

import (
	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/shop/models"
)

func MapGrpcModelToProductModel(product *proto.Product) *models.Product {
	return &models.Product{
		Id:     product.GetId(),
		Name:   product.GetName(),
		Price:  product.GetPrice(),
		ShopId: product.GetShopId(),
	}
}
