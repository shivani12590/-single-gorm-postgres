package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"postgres/gormPostgres/controller"
)

func Handler() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/{userId}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users", controller.GetAllUser).Methods("GET")
	r.HandleFunc("/users/{userId}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", controller.DeleteAllUser).Methods("DELETE")

	http.ListenAndServe(":2020", r)
}
