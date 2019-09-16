package controllers

import (
	"encoding/json"
	"github.com/kurabi1966/GoMicroService/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	//log.Printf("about to process user_id %v", userIdParam)
	userId , err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("userId must be a number"))
		return
	}
	user, errNotFound := services.GetUser(userId)
	if errNotFound != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(errNotFound.Error()))
		return
	}

	jsonValue, errBadFormat := json.Marshal(user)
	if errBadFormat != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(errBadFormat.Error()))
		return
	}
	resp.Write(jsonValue)
}
