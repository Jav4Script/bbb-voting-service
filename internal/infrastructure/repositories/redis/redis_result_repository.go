package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"

	"bbb-voting-service/internal/domain"
	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"
	"bbb-voting-service/internal/infrastructure/mappers"
)

type RedisResultRepository struct {
	Client *redis.Client
}

// NewRedisResultRepository creates a new instance of RedisResultRepository.
func NewRedisResultRepository(client *redis.Client) *RedisResultRepository {
	return &RedisResultRepository{Client: client}
}

// GetPartialResults retrieves the partial voting results including participant details.
func (repository *RedisResultRepository) GetPartialResults(ctx context.Context) ([]entities.PartialResult, error) {
	// Fetch all vote counts from Redis hash
	partialResultsMap, err := repository.Client.HGetAll(ctx, domain.PartialResultsKey).Result()
	if err != nil {
		log.Printf("Error fetching partial results: %v", err)
		return nil, errors.NewInfrastructureError("Error fetching partial results from Redis")
	}

	// Prepare to collect partial results
	var partialResults []entities.PartialResult

	// Process each vote and fetch corresponding participant details
	for id, votesStr := range partialResultsMap {
		participant, err := repository.getParticipant(ctx, "participant:"+id)
		if err != nil {
			log.Printf("Error fetching participant details for ID %s: %v", id, err)
			continue // Skip participants with errors
		}

		// Map participant details into a partial result object
		partialResult := mappers.ToPartialResultFromParticipant(id, participant, votesStr)
		partialResults = append(partialResults, partialResult)
	}

	return partialResults, nil
}

// getParticipant retrieves participant details from Redis.
func (repository *RedisResultRepository) getParticipant(ctx context.Context, key string) (entities.Participant, error) {
	var participant entities.Participant

	// Fetch participant data from Redis
	data, err := repository.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return participant, errors.NewInfrastructureError("Participant not found")
	} else if err != nil {
		log.Printf("Error fetching participant with key %s: %v", key, err)
		return participant, errors.NewInfrastructureError("Error fetching participant from Redis")
	}

	// Deserialize the participant data
	if err := msgpack.Unmarshal([]byte(data), &participant); err != nil {
		log.Printf("Error unmarshalling participant data for key %s: %v", key, err)
		return participant, errors.NewInfrastructureError("Error unmarshalling participant data")
	}

	return participant, nil
}

// UpdatePartialResults increments the vote count for a participant and stores their details.
func (repository *RedisResultRepository) UpdatePartialResults(ctx context.Context, vote entities.Vote, participant entities.Participant) error {
	pipe := repository.Client.TxPipeline()

	// Check if the vote has already been processed
	existsCmd := pipe.SIsMember(ctx, domain.VoteSetKey, vote.ID)

	// Increment the vote count for the participant
	pipe.HIncrBy(ctx, domain.PartialResultsKey, vote.ParticipantID, 1)

	// Serialize and store participant details
	participantKey := "participant:" + vote.ParticipantID
	participantData, err := msgpack.Marshal(participant)
	if err != nil {
		log.Printf("Error serializing participant data: %v", err)
		return errors.NewInfrastructureError("Error serializing participant data")
	}
	pipe.Set(ctx, participantKey, participantData, 0)

	// Add the vote to the set of processed votes
	pipe.SAdd(ctx, domain.VoteSetKey, vote.ID)

	// Execute the pipeline
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("Error updating partial results in Redis: %v", err)
		return errors.NewInfrastructureError("Error updating partial results in Redis")
	}

	// Verify if the vote was already processed
	if exists, _ := existsCmd.Result(); exists {
		return errors.NewBusinessError("Vote already exists", 409)
	}

	return nil
}

// UpdateCacheWithFinalResults updates Redis with final voting results.
func (repository *RedisResultRepository) UpdateCacheWithFinalResults(ctx context.Context, finalResults entities.FinalResults) error {
	pipe := repository.Client.TxPipeline()

	// Process and store each participant's final result
	for _, result := range finalResults.ParticipantResults {
		participantKey, participantData, resultID, resultVotes := mappers.ToRedisData(result)

		// Serialize participant details
		data, err := msgpack.Marshal(participantData)
		if err != nil {
			log.Printf("Error serializing participant data: %v", err)
			return errors.NewInfrastructureError("Error serializing participant data")
		}

		// Store participant details and vote counts in Redis
		pipe.Set(ctx, participantKey, data, 0)
		pipe.HSet(ctx, domain.PartialResultsKey, resultID, resultVotes)
	}

	// Execute the pipeline
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("Error updating final results in Redis: %v", err)
		return errors.NewInfrastructureError("Error updating final results in Redis")
	}

	return nil
}
