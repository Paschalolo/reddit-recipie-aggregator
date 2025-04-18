package repository

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
)

type Repository interface {
	AddOne(ctx context.Context, entry *pkg.Entry) error
}
