package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"postgres/gormPostgres/connection"
	"postgres/gormPostgres/model"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	db := connection.Connection()
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(response).Encode(user.ID)
	return

}
func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User
	db := connection.Connection()
	params := mux.Vars(request)
	db.First(&user, params["userId"])
	json.NewEncoder(response).Encode(user)
	return

}
func GetAllUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var Users []model.User

	db := connection.Connection()
	db.Find(&Users)
	json.NewEncoder(response).Encode(Users)
	return

}
func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	db := connection.Connection()
	var user model.User
	db.First(&user, params["userId"])
	json.NewDecoder(request.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(response).Encode(user)
	return

}
func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.Connection()
	params := mux.Vars(request)
	var user model.User
	db.First(&user, params["userId"])
	db.Delete(&user)
	json.NewEncoder(response).Encode("deleted successfully")
	return

}
func DeleteAllUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "appplication/json")
	var Users []model.User

	db := connection.Connection()
	db.Find(&Users)
	db.Delete(Users)
	json.NewEncoder(response).Encode("deleted all successfully")
	return

}
