package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var allowedOrigins = map[string]bool{
	"http://localhost:8080": true,
}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return allowedOrigins[origin]
	},
}

var mu sync.Mutex
var messageHistory []string

type Response struct {
	Message string   `json:"message"`
	History []string `json:"history,omitempty"`
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		http.Error(w,"Failed to upgrade to WebSocket",http.StatusBadRequest)
		return
	}
	go handleWebSocket(conn)
}

func handleWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		message := string(msg)

		if message == ""{
			log.Println("Received an empty message. Please enter a valid message")
			response := Response{Message: "Empty message received,please send a valid message."}
			responseJSON , _ := json.Marshal(response)
			err = conn.WriteMessage(websocket.TextMessage,responseJSON)
			if err != nil {
				log.Println("Error reading message" , err)
				break
			}
			continue
		}

		if message == "GET_HISTORY" {
			mu.Lock()
			history := messageHistory
			mu.Unlock()
			response := Response{Message: "History", History: history}
			responseJSON, err := json.Marshal(response)
			if err != nil {
				log.Println("Error marshalling message" , err)
				break
			}
			err = conn.WriteMessage(websocket.TextMessage, responseJSON)
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		} else {
			reversed := reverseString(message)

			mu.Lock()
			if len(messageHistory) >= 5 {
				messageHistory = messageHistory[1:]
			}
			messageHistory = append(messageHistory, message)
			mu.Unlock()

			response := Response{Message: reversed}
			responseJSON, err := json.Marshal(response)
			if err != nil {
				log.Println("Error marshalling message" , err)
				break
			}

			err = conn.WriteMessage(websocket.TextMessage, responseJSON)
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}

		}
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("../Client/index.html")
	r.Static("/static", "../Client")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ws", func(ctx *gin.Context) {
		wsHandler(ctx.Writer, ctx.Request)
	})

	log.Println("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
