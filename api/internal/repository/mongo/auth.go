package mongo

import (
	"context"
	"log"
	"os"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type AuthRepository struct {
	collection *mongo.Collection
}

var _ repository.AuthRepo = (*AuthRepository)(nil)

func NewAuthMongoDB(client *mongo.Client) *AuthRepository {
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MONGODB ")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	return &AuthRepository{collection: collection}
}

func (r *AuthRepository) FindUser(ctx context.Context, username string, hashPassword string) error {
	filter := bson.M{
		"password": hashPassword,
		"username": username,
	}
	log.Println(username, hashPassword)
	cursor := r.collection.FindOne(ctx, filter)
	if cursor.Err() != nil {
		return repository.ErrAuthUser
		// return cursor.Err()
	}
	return nil
}

func (r *AuthRepository) checkExist(ctx context.Context, username string, hashPassword string) error {
	filter := bson.M{
		"username": username,
	}
	log.Println(username, hashPassword)
	cursor := r.collection.FindOne(ctx, filter)
	if cursor.Err() != nil {
		return repository.ErrAuthUser
		// return cursor.Err()
	}
	return nil
}

func (r *AuthRepository) AddUser(ctx context.Context, user *pkg.AuthUser) error {
	if err := r.checkExist(ctx, user.Username, user.Password); err == nil {
		return repository.ErrUserExist
	}
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) AddBulkAuthUser(ctx context.Context, users *[]pkg.AuthUser) error {
	log.Println("Adding to mongo db  ")
	_, err := r.collection.InsertMany(context.Background(), *users)
	if err != nil {
		return err
	}
	return nil
}
