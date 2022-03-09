package services

import (
	"context"

	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/user/models"
	"github.com/naim6246/grpc-GO/user/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	shopClinet     *proto.ShopServiceClient
	productClinet  *proto.ProductServiceClient
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	connToshop, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	shopClinet := proto.NewShopServiceClient(connToshop)

	connToProduct, err := grpc.Dial("localhost:4041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	productClinet := proto.NewProductServiceClient(connToProduct)

	return &UserService{
		shopClinet:     &shopClinet,
		productClinet:  &productClinet,
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	return u.userRepository.Create(user)
}

func (u *UserService) GetUserById(userId int32) (*models.User, error) {
	return u.userRepository.GetUserById(userId)
}

func (u *UserService) GetAllUser() ([]*models.User, error) {
	return u.userRepository.GetAll()
}

func (u *UserService) GetUserShopDetails(userId int) (*models.Shop, error) {
	res, err := u.GetShopByOwnerId(context.TODO(), &proto.ShopByOwnerId{
		OwnerId: int32(userId),
	})
	if err != nil {
		return nil, err
	}
	return &models.Shop{
		Id:      res.Id,
		Name:    res.Name,
		OwnerId: res.OwnerId,
	}, nil
}
