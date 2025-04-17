package utils

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/Paschalolo/reddit-recipie-aggregator/docs"
	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

func AddBulkRecipe(repo repository.Repository) {
	file, err := os.ReadFile("recipes.json")
	if err != nil {
		log.Fatal("could not read file", err.Error())
	}
	var recipe []pkg.Recipe
	if err := json.Unmarshal(file, &recipe); err != nil {
		log.Fatalln("could not read file ", err.Error())
	}
	if err := repo.BulkAddRecipe(&recipe); err != nil {
		log.Fatalln("could not save file ", err.Error())
	}
}
