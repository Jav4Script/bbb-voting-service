package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	entities "bbb-voting-service/internal/domain/entities"

	"github.com/dchest/captcha"
	"github.com/go-redis/redis/v8"
)

type CaptchaService struct {
	redisClient *redis.Client
}

func NewCaptchaService(redisClient *redis.Client) *CaptchaService {
	return &CaptchaService{
		redisClient: redisClient,
	}
}

func (captchaService *CaptchaService) GenerateCaptcha() entities.Captcha {
	captchaID := captcha.New()
	captchaURL := fmt.Sprintf("/captcha/%s", captchaID)

	return entities.Captcha{ID: captchaID, Image: captchaURL}
}

func (captchaService *CaptchaService) ServeCaptcha(w http.ResponseWriter, r *http.Request, captchaID string) {
	captcha.WriteImage(w, captchaID, 240, 80)
}

func (captchaService *CaptchaService) ValidateCaptcha(captchaID, captchaSolution string) (string, bool) {
	if captcha.VerifyString(captchaID, captchaSolution) {
		token := generateCaptchaToken()
		err := captchaService.redisClient.Set(context.Background(), token, "true", 10*time.Minute).Err()
		if err != nil {
			log.Printf("Error storing token in Redis: %v", err)
			return "", false
		}
		log.Printf("Token stored in Redis: %s", token)
		return token, true
	}
	return "", false
}

func (captchaService *CaptchaService) ValidateCaptchaToken(token string) bool {
	val, err := captchaService.redisClient.Get(context.Background(), token).Result()
	if err == redis.Nil || val != "true" {
		log.Printf("Token not found or invalid in Redis: %s", token)
		return false
	} else if err != nil {
		log.Printf("Error retrieving token from Redis: %v", err)
		panic(err) // handle error appropriately in production code
	}

	// Delete the token after validation to prevent reuse
	err = captchaService.redisClient.Del(context.Background(), token).Err()
	if err != nil {
		log.Printf("Error deleting token from Redis: %v", err)
	}
	return true
}

func generateCaptchaToken() string {
	// Generate a secure random token
	tokenBytes := make([]byte, 32)

	if _, err := rand.Read(tokenBytes); err != nil {
		panic(err) // handle error appropriately in production code
	}

	return hex.EncodeToString(tokenBytes)
}
