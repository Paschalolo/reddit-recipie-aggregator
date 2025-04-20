package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
)

type Application struct {
	db    repository.Repository
	queue repository.Queue
}
type App interface {
	AppConsumer
	AppProducer
}
type AppProducer interface {
	PushToQueue(request *pkg.Request) error
	FindRecipes(ctx context.Context) (*[]pkg.Entry, error)
}
type AppConsumer interface {
	ConsumeQueue(ctx context.Context)
	InsertOne(ctx context.Context, url string) error
}

var _ App = (*Application)(nil)

func NewApp(db repository.Repository, Queue repository.Queue) *Application {
	return &Application{db: db, queue: Queue}
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

func (a *Application) FindRecipes(ctx context.Context) (*[]pkg.Entry, error) {
	return a.db.FindAll(ctx)
}
func (a *Application) PushToQueue(request *pkg.Request) error {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	if err := a.queue.Publish(data); err != nil {
		return err
	}
	return nil
}

func (a *Application) ConsumeQueue(ctx context.Context) {
	forever := make(chan bool)
	msgs, err := a.queue.Subscribe()
	if err != nil {
		log.Fatal("error in the quue ")
	}
	go func(a App) {
		for msg := range msgs {
			log.Printf("Recieved a message :%s ", msg.Body)
			var request pkg.Request
			if err := json.Unmarshal(msg.Body, &request); err != nil {
				log.Fatal(err)
			}
			a.InsertOne(ctx, request.URL)
		}
	}(a)
	log.Print("Waiting for message press ctrl + C to exit ")
	<-forever

}
