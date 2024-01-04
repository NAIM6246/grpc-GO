package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

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
	// to check for prometheus alertmanager alert
	router.Post("/alert", h.updateConfig)
	var port string = "8082"
	if val, exists := os.LookupEnv("PORDUCT_SERVICE_PORT"); exists {
		port = val
	}

	fmt.Println("Product Api server is running on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	models.Wg.Done()
}

func (u *ProductHandler) updateConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println("alert received")
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("error while getting cluster config, error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %s\n", err.Error())
		os.Exit(1)
	}

	namespace := "default"

	cm, err := clientset.CoreV1().ConfigMaps(namespace).Get(r.Context(), "product-cm", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting configmaps: %s\n", err.Error())
		return
	}
	fmt.Println("old configMap: ", cm.Data)
	cm.Data["product-price"] = "888"

	ucm, err := clientset.CoreV1().ConfigMaps(namespace).Update(r.Context(), cm, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println("error while updating configMap", "error: ", err)
		return
	}
	fmt.Println("updated configMap: ", ucm.Name, "data: ", ucm.Data)
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	path := "/config/product-price"
	if price, err := param.GetPriceFromConfig(path); err == nil {
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
