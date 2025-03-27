package main

import "flemming-pr/gray/gray"

func main() {
	sendInt()
	sendString()
	sendStringWithColor()
	sendStruct()
	sendWithAdditional()
	sendMap()
}

func sendInt() {
	gray.Send(42)
}

func sendString() {
	gray.Send("Hello World")
}

func sendStringWithColor() {
	gray.WithColor("green").Send("Hello World")
	gray.WithColor("red").Send("Hello World")
	gray.WithColor("yellow").Send("Hello World")
}

type TestStruct struct {
	Title  string
	Amount int
}

func sendStruct() {
	testStruct := TestStruct{
		Title:  "Hello world",
		Amount: 4223,
	}

	gray.Send(testStruct)
}

func sendWithAdditional() {
	gray.WithAdditionals("id", "6f54a18f-0af5-449d-94ba-fa2b45fc912f").Send("Test")

	testMap := map[string]any{
		"foo": "bar",
		"23":  42,
	}

	gray.WithColor("green").WithAdditionals("foo", testMap).Send("Hello World")
}

func sendWithAdditionalsAndColor() {
	gray.WithColor("green").WithAdditionals("foo", "bar").Send("Hello World")
}

func sendMap() {
	newMap := []any{
		"Hello",
		"World",
		23,
		42,
		23.235356467,
	}

	gray.Send(newMap)
}
