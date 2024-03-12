// This file contains all the types and methods used for handling the chats
package chat

import "time"

// Method for initialising the chat with the male and female personalities
// The messages will be generated according to the personalities passed
func InitializeChat(malePersonality string, femalePersonality string) ChatThread {
	var maleSystemPrompt string = "You have a personality similar to " + malePersonality + ". Please give replies accordingly."
	var femaleSystemPrompt string = "You have a personality similar to " + femalePersonality + ". Please give replies accordingly."

	var thread ChatThread = CreateChatThread(malePersonality, femalePersonality)

	var maleUserPrompt string = "Generate an interesting message for starting a conversation"

	var maleEntry ChatEntry = CreateChatEntry(malePersonality, maleUserPrompt, maleSystemPrompt)

	thread.Conversation = append(thread.Conversation, maleEntry)

	var femaleUserPrompt string = "Generate an interesting message for starting a conversation"

	var femaleEntry ChatEntry = CreateChatEntry(femalePersonality, femaleUserPrompt, femaleSystemPrompt)

	thread.Conversation = append(thread.Conversation, femaleEntry)

	return thread
}

// Method for generating the reply of a message based on the previous message with a time delay
func AddNextMessage(thread ChatThread, delay int) {
	var chatLength int = len(thread.Conversation)
	var lastChat ChatEntry = thread.Conversation[chatLength-1]

	var lastMessage string = lastChat.Message
	var user_prompt string = "Generate a reply for : " + lastMessage + ". Keep the conversation interesting"
	var sender string
	var newEntry ChatEntry

	if lastChat.Sender == thread.FirstPerson {
		sender = thread.SecondPerson
	} else {
		sender = thread.FirstPerson
	}

	newEntry = CreateChatEntry(
		sender,
		user_prompt,
		lastChat.SystemPrompt,
	)

	var replyDelay int = delay * int(time.Second)

	time.Sleep(time.Duration(replyDelay))

	thread.Conversation = append(thread.Conversation, newEntry)
}
