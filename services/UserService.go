package services

import (
	"UserService/models"
	"UserService/repository"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var usercollection = repository.DBConnection().Database("UserService").Collection("User")

func UserService(person models.User) *mongo.InsertOneResult {
	insertedUser, err := usercollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
