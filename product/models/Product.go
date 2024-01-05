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

type AlertReq struct {
	GroupKey          string
	Status            string
	GroupLabels       map[string]interface{}
	CommonLabels      map[string]interface{}
	CommonAnnotations map[string]interface{}
	Alerts            []Alert
}

type Alert struct {
	Labels      map[string]interface{}
	Annotations map[string]interface{}
}
