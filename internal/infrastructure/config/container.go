package config

import (
	usecases "bbb-voting-service/internal/application/usecases"
	consumer "bbb-voting-service/internal/infrastructure/consumer"
	controllers "bbb-voting-service/internal/infrastructure/controllers"
	producer "bbb-voting-service/internal/infrastructure/producer"
	postgres "bbb-voting-service/internal/infrastructure/repositories/postgres"
	redis_repository "bbb-voting-service/internal/infrastructure/repositories/redis"
	services "bbb-voting-service/internal/infrastructure/services"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type Container struct {
	DB                       *gorm.DB
	RedisClient              *redis.Client
	RabbitMQChannel          *amqp.Channel
	VoteRepository           *postgres.PostgresVoteRepository
	ParticipantRepository    *postgres.ParticipantRepository
	RedisRepository          *redis_repository.RedisRepository
	RabbitMQRepository       *producer.RabbitMQProducer
	ProcessVoteUsecase       *usecases.ProcessVoteUsecase
	CastVoteUsecase          *usecases.CastVoteUsecase
	CreateParticipantUsecase *usecases.CreateParticipantUsecase
	GetParticipantsUsecase   *usecases.GetParticipantsUsecase
	DeleteParticipantUsecase *usecases.DeleteParticipantUsecase
	GetParticipantUsecase    *usecases.GetParticipantUsecase
	GetPartialResultsUsecase *usecases.GetPartialResultsUsecase
	GetFinalResultsUseCase   *usecases.GetFinalResultsUsecase
	CaptchaService           *services.CaptchaService
	CaptchaController        *controllers.CaptchaController
	ParticipantController    *controllers.ParticipantController
	VoteController           *controllers.VoteController
	ResultController         *controllers.ResultController
	RabbitMQConsumer         *consumer.RabbitMQConsumer
}
