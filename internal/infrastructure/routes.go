package infrastructure

import (
	"net/http"

	"bbb-voting-service/internal/infrastructure/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ConfigureRoutes configures the routes for the application
func ConfigureRoutes(router *gin.Engine, captchaController *controllers.CaptchaController, participantController *controllers.ParticipantController, voteController *controllers.VoteController, resultController *controllers.ResultController) {
	// Health check endpoint
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CAPTCHA endpoints
	router.GET("/generate-captcha", captchaController.GenerateCaptcha)
	router.GET("/captcha/:captcha_id", captchaController.ServeCaptcha)
	router.POST("/validate-captcha", captchaController.ValidateCaptcha)

	// Participant endpoints
	router.GET("/participants", participantController.GetParticipants)
	router.GET("/participants/:id", participantController.GetParticipant)
	router.POST("/participants", participantController.CreateParticipant)
	router.DELETE("/participants/:id", participantController.DeleteParticipant)

	// Vote endpoints
	router.POST("/votes", voteController.CastVote)

	// Result endpoints
	router.GET("/results/partial", resultController.GetPartialResults)
	router.GET("/results/final", resultController.GetFinalResults)

}
