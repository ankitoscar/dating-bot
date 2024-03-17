package chat

// Method to parse the response recieved from the OpenAI API
func ParseRespone(response ChatResponse) []string {
	var replies []string

	var reply string = response.Choices[0].Message.Content

	replies = append(replies, reply)

	return replies
}
