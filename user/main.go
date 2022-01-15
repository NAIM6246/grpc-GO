package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/proto"
	"google.golang.org/grpc"
)

func main() {
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
		json.NewEncoder(rw).Encode(res)
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
		fmt.Println(req.ShopId)
		res, err := productClinet.GetShopProductsByShopId(r.Context(), &req)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(res)
	})

	fmt.Println("Running client on port : 8081")
	http.ListenAndServe(":8081", router)
}
