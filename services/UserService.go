package services

import (
	"UserService/models"
	"UserService/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var usercollection = repository.DBConnection().Database("UserService").Collection("User")

func InsertUser(person models.User) *mongo.InsertOneResult {
	insertedUser, err := usercollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func GetUserDetails(person models.User) primitive.M {
	var result primitive.M
	er := usercollection.FindOne(context.TODO(), bson.D{{"id", person.Id}}).Decode(&result)

	if er != nil {
		log.Fatal(er)
	}
	return result
}
