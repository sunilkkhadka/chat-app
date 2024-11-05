package routes

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/sunilkkhadka/chat-app/internal/server"
)

type RequestPayload struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"message"`
	ChatType   string `json:"chat_type"`
}

type RequestId struct {
	SenderID string `json:"sender_id"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

var clients = make(map[string]*websocket.Conn)
var broadcast = make(chan RequestPayload)
var mutex = &sync.Mutex{}

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

	var senderId RequestId
	err = conn.ReadJSON(&senderId)
	if err != nil {
		log.Printf("error: %+v", err)
		return
	}

	mutex.Lock()
	clients[senderId.SenderID] = conn
	mutex.Unlock()

	fmt.Printf("\nCLIENTS = %+v\n", clients)

	status := StatusResponse{
		Status: "You Are Online",
	}
	err = conn.WriteJSON(status)
	if err != nil {
		log.Fatal("Error occured", err)
	}

	go ListenMessages(conn, senderId.SenderID)

}

func ListenMessages(conn *websocket.Conn, id string) {
	var payload RequestPayload

	defer func() {
		conn.Close()
		mutex.Lock()
		delete(clients, id)
		mutex.Unlock()
	}()

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			mutex.Lock()
			delete(clients, id)
			mutex.Unlock()
			break
		}
		broadcast <- payload
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		switch msg.ChatType {
		case "private":
			if recipientConn, ok := clients[msg.ReceiverID]; ok {
				recipientConn.WriteJSON(msg)
			}
		}
	}
}

func ConfigureRoutes(server *server.Server) {

	server.Gin.GET("/", handleWebSocket)

}
