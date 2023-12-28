package services

import (
	"context"

	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/shop/mapper"
	"github.com/naim6246/grpc-GO/shop/models"
	"github.com/naim6246/grpc-GO/shop/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShopService struct {
	userClient     *proto.UserServiceClient
	productClient  *proto.ProductServiceClient
	shopRepository *repositories.ShopRepository
	proto.UnimplementedShopServiceServer
}

func NewShopService(shopRepository *repositories.ShopRepository) *ShopService {
	connToProduct, err := grpc.Dial("localhost:4041", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	productClinet := proto.NewProductServiceClient(connToProduct)

	connToUser, err := grpc.Dial("localhost:4042", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient := proto.NewUserServiceClient(connToUser)
	return &ShopService{
		userClient:     &userClient,
		productClient:  &productClinet,
		shopRepository: shopRepository,
	}
}

func (s *ShopService) Create(shop *models.Shop) (*models.Shop, error) {
	return s.shopRepository.Create(shop)
}
func (s *ShopService) GetShopByID(id int32) (*models.Shop, error) {
	return s.shopRepository.GetById(id)
}
func (s *ShopService) GetAllShops() ([]*models.Shop, error) {
	return s.shopRepository.GetAll()
}
func (s *ShopService) GetShopByOwnerID(id int32) ([]*models.Shop, error) {
	return s.shopRepository.GetByFilter("owner_id=?", id)
}

func (s *ShopService) GetShopDetails(shopId int32, ctx context.Context) (*models.ShopDetails, error) {
	shop, err := s.shopRepository.GetById(shopId)
	if err != nil {
		return nil, err
	}
	user, err := s.GetUser(ctx, &proto.ReqUser{
		Id: shop.OwnerId,
	})
	if err != nil {
		return nil, err
	}
	return &models.ShopDetails{
		Shop: shop,
		User: &models.UserDto{
			Id:   user.GetId(),
			Name: user.GetName(),
		},
	}, nil
}

func (s *ShopService) GetShopProduts(shopId int) ([]*models.Product, error) {
	products, err := s.GetShopProductsByShopId(context.TODO(), &proto.ReqShopProducts{
		ShopId: int32(shopId),
	})
	if err != nil {
		return nil, err
	}
	var shopProducts []*models.Product
	for _, product := range products.Products {
		shopProducts = append(shopProducts, mapper.MapGrpcModelToProductModel(product))
	}
	return shopProducts, nil
}
