package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Project struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"Description"`
	CreatedAt   time.Time          `json:"createdAt"`
	Owner       primitive.ObjectID `json:"owner"`
}
