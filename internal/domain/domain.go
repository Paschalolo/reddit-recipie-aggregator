package domain

import (
	"time"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"github.com/rs/xid"
)

func AddRecipe(recipe *pkg.Recipe) (*pkg.Recipe, error) {
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	return recipe, nil
}
