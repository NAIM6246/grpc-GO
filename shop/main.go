package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Shop struct{}

var shops []*proto.Shop

type ShopDetails struct {
	Shop  *proto.Shop
	Owner *proto.ResUser
}

func main() {

	//running shops api in another go routine on port 8082
	go func() {
		router := chi.NewRouter()
		router.Get("/shops", func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(shops)
		})

		router.Get("/shop/{shopID}/details", func(rw http.ResponseWriter, r *http.Request) {
			shopId := param.Int(r, "shopID")
			shopDetail := &ShopDetails{
				Shop:  nil,
				Owner: nil,
			}
			for _, s := range shops {
				if shopId == int(s.Id) {
					shopDetail.Shop = s
				}
			}
			if shopDetail.Shop == nil {
				rw.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(rw).Encode("shop not found")
				return
			}

			//creating userClient for grpc
			connToUser, err := grpc.Dial("localhost:4042", grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			userClient := proto.NewUserServiceClient(connToUser)
			req := proto.ReqUser{
				Id: shopDetail.Shop.OwnerId,
			}
			userInfo, err := userClient.GetUser(r.Context(), &req)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(rw).Encode("shop owner not found")
				return
			}
			shopDetail.Owner = userInfo

			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(shopDetail)
		})

		fmt.Println("http serving for shops on port : 8083")
		http.ListenAndServe(":8083", router)
	}()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterShopServiceServer(srv, &Shop{})
	reflection.Register(srv)

	fmt.Println("serving on port : 4040")

	//generating demo shops
	for i := 01; i < 10; i++ {
		shops = append(shops, &proto.Shop{
			Id:      int32(i),
			Name:    fmt.Sprintf("MyShop %d", i),
			OwnerId: int32(10 - i),
		})
	}

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *Shop) GetAllShop(ctx context.Context, in *proto.ReqAllShop) (*proto.AllShop, error) {

	return &proto.AllShop{
		Shop: shops,
	}, nil
}
func (s *Shop) GetShopByID(ctx context.Context, in *proto.ShopByID) (*proto.Shop, error) {
	id := in.GetShopId()
	for _, s := range shops {
		if s.Id == id {
			return s, nil
		}
	}
	return nil, errors.New("no shop found")
}
