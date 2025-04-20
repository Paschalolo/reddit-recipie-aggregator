package internal

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/mongo"
	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/internal/repository/queue"
)

func Consumer() {
	repo := mongo.NewMongo()
	Queue := queue.NewQueue()
	App := NewApp(repo, Queue)
	App.ConsumeQueue(context.Background())

}
