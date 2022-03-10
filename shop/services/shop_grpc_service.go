package services

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *ShopService) StartGrpcShopServer() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterShopServiceServer(srv, s)
	reflection.Register(srv)
	fmt.Println("serving on port : 4040")
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *ShopService) GetUser(ctx context.Context, in *proto.ReqUser) (*proto.ResUser, error) {
	return (*s.userClient).GetUser(ctx, in)
}

func (s *ShopService) GetShopByOwnerId(ctx context.Context, in *proto.ShopByOwnerId) (*proto.Shop, error) {
	id := in.GetOwnerId()
	shop, err := s.GetShopByOwnerID(id)
	if err != nil {
		return nil, err
	}
	return &proto.Shop{
		Id:      shop.Id,
		Name:    shop.Name,
		OwnerId: shop.OwnerId,
	}, nil
}

func (s *ShopService) GetShopProductsByShopId(ctx context.Context, in *proto.ReqShopProducts) (*proto.ShopProducts, error) {
	// return (*s.productClient).GetShopProductsByShopId(ctx, in)
	productStream, err := (*s.productClient).GetShopProductsByShopId(ctx, in)
	if err != nil {
		return nil, err
	}
	var shopProducts []*proto.Product
	count := 1
	for {
		product, err := productStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while receiving a product")
			// break
		}
		fmt.Println("Product received ", count)
		count++
		shopProducts = append(shopProducts, product)
	}
	return &proto.ShopProducts{
		Products: shopProducts,
	}, nil
}
