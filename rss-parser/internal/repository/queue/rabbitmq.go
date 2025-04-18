package queue

import (
	"log"
	"os"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	conn    *ampq.Connection
	channel *ampq.Channel
}

func NewQueue() *Queue {
	ampqConnection, err := ampq.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}
	channelAmpq, err := ampqConnection.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &Queue{conn: ampqConnection, channel: channelAmpq}
}

func (q *Queue) CloseConnection() {
	q.conn.Close()
}

func (q *Queue) Publish(data []byte) error {
	if err := q.channel.Publish("", os.Getenv("RABBITMQ_QUEUE"), false, false, ampq.Publishing{
		ContentType: "application/json",
		Body:        data,
	}); err != nil {
		return err
	}
	return nil
}

func (q *Queue) Subscribe() (<-chan ampq.Delivery, error) {
	msgs, err := q.channel.Consume(os.Getenv("RABBITMQ_QUEUE"), "", false, false, false, false, nil)
	return msgs, err
}
