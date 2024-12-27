//go:build wireinject
// +build wireinject

package config

import (
	usecases "bbb-voting-service/internal/application/usecases"
	message_producer "bbb-voting-service/internal/domain/producer"
	repositories "bbb-voting-service/internal/domain/repositories"
	consumer "bbb-voting-service/internal/infrastructure/consumer"
	controllers "bbb-voting-service/internal/infrastructure/controllers"
	producer "bbb-voting-service/internal/infrastructure/producer"
	postgres "bbb-voting-service/internal/infrastructure/repositories/postgres"
	redis_repository "bbb-voting-service/internal/infrastructure/repositories/redis"
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
		redis_repository.NewRedisRepository,
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
		usecases.NewCastVoteUsecase,
		usecases.NewProcessVoteUsecase,
		usecases.NewGetPartialResultsUsecase,
		usecases.NewGetFinalResultsUseCase,
		services.NewCaptchaService,
		wire.Bind(new(repositories.InMemoryRepository), new(*redis_repository.RedisRepository)),
		wire.Bind(new(message_producer.MessageProducer), new(*producer.RabbitMQProducer)),
		wire.Bind(new(repositories.ParticipantRepository), new(*postgres.ParticipantRepository)),
		wire.Bind(new(repositories.VoteRepository), new(*postgres.PostgresVoteRepository)),
		wire.Struct(new(Container), "*"),
	)
	return &Container{}, nil
}
