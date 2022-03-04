package main

import (
	"UserService/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	route := mux.NewRouter()

	s := route.PathPrefix("/api").Subrouter()
	s.HandleFunc("/createUser", controllers.CreateUser).Methods("POST")
	s.HandleFunc("/getUserById", controllers.GetUserById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", s))
}
