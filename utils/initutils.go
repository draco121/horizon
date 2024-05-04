package utils

import (
	"context"
	"errors"
	"github.com/draco121/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRootUser(db *mongo.Database) {
	var user models.User
	err := db.Collection("users").FindOne(context.Background(), bson.M{"email": "superuser"}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		user = models.User{
			ID:        primitive.NewObjectID(),
			Email:     "superuser",
			FirstName: "super",
			Role:      "root",
			LastName:  "user",
		}
		password, err := HashPassword("superuser")
		if err != nil {
			panic("unable to initialize superuser")
		} else {
			user.Password = password
			_, err := db.Collection("users").InsertOne(context.Background(), user)
			if err != nil {
				panic("unable to initialize superuser")
			}
		}
	}
}
