package redis

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"
	"bbb-voting-service/internal/domain/errors"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"
)

type RedisParticipantRepository struct {
	Client *redis.Client
}

func NewRedisParticipantRepository(client *redis.Client) *RedisParticipantRepository {
	return &RedisParticipantRepository{Client: client}
}

// Save stores or updates a participant in Redis.
func (repository *RedisParticipantRepository) Save(context context.Context, participant entities.Participant) error {
	// Serialize the participant using msgpack
	participantData, err := msgpack.Marshal(participant)
	if err != nil {
		log.Printf("Error marshaling participant %s: %v", participant.ID, err)
		return errors.NewInfrastructureError("Failed to marshal participant data")
	}

	// Save the serialized data into the Redis hash
	err = repository.Client.HSet(context, "participants", participant.ID, participantData).Err()
	if err != nil {
		log.Printf("Error saving participant %s to Redis: %v", participant.ID, err)
		return errors.NewInfrastructureError("Failed to save participant to Redis")
	}

	return nil
}

// FindAll retrieves all participants from Redis.
func (repository *RedisParticipantRepository) FindAll(context context.Context) ([]entities.Participant, error) {
	// Get all participants from the Redis hash
	participantsData, err := repository.Client.HGetAll(context, "participants").Result()
	if err != nil {
		log.Printf("Error retrieving all participants from Redis: %v", err)
		return nil, errors.NewInfrastructureError("Failed to retrieve participants from Redis")
	}

	var participants []entities.Participant
	for _, data := range participantsData {
		var participant entities.Participant
		err := msgpack.Unmarshal([]byte(data), &participant)
		if err != nil {
			log.Printf("Error deserializing participant data: %v", err)
			continue // Skip invalid entries
		}
		participants = append(participants, participant)
	}

	return participants, nil
}

// FindByID retrieves a participant by ID from Redis.
func (repository *RedisParticipantRepository) FindByID(context context.Context, id string) (entities.Participant, error) {
	// Fetch the participant data from Redis
	data, err := repository.Client.HGet(context, "participants", id).Result()
	if err == redis.Nil {
		return entities.Participant{}, errors.NewNotFoundError("Participant not found")
	}
	if err != nil {
		log.Printf("Error retrieving participant %s from Redis: %v", id, err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to retrieve participant")
	}

	// Deserialize the participant data
	var participant entities.Participant
	err = msgpack.Unmarshal([]byte(data), &participant)
	if err != nil {
		log.Printf("Error deserializing participant %s data: %v", id, err)
		return entities.Participant{}, errors.NewInfrastructureError("Failed to deserialize participant data")
	}

	return participant, nil
}

// Delete removes a participant by ID from Redis.
func (repository *RedisParticipantRepository) Delete(context context.Context, id string) error {
	// Remove the participant from the Redis hash
	err := repository.Client.HDel(context, "participants", id).Err()
	if err != nil {
		log.Printf("Error deleting participant %s from Redis: %v", id, err)
		return errors.NewInfrastructureError("Failed to delete participant from Redis")
	}

	return nil
}
