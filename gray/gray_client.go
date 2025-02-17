package gray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Payload     interface{}            `json:"payload"`
	Metadata    map[string]interface{} `json:"metadata"`
	Additionals map[string]interface{} `json:"additionals"`
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

	// Check if we are in a docker container
	url := "http://localhost:23517"
	if _, err := os.Stat("/.dockerenv"); err == nil {
		url = "http://host.docker.internal:23517"
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println("Gray: could not send message", err)
	}
}

func (m Message) WithColor(color string) Message {
	if m.Metadata == nil {
		m.Metadata = map[string]interface{}{}
	}

	m.Metadata["color"] = color

	return m
}

func (m Message) WithAdditionals(key string, value interface{}) Message {
	if m.Additionals == nil {
		m.Additionals = map[string]interface{}{}
	}

	m.Additionals[key] = value

	return m
}

func WithColor(color string) Message {
	message := Message{
		Metadata: map[string]interface{}{
			"color": color,
		},
	}

	return message
}

func WithAdditionals(key string, value interface{}) Message {
	message := Message{
		Additionals: map[string]interface{}{
			key: value,
		},
	}

	return message
}

func Send(payload interface{}) {
	message := Message{}

	message.Send(payload)
}
