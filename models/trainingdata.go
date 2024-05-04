package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FAQS struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
type Files struct {
	FileName  string             `json:"fileName"`
	FileId    primitive.ObjectID `json:"fileId"`
	Extension string             `json:"extension"`
}
type TrainingData struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ProjectId   primitive.ObjectID `json:"projectId"`
	BotId       primitive.ObjectID `json:"botId"`
	Files       []Files            `json:"files"`
	Description string             `json:"description"`
	Greeting    string             `json:"greeting"`
	Persona     string             `json:"persona"`
	QA          []FAQS             `json:"qa"`
	Owner       primitive.ObjectID `json:"owner"`
}
