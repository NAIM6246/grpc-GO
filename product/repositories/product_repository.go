package repositories

import (
	"github.com/naim6246/grpc-GO/product/conn"
	"github.com/naim6246/grpc-GO/product/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *conn.DB) *ProductRepository {
	return &ProductRepository{
		db: db.DB,
	}
}

func (repo *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) GetById(id int32) (*models.Product, error) {
	var product models.Product
	if err := repo.db.Where("id=?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) GetAllByFilter(filter interface{}, args ...interface{}) ([]*models.Product, error) {
	var products []*models.Product
	if err := repo.db.Where(filter, args...).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
