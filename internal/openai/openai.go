package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAiClient struct {
	client *openai.Client
}

func NewClient(token string) *OpenAiClient {
	client := openai.NewClient(token)

	return &OpenAiClient{
		client: client,
	}
}

func (c *OpenAiClient) Complete(instructions, input string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: instructions,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
