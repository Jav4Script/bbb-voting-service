package repositories

import domain "bbb-voting-service/internal/domain/entities"

type ParticipantRepository interface {
	Delete(participant domain.Participant) error
	FindAll() ([]domain.Participant, error)
	FindByID(id string) (domain.Participant, error)
	FindByName(name string) (domain.Participant, error)
	Save(participant domain.Participant) (domain.Participant, error)
}
