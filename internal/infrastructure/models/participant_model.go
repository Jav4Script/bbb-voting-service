package models

import (
	"time"

	entities "bbb-voting-service/internal/domain/entities"

	"github.com/google/uuid"
)

type ParticipantModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	Gender    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ParticipantModel) TableName() string {
	return "participants"
}

func ToDomainParticipant(participantModel ParticipantModel) entities.Participant {
	return entities.Participant{
		ID:        participantModel.ID.String(),
		Name:      participantModel.Name,
		Age:       participantModel.Age,
		Gender:    participantModel.Gender,
		CreatedAt: participantModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt: participantModel.UpdatedAt.Format(time.RFC3339),
	}
}

func ToModelParticipant(participant entities.Participant) ParticipantModel {
	id, _ := uuid.Parse(participant.ID)
	createdAt, _ := time.Parse(time.RFC3339, participant.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, participant.UpdatedAt)

	return ParticipantModel{
		ID:        id,
		Name:      participant.Name,
		Age:       participant.Age,
		Gender:    participant.Gender,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
