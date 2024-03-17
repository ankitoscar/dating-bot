package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/chat"
)

func main() {

	fmt.Println("Welcome to a simple dating simulator.\nHere you can choose two personalities, one male and female, and observe how they would chat with each other.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter male personality")
	male, _ := reader.ReadString('\n')
	male = strings.TrimSpace(male)

	fmt.Println("Enter female personlity")
	female, _ := reader.ReadString('\n')
	female = strings.TrimSpace(female)

	var thread chat.ChatThread = chat.InitializeChat(male, female)

	for {
		chat.AddNextMessage(&thread, 5)
		time.Sleep(25 * time.Second)
	}
}
