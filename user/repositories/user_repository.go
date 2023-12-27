package repositories

import (
	"github.com/naim6246/grpc-GO/user/conn"
	"github.com/naim6246/grpc-GO/user/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *conn.DB) *UserRepository {
	return &UserRepository{
		db: db.DB,
	}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserById(id int32) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
