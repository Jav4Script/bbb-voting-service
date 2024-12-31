package repositories

import domain "bbb-voting-service/internal/domain/entities"

type InMemoryParticipantRepository interface {
	Save(participant domain.Participant) error
	FindAll() ([]domain.Participant, error)
	FindByID(id string) (domain.Participant, error)
	Delete(id string) error
}
