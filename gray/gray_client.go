package gray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Payload  interface{}            `json:"payload"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (m Message) Send(payload interface{}) {
	m.Payload = payload
	if m.Metadata == nil {
		m.Metadata = map[string]interface{}{}
	}

	m.Metadata["name"] = fmt.Sprintf("%T", m.Payload)

	// Send to server
	body, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Gray: could not marshal message", err)
		return
	}

	http.Post("http://localhost:23517", "application/json", bytes.NewReader(body))
}

func WithColor(color string) Message {
	message := Message{
		Metadata: map[string]interface{}{
			"color": color,
		},
	}

	return message
}

func Send(payload interface{}) {
	message := Message{}

	message.Send(payload)
}
