package models

type User struct {
	Id       int32
	Name     string
	Password string
}

func UsersTable() string {
	return "users"
}

type UserDto struct {
	Id   int32
	Name string
}
