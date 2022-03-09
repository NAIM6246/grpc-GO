package services

import (
	"context"
	"fmt"
	"net"

	"github.com/naim6246/grpc-GO/product/mapper"
	"github.com/naim6246/grpc-GO/product/models"
	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (p *ProductService) StartProductGrpcServer() {
	lister, err := net.Listen("tcp", ":4041")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterProductServiceServer(server, p)
	reflection.Register(server)

	fmt.Println("serving product grpc service on port: 4041")
	if err = server.Serve(lister); err != nil {
		panic(err)
	}
	models.Wg.Done()
}

func (p *ProductService) GetShopProductsByShopId(ctx context.Context, in *proto.ReqShopProducts) (*proto.ShopProducts, error) {
	products, err := p.GetShopProducts(in.GetShopId())
	if err != nil {
		return nil, err
	}
	var shopProducts []*proto.Product
	for _, product := range products {
		shopProducts = append(shopProducts, mapper.MapProductToGrpcModel(product))
	}
	return &proto.ShopProducts{
		Products: shopProducts,
	}, nil
}
