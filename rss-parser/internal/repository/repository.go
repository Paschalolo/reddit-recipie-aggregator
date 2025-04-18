package repository

import (
	"context"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
	ampq "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	AddOne(ctx context.Context, entry *pkg.Entry) error
}

type Queue interface {
	Publish(data []byte) error
	Subscribe() (<-chan ampq.Delivery, error)
}
