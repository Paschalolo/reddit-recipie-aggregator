package repository

import (
	"context"
	"errors"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

var (
	ErrNotFound        = errors.New("id not found in repo ")
	ErrSearchNotFound  = errors.New("no tags found of that category ")
	ErrNothingToUpdate = errors.New("there is nothing to update ")
	ErrCache           = errors.New("cannot cache in cache ")
	ErrNotInCache      = errors.New("cache is empty  ")
	ErrAuthUser        = errors.New("invalid username or password ")
)

// This is The Repository interface implementation
type Repository interface {
	// AddRecipe adds a pecipe to the data
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	GetRecipe(ctx context.Context) (*[]pkg.Recipe, error)
	GetOneRecipe(ctx context.Context, id string) (*pkg.Recipe, error)
	BulkAddRecipe(*[]pkg.Recipe) error
	UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error)
	DeleteRecipe(ctx context.Context, id string) bool
	SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error)
}

// This is the Cache Repository impplemntation
// Cache is currently set to have a lifetime of 20 minutes
// except the Delete function is invoked and cache is cleared
type CacheRepo interface {

	// Get returns a string of cached data and error
	// if Cache is not available Get will automatically make a call to the database
	Get(ctx context.Context, key string) (string, error)

	// Delete removes thecache located in the key position
	Delete(ctx context.Context, key string)
}

type AuthRepo interface {
	FindUser(ctx context.Context, username string, hashPassword string) error
	AddBulkAuthUser(ctx context.Context, users *[]pkg.AuthUser) error
}
