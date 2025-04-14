package repository

import (
	"context"
	"errors"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

var (
	ErrNotFound       = errors.New("id not found in repo ")
	ErrSearchNotFound = errors.New("no tags found of that category ")
)

type Repository interface {
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	GetRecipe(ctx context.Context) (*[]pkg.Recipe, error)
	BulkAddRecipe(*[]pkg.Recipe) error
	UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error)
	DeleteRecipe(ctx context.Context, id string) bool
	SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error)
}
