package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var users []*proto.ResUser
var wg sync.WaitGroup

type User struct{}

func main() {

	//generating demo users
	for i := 01; i < 10; i++ {
		users = append(users, &proto.ResUser{
			Id:   int32(i),
			Name: fmt.Sprintf("MyShop %d", 10-i),
		})
	}
	go startGrpcUserService()
	wg.Add(1)
	go startUsersApi()
	wg.Add(1)
	wg.Wait()
}

func (u *User) GetUser(ctx context.Context, in *proto.ReqUser) (*proto.ResUser, error) {
	userId := in.GetId()
	for _, u := range users {
		if userId == u.Id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func startGrpcUserService() {
	listener, err := net.Listen("tcp", ":4042")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, &User{})
	reflection.Register(srv)

	fmt.Println("serving grpc user service on port : 4042")
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
	wg.Done()
}

func startUsersApi() {

	connToshop, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	shopClinet := proto.NewShopServiceClient(connToshop)

	connToProduct, err := grpc.Dial("localhost:4041", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	productClinet := proto.NewProductServiceClient(connToProduct)

	router := chi.NewRouter()

	//getAllShop
	router.Get("/shops", func(rw http.ResponseWriter, r *http.Request) {
		req := proto.ReqAllShop{}
		res, err := shopClinet.GetAllShop(r.Context(), &req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res.Shop)
	})

	//getShopById
	router.Get("/shop/{shopID}", func(rw http.ResponseWriter, r *http.Request) {
		shopId := param.Int(r, "shopID")
		req := proto.ShopByID{
			ShopId: int32(shopId),
		}
		res, err := shopClinet.GetShopByID(r.Context(), &req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res)
	})

	//getShopProductsByShopId
	router.Get("/shop/{shopID}/products", func(rw http.ResponseWriter, r *http.Request) {
		shopID := param.Int(r, "shopID")
		req := proto.ReqShopProducts{
			ShopId: int32(shopID),
		}
		res, err := productClinet.GetShopProductsByShopId(r.Context(), &req)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res.Products)
	})

	fmt.Println("Running client on port : 8081")
	http.ListenAndServe(":8081", router)
	wg.Done()
}
