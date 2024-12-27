package repositories

import (
	entities "bbb-voting-service/internal/domain/entities"
	models "bbb-voting-service/internal/infrastructure/models"

	"gorm.io/gorm"
)

type ParticipantRepository struct {
	DB *gorm.DB
}

func NewParticipantRepository(database *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{DB: database}
}

func (repository *ParticipantRepository) Delete(participant entities.Participant) error {
	participantModel := models.ToModelParticipant(participant)

	err := repository.DB.Delete(&participantModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *ParticipantRepository) FindAll() ([]entities.Participant, error) {
	var participantModels []models.ParticipantModel

	err := repository.DB.Find(&participantModels).Error
	if err != nil {
		return nil, err
	}

	participants := make([]entities.Participant, len(participantModels))
	for i, participantModel := range participantModels {
		participants[i] = models.ToDomainParticipant(participantModel)
	}

	return participants, nil
}

func (repository *ParticipantRepository) FindByID(id string) (entities.Participant, error) {
	var participantModel models.ParticipantModel

	err := repository.DB.First(&participantModel, "id = ?", id).Error
	if err != nil {
		return entities.Participant{}, err
	}

	return models.ToDomainParticipant(participantModel), nil
}

func (repository *ParticipantRepository) FindByName(name string) (entities.Participant, error) {
	var participantModel models.ParticipantModel

	err := repository.DB.First(&participantModel, "name = ?", name).Error
	if err != nil {
		return entities.Participant{}, err
	}

	return models.ToDomainParticipant(participantModel), nil
}

func (repository *ParticipantRepository) Save(participant entities.Participant) (entities.Participant, error) {
	participantModel := models.ToModelParticipant(participant)

	err := repository.DB.Create(&participantModel).Error
	if err != nil {
		return entities.Participant{}, err
	}

	// Map the generated ID back to the participant entity
	participant.ID = participantModel.ID.String()

	return participant, nil
}
