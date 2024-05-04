package models

import (
	"github.com/draco121/common/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Email     string             `json:"email"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Password  string             `json:"password"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Role      constants.Role     `json:"role"`
}
