package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/sunilkkhadka/chat-app/internal/config"
	"github.com/sunilkkhadka/chat-app/internal/routes"
	"github.com/sunilkkhadka/chat-app/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Start(config.NewConfig())
}

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	app.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.ConfigureRoutes(app)

	fmt.Println("PORT = ", cfg.HTTP.Port)

	err := app.Run(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port is already in use")
	}
}
