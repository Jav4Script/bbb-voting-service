package repositories

import (
	"time"

	entities "bbb-voting-service/internal/domain/entities"

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
	return repository.DB.Create(&vote).Error
}

func (repository *PostgresVoteRepository) CountTotalVotes() (int, error) {
	var count int64
	err := repository.DB.Model(&entities.Vote{}).Count(&count).Error
	return int(count), err
}

func (repository *PostgresVoteRepository) CountVotesByParticipant(participantID string) (int, error) {
	var count int64
	err := repository.DB.Model(&entities.Vote{}).Where("participant_id = ?", participantID).Count(&count).Error
	return int(count), err
}

func (repository *PostgresVoteRepository) CountVotesByHour(sessionID string) (map[time.Time]int, error) {
	var results []struct {
		Hour  time.Time
		Count int
	}
	err := repository.DB.Model(&entities.Vote{}).
		Select("date_trunc('hour', timestamp) as hour, count(*) as count").
		Group("hour").
		Where("session_id = ?", sessionID).
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

func (repository *PostgresVoteRepository) GetFinalResults() (map[string]int, error) {
	var results []struct {
		ParticipantID string
		Count         int
	}
	err := repository.DB.Model(&entities.Vote{}).
		Select("participant_id, count(*) as count").
		Group("participant_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	finalResults := make(map[string]int)
	for _, result := range results {
		finalResults[result.ParticipantID] = result.Count
	}

	return finalResults, nil
}