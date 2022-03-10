package services

import (
	"github.com/naim6246/grpc-GO/product/models"
	"github.com/naim6246/grpc-GO/product/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return p.productRepository.Create(product)
}

func (p *ProductService) GetProductById(id int32) (*models.Product, error) {
	return p.productRepository.GetById(id)
}

func (p *ProductService) GetShopProducts(shopId int32) ([]*models.Product, error) {
	return p.productRepository.GetAllByFilter("shop_id=?", shopId)
}

func (p *ProductService) GetAllProducts() ([]*models.Product, error) {
	return p.productRepository.GetAll()
}
