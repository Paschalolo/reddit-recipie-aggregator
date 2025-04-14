package application

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/domain"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

type App interface {
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	ListRecipe(ctx context.Context) (*[]pkg.Recipe, error)
}
type Application struct {
	repo repository.Repository
}

var _ App = (*Application)(nil)

func New(Repo repository.Repository) *Application {
	return &Application{repo: Repo}
}
func (r *Application) AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	rec, err := domain.AddRecipe(recipe)
	if err != nil {
		return nil, err
	}
	if _, err := r.repo.AddRecipe(ctx, rec); err != nil {
		return nil, err
	}
	return rec, nil
}
func (r *Application) ListRecipe(ctx context.Context) (*[]pkg.Recipe, error) {
	list, err := r.repo.GetRecipe(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
