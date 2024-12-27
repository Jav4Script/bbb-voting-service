package services

import (
	"fmt"
	"net/http"

	entities "bbb-voting-service/internal/domain/entities"

	"github.com/dchest/captcha"
)

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

func (cs *CaptchaService) GenerateCaptcha() entities.Captcha {
	captchaID := captcha.New()
	captchaURL := fmt.Sprintf("/captcha/%s", captchaID)

	return entities.Captcha{ID: captchaID, Image: captchaURL}
}

func (cs *CaptchaService) ServeCaptcha(w http.ResponseWriter, r *http.Request, captchaID string) {
	captcha.WriteImage(w, captchaID, 240, 80)
}

func (cs *CaptchaService) ValidateCaptcha(captchaID, captchaSolution string) bool {
	return captcha.VerifyString(captchaID, captchaSolution)
}
