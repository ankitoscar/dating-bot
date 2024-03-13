package main

import (
	"fmt"

	"example.com/chat"
)

func main() {

	fmt.Println("Welcome to a simple dating simulator.\nHere you can choose two personalities, one male and female, and observe how they would chat with each other.")

	var male, female string
	fmt.Print("Enter the male personality: ")
	fmt.Scan(&male)
	fmt.Print("Enter the female personality: ")
	fmt.Scan(&female)

	var thread chat.ChatThread = chat.InitializeChat(male, female)

	for {
		chat.AddNextMessage(thread, 3600)
	}
}
