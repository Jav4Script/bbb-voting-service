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
		redis.NewRedisParticipantRepository,
		redis.NewRedisResultRepository,
		postgres.NewPostgresParticipantRepository,
		postgres.NewPostgresResultRepository,
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
		InitCreateParticipantUsecase,
		InitGetParticipantUsecase,
		InitGetParticipantsUsecase,
		InitDeleteParticipantUsecase,
		InitCastVoteUsecase,
		votes.NewProcessVoteUsecase,
		results.NewGetPartialResultsUsecase,
		results.NewGetFinalResultsUsecase,
		services.NewCaptchaService,
		InitSyncParticipantsCacheUsecase,
		InitSyncResultsCacheUsecase,
		InitSyncCacheJob,
		InitCron,
		wire.Bind(new(domainProducer.MessageProducer), new(*producer.RabbitMQProducer)),
		wire.Bind(new(domainRepositories.InMemoryParticipantRepository), new(*redis.RedisParticipantRepository)),
		wire.Bind(new(domainRepositories.InMemoryResultRepository), new(*redis.RedisResultRepository)),
		wire.Bind(new(domainRepositories.ParticipantRepository), new(*postgres.PostgresParticipantRepository)),
		wire.Bind(new(domainRepositories.ResultRepository), new(*postgres.PostgresResultRepository)),
		wire.Bind(new(domainServices.CaptchaService), new(*services.CaptchaService)),
		wire.Struct(new(Container), "*"),
	)
	return &Container{}, nil
}

func InitCreateParticipantUsecase(
	participantRepository domainRepositories.ParticipantRepository,
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
) *participants.CreateParticipantUsecase {
	return participants.NewCreateParticipantUsecase(participantRepository, inMemoryParticipantRepository)
}

func InitGetParticipantUsecase(
	participantRepository domainRepositories.ParticipantRepository,
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
) *participants.GetParticipantUsecase {
	return participants.NewGetParticipantUsecase(participantRepository, inMemoryParticipantRepository)
}

func InitGetParticipantsUsecase(
	participantRepository domainRepositories.ParticipantRepository,
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
) *participants.GetParticipantsUsecase {
	return participants.NewGetParticipantsUsecase(participantRepository, inMemoryParticipantRepository)
}

func InitDeleteParticipantUsecase(
	participantRepository domainRepositories.ParticipantRepository,
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
) *participants.DeleteParticipantUsecase {
	return participants.NewDeleteParticipantUsecase(participantRepository, inMemoryParticipantRepository)
}

func InitCastVoteUsecase(
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
	inMemoryResultRepository domainRepositories.InMemoryResultRepository,
	domainProducer domainProducer.MessageProducer,
) *votes.CastVoteUsecase {
	voteQueue := os.Getenv("VOTE_QUEUE")

	return votes.NewCastVoteUsecase(inMemoryParticipantRepository, inMemoryResultRepository, domainProducer, voteQueue)
}

func InitSyncCacheJob(
	syncParticipantsCacheUsecase *cache.SyncParticipantsCacheUsecase,
	syncResultsCacheUsecase *cache.SyncResultsCacheUsecase,
) *jobs.SyncCacheJob {
	return jobs.NewSyncCacheJob(syncResultsCacheUsecase, syncParticipantsCacheUsecase)
}

func InitSyncParticipantsCacheUsecase(
	participantRepository domainRepositories.ParticipantRepository,
	inMemoryParticipantRepository domainRepositories.InMemoryParticipantRepository,
) *cache.SyncParticipantsCacheUsecase {
	return cache.NewSyncParticipantsCacheUsecase(participantRepository, inMemoryParticipantRepository)
}

func InitSyncResultsCacheUsecase(
	getFinalResultsUsecase *results.GetFinalResultsUsecase,
	inMemoryResultRepository domainRepositories.InMemoryResultRepository,
) *cache.SyncResultsCacheUsecase {
	return cache.NewSyncResultsCacheUsecase(getFinalResultsUsecase, inMemoryResultRepository)
}
