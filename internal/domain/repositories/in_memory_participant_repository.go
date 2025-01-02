package repositories

import (
	"bbb-voting-service/internal/domain/entities"
	"context"
)

type InMemoryParticipantRepository interface {
	Save(context context.Context, participant entities.Participant) error
	FindAll(context context.Context) ([]entities.Participant, error)
	FindByID(context context.Context, id string) (entities.Participant, error)
	Delete(context context.Context, id string) error
}
