package chat

import "time"

type ChatThread struct {
	FirstPerson  string
	SecondPerson string
	Conversation []ChatEntry
}

type ChatEntry struct {
	Sender  string
	Message string
	Time    time.Time
	Status  string
}

func CreateChatThread(male string, female string) ChatThread {
	var conversation []ChatEntry
	return ChatThread{
		FirstPerson:  male,
		SecondPerson: female,
		Conversation: conversation,
	}
}

func CreateChatMessage(user_prompt string, system_prompt string) string {
	var response ChatResponse = GenrateResponse(user_prompt, system_prompt)

	var replies []string = ParseRespone(response)

	return replies[0]
}

func CreateChatEntry(sender string, user_prompt string, system_prompt string) ChatEntry {
	var message string = CreateChatMessage(user_prompt, system_prompt)

	return ChatEntry{
		Sender:  sender,
		Message: message,
		Time:    time.Now(),
		Status:  "Sent",
	}
}
