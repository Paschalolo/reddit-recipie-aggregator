package memory

import (
	"context"
	"slices"
	"sync"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

type Repository struct {
	sync.RWMutex
	Recipe []pkg.Recipe
}

var _ repository.Repository = (*Repository)(nil)

func NewRepository() *Repository {
	return &Repository{
		Recipe: []pkg.Recipe{},
	}
}

func (r *Repository) AddRecipe(_ context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	r.Lock()
	defer r.Unlock()
	r.Recipe = append(r.Recipe, *recipe)
	return recipe, nil
}

func (r *Repository) GetRecipe(_ context.Context) (*[]pkg.Recipe, error) {
	r.RLock()
	defer r.RUnlock()
	return &r.Recipe, nil
}
func (r *Repository) GetOneRecipe(ctx context.Context, id string) (*pkg.Recipe, error) {
	r.RLock()
	defer r.RUnlock()
	for _, recipe := range r.Recipe {
		if recipe.ID == id {
			return &recipe, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (r *Repository) BulkAddRecipe(Recipes *[]pkg.Recipe) error {
	r.Lock()
	defer r.Unlock()
	r.Recipe = append(r.Recipe, *Recipes...)
	return nil
}

func (r *Repository) UpdateRecipe(_ context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	r.Lock()
	defer r.Unlock()
	for i, uprecipe := range r.Recipe {
		if uprecipe.ID == id {
			if uprecipe.Name != recipe.Name {
				r.Recipe[i].Name = recipe.Name
			}
			if recipe.Ingredients != nil {
				r.Recipe[i].Ingredients = recipe.Ingredients
			}
			if recipe.Instructions != nil {
				r.Recipe[i].Instructions = recipe.Instructions
			}
			if recipe.Tags != nil {
				r.Recipe[i].Tags = recipe.Tags
			}

			return &r.Recipe[i], nil
		}
	}
	return nil, repository.ErrNotFound
}

func (r *Repository) DeleteRecipe(_ context.Context, id string) bool {
	r.Lock()
	defer r.Unlock()
	for i, recipe := range r.Recipe {
		if recipe.ID == id {
			r.Recipe = slices.Delete(r.Recipe, i, i+1)
			return true
		}
	}
	return false
}

func (r *Repository) SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error) {
	r.RLock()
	defer r.RUnlock()
	var search []pkg.Recipe

	for i, recipe := range r.Recipe {
		for _, t := range recipe.Tags {
			if t == tag {
				search = append(search, r.Recipe[i])
			}
		}
	}
	if len(search) < 1 {
		return nil, repository.ErrSearchNotFound
	}
	return &search, nil

}
