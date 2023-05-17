package routes

import (
	"github.com/gin-gonic/gin"
	"modelo-rest-go/internal/api"
)

func HealthRoute(s *gin.Engine) {
	health := new(api.HealthController)
	s.GET("/health", health.Status)
}
