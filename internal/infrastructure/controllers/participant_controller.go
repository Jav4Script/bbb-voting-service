package controllers

import (
	"log"
	"net/http"

	usecase "bbb-voting-service/internal/application/usecases"
	dtos "bbb-voting-service/internal/domain/dtos"
	mappers "bbb-voting-service/internal/infrastructure/mappers"

	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	GetParticipantsUsecase   *usecase.GetParticipantsUsecase
	GetParticipantUsecase    *usecase.GetParticipantUsecase
	CreateParticipantUsecase *usecase.CreateParticipantUsecase
	DeleteParticipantUsecase *usecase.DeleteParticipantUsecase
}

func NewParticipantController(participantsUsecase *usecase.GetParticipantsUsecase, participantUsecase *usecase.GetParticipantUsecase, createParticipantUsecase *usecase.CreateParticipantUsecase, deleteParticipantUsecase *usecase.DeleteParticipantUsecase) *ParticipantController {
	return &ParticipantController{
		GetParticipantsUsecase:   participantsUsecase,
		GetParticipantUsecase:    participantUsecase,
		CreateParticipantUsecase: createParticipantUsecase,
		DeleteParticipantUsecase: deleteParticipantUsecase,
	}
}

// GetParticipants godoc
// @Summary Get Participants
// @Description Retrieves all participants
// @Tags participants
// @Accept  json
// @Produce  json
// @Success 200 {object} []map[string]interface{}
// @Router /participants [get]
func (controller *ParticipantController) GetParticipants(context *gin.Context) {
	participants, err := controller.GetParticipantsUsecase.Execute()
	if err != nil {
		log.Printf("Error retrieving participants: %v", err)
		context.Error(err)
		return
	}

	participantMaps := mappers.ToParticipantMaps(participants)
	context.JSON(http.StatusOK, participantMaps)
}

// GetParticipant godoc
// @Summary Get Participant
// @Description Retrieves a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 200 {object} map[string]interface{}
// @Router /participants/{id} [get]
func (controller *ParticipantController) GetParticipant(context *gin.Context) {
	id := context.Param("id")
	participant, err := controller.GetParticipantUsecase.Execute(id)
	if err != nil {
		log.Printf("Error retrieving participant: %v", err)
		context.Error(err)
		return
	}

	participantMap := mappers.ToParticipantMap(participant)
	context.JSON(http.StatusOK, participantMap)
}

// CreateParticipant godoc
// @Summary Create Participant
// @Description Creates a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param participant body dtos.CreateParticipantDTO true "Participant"
// @Success 201 {object} map[string]interface{}
// @Router /participants [post]
func (controller *ParticipantController) CreateParticipant(context *gin.Context) {
	var dto dtos.CreateParticipantDTO
	if err := context.ShouldBindJSON(&dto); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	participantEntity := mappers.FromCreateParticipantDTO(dto)

	participant, err := controller.CreateParticipantUsecase.Execute(participantEntity)
	if err != nil {
		log.Printf("Error creating participant: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create participant"})
		return
	}

	context.JSON(http.StatusCreated, participant)
}

// DeleteParticipant godoc
// @Summary Delete Participant
// @Description Deletes a participant
// @Tags participants
// @Accept  json
// @Produce  json
// @Param id path string true "Participant ID"
// @Success 204
// @Router /participants/{id} [delete]
func (controller *ParticipantController) DeleteParticipant(context *gin.Context) {
	id := context.Param("id")
	err := controller.DeleteParticipantUsecase.Execute(id)
	if err != nil {
		log.Printf("Error deleting participant: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete participant"})
		return
	}

	context.Status(http.StatusNoContent)
}
