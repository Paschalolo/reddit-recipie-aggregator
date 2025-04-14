package repository

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

type Repository interface {
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	GetRecipe(ctx context.Context) (*[]pkg.Recipe, error)
}
