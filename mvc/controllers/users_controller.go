package controllers

import (
	"encoding/json"
	"github.com/kurabi1966/GoMicroService/mvc/utils"
	"net/http"
	"strconv"
	"github.com/kurabi1966/GoMicroService/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	//lo  g.Printf("about to process user_id %v", userIdParam)
	userId, errBadRequest := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if errBadRequest != nil {
		userErr := &utils.ApplicationError{
			Message: "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		resp.WriteHeader(userErr.StatusCode)
		jsonValue, _ := json.Marshal(userErr)
		resp.Write(jsonValue)
		return
	}
	user, errNotFound := services.GetUser(userId)
	if errNotFound != nil {
		resp.WriteHeader(errNotFound.StatusCode)
		jsonValue, _ := json.Marshal(errNotFound)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	//resp.Header().Set("Content-Type", "text/json")
	resp.Write(jsonValue)
}
