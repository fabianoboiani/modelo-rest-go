package server

import (
	"github.com/gin-gonic/gin"
	"modelo-rest-go/internal/api/routes"
	"modelo-rest-go/internal/configs"
)

func Init() {
	conf := configs.GetConfig()
	s := gin.New()
	s.Use(gin.Logger())
	s.Use(gin.Recovery())

	//add routes here
	routes.ToolsRoute(s)
	routes.HealthRoute(s)
	routes.BookRoute(s)

	s.Run(conf.GetString("server.port"))
}
