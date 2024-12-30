package middlewares

import (
	"log"
	"net/http"

	"bbb-voting-service/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			switch e := err.(type) {
			case *errors.BusinessError:
				log.Printf("Business error: %s", e.Message)
				c.JSON(e.StatusCode, gin.H{"error": e.Message})
			case *errors.InfrastructureError:
				log.Printf("Infrastructure error: %s", e.Message)
				c.JSON(e.StatusCode, gin.H{"error": e.Message})
			default:
				log.Printf("Unexpected error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error"})
			}
		}
	}
}
