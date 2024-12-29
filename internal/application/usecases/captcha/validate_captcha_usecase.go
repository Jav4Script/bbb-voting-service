package captcha

import (
	"bbb-voting-service/internal/domain/services"
)

type ValidateCaptchaUsecase struct {
	CaptchaService services.CaptchaService
}

func NewValidateCaptchaUsecase(captchaService services.CaptchaService) *ValidateCaptchaUsecase {
	return &ValidateCaptchaUsecase{CaptchaService: captchaService}
}

func (usecase *ValidateCaptchaUsecase) Execute(captchaID, captchaSolution string) (string, bool) {
	return usecase.CaptchaService.ValidateCaptcha(captchaID, captchaSolution)
}
