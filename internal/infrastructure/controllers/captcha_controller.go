package controllers

import (
	"log"
	"net/http"

	"bbb-voting-service/internal/domain/dtos"
	"bbb-voting-service/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	CaptchaService *services.CaptchaService
}

func NewCaptchaController(captchaService *services.CaptchaService) *CaptchaController {
	return &CaptchaController{CaptchaService: captchaService}
}

// GenerateCaptcha godoc
// @Summary Generate CAPTCHA
// @Description Generates a new CAPTCHA
// @Tags captcha
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "captcha_id and captcha_image"
// @Router /generate-captcha [get]
func (captchaController *CaptchaController) GenerateCaptcha(context *gin.Context) {
	captcha := captchaController.CaptchaService.GenerateCaptcha()
	context.JSON(http.StatusOK, captcha)
}

// ServeCaptcha godoc
// @Summary Serve CAPTCHA
// @Description Serves the CAPTCHA image
// @Tags captcha
// @Accept  json
// @Produce  json
// @Param captcha_id path string true "CAPTCHA ID"
// @Success 200 {object} map[string]interface{} "captcha_id and captcha_image"
// @Failure 404 {object} map[string]string "CAPTCHA not found"
// @Router /captcha/{captcha_id} [get]
func (captchaController *CaptchaController) ServeCaptcha(context *gin.Context) {
	captchaID := context.Param("captcha_id")
	captchaController.CaptchaService.ServeCaptcha(context.Writer, context.Request, captchaID)
}

// ValidateCaptcha godoc
// @Summary Validate CAPTCHA
// @Description Validates the CAPTCHA solution
// @Tags captcha
// @Accept  json
// @Produce  json
// @Param captcha body dtos.ValidateCaptchaDTO true "CAPTCHA ID and solution"
// @Success 200 {object} map[string]string "CAPTCHA validated successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 403 {object} map[string]string "Invalid CAPTCHA"
// @Router /validate-captcha [post]
func (captchaController *CaptchaController) ValidateCaptcha(context *gin.Context) {
	var request dtos.ValidateCaptchaDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, valid := captchaController.CaptchaService.ValidateCaptcha(request.CaptchaID, request.CaptchaSolution)
	if !valid {
		log.Printf("Invalid CAPTCHA for ID: %s", request.CaptchaID)
		context.JSON(http.StatusForbidden, gin.H{"error": "Invalid CAPTCHA"})
		return
	}

	context.Header("X-Captcha-Token", token)
	context.JSON(http.StatusOK, gin.H{"message": "CAPTCHA validated successfully"})
}
