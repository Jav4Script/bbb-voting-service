package redis

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"bbb-voting-service/internal/domain"
	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"

	"github.com/go-redis/redis/v8"
)

type RedisResultRepository struct {
	Client *redis.Client
}

func NewRedisResultRepository(client *redis.Client) *RedisResultRepository {
	return &RedisResultRepository{Client: client}
}

// GetPartialResults - Fetch all partial results from the Redis hash
func (repository *RedisResultRepository) GetPartialResults() ([]entities.PartialResult, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch all fields from the hash
	partialResultsMap, err := repository.Client.HGetAll(ctx, domain.PartialResultsKey).Result()
	if err != nil {
		log.Printf("Failed to fetch partial results from Redis: %v", err)
		return nil, errors.NewInfrastructureError("Failed to fetch partial results from Redis")
	}

	// Convert the hash to a slice of PartialResult
	partialResults := make([]entities.PartialResult, 0, len(partialResultsMap))
	for id, votesStr := range partialResultsMap {
		partialResults = append(partialResults, entities.PartialResult{
			ID:    id,
			Votes: parseVotes(votesStr),
		})
	}

	return partialResults, nil
}

// UpdatePartialResults - Increment the vote count for a participant
func (repository *RedisResultRepository) UpdatePartialResults(vote entities.Vote, participant entities.Participant) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the vote already exists
	log.Printf("Checking if vote exists: %s", vote.ID)
	exists, err := repository.Client.SIsMember(ctx, domain.VoteSetKey, vote.ID).Result()
	if err != nil {
		log.Printf("Error checking vote existence in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to check vote existence in Redis")
	}

	if exists {
		log.Printf("Vote already exists: %s", vote.ID)
		return errors.NewBusinessError("Vote already exists", http.StatusConflict)
	}

	// Increment the vote count for the participant in the hash
	log.Printf("Incrementing vote count for participant: %s", vote.ParticipantID)
	err = repository.Client.HIncrBy(ctx, domain.PartialResultsKey, vote.ParticipantID, 1).Err()
	if err != nil {
		log.Printf("Error incrementing vote count in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to increment vote count in Redis")
	}

	// Add the vote ID to the set of processed votes
	log.Printf("Adding vote ID to Redis set: %s", vote.ID)
	err = repository.Client.SAdd(ctx, domain.VoteSetKey, vote.ID).Err()
	if err != nil {
		log.Printf("Error adding vote ID to Redis set: %v", err)
		return errors.NewInfrastructureError("Failed to add vote ID to Redis set")
	}

	return nil
}

// UpdateCacheWithFinalResults - Overwrite the Redis hash with final results
func (repository *RedisResultRepository) UpdateCacheWithFinalResults(finalResults entities.FinalResults) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare the hash fields for Redis
	partialResults := make(map[string]interface{})
	for _, participant := range finalResults.ParticipantResults {
		partialResults[participant.ID] = participant.Votes
	}

	// Overwrite the hash in Redis with the new final results
	err := repository.Client.HMSet(ctx, domain.PartialResultsKey, partialResults).Err()
	if err != nil {
		log.Printf("Error saving final results to Redis: %v", err)
		return errors.NewInfrastructureError("Failed to save final results to Redis")
	}

	return nil
}

// parseVotes - Helper function to parse votes from a string to an integer
func parseVotes(votesStr string) int {
	votes, err := strconv.Atoi(votesStr)
	if err != nil {
		log.Printf("Error parsing votes string '%s': %v", votesStr, err)
		return 0
	}
	return votes
}
