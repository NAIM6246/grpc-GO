package models

import (
	"fmt"
	"sync"

	"github.com/naim6246/grpc-GO/proto"
)

var Users []*proto.ResUser
var Wg sync.WaitGroup

func init() {
	//generating demo users
	for i := 01; i < 10; i++ {
		Users = append(Users, &proto.ResUser{
			Id:   int32(i),
			Name: fmt.Sprintf("MyShop %d", 10-i),
		})
	}
}
