package domain

import (
	"fmt"
	"github.com/kurabi1966/GoMicroService/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		1: &User{1, "Ammar", "Kurabi", "ammar@kurabi.net"},
		2: &User{2, "Ali", "Omar", "ali@kurabi.net"},
	}
)
func GetUser(userId int64) (*User, *utils.ApplicationError){
	 user := users[userId]
	 if user == nil {
	 	return nil, &utils.ApplicationError{
	 		Message:fmt.Sprintf("user %v was not found", userId),
	 		StatusCode:http.StatusNotFound,
	 		Code:  "not_found",
	 	}
	 }
	 return user, nil
}