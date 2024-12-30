package mappers

import (
	entities "bbb-voting-service/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

func ToParticipantResult(participantID string, name string, age int, gender string, votes int, createdAt time.Time, updatedAt time.Time) entities.ParticipantResult {
	return entities.ParticipantResult{
		ID:        participantID,
		Name:      name,
		Age:       age,
		Gender:    gender,
		Votes:     votes,
		CreatedAt: createdAt.Format(time.RFC3339),
		UpdatedAt: updatedAt.Format(time.RFC3339),
	}
}

func ToParticipantResultFromStruct(data struct {
	ParticipantID uuid.UUID
	Name          string
	Age           int
	Gender        string
	Votes         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}) entities.ParticipantResult {
	return entities.ParticipantResult{
		ID:        data.ParticipantID.String(),
		Name:      data.Name,
		Age:       data.Age,
		Gender:    data.Gender,
		Votes:     data.Votes,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
}
