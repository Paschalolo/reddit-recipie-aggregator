package repository

import (
	"context"
	"errors"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

var (
	ErrNotFound = errors.New("id not found in repo ")
)

type Repository interface {
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	GetRecipe(ctx context.Context) (*[]pkg.Recipe, error)
	BulkAddRecipe(*[]pkg.Recipe) error
	UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error)
}
