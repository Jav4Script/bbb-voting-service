package config

import (
	"bbb-voting-service/internal/application/usecases/captcha"
	"bbb-voting-service/internal/application/usecases/participants"
	"bbb-voting-service/internal/application/usecases/results"
	"bbb-voting-service/internal/application/usecases/votes"
	"bbb-voting-service/internal/infrastructure/consumer"
	"bbb-voting-service/internal/infrastructure/controllers"
	"bbb-voting-service/internal/infrastructure/jobs"
	producer "bbb-voting-service/internal/infrastructure/producer"
	postgres "bbb-voting-service/internal/infrastructure/repositories/postgres"
	redisRepository "bbb-voting-service/internal/infrastructure/repositories/redis"
	"bbb-voting-service/internal/infrastructure/services"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type Container struct {
	DB                            *gorm.DB
	RedisClient                   *redis.Client
	RabbitMQChannel               *amqp.Channel
	ParticipantRepository         *postgres.PostgresParticipantRepository
	ResultRepository              *postgres.PostgresResultRepository
	InMemoryParticipantRepository *redisRepository.RedisParticipantRepository
	InMemoryResultRepository      *redisRepository.RedisResultRepository
	RabbitMQProducer              *producer.RabbitMQProducer
	RabbitMQConsumer              *consumer.RabbitMQConsumer
	SyncCacheJob                  *jobs.SyncCacheJob
	Cron                          *cron.Cron
	GenerateCaptchaUsecase        *captcha.GenerateCaptchaUsecase
	ServeCaptchaUsecase           *captcha.ServeCaptchaUsecase
	ValidateCaptchaUsecase        *captcha.ValidateCaptchaUsecase
	ValidateCaptchaTokenUsecase   *captcha.ValidateCaptchaTokenUsecase
	CreateParticipantUsecase      *participants.CreateParticipantUsecase
	GetParticipantsUsecase        *participants.GetParticipantsUsecase
	DeleteParticipantUsecase      *participants.DeleteParticipantUsecase
	GetParticipantUsecase         *participants.GetParticipantUsecase
	CastVoteUsecase               *votes.CastVoteUsecase
	ProcessVoteUsecase            *votes.ProcessVoteUsecase
	GetPartialResultsUsecase      *results.GetPartialResultsUsecase
	GetFinalResultsUseCase        *results.GetFinalResultsUsecase
	CaptchaService                *services.CaptchaService
	CaptchaController             *controllers.CaptchaController
	ParticipantController         *controllers.ParticipantController
	VoteController                *controllers.VoteController
	ResultController              *controllers.ResultController
}
