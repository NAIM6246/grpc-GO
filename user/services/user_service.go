package services

import (
	"errors"

	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/user/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	shopClinet    *proto.ShopServiceClient
	productClinet *proto.ProductServiceClient
}

func NewUserService() *UserService {
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
		shopClinet:    &shopClinet,
		productClinet: &productClinet,
	}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	models.Users = append(models.Users, &proto.ResUser{
		Id:   user.Id,
		Name: user.Name,
	})
	return user, nil
}

func (u *UserService) GetUserById(userId int32) (*proto.ResUser, error) {
	// userId := in.GetId()
	for _, u := range models.Users {
		if userId == u.Id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *UserService) GetAllUser() ([]*proto.ResUser, error) {
	return models.Users, nil
}
