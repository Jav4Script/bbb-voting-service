package models

import (
	entities "bbb-voting-service/internal/domain/entities"

	"github.com/google/uuid"
)

type ParticipantModel struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name  string    `gorm:"not null"`
	Votes int       `gorm:"not null"`
}

func (ParticipantModel) TableName() string {
	return "participants"
}

func ToDomainParticipant(participantModel ParticipantModel) entities.Participant {
	return entities.Participant{
		ID:    participantModel.ID.String(),
		Name:  participantModel.Name,
		Votes: participantModel.Votes,
	}
}

func ToModelParticipant(participant entities.Participant) ParticipantModel {
	id, _ := uuid.Parse(participant.ID)

	return ParticipantModel{
		ID:    id,
		Name:  participant.Name,
		Votes: participant.Votes,
	}
}
