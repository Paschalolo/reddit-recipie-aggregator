producer  :
	MONGO_URI=mongodb://myuser:mypassword@localhost:27017/  MONGO_DATABASE=demo RABBITMQ_URI="amqp://user:password@localhost:5672" RABBITMQ_QUEUE=rss_feed_queue go run cmd/producer/*.go

consumer  :
	MONGO_URI=mongodb://myuser:mypassword@localhost:27017/  MONGO_DATABASE=demo RABBITMQ_URI="amqp://user:password@localhost:5672" RABBITMQ_QUEUE=rss_feed_queue go run cmd/consumer/*.go