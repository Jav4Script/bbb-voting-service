package services

import (
	"net/http"

	entities "bbb-voting-service/internal/domain/entities"
)

type CaptchaService interface {
	GenerateCaptcha() entities.Captcha
	ServeCaptcha(w http.ResponseWriter, r *http.Request, captchaID string)
	ValidateCaptcha(captchaID, captchaSolution string) (string, bool)
	ValidateCaptchaToken(token string) bool
}
