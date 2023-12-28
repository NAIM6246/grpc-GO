package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/naim6246/grpc-GO/param"
	"github.com/naim6246/grpc-GO/product/models"
	"github.com/naim6246/grpc-GO/product/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) Handler() {
	router := chi.NewRouter()

	router.Route("/products", func(router chi.Router) {
		router.Post("/", h.createProduct)
		router.Get("/{id}", h.getProductById)
		router.Get("/", h.getAllProducts)
	})
	var port string = "8082"
	if val, exists := os.LookupEnv("PORDUCT_SERVICE_PORT"); exists {
		port = val
	}

	fmt.Println("Product Api server is running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	models.Wg.Done()
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if price, err := param.LookupEnvInt64("PRODUCT_PRICE"); err == nil {
		product.Price = price
	}

	pr, err := h.productService.CreateProduct(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pr)
}

func (h *ProductHandler) getProductById(w http.ResponseWriter, r *http.Request) {
	id := param.Int(r, "id")
	product, err := h.productService.GetProductById(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
