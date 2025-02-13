package main

import "flemming-pr/gray/gray"

func main() {
	sendInt()
	sendString()
	sendStringWithColor()
	sendStruct()
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

func sendMap() {
	newMap := []interface{}{
		"Hello",
		"World",
		23,
		42,
		23.235356467,
	}

	gray.Send(newMap)
}
