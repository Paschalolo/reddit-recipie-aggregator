package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"

	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func AddAuthUser(client *mongo.Client) {
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	file, err := os.ReadFile("authuser.json")
	if err != nil {
		log.Fatal("could not read file", err.Error())
	}
	var Users []pkg.AuthUser
	if err := json.Unmarshal(file, &Users); err != nil {
		log.Fatalln("could not read file ", err.Error())
	}
	log.Println("Adding to mongo db  ")

	hash := sha256.New()
	for i := range Users {
		hashedSting := hash.Sum([]byte(Users[i].Password))
		Users[i].Password = hex.EncodeToString(hashedSting)
	}
	log.Println(Users)
	_, err = collection.InsertMany(context.Background(), Users)
	if err != nil {
		log.Fatalln("error in database ", err)
	}

}
