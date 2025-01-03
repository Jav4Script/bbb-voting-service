package redis

import (
	"context"
	"log"
	"strconv"

	"bbb-voting-service/internal/domain/entities"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

type RedisResultRepository struct {
	Client *redis.Client
}

func NewRedisResultRepository(client *redis.Client) *RedisResultRepository {
	return &RedisResultRepository{Client: client}
}

// GetPartialResults retrieves a list of PartialResult, including participant details and votes.
func (redisResultRepository *RedisResultRepository) GetPartialResults(context context.Context) ([]entities.PartialResult, error) {
	// Fetch all vote counts from Redis
	partialResultsMap, err := redisResultRepository.Client.HGetAll(context, "partial_results").Result()
	if err != nil {
		log.Printf("Error fetching partial results: %v", err)
		return nil, err
	}

	// Prepare the result list
	partialResults := make([]entities.PartialResult, 0, len(partialResultsMap))
	for id, votesStr := range partialResultsMap {
		// Retrieve participant details
		participantData, err := redisResultRepository.Client.HGet(context, "participants", id).Result()
		if err == redis.Nil {
			log.Printf("Participant with ID %s not found", id)
			continue
		}
		if err != nil {
			log.Printf("Error fetching participant %s: %v", id, err)
			continue
		}

		// Deserialize participant data
		var participant entities.Participant
		if err := msgpack.Unmarshal([]byte(participantData), &participant); err != nil {
			log.Printf("Error unmarshalling participant %s: %v", id, err)
			continue
		}

		// Parse votes and create PartialResult
		votes, _ := strconv.Atoi(votesStr)
		partialResults = append(partialResults, entities.PartialResult{
			ID:     participant.ID,
			Name:   participant.Name,
			Age:    participant.Age,
			Gender: participant.Gender,
			Votes:  votes,
		})
	}

	return partialResults, nil
}

// UpdatePartialResults increments the vote count and updates participant details.
func (redisResultRepository *RedisResultRepository) UpdatePartialResults(context context.Context, vote entities.Vote, participant entities.Participant) error {
	pipe := redisResultRepository.Client.TxPipeline()

	// Increment votes
	pipe.HIncrBy(context, "partial_results", vote.ParticipantID, 1)

	// Serialize participant and store in Redis
	participantData, err := msgpack.Marshal(participant)
	if err != nil {
		log.Printf("Error serializing participant %s: %v", participant.ID, err)
		return err
	}
	pipe.HSet(context, "participants", participant.ID, participantData)

	// Execute transaction
	_, err = pipe.Exec(context)
	if err != nil {
		log.Printf("Error updating partial results: %v", err)
		return err
	}

	return nil
}

func (redisResultRepository *RedisResultRepository) UpdateCacheWithFinalResults(ctx context.Context, finalResults entities.FinalResults) error {
	pipe := redisResultRepository.Client.TxPipeline()

	for _, result := range finalResults.ParticipantResults {
		participantKey := "participant:" + result.ID

		// Serializar os dados do participante final
		participantData, err := msgpack.Marshal(result)
		if err != nil {
			log.Printf("Error serializing participant %s: %v", result.ID, err)
			return err
		}

		// Atualizar os resultados no Redis
		pipe.Set(ctx, participantKey, participantData, 0)
		pipe.HSet(ctx, "final_results", result.ID, result.Votes)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Printf("Error updating final results cache: %v", err)
		return err
	}

	return nil
}
