package chat

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
