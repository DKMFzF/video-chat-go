package rtc

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

type SignalMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func HandleWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if peerConnection != nil {
		log.Println("PeerConnection init error:", err)
		return
	}

	peerConnection.OnTrack(func(tr *webrtc.TrackRemote, r *webrtc.RTPReceiver) {
		log.Printf("Track received: %s", tr.Kind().String())
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocker read error:", err)
			break
		}

		var signal SignalMessage
		if err := json.Unmarshal(msg, &signal); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}

		switch signal.Type {
		case "offer":
			var sdp webrtc.SessionDescription
			_ = json.Unmarshal(signal.Data, &sdp)
			_ = peerConnection.SetRemoteDescription(sdp)

			answer, _ := peerConnection.CreateAnswer(nil)
			_ = peerConnection.SetLocalDescription(answer)

			answerJSON, _ := json.Marshal(answer)
			conn.WriteJSON(SignalMessage{Type: "answer", Data: answerJSON})

		case "candidate":
			var candidate webrtc.ICECandidateInit
			_ = json.Unmarshal(signal.Data, &candidate)
			_ = peerConnection.AddICECandidate(candidate)
		}
	}
}
