package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User {
		1: &User{Id:1, FirstName:"Ammar", LastName:"Kurabi", Email:"ammar@kurabi.net"},
		2: &User{Id:2, FirstName:"Ali", LastName:"Omar", Email:"ali@kurabi.net"},
	}
)
func GetUser(userId int64) (*User, error){
	 user := users[userId]
	 if user == nil {
	 	return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
	 }

	 //return &User{Id: user.Id, FirstName:user.FirstName, LastName:user.LastName, Email: user.Email}, nil
	 return user, nil
}