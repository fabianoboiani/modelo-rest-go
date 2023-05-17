package main

import (
	"flag"
	"fmt"
	"modelo-rest-go/internal/configs"
	"modelo-rest-go/internal/db"
	"modelo-rest-go/internal/server"
	"os"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	configs.Init(*environment)
	db.ConnectDB()
	server.Init()
}
