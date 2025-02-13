package main

import (
	"encoding/json"
	"flemming-pr/gray/gray"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var defaultStyle = lipgloss.NewStyle().
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(4).
	PaddingRight(4)

var headerStyle = lipgloss.NewStyle()

var welcomeStyle = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Foreground(lipgloss.Color("#CCCCCC")).
	Background(lipgloss.Color("#dd21f2")).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(4).
	PaddingRight(4).
	MarginTop(1).
	MarginBottom(1)

func main() {
	fmt.Println(welcomeStyle.Render("starting gray!"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := gray.Message{}

		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			fmt.Println("Can't decode message", err)
			return
		}

		renderMessage(message)
	})

	log.Fatal(http.ListenAndServe(":23517", nil))
}

func renderMessage(message gray.Message) {
	payloadJson, _ := json.MarshalIndent(message.Payload, "", "  ")

	color, ok := message.Metadata["color"]
	if !ok {
		color = "#FAFAFA"
	}

	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(getColorHex(color.(string)))))

	table.Headers(
		headerStyle.Render(time.Now().Format(time.TimeOnly)),
		headerStyle.Render(message.Metadata["name"].(string)),
	)

	table.Row("", defaultStyle.Render(string(payloadJson)))

	fmt.Println(table)
}

func getColorHex(color string) string {
	switch color {
	case "green":
		return "#50FA7B"
	case "red":
		return "#FF5555"
	case "blue":
		return "#8BE9FD"
	case "yellow":
		return "#F1FA8C"
	case "purple":
		return "#BD93F9"
	case "cyan":
		return "#8BE9FD"
	default:
		return "#FAFAFA"
	}
}
