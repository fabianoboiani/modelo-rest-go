package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "modelo-rest-go/api"
)

func ToolsRoute(s *gin.Engine) {
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
