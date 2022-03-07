package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/shop/models"
	"github.com/naim6246/grpc-GO/shop/services"
)

var Wg sync.WaitGroup

type ShopHandler struct {
	shopService *services.ShopService
}

func NewShopHandler(shopService *services.ShopService) *ShopHandler {
	return &ShopHandler{
		shopService: shopService,
	}
}

func (h *ShopHandler) Handler() {
	router := chi.NewRouter()
	router.Get("/shop", h.getAllShop)
	router.Post("/shop", h.createShop)
	router.Get("/shop/{shopId}", h.getShopById)
	router.Get("/shop/{shopId}/details", h.getShopDetails)
	router.Get("/shop/{shopId}/owner", h.getShopbyOwner)

	fmt.Println("serving api server on port: 8083")
	http.ListenAndServe(":8083", router)
	Wg.Done()
}

func (h *ShopHandler) getShopbyOwner(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "shopId")
	shop, err := h.shopService.GetShopByOwnerID(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

func (h *ShopHandler) createShop(w http.ResponseWriter, r *http.Request) {
	var shop models.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	createdShop, err := h.shopService.Create(&shop)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdShop)
}

func (h *ShopHandler) getShopById(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "shopId")
	shop, err := h.shopService.GetShopByID(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

func (h *ShopHandler) getAllShop(w http.ResponseWriter, r *http.Request) {
	shops, err := h.shopService.GetAllShops()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shops)
}

func (h *ShopHandler) getShopDetails(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "shopId")
	shop, err := h.shopService.GetShopDetails(int32(id), r.Context())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}
