package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"video-chat/internal/services/rtc"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Initial message read error:", err)
		conn.Close()
		return
	}

	var joinMsg map[string]interface{}
	if err := json.Unmarshal(msg, &joinMsg); err != nil {
		log.Println("JSON unmarshal error:", err)
		conn.Close()
		return
	}

	clientID := joinMsg["id"].(string)
	client := &rtc.Client{ID: clientID, Conn: conn}
	rtc.AddClient(client)

	defer func() {
		rtc.RemoveClient(clientID)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		var signalMsg map[string]interface{}
		if err := json.Unmarshal(msg, &signalMsg); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}
		rtc.BroadcastMessage(clientID, signalMsg)
	}
}
