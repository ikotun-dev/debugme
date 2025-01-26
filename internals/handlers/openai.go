package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func InitOpenAI() {
	Client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	param := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("What kind of houseplant is easy to take care of?"),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	}

	ctx := context.Background()

	completion, err := Client.Chat.Completions.New(ctx, param)
	if err != nil {
		fmt.Println(err)
	}

	param.Messages.Value = append(param.Messages.Value, completion.Choices[0].Message)
	fmt.Println(completion.Choices[0].Message.Content)
	param.Messages.Value = append(param.Messages.Value, openai.UserMessage("How big are those?"))

	// continue the conversation
	completion, err = Client.Chat.Completions.New(ctx, param)

	/*
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
	*/
}
