package models

import (
	"fmt"
	"sync"

	"github.com/naim6246/grpc-GO/proto"
)

type User struct{
	Id int32
	Name string
	Password string
}

func UsersTable() string {
	return "users"
}

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


