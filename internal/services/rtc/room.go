package rtc

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type Room struct {
	Clients map[string]*Client
	mu      sync.Mutex
}

var room = &Room{
	Clients: make(map[string]*Client),
}

func AddClient(client *Client) {
	room.mu.Lock()
	defer room.mu.Unlock()
	room.Clients[client.ID] = client
	BroadcastJoin(client.ID)
}

func RemoveClient(clientID string) {
	room.mu.Lock()
	defer room.mu.Unlock()
	delete(room.Clients, clientID)
}

func BroadcastJoin(clientID string) {
	for id, c := range room.Clients {
		if id != clientID {
			c.Conn.WriteJSON(map[string]interface{}{
				"type": "join",
				"from": clientID,
			})
		}
	}
}

func BroadcastMessage(from string, msg map[string]interface{}) {
	room.mu.Lock()
	defer room.mu.Unlock()
	toID, _ := msg["to"].(string)
	if toID != "" {
		if c, ok := room.Clients[toID]; ok {
			c.Conn.WriteJSON(msg)
		}
	} else {
		for id, c := range room.Clients {
			if id != from {
				c.Conn.WriteJSON(msg)
			}
		}
	}
}
