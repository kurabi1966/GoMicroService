package services

import (
	domain "github.com/kurabi1966/GoMicroService/mvc/domain"
	"github.com/kurabi1966/GoMicroService/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}