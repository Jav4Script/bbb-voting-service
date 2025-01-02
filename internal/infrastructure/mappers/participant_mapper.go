package mappers

import (
	"bbb-voting-service/internal/domain/dtos"
	entities "bbb-voting-service/internal/domain/entities"
)

func FromCreateParticipantDTO(createParticipantDTO dtos.CreateParticipantDTO) entities.Participant {
	return entities.Participant{
		Name:   createParticipantDTO.Name,
		Age:    createParticipantDTO.Age,
		Gender: createParticipantDTO.Gender,
	}
}

func FromParticipantResult(participantResult entities.ParticipantResult) entities.Participant {
	return entities.Participant{
		ID:        participantResult.ID,
		Name:      participantResult.Name,
		Age:       participantResult.Age,
		CreatedAt: participantResult.CreatedAt,
		UpdatedAt: participantResult.UpdatedAt,
	}
}
