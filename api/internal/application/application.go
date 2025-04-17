package application

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/domain"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

type App interface {
	AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error)
	ListRecipe(ctx context.Context) (*[]pkg.Recipe, error)
	ListOneRecipe(ctx context.Context, id string) (*pkg.Recipe, error)
	UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error)
	DeleteRecipe(ctx context.Context, id string) bool
	SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error)
}
type Application struct {
	repo  repository.Repository
	cache repository.CacheRepo
}

var _ App = (*Application)(nil)

func New(Repo repository.Repository, cache repository.CacheRepo) *Application {
	return &Application{repo: Repo, cache: cache}
}
func (r *Application) AddRecipe(ctx context.Context, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context, cache repository.CacheRepo) {
		wg.Done()
		cache.Delete(ctx, "recipes")
	}(ctx, r.cache)
	rec, err := domain.AddRecipe(recipe)
	if err != nil {
		return nil, err
	}
	if _, err := r.repo.AddRecipe(ctx, rec); err != nil {
		return nil, err
	}
	wg.Wait()
	return rec, nil
}
func (r *Application) ListRecipe(ctx context.Context) (*[]pkg.Recipe, error) {
	data, err := r.cache.Get(ctx, "recipes")
	if err != nil {
		if errors.Is(err, repository.ErrNotInCache) {
			log.Println("not in cache ")
		} else {
			log.Println("Redis internal error")
		}
	} else {
		var recipes []pkg.Recipe
		if err = json.Unmarshal([]byte(data), &recipes); err != nil {
			log.Println("error marshalling cache")
		} else {
			log.Println("Reading from cache")
			return &recipes, nil
		}
	}

	list, err := r.repo.GetRecipe(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *Application) ListOneRecipe(ctx context.Context, id string) (*pkg.Recipe, error) {
	list, err := r.repo.GetOneRecipe(ctx, id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *Application) UpdateRecipe(ctx context.Context, id string, recipe *pkg.Recipe) (*pkg.Recipe, error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context, cache repository.CacheRepo) {
		wg.Done()
		cache.Delete(ctx, "recipes")
	}(ctx, r.cache)
	rec, err := domain.UpdateRecipe(recipe)
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return r.repo.UpdateRecipe(ctx, id, rec)

}

func (r *Application) DeleteRecipe(ctx context.Context, id string) bool {
	return r.repo.DeleteRecipe(ctx, id)
}

func (r *Application) SearchRecipe(ctx context.Context, tag string) (*[]pkg.Recipe, error) {
	return r.repo.SearchRecipe(ctx, tag)
}
