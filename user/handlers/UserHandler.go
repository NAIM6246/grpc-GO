package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/user/models"
	"github.com/naim6246/grpc-GO/user/services"
)

type UserHandler struct {
	// shopClinet    *proto.ShopServiceClient
	// productClinet *proto.ProductServiceClient
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	// connToshop, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }
	// shopClinet := proto.NewShopServiceClient(connToshop)

	// connToProduct, err := grpc.Dial("localhost:4041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	panic(err)
	// }
	// productClinet := proto.NewProductServiceClient(connToProduct)

	return &UserHandler{
		// shopClinet:    &shopClinet,
		// productClinet: &productClinet,
		userService: userService,
	}
}

func (u *UserHandler) Handler() {
	//creating a http router
	router := chi.NewRouter()

	router.Post("/", u.createUser)
	router.Get("/user/{userId}/shop", u.getShopByUserId)
	router.Get("/users",u.getAllUser)

	//getAllShop
	// router.Get("/shops", u.getAllShop)

	//getShopById
	// router.Get("/shop/{shopID}", u.getShopById)

	//getShopProductsByShopId
	// router.Get("/shop/{shopID}/products", u.getShopProductsByShopId)

	fmt.Println("Running client on port : 8081")
	http.ListenAndServe(":8081", router)
	models.Wg.Done()
}

func (u *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	createdUser, err := u.userService.CreateUser(&user)
	if err != nil {		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (u *UserHandler) getShopByUserId(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "userId")
	req := proto.ShopByOwnerId{
		OwnerId: int32(id),
	}
	res, err := u.userService.GetShopByOwnerId(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (u *UserHandler) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// func (u *UserHandler) GetUser(ctx context.Context, in *proto.ReqUser) (*proto.ResUser, error) {
// 	userId := in.GetId()
// 	for _, u := range models.Users {
// 		if userId == u.Id {
// 			return u, nil
// 		}
// 	}
// 	return nil, errors.New("user not found")
// }

// func (u *UserHandler) StartGrpcUserService() {
// 	listener, err := net.Listen("tcp", ":4042")
// 	if err != nil {
// 		panic(err)
// 	}
// 	srv := grpc.NewServer()
// 	proto.RegisterUserServiceServer(srv, u)
// 	reflection.Register(srv)
// 	fmt.Println("serving grpc user server on port : 4042")
// 	if err = srv.Serve(listener); err != nil {
// 		panic(err)
// 	}
// 	models.Wg.Done()
// }


// func (u *UserHandler) getAllShop(w http.ResponseWriter, r *http.Request) {
// 	req := proto.ReqAllShop{}
// 	res, err := (*u.shopClinet).GetAllShop(r.Context(), &req)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(res.Shop)
// }

// func (u *UserHandler) getShopById(w http.ResponseWriter, r *http.Request) {
// 	shopId := param.Int(r, "shopID")
// 	req := proto.ShopByID{
// 		ShopId: int32(shopId),
// 	}
// 	res, err := (*u.shopClinet).GetShopByID(r.Context(), &req)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(res)
// }

// func (u *UserHandler) getShopProductsByShopId(w http.ResponseWriter, r *http.Request) {
// 	shopID := param.Int(r, "shopID")
// 	req := proto.ReqShopProducts{
// 		ShopId: int32(shopID),
// 	}
// 	res, err := (*u.productClinet).GetShopProductsByShopId(r.Context(), &req)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(res.Products)
// }
