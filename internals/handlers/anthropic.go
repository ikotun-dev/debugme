package handlers

import "os"

func InitAnthropic() {
	anthropicAPIkey := os.Getenv("ANTHROPIC_API_KEY")

	if anthropicAPIkey == "" {
		panic("Anthropic API Key is required")
	}
}
