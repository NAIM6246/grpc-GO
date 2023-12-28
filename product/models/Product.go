package models

import "sync"

type Product struct {
	Id     int32
	Name   string
	Price  int64
	ShopId int32
}

func ProductTabelName() string {
	return "products"
}

var Wg sync.WaitGroup
