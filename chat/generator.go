package chat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Id                string `json:"id"`
	Object            string `json:object`
	Created           int    `json:created`
	Model             string `json:model`
	SystemFingerprint string `json:system_fingerprint`
	Choices           []struct {
		Index        int           `json:"index"`
		Message      []ChatMessage `json:"message"`
		logprobs     bool          `json:"logprobs"`
		FinishReason string        `json:finish_reason`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:prompt_tokens`
		CompletionTokens int `json:completion_tokens`
		TotalTokens      int `json:total_tokens`
	} `json:"usage"`
}

func getApiKey() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("OPENAI_API_KEY")
}

func GenrateResponse(user_prompt string, system_prompt string) ChatResponse {
	var api_key string = getApiKey()

	client := http.Client{}

	var auth_token string = "Bearer " + api_key

	var requestBody = ChatRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{
				"system",
				system_prompt,
			},
			{
				"user",
				user_prompt,
			},
		},
	}

	marshalled, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error in marshalling teacher: %s", err)
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(marshalled))
	if err != nil {
		log.Fatalf("Error in making request %s", err)
	}

	request.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {auth_token},
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error in making request %s", err)
	}

	log.Printf("Status Code: %d", response.StatusCode)

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	var result ChatResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Fatalf("Cannot unmarshal JSON due to %s", err)
	}

	return result
}
