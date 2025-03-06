package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func InitOpenAI(errorMessage string) string {
	Client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	param := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Explain this error message in simple terms: " + errorMessage),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	}

	ctx := context.Background()

	completion, err := Client.Chat.Completions.New(ctx, param)
	if err != nil {
		fmt.Println("Error calling OpenAI:", err)
		return "Failed to fetch explanation from OpenAI."
	}

	return completion.Choices[0].Message.Content
}
