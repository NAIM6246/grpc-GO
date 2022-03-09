package models

import (
	"sync"
)

type User struct {
	Id       int32
	Name     string
	Password string
}

func UsersTable() string {
	return "users"
}

var Wg sync.WaitGroup
