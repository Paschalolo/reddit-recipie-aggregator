package mongo

import (
	"context"
	"log"
	"os"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository struct {
	collection *mongo.Collection
}

var _ repository.Repository = (*Repository)(nil)

func NewMongoDB() *Repository {
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MONGODB ")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	return &Repository{collection: collection}
}

func (r *Repository) AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	_, err := r.collection.InsertOne(ctx, recipe)
	if err != nil {
		return nil, err
	}
	return recipe, nil

}

func (r *Repository) GetRecipe(ctx context.Context) (*[]pkg.Recipe, error) {
	filter := bson.M{}
	var result = []pkg.Recipe{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (r *Repository) GetOneRecipe(ctx context.Context, id string) (*pkg.Recipe, error) {
	var result pkg.Recipe
	filter := bson.M{
		"_id": id,
	}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, repository.ErrNotFound
		// return nil, fmt.Errorf("hey its me %v", err)
	}
	return &result, nil
}

func (r *Repository) BulkAddRecipe(Recipes *[]pkg.Recipe) error {
	_, err := r.collection.InsertMany(context.Background(), *Recipes)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{}
	set := bson.M{}

	if recipe.Name != "" {
		set["name"] = recipe.Name
	}
	if recipe.Ingredients != nil {
		set["ingredients"] = recipe.Ingredients
	}
	if recipe.Instructions != nil {
		set["instructions"] = recipe.Instructions
	}
	if recipe.Tags != nil {
		set["tags"] = recipe.Tags
	}

	// Only add the $set operation if there are fields to update
	if len(set) > 0 {
		update["$set"] = set
	} else {
		return nil, repository.ErrNothingToUpdate
	}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return r.GetOneRecipe(ctx, id)

}

func (r *Repository) DeleteRecipe(ctx context.Context, id string) bool {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err == nil
}

func (r *Repository) SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error) {
	filter := bson.M{
		"tags": bson.M{
			"$in": []string{tag},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, repository.ErrSearchNotFound
	}
	defer cursor.Close(ctx)
	// Unpacks the cursor into a slice
	var results []pkg.Recipe
	if err = cursor.All(ctx, &results); err != nil {
		return nil, repository.ErrSearchNotFound
	}
	if len(results) < 1 {
		return nil, repository.ErrSearchNotFound
	}
	return &results, nil
}
