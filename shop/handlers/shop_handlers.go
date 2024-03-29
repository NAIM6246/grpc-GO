package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	router.Route("/shops", func(r chi.Router) {
		r.Post("/", h.createShop)
		r.Get("/", h.getAllShop)

		r.Route("/{shopId}", func(r2 chi.Router) {
			r2.Get("/", h.getShopById)
			r2.Get("/details", h.getShopDetails)
			r2.Get("/products", h.getShopProducts)
		})
	})

	var port string = "8083"
	if val, exists := os.LookupEnv("SHOP_SERVICE_PORT"); exists {
		port = val
	}

	fmt.Println("serving api server on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	Wg.Done()
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
		fmt.Println(err)
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

func (h *ShopHandler) getShopProducts(w http.ResponseWriter, r *http.Request) {
	shopId := param.Int(r, "shopId")
	products, err := h.shopService.GetShopProduts(shopId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
