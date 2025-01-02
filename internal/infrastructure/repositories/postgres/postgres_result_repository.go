package repositories

import (
	"time"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/infrastructure/mappers"
	"bbb-voting-service/internal/infrastructure/models"

	"gorm.io/gorm"
)

type PostgresResultRepository struct {
	DB *gorm.DB
}

func NewPostgresResultRepository(database *gorm.DB) *PostgresResultRepository {
	return &PostgresResultRepository{DB: database}
}

func (repository *PostgresResultRepository) Save(vote entities.Vote) error {
	voteModel := models.ToModelVote(vote)
	return repository.DB.Create(&voteModel).Error
}

func (repository *PostgresResultRepository) CountTotalVotes() (int, error) {
	var count int64
	err := repository.DB.Model(&models.VoteModel{}).Count(&count).Error
	return int(count), err
}

func (repository *PostgresResultRepository) CountVotesByParticipant(participantID string) (int, error) {
	var count int64
	err := repository.DB.Model(&models.VoteModel{}).Where("participant_id = ?", participantID).Count(&count).Error
	return int(count), err
}

func (repository *PostgresResultRepository) CountVotesByHour() (map[time.Time]int, error) {
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

func (repository *PostgresResultRepository) GetParticipantResults() (map[string]entities.ParticipantResult, error) {
	var results []entities.ParticipantResult

	err := repository.DB.Model(&models.VoteModel{}).
		Select("participant_id as id, participants.name, participants.age, participants.gender, participants.created_at, participants.updated_at, count(*) as votes").
		Joins("left join participants on participants.id = votes.participant_id").
		Group("participant_id, participants.name, participants.age, participants.gender, participants.created_at, participants.updated_at").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	finalResults := make(map[string]entities.ParticipantResult)
	for _, result := range results {
		finalResults[result.ID] = result
	}

	return finalResults, nil
}

func (repository *PostgresResultRepository) GetFinalResults() (entities.FinalResults, error) {
	// Get total votes
	totalVotes, err := repository.CountTotalVotes()
	if err != nil {
		return entities.FinalResults{}, err
	}

	// Get votes by participant
	participantResults, err := repository.GetParticipantResults()
	if err != nil {
		return entities.FinalResults{}, err
	}

	// Get votes by hour
	votesByHour, err := repository.CountVotesByHour()
	if err != nil {
		return entities.FinalResults{}, err
	}

	// Aggregate results using the mapper
	finalResultsEntity := mappers.ToFinalResults(participantResults, totalVotes, votesByHour)

	return finalResultsEntity, nil
}
