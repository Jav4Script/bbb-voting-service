package models

import (
	"time"

	entities "bbb-voting-service/internal/domain/entities"

	"github.com/google/uuid"
)

type VoteModel struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ParticipantID uuid.UUID `gorm:"type:uuid;not null"`
	VoterID       string    `gorm:"not null"`
	IPAddress     string    `gorm:"not null"`
	UserAgent     string    `gorm:"not null"`
	Region        string    `gorm:"not null"`
	Device        string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (VoteModel) TableName() string {
	return "votes"
}

func ToDomainVote(voteModel VoteModel) entities.Vote {
	return entities.Vote{
		ID:            voteModel.ID.String(),
		ParticipantID: voteModel.ParticipantID.String(),
		VoterID:       voteModel.VoterID,
		IPAddress:     voteModel.IPAddress,
		UserAgent:     voteModel.UserAgent,
		Region:        voteModel.Region,
		Device:        voteModel.Device,
		CreatedAt:     voteModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     voteModel.UpdatedAt.Format(time.RFC3339),
	}
}

func ToModelVote(vote entities.Vote) VoteModel {
	id, _ := uuid.Parse(vote.ID)
	participantID, _ := uuid.Parse(vote.ParticipantID)
	createdAt, _ := time.Parse(time.RFC3339, vote.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, vote.UpdatedAt)

	return VoteModel{
		ID:            id,
		ParticipantID: participantID,
		VoterID:       vote.VoterID,
		IPAddress:     vote.IPAddress,
		UserAgent:     vote.UserAgent,
		Region:        vote.Region,
		Device:        vote.Device,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
