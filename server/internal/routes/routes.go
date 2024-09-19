package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/sunilkkhadka/chat-app/internal/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Echo the message back to the client
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func ConfigureRoutes(server *server.Server) {

	server.Gin.GET("/", handleWebSocket)

}
