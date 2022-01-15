package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Product struct{}

var products []*proto.Product

func main() {
	listener, err := net.Listen("tcp", ":4041")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterProductServiceServer(srv, &Product{})
	reflection.Register(srv)

	fmt.Println("serving products on port : 4041")

	//generating demo products
	i, j := 0, 0
	for j = 01; j < 5; j++ {
		for i = 01; i < 5; i++ {
			products = append(products, &proto.Product{
				Id:     int32(i),
				Name:   fmt.Sprintf("MyShop %d", j),
				ShopId: int32(j),
				Price:  int32(1560),
			})
		}
	}

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}

}

func (p *Product) GetShopProductsByShopId(ctx context.Context, in *proto.ReqShopProducts) (*proto.ShopProducts, error) {
	shopId := in.GetShopId()
	var prdts []*proto.Product
	for _, pr := range products {
		fmt.Println(shopId, " hey ", pr.ShopId)
		if pr.ShopId == shopId {
			prdts = append(prdts, pr)
		}
	}
	if len(prdts) != 0 {
		return &proto.ShopProducts{
			Products: prdts,
		}, nil
	}
	return nil, errors.New("no products found under this shop id")
}
