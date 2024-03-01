package chat

func ParseRespone(response ChatResponse) []string {
	var replies []string

	for _, choice := range response.Choices {
		for _, message := range choice.Message {
			replies = append(replies, message.Content)
		}
	}

	return replies
}