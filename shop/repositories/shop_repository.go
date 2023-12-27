package repositories

import (
	"github.com/naim6246/grpc-GO/shop/conn"
	"github.com/naim6246/grpc-GO/shop/models"
	"gorm.io/gorm"
)

type ShopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *conn.DB) *ShopRepository {
	return &ShopRepository{
		db: db.Table(models.ShopTableName()),
	}
}

func (repo *ShopRepository) Create(shop *models.Shop) (*models.Shop, error) {
	if err := repo.db.Create(shop).Error; err != nil {
		return nil, err
	}
	return shop, nil
}

func (repo *ShopRepository) GetAll() ([]*models.Shop, error) {
	var shops []*models.Shop
	if err := repo.db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (repo *ShopRepository) GetById(id int32) (*models.Shop, error) {
	var shop models.Shop
	if err := repo.db.Where("id=?", id).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (repo *ShopRepository) GetByFilter(filter interface{}, args ...interface{}) (*models.Shop, error) {
	var shop models.Shop
	if err := repo.db.Where(filter, args...).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

// func (repo *ShopRepository) GetShopByOwnerId(id int32) (*models.Shop, error) {
// 	var shop models.Shop
// 	if err := repo.db.Where("owner_id=?", id).First(&shop).Error; err != nil {
// 		return nil, err
// 	}
// 	return &shop, nil
// }
