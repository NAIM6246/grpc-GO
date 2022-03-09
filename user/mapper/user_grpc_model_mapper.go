package mapper

import (
	"github.com/naim6246/grpc-GO/proto"
	"github.com/naim6246/grpc-GO/user/models"
)

func MapUserToGrpcModel(user *models.User) *proto.ResUser {
	return &proto.ResUser{
		Id:   user.Id,
		Name: user.Name,
	}
}
