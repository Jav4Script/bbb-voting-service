package models

import (
	"time"

	entities "bbb-voting-service/internal/domain/entities"

	"github.com/google/uuid"
)

type VoteModel struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ParticipantID      string    `gorm:"not null"`
	VoterID            string    `gorm:"not null"`
	Timestamp          time.Time `gorm:"not null"`
	IPAddress          string    `gorm:"not null"`
	UserAgent          string    `gorm:"not null"`
	Region             string    `gorm:"not null"`
	Device             string    `gorm:"not null"`
	IsCaptchaCompleted bool      `gorm:"not null"`
}

func (VoteModel) TableName() string {
	return "votes"
}

func ToDomainVote(voteModel VoteModel) entities.Vote {
	return entities.Vote{
		ID:                 voteModel.ID.String(),
		ParticipantID:      voteModel.ParticipantID,
		VoterID:            voteModel.VoterID,
		Timestamp:          voteModel.Timestamp,
		IPAddress:          voteModel.IPAddress,
		UserAgent:          voteModel.UserAgent,
		Region:             voteModel.Region,
		Device:             voteModel.Device,
		IsCaptchaCompleted: voteModel.IsCaptchaCompleted,
	}
}

func ToModelVote(vote entities.Vote) VoteModel {
	id, _ := uuid.Parse(vote.ID)

	return VoteModel{
		ID:                 id,
		ParticipantID:      vote.ParticipantID,
		VoterID:            vote.VoterID,
		Timestamp:          vote.Timestamp,
		IPAddress:          vote.IPAddress,
		UserAgent:          vote.UserAgent,
		Region:             vote.Region,
		Device:             vote.Device,
		IsCaptchaCompleted: vote.IsCaptchaCompleted,
	}
}
