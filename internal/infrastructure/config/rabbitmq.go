package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func InitRabbitMQ() *amqp.Channel {
	rabbitMQUser := getEnv("RABBITMQ_USER")
	rabbitMQPassword := getEnv("RABBITMQ_PASSWORD")
	rabbitMQHost := getEnv("RABBITMQ_HOST")
	rabbitMQPort := getEnv("RABBITMQ_PORT")
	rabbitMQVHost := getEnv("RABBITMQ_VHOST")

	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", rabbitMQUser, rabbitMQPassword, rabbitMQHost, rabbitMQPort, rabbitMQVHost)

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	_, err = channel.QueueDeclare(
		"vote_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	return channel
}
