package mongo

import (
	"context"
	"log"
	"os"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoRepository struct {
	collection *mongo.Collection
}

var _ repository.Repository = (*MongoRepository)(nil)

func NewMongo() *MongoRepository {
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MONGODB ")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes2")
	return &MongoRepository{collection: collection}
}

func (m *MongoRepository) AddOne(ctx context.Context, entry *pkg.Entry) error {
	if _, err := m.collection.InsertOne(ctx, entry); err != nil {
		return err
	}
	return nil
}
