package bots

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type Bot interface {
	Dialogue(context.Context, string) (string, error)
}

type ChatGptBot struct {
	gpt *gogpt.Client
}

func NewChatGptBot(gpt *gogpt.Client) *ChatGptBot {
	return &ChatGptBot{gpt}
}

func (r *ChatGptBot) Dialogue(ctx context.Context, message string) (string, error) {
	req := gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: []gogpt.ChatCompletionMessage{
			{Role: "user", Content: message},
		},
	}

	res, err := r.gpt.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
