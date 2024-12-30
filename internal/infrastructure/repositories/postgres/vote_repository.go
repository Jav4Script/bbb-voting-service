package repositories

import (
	"time"

	entities "bbb-voting-service/internal/domain/entities"
	mappers "bbb-voting-service/internal/infrastructure/mappers"
	models "bbb-voting-service/internal/infrastructure/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresVoteRepository struct {
	DB *gorm.DB
}

func NewPostgresVoteRepository(database *gorm.DB) *PostgresVoteRepository {
	return &PostgresVoteRepository{
		DB: database,
	}
}

func (repository *PostgresVoteRepository) Save(vote entities.Vote) error {
	voteModel := models.ToModelVote(vote)

	err := repository.DB.Create(&voteModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostgresVoteRepository) CountTotalVotes() (int, error) {
	var count int64
	err := repository.DB.Model(&models.VoteModel{}).Count(&count).Error
	return int(count), err
}

func (repository *PostgresVoteRepository) CountVotesByParticipant(participantID string) (int, error) {
	var count int64
	err := repository.DB.Model(&models.VoteModel{}).Where("participant_id = ?", participantID).Count(&count).Error
	return int(count), err
}

func (repository *PostgresVoteRepository) CountVotesByHour() (map[time.Time]int, error) {
	var results []struct {
		Hour  time.Time
		Count int
	}
	err := repository.DB.Model(&models.VoteModel{}).
		Select("date_trunc('hour', created_at) as hour, count(*) as count").
		Group("hour").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	votesByHour := make(map[time.Time]int)
	for _, result := range results {
		votesByHour[result.Hour] = result.Count
	}

	return votesByHour, nil
}

func (repository *PostgresVoteRepository) GetParticipantResults() (map[string]entities.ParticipantResult, error) {
	var results []struct {
		ParticipantID uuid.UUID
		Name          string
		Age           int
		Gender        string
		Votes         int
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	err := repository.DB.Model(&models.VoteModel{}).
		Select("participant_id, participants.name, participants.age, participants.gender, participants.created_at, participants.updated_at, count(*) as votes").
		Joins("left join participants on participants.id = votes.participant_id").
		Group("participant_id, participants.name, participants.age, participants.gender, participants.created_at, participants.updated_at").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	finalResults := make(map[string]entities.ParticipantResult)
	for _, result := range results {
		finalResults[result.ParticipantID.String()] = mappers.ToParticipantResultFromStruct(result)
	}

	return finalResults, nil
}
