package queue

import "github.com/streadway/amqp"

type RabbitMQProducer struct {
	Channel *amqp.Channel
}

func NewRabbitMQProducer(channel *amqp.Channel) *RabbitMQProducer {
	return &RabbitMQProducer{Channel: channel}
}

func (producer *RabbitMQProducer) Publish(queueName string, message []byte) error {
	return producer.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}
