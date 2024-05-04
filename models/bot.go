package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bot struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ProjectId   primitive.ObjectID `json:"projectId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
	Owner       primitive.ObjectID `json:"owner"`
}
