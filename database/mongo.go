package database

import (
	"context"
	"github.com/draco121/common/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(mongoUri string) *mongo.Client {
	utils.Logger.Info("initializing mongo db connection")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		utils.Logger.Fatal("failed to connect to mongo db")
		log.Fatal(err)
	}
	utils.Logger.Info("connected to mongo db")
	return client
}
