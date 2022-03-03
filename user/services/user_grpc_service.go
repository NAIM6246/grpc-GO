package services

import (
	"context"
	"fmt"
	"net"

	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/user/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (u *UserService) StartGrpcUserService() {
	listener, err := net.Listen("tcp", ":4042")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, u)
	reflection.Register(srv)

	fmt.Println("serving grpc user server on port : 4042")
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
	models.Wg.Done()
}

func (u *UserService) GetUser(ctx context.Context, in *proto.ReqUser) (*proto.ResUser, error) {
	return u.GetUserById(in.GetId())
}


func (u *UserService) GetShopByOwnerId(ctx context.Context, in *proto.ShopByOwnerId) (*proto.Shop, error) {
	return (*u.shopClinet).GetShopByOwnerId(ctx, in)
}
