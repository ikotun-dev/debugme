package handlers

import (
	"context"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Message struct {
	Message  string `json:"message"`
	UserType string `json:"user_type"`
}

func InitOpenAI() {
	Client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	chatCompletion, err := Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say this is a test"),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}
