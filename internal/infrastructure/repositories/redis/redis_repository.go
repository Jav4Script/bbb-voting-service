package repositories

import (
	"context"
	"encoding/json"

	entities "bbb-voting-service/internal/domain/entities"

	redis "github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{Client: client}
}

func (r *RedisRepository) GetPartialResults() (map[string]int, error) {
	ctx := context.Background()
	result, err := r.Client.Get(ctx, "partial_results").Result()
	if err != nil {
		return nil, err
	}

	var partialResults map[string]int
	err = json.Unmarshal([]byte(result), &partialResults)
	if err != nil {
		return nil, err
	}

	return partialResults, nil
}

func (r *RedisRepository) SavePartialVote(vote entities.Vote) error {
	ctx := context.Background()
	voteData, err := json.Marshal(vote)
	if err != nil {
		return err
	}

	return r.Client.LPush(ctx, "partial_votes", voteData).Err()
}
