package controllers

import (
	"log"
	"net/http"

	"bbb-voting-service/internal/application/usecases/votes"
	"bbb-voting-service/internal/domain/dtos"
	"bbb-voting-service/internal/infrastructure/mappers"
	"bbb-voting-service/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

type VoteController struct {
	CastVoteUsecase *votes.CastVoteUsecase
	CaptchaService  *services.CaptchaService
}

func NewVoteController(castVoteUseCase *votes.CastVoteUsecase, captchaService *services.CaptchaService) *VoteController {
	return &VoteController{
		CastVoteUsecase: castVoteUseCase,
		CaptchaService:  captchaService,
	}
}

// CastVote godoc
// @Summary Cast Vote
// @Description Casts a vote for a participant
// @Tags vote
// @Accept  json
// @Produce  json
// @Param vote body dtos.CastVoteDTO true "Vote details"
// @Param X-Captcha-Token header string true "CAPTCHA validation token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 403 {object} map[string]string "Invalid CAPTCHA token"
// @Router /votes [post]
func (controller *VoteController) CastVote(context *gin.Context) {
	var voteDTO dtos.CastVoteDTO
	if err := context.ShouldBindJSON(&voteDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate CAPTCHA token from header
	captchaToken := context.GetHeader("X-Captcha-Token")
	if captchaToken == "" {
		log.Printf("Missing CAPTCHA token")
		context.JSON(http.StatusForbidden, gin.H{"error": "Missing CAPTCHA token"})
		return
	}

	if !controller.CaptchaService.ValidateCaptchaToken(captchaToken) {
		log.Printf("Invalid CAPTCHA token: %s", captchaToken)
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid CAPTCHA token"})
		return
	}

	vote := mappers.FromCastVoteDTO(voteDTO)

	if err := controller.CastVoteUsecase.Execute(vote); err != nil {
		log.Printf("Error casting vote: %v", err)
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}
