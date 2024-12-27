package consumer

import (
	"encoding/json"
	"log"

	usecase "bbb-voting-service/internal/application/usecases"
	entities "bbb-voting-service/internal/domain/entities"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	Channel            *amqp.Channel
	ProcessVoteUsecase *usecase.ProcessVoteUsecase
}

func NewRabbitMQConsumer(channel *amqp.Channel, processVoteUsecase *usecase.ProcessVoteUsecase) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		Channel:            channel,
		ProcessVoteUsecase: processVoteUsecase,
	}
}

func (consumer *RabbitMQConsumer) ConsumeVotes() {
	msgs, err := consumer.Channel.Consume(
		"vote_queue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var vote entities.Vote
			err := json.Unmarshal(d.Body, &vote)
			if err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}

			err = consumer.ProcessVoteUsecase.Execute(vote)
			if err != nil {
				log.Printf("Error processing vote: %v", err)
				continue
			}

			log.Printf("Vote processed: %v", vote)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
