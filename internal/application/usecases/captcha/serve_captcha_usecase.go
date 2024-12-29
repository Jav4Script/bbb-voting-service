package captcha

import (
	"net/http"

	"bbb-voting-service/internal/domain/services"
)

type ServeCaptchaUsecase struct {
	CaptchaService services.CaptchaService
}

func NewServeCaptchaUsecase(captchaService services.CaptchaService) *ServeCaptchaUsecase {
	return &ServeCaptchaUsecase{CaptchaService: captchaService}
}

func (usecase *ServeCaptchaUsecase) Execute(reponseWriter http.ResponseWriter, request *http.Request, captchaID string) {
	usecase.CaptchaService.ServeCaptcha(reponseWriter, request, captchaID)
}
