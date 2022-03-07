package models

type Shop struct {
	Id      int32
	Name    string
	OwnerId int32
}

func ShopTableName() string {
	return "shops"
}

type ShopDetails struct {
	Shop *Shop
	User *UserDto
}
