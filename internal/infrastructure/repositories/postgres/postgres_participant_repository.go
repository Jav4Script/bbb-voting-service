package repositories

import (
	entities "bbb-voting-service/internal/domain/entities"
	models "bbb-voting-service/internal/infrastructure/models"

	"gorm.io/gorm"
)

type PostgresParticipantRepository struct {
	DB *gorm.DB
}

func NewPostgresParticipantRepository(database *gorm.DB) *PostgresParticipantRepository {
	return &PostgresParticipantRepository{DB: database}
}

func (repository *PostgresParticipantRepository) Delete(participant entities.Participant) error {
	participantModel := models.ToModelParticipant(participant)

	err := repository.DB.Delete(&participantModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostgresParticipantRepository) FindAll() ([]entities.Participant, error) {
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

func (repository *PostgresParticipantRepository) FindByID(id string) (entities.Participant, error) {
	var participantModel models.ParticipantModel

	err := repository.DB.First(&participantModel, "id = ?", id).Error
	if err != nil {
		return entities.Participant{}, err
	}

	return models.ToDomainParticipant(participantModel), nil
}

func (repository *PostgresParticipantRepository) FindByName(name string) (entities.Participant, error) {
	var participantModel models.ParticipantModel

	err := repository.DB.First(&participantModel, "name = ?", name).Error
	if err != nil {
		return entities.Participant{}, err
	}

	return models.ToDomainParticipant(participantModel), nil
}

func (repository *PostgresParticipantRepository) Save(participant entities.Participant) (entities.Participant, error) {
	participantModel := models.ToModelParticipant(participant)

	err := repository.DB.Create(&participantModel).Error
	if err != nil {
		return entities.Participant{}, err
	}

	participantEntity := models.ToDomainParticipant(participantModel)

	return participantEntity, nil
}
