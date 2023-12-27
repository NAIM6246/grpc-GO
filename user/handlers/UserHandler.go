package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/user/models"
	"github.com/naim6246/grpc-GO/user/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) Handler() {
	//creating a http router
	router := chi.NewRouter()

	router.Post("/", u.createUser)
	router.Get("/{userId}", u.getUserById)
	router.Get("/user/{userId}/shop", u.getShopByUserId)
	router.Get("/users", u.getAllUser)

	var port string = "8081"
	if val, exists := os.LookupEnv("PORDUCT_SERVICE_PORT"); exists {
		port = val
	}

	fmt.Println("Running client on port : ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (u *UserHandler) getShopByUserId(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "userId")
	res, err := u.userService.GetUserShopDetails(id)
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

func (u *UserHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "userId")
	user, err := u.userService.GetUserById(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
