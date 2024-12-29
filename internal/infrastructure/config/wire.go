//go:build wireinject
// +build wireinject

package config

import (
	"os"

	"bbb-voting-service/internal/application/usecases/cache"
	"bbb-voting-service/internal/application/usecases/captcha"
	"bbb-voting-service/internal/application/usecases/participants"
	"bbb-voting-service/internal/application/usecases/results"
	"bbb-voting-service/internal/application/usecases/votes"
	domainProducer "bbb-voting-service/internal/domain/producer"
	domainRepositories "bbb-voting-service/internal/domain/repositories"
	domainServices "bbb-voting-service/internal/domain/services"
	consumer "bbb-voting-service/internal/infrastructure/consumer"
	controllers "bbb-voting-service/internal/infrastructure/controllers"
	"bbb-voting-service/internal/infrastructure/jobs"
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
		redis.NewRedisRepository,
		postgres.NewParticipantRepository,
		postgres.NewPostgresVoteRepository,
		producer.NewRabbitMQProducer,
		consumer.NewRabbitMQConsumer,
		controllers.NewCaptchaController,
		controllers.NewParticipantController,
		controllers.NewVoteController,
		controllers.NewResultController,
		captcha.NewGenerateCaptchaUsecase,
		captcha.NewServeCaptchaUsecase,
		captcha.NewValidateCaptchaUsecase,
		captcha.NewValidateCaptchaTokenUsecase,
		participants.NewCreateParticipantUsecase,
		participants.NewGetParticipantsUsecase,
		participants.NewGetParticipantUsecase,
		participants.NewDeleteParticipantUsecase,
		InitCastVoteUsecase,
		votes.NewProcessVoteUsecase,
		results.NewGetPartialResultsUsecase,
		results.NewGetFinalResultsUseCase,
		services.NewCaptchaService,
		InitSyncCacheUsecase,
		jobs.NewSyncCacheJob,
		InitCron,
		wire.Bind(new(domainProducer.MessageProducer), new(*producer.RabbitMQProducer)),
		wire.Bind(new(domainRepositories.InMemoryRepository), new(*redis.RedisRepository)),
		wire.Bind(new(domainRepositories.ParticipantRepository), new(*postgres.ParticipantRepository)),
		wire.Bind(new(domainRepositories.VoteRepository), new(*postgres.PostgresVoteRepository)),
		wire.Bind(new(domainServices.CaptchaService), new(*services.CaptchaService)),
		wire.Struct(new(Container), "*"),
	)
	return &Container{}, nil
}

func InitCastVoteUsecase(
	inMemoryRepository domainRepositories.InMemoryRepository,
	domainProducer domainProducer.MessageProducer,
	participantRepository domainRepositories.ParticipantRepository,
) *votes.CastVoteUsecase {
	voteQueue := os.Getenv("VOTE_QUEUE")

	return votes.NewCastVoteUsecase(inMemoryRepository, domainProducer, participantRepository, voteQueue)
}

func InitSyncCacheUsecase(
	voteRepository domainRepositories.VoteRepository,
	inMemoryRepository domainRepositories.InMemoryRepository,
) *cache.SyncCacheUsecase {
	return cache.NewSyncCacheUsecase(voteRepository, inMemoryRepository)
}
