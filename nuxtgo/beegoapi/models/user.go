package models

import (
	"errors"
)

var (
	UserList []User
)

func init() {
	u1 := User{"1", "Saleforce"}
	u2 := User{"2", "TypeScript"}
	u3 := User{"3", "Docker"}
	UserList = append(UserList, u1, u2, u3)
}

type User struct {
	Id       string
	Username string
}

func GetUser(uid string) (u User, err error) {
	for _, user := range UserList {
		if user.Id == uid {
			u = user
			return u, nil
		}
	}

	return u, errors.New("User not exists")
}

func GetAllUsers() []User {
	return UserList
}
