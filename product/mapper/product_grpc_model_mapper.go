package mapper

import (
	"github.com/naim6246/grpc-GO/product/models"
	"github.com/naim6246/grpc-GO/proto"
)

func MapProductToGrpcModel(product *models.Product) *proto.Product {
	return &proto.Product{
		Id:     product.Id,
		Name:   product.Name,
		Price:  product.Price,
		ShopId: product.ShopId,
	}
}
