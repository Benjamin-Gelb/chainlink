package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIChatCompletion struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

var logger log.Logger = *log.Default()

func main() {

	// client := &http.Client{}

	// openaiApi := "https://api.openai.com/v1/chat/completions"
	token := ""
	chat := NewChatOpenAI(token)
	messages := []Message{
		{
			Role:    "system",
			Content: "You are a character in Breaking Bad having a conversation with Jesse and Walter White",
		},
		{
			Role:    "user",
			Content: "Say my name...",
		},
	}
	output, err := chat.Input(messages)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*output)
	}
}

type Choice struct {
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
	LogProbs     interface{} `json:"logprobs"`
	Message      Message     `json:"message"`
}

type ChatCompletion struct {
	Id                string         `json:"id"`
	Object            string         `json:"object"`
	Created           int            `json:"created"`
	Model             string         `json:"model"`
	Choices           []Choice       `json:"choices"`
	Usage             map[string]int `json:"usage"`
	SystemFingerprint string         `json:"system_fingerprint"`
}

type LLMResponse struct {
	Response []Message // The LLM's Response, what you want the next Link in your Chain to receive as an input.
	// Miscilaneous data associated with the LLM response.
}

type Model interface {
	GetEndpoint() string
	GetApiKey() string
	Input([]Message) (*LLMResponse, error)
}

// type OpenAI struct {
// 	OpenAIApiKey string
// 	Endpoint     string
// 	Model        string `json:"model"`
// 	Client       *http.Client
// }

type ChatOpenAI struct {
	OpenAIApiKey string
	Endpoint     string
	Model        string
	Client       *http.Client
}

func (m ChatOpenAI) GetApiKey() string {
	return m.OpenAIApiKey
}

func (m ChatOpenAI) GetEndpoint() string {
	return m.Endpoint
}

func MakeRequest(m Model, body []byte) (*http.Request, error) {
	requestBody := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", m.GetEndpoint(), requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.GetApiKey())
	return req, nil
}

func (m *ChatOpenAI) Input(messages []Message) (*LLMResponse, error) {
	openAIChatCompletion := OpenAIChatCompletion{
		Model:    m.Model,
		Messages: messages,
	}

	body, err := json.Marshal(openAIChatCompletion)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	req, err := MakeRequest(m, body)
	if err != nil {
		return nil, err
	}
	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, errors.New(string(respBody))
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	chatCompletion := &ChatCompletion{}

	err = json.Unmarshal(respBody, chatCompletion)
	if err != nil {
		return nil, err
	}
	var outputs []Message
	for _, choice := range chatCompletion.Choices {
		outputs = append(outputs, choice.Message)
	}
	return &LLMResponse{
		Response: outputs,
	}, nil
}

func NewChatOpenAI(token string) ChatOpenAI {
	client := &http.Client{}
	return ChatOpenAI{
		OpenAIApiKey: token,
		Endpoint:     "https://api.openai.com/v1/chat/completions",
		Model:        "gpt-4-turbo",
		Client:       client,
	}
}

type Link struct {
}
