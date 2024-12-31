package redis

import (
	"context"
	"log"
	"time"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5" // Library for efficient serialization
)

type RedisParticipantRepository struct {
	Client *redis.Client
}

func NewRedisParticipantRepository(client *redis.Client) *RedisParticipantRepository {
	return &RedisParticipantRepository{Client: client}
}

func (repository *RedisParticipantRepository) Save(participant entities.Participant) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Serialize participant using MessagePack
	participantData, err := msgpack.Marshal(participant)
	if err != nil {
		log.Printf("Error marshaling participant data: %v", err)
		return errors.NewInfrastructureError("Failed to marshal participant data")
	}

	// Save participant to Redis with a TTL of 24 hours
	const participantTTL = 24 * time.Hour
	err = repository.Client.Set(ctx, "participant:"+participant.ID, participantData, participantTTL).Err()
	if err != nil {
		log.Printf("Error saving participant in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to save participant in Redis")
	}

	return nil
}

func (repository *RedisParticipantRepository) FindAll() ([]entities.Participant, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor uint64
	var participants []entities.Participant

	for {
		// Use SCAN to fetch keys incrementally
		keys, newCursor, err := repository.Client.Scan(ctx, cursor, "participant:*", 100).Result()
		if err != nil {
			log.Printf("Error scanning participant keys from Redis: %v", err)
			return nil, errors.NewInfrastructureError("Failed to scan participant keys from Redis")
		}

		if len(keys) > 0 {
			// Use MGET to fetch multiple keys in one call
			participantDataList, err := repository.Client.MGet(ctx, keys...).Result()
			if err != nil {
				log.Printf("Error getting participants from Redis: %v", err)
				return nil, errors.NewInfrastructureError("Failed to get participants from Redis")
			}

			// Deserialize data returned by MGET
			for _, data := range participantDataList {
				if data == nil {
					continue // Skip nonexistent keys
				}

				var participant entities.Participant
				err = msgpack.Unmarshal([]byte(data.(string)), &participant)
				if err != nil {
					log.Printf("Error unmarshaling participant data: %v", err)
					return nil, errors.NewInfrastructureError("Failed to unmarshal participant data")
				}
				participants = append(participants, participant)
			}
		}

		// Update the cursor
		cursor = newCursor
		if cursor == 0 { // SCAN completed
			break
		}
	}

	return participants, nil
}

func (repository *RedisParticipantRepository) FindByID(id string) (entities.Participant, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch participant by ID
	participantData, err := repository.Client.Get(ctx, "participant:"+id).Result()
	if err == redis.Nil {
		return entities.Participant{}, errors.NewNotFoundError("Participant not found")
	} else if err != nil {
		log.Printf("Error getting participant from Redis: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to get participant from Redis")
	}

	// Deserialize the participant
	var participant entities.Participant
	err = msgpack.Unmarshal([]byte(participantData), &participant)
	if err != nil {
		log.Printf("Error unmarshaling participant data: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to unmarshal participant data")
	}

	return participant, nil
}

func (repository *RedisParticipantRepository) Delete(id string) error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete participant from Redis
	err := repository.Client.Del(ctx, "participant:"+id).Err()
	if err != nil {
		log.Printf("Error deleting participant from Redis: %v", err)
		return errors.NewInfrastructureError("Failed to delete participant from Redis")
	}

	return nil
}
