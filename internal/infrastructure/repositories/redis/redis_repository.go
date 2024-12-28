package repositories

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"bbb-voting-service/internal/domain"
	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/infrastructure/mappers"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{Client: client}
}

func (redisRepository *RedisRepository) GetPartialResults() ([]entities.PartialResult, error) {
	ctx := context.Background()
	result, err := redisRepository.Client.Get(ctx, domain.PartialResultsKey).Result()
	if err == redis.Nil {
		log.Printf("Partial results not found in Redis: %v", err)
		return nil, errors.ErrorNotFound
	} else if err != nil {
		log.Printf("Failed to get partial results from Redis: %v", err)
		return nil, errors.NewInfrastructureError("Failed to get partial results from Redis")
	}

	var partialResults []entities.PartialResult
	err = json.Unmarshal([]byte(result), &partialResults)
	if err != nil {
		log.Printf("Failed to unmarshal partial results: %v", err)
		return nil, errors.NewInfrastructureError("Failed to unmarshal partial results")
	}

	return partialResults, nil
}

func (redisRepository *RedisRepository) UpdatePartialResults(vote entities.Vote, participant entities.Participant) error {
	ctx := context.Background()

	// Check if the vote already exists
	log.Printf("Checking if vote exists: %s", vote.ID) // Adicione este log
	exists, err := redisRepository.Client.SIsMember(ctx, domain.VoteSetKey, vote.ID).Result()
	if err != nil {
		log.Printf("Error checking vote existence in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to check vote existence in Redis")
	}
	if exists {
		log.Printf("Vote already exists: %s", vote.ID)
		return errors.NewBusinessError("Vote already exists", http.StatusConflict)
	}

	// Get current partial results
	partialResults, err := redisRepository.GetPartialResults()
	if err != nil && err != errors.ErrorNotFound {
		log.Printf("Error getting partial results from Redis: %v", err)
		return errors.NewInfrastructureError("Failed to get partial results from Redis")
	}

	// Update the partial results with the new vote
	updated := false
	for i, result := range partialResults {
		if result.ParticipantID == vote.ParticipantID {
			partialResults[i].Votes++
			updated = true
			break
		}
	}

	if !updated {
		partialResults = append(partialResults, mappers.ToPartialResult(participant))
	}

	partialResultsData, err := json.Marshal(partialResults)
	if err != nil {
		log.Printf("Error marshaling partial results data: %v", err)
		return errors.NewInfrastructureError("Failed to marshal partial results data")
	}

	// Save the updated partial results in Redis
	err = redisRepository.Client.Set(ctx, domain.PartialResultsKey, partialResultsData, 0).Err()
	if err != nil {
		log.Printf("Error saving partial results in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to save partial results in Redis")
	}

	// Add the vote ID to the set of processed votes
	log.Printf("Adding vote ID to Redis set: %s", vote.ID) // Adicione este log
	err = redisRepository.Client.SAdd(ctx, domain.VoteSetKey, vote.ID).Err()
	if err != nil {
		log.Printf("Error adding vote ID to Redis set: %v", err)
		return errors.NewInfrastructureError("Failed to add vote ID to Redis set")
	}

	return nil
}
