package redis

import (
	"context"
	"encoding/json"
	"log"

	entities "bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"

	"github.com/go-redis/redis/v8"
)

type RedisParticipantRepository struct {
	Client *redis.Client
}

func NewRedisParticipantRepository(client *redis.Client) *RedisParticipantRepository {
	return &RedisParticipantRepository{Client: client}
}

func (repository *RedisParticipantRepository) Save(participant entities.Participant) error {
	ctx := context.Background()
	participantData, err := json.Marshal(participant)
	if err != nil {
		log.Printf("Error marshaling participant data: %v", err)
		return errors.NewInfrastructureError("Failed to marshal participant data")
	}

	err = repository.Client.Set(ctx, "participant:"+participant.ID, participantData, 0).Err()
	if err != nil {
		log.Printf("Error saving participant in Redis: %v", err)
		return errors.NewInfrastructureError("Failed to save participant in Redis")
	}

	return nil
}

func (repository *RedisParticipantRepository) FindAll() ([]entities.Participant, error) {
	ctx := context.Background()
	keys, err := repository.Client.Keys(ctx, "participant:*").Result()
	if err != nil {
		log.Printf("Error getting participant keys from Redis: %v", err)
		return nil, errors.NewInfrastructureError("Failed to get participant keys from Redis")
	}

	participants := make([]entities.Participant, 0)
	for _, key := range keys {
		participantData, err := repository.Client.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Error getting participant from Redis: %v", err)
			return nil, errors.NewInfrastructureError("Failed to get participant from Redis")
		}

		var participant entities.Participant
		err = json.Unmarshal([]byte(participantData), &participant)
		if err != nil {
			log.Printf("Error unmarshaling participant data: %v", err)
			return nil, errors.NewInfrastructureError("Failed to unmarshal participant data")
		}

		participants = append(participants, participant)
	}

	return participants, nil
}

func (repository *RedisParticipantRepository) FindByID(id string) (entities.Participant, error) {
	ctx := context.Background()
	participantData, err := repository.Client.Get(ctx, "participant:"+id).Result()
	if err == redis.Nil {
		return entities.Participant{}, errors.NewNotFoundError("Participant not found")
	} else if err != nil {
		log.Printf("Error getting participant from Redis: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to get participant from Redis")
	}

	var participant entities.Participant
	err = json.Unmarshal([]byte(participantData), &participant)
	if err != nil {
		log.Printf("Error unmarshaling participant data: %v", err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to unmarshal participant data")
	}

	return participant, nil
}

func (repository *RedisParticipantRepository) Delete(id string) error {
	ctx := context.Background()
	err := repository.Client.Del(ctx, "participant:"+id).Err()
	if err != nil {
		log.Printf("Error deleting participant from Redis: %v", err)
		return errors.NewInfrastructureError("Failed to delete participant from Redis")
	}

	return nil
}
