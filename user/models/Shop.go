package models

type Shop struct {
	Id      int32
	Name    string
	OwnerId int32
}

func ShopTable() string {
	return "shops"
}
