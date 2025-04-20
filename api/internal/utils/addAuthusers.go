package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
)

func AddAuthUser(repo repository.AuthRepo) {
	file, err := os.ReadFile("authuser.json")
	if err != nil {
		log.Fatal("could not read file", err.Error())
	}
	var Users []pkg.AuthUser
	if err := json.Unmarshal(file, &Users); err != nil {
		log.Fatalln("could not read file ", err.Error())
	}
	hash := sha256.New()
	for i := range Users {
		hashedSting := hash.Sum([]byte(Users[i].Password))
		Users[i].Password = hex.EncodeToString(hashedSting)
	}
	err = repo.AddBulkAuthUser(context.Background(), &Users)
	if err != nil {
		log.Fatal("Issue with putting bulk in a database")
	}

}
