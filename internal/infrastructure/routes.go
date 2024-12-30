package infrastructure

import (
	"net/http"

	"bbb-voting-service/internal/infrastructure/controllers"
	"bbb-voting-service/internal/infrastructure/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ConfigureRoutes configures the routes for the application
func ConfigureRoutes(router *gin.Engine, captchaController *controllers.CaptchaController, participantController *controllers.ParticipantController, voteController *controllers.VoteController, resultController *controllers.ResultController) {
	router.Use(middlewares.ErrorHandler())

	// Health check endpoint
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Version 1 routes
	v1 := router.Group("/v1")
	{
		// CAPTCHA endpoints
		v1.GET("/generate-captcha", captchaController.GenerateCaptcha)
		v1.GET("/captcha/:captcha_id", captchaController.ServeCaptcha)
		v1.POST("/validate-captcha", captchaController.ValidateCaptcha)

		// Participant endpoints
		v1.GET("/participants", participantController.GetParticipants)
		v1.GET("/participants/:id", participantController.GetParticipant)
		v1.POST("/participants", participantController.CreateParticipant)
		v1.DELETE("/participants/:id", participantController.DeleteParticipant)

		// Vote endpoints
		v1.POST("/votes", voteController.CastVote)

		// Result endpoints
		v1.GET("/results/partial", resultController.GetPartialResults)
		v1.GET("/results/final", resultController.GetFinalResults)
	}
}
