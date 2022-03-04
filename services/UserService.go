package services

import (
	"UserService/models"
	"UserService/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateUserDetails(person models.User) *mongo.SingleResult {
	filter := bson.D{{"id", person.Id}}
	after := options.After

	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"address", person.Address}}}}
	return usercollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)
}

func DeleteUser(person models.User) *mongo.DeleteResult {
	filter := bson.D{{"id", person.Id}}
	opts := options.Delete().SetCollation(&options.Collation{})
	removedUser, err := usercollection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
	}
	return removedUser
}
