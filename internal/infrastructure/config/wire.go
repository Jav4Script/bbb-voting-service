//go:build wireinject
// +build wireinject

package config

import (
	"os"

	usecases "bbb-voting-service/internal/application/usecases"
	messageProducer "bbb-voting-service/internal/domain/producer"
	repositories "bbb-voting-service/internal/domain/repositories"
	consumer "bbb-voting-service/internal/infrastructure/consumer"
	controllers "bbb-voting-service/internal/infrastructure/controllers"
	producer "bbb-voting-service/internal/infrastructure/producer"
	postgres "bbb-voting-service/internal/infrastructure/repositories/postgres"
	redis "bbb-voting-service/internal/infrastructure/repositories/redis"
	services "bbb-voting-service/internal/infrastructure/services"

	"github.com/google/wire"
)

func InitializeContainer() (*Container, error) {
	wire.Build(
		InitDB,
		InitRedis,
		InitRabbitMQ,
		postgres.NewParticipantRepository,
		postgres.NewPostgresVoteRepository,
		redis.NewRedisRepository,
		producer.NewRabbitMQProducer,
		controllers.NewCaptchaController,
		controllers.NewParticipantController,
		controllers.NewVoteController,
		controllers.NewResultController,
		consumer.NewRabbitMQConsumer,
		usecases.NewCreateParticipantUsecase,
		usecases.NewGetParticipantsUsecase,
		usecases.NewGetParticipantUsecase,
		usecases.NewDeleteParticipantUsecase,
		InitCastVoteUsecase,
		usecases.NewProcessVoteUsecase,
		usecases.NewGetPartialResultsUsecase,
		usecases.NewGetFinalResultsUseCase,
		services.NewCaptchaService,
		wire.Bind(new(repositories.InMemoryRepository), new(*redis.RedisRepository)),
		wire.Bind(new(messageProducer.MessageProducer), new(*producer.RabbitMQProducer)),
		wire.Bind(new(repositories.ParticipantRepository), new(*postgres.ParticipantRepository)), // Adicione esta linha
		wire.Bind(new(repositories.VoteRepository), new(*postgres.PostgresVoteRepository)),
		wire.Struct(new(Container), "*"),
	)
	return &Container{}, nil
}

func InitCastVoteUsecase(
	inMemoryRepository repositories.InMemoryRepository,
	messageProducer messageProducer.MessageProducer,
	participantRepository repositories.ParticipantRepository,
) *usecases.CastVoteUsecase {
	voteQueue := os.Getenv("VOTE_QUEUE")

	return usecases.NewCastVoteUsecase(inMemoryRepository, messageProducer, participantRepository, voteQueue)
}
