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
	s.HandleFunc("/updateUser", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/removeUser", controllers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", s))
}
