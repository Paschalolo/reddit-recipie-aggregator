package memory

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

type Repository struct {
	Recipe []pkg.Recipe
}

var _ repository.Repository = (*Repository)(nil)

func NewRepository() *Repository {
	return &Repository{
		Recipe: []pkg.Recipe{},
	}
}

func (r *Repository) AddRecipe(_ context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	r.Recipe = append(r.Recipe, *recipe)
	return recipe, nil
}

func (r *Repository) GetRecipe(_ context.Context) (*[]pkg.Recipe, error) {
	return &r.Recipe, nil
}
