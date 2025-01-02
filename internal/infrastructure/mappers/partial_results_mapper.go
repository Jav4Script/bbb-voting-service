package mappers

import (
	entities "bbb-voting-service/internal/domain/entities"
	"strconv"
)

func ToPartialResult(id, name, ageStr, gender, votesStr string) entities.PartialResult {
	age, _ := strconv.Atoi(ageStr)
	votes, _ := strconv.Atoi(votesStr)

	return entities.PartialResult{
		ID:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
		Votes:  votes,
	}
}

func FromPartialResult(partialResult entities.PartialResult) map[string]interface{} {
	return map[string]interface{}{
		"id":     partialResult.ID,
		"name":   partialResult.Name,
		"age":    strconv.Itoa(partialResult.Age),
		"gender": partialResult.Gender,
		"votes":  strconv.Itoa(partialResult.Votes),
	}
}

func ToPartialResultFromParticipant(id string, participant entities.Participant, votesStr string) entities.PartialResult {
	return ToPartialResult(
		id,
		participant.Name,
		strconv.Itoa(participant.Age),
		participant.Gender,
		votesStr,
	)
}

func ToRedisData(result entities.ParticipantResult) (string, map[string]interface{}, string, int) {
	participantKey := "participant:" + result.ID

	participantData := FromPartialResult(entities.PartialResult{
		ID:     result.ID,
		Name:   result.Name,
		Age:    result.Age,
		Gender: result.Gender,
		Votes:  result.Votes,
	})

	return participantKey, participantData, result.ID, result.Votes
}
