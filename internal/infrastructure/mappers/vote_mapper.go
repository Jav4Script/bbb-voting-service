package mappers

import (
	"bbb-voting-service/internal/domain/dtos"
	entities "bbb-voting-service/internal/domain/entities"

	"github.com/google/uuid"
)

func FromCastVoteDTO(castVoteDTO dtos.CastVoteDTO) entities.Vote {
	return entities.Vote{
		ID:            uuid.New().String(),
		ParticipantID: castVoteDTO.ParticipantID,
		VoterID:       castVoteDTO.VoterID,
		IPAddress:     castVoteDTO.IPAddress,
		UserAgent:     castVoteDTO.UserAgent,
		Region:        castVoteDTO.Region,
		Device:        castVoteDTO.Device,
	}
}
