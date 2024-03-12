// This file contains the types and methods necessary for making API calls to the OpenAI API
// for generating messages and their replies between the two personalities
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

// Type for the message to be passed to the API
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

/// Type for the request to be placed to the API
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// Type for the response which will be recieved by the API
type ChatResponse struct {
	Id                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index        int           `json:"index"`
		Message      []ChatMessage `json:"message"`
		Logprobs     bool          `json:"logprobs"`
		FinishReason string        `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Method to obtain the API secret key from the dotenv file
func getApiKey() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("OPENAI_API_KEY")
}

// Method to make a request to the OpenAI API and return its response in the form of ChatResponse
// user_prompt -> string : This prompt will be based on the messages in the chat
// system_prompt -> string : This prompt will define the personality of the chat members
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
	if err != nil {
		log.Fatalf("Error in parsing response %s", err)
	}

	var result ChatResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		log.Fatalf("Cannot unmarshal JSON due to %s", err)
	}

	return result
}
