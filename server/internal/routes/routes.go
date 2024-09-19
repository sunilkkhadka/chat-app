package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/chat-app/internal/server"
)

func ConfigureRoutes(server *server.Server) {

	server.Gin.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"message": "running..",
		})
	})

}
