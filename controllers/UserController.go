package controllers

import (
	"UserService/models"
	"UserService/services"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person models.User
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		log.Fatal(err)
	}
	insertedUser := services.InsertUser(person)
	log.Println(insertedUser)
	json.NewEncoder(w).Encode(insertedUser.InsertedID)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person models.User
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		log.Fatal(err)
	}
	userDetails := services.GetUserDetails(person)

	log.Println(userDetails)
	json.NewEncoder(w).Encode(userDetails)
}
