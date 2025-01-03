package redis

import (
	"context"
	"log"

	"bbb-voting-service/internal/domain/entities"

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
func (redisParticipantRepository *RedisParticipantRepository) Save(context context.Context, participant entities.Participant) error {
	participantData, err := msgpack.Marshal(participant)
	if err != nil {
		log.Printf("Error serializing participant %s: %v", participant.ID, err)
		return err
	}

	err = redisParticipantRepository.Client.HSet(context, "participants", participant.ID, participantData).Err()
	if err != nil {
		log.Printf("Error saving participant %s: %v", participant.ID, err)
		return err
	}

	return nil
}

// FindByID retrieves a participant by ID from Redis.
func (redisParticipantRepository *RedisParticipantRepository) FindByID(context context.Context, id string) (entities.Participant, error) {
	data, err := redisParticipantRepository.Client.HGet(context, "participants", id).Result()
	if err == redis.Nil {
		log.Printf("Participant with ID %s not found", id)
		return entities.Participant{}, nil
	}
	if err != nil {
		log.Printf("Error fetching participant %s: %v", id, err)
		return entities.Participant{}, err
	}

	var participant entities.Participant
	if err := msgpack.Unmarshal([]byte(data), &participant); err != nil {
		log.Printf("Error unmarshalling participant %s: %v", id, err)
		return entities.Participant{}, err
	}

	return participant, nil
}

// FindAll retrieves all participants from Redis.
func (redisParticipantRepository *RedisParticipantRepository) FindAll(context context.Context) ([]entities.Participant, error) {
	participantsData, err := redisParticipantRepository.Client.HGetAll(context, "participants").Result()
	if err != nil {
		log.Printf("Error fetching participants: %v", err)
		return nil, err
	}

	participants := make([]entities.Participant, 0, len(participantsData))
	for id, data := range participantsData {
		var participant entities.Participant
		if err := msgpack.Unmarshal([]byte(data), &participant); err != nil {
			log.Printf("Error unmarshalling participant %s: %v", id, err)
			continue
		}
		participants = append(participants, participant)
	}

	return participants, nil
}

// Delete removes a participant by ID from Redis.
func (redisParticipantRepository *RedisParticipantRepository) Delete(context context.Context, id string) error {
	err := redisParticipantRepository.Client.HDel(context, "participants", id).Err()
	if err != nil {
		log.Printf("Error deleting participant %s: %v", id, err)
		return err
	}

	return nil
}
