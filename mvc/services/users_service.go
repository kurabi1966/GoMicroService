package services

import domain "github.com/kurabi1966/GoMicroService/mvc/domain"

func GetUser(userId int64) (*domain.User, error){
	return domain.GetUser(userId)
}