package models

import "sync"

type Product struct {
	Id     int32
	Name   string
	Price  int32
	ShopId int32
}

func ProductTabelName() string {
	return "products"
}

var Wg sync.WaitGroup
