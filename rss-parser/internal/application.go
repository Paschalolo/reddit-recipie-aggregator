package internal

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository"
)

type Application struct {
	db repository.Repository
}

type App interface {
	InsertOne(ctx context.Context, url string) error
}

var _ App = (*Application)(nil)

func NewApp(db repository.Repository) *Application {
	return &Application{db: db}
}
func (a *Application) InsertOne(ctx context.Context, url string) error {
	entries, err := GetFeedEntries(url)
	if err != nil {
		return err
	}
	for _, entry := range *entries {
		if err := a.db.AddOne(ctx, &entry); err != nil {
			return err
		}
	}
	return nil
}
