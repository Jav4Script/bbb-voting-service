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

func ToParticipantMap(participant entities.Participant) map[string]interface{} {
	return map[string]interface{}{
		"id":         participant.ID,
		"name":       participant.Name,
		"age":        participant.Age,
		"gender":     participant.Gender,
		"created_at": participant.CreatedAt,
		"updated_at": participant.UpdatedAt,
	}
}

func ToParticipantMaps(participants []entities.Participant) []map[string]interface{} {
	participantMaps := make([]map[string]interface{}, len(participants))
	for i, participant := range participants {
		participantMaps[i] = ToParticipantMap(participant)
	}
	return participantMaps
}
