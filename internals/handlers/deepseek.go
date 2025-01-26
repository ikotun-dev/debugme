package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func InitDeepSeek() {
	deepSeekAPIKey := os.Getenv("DEEPSEEK")
	if deepSeekAPIKey == "" {
		panic("DeepSeek API Key is required")
	}
}

func ChatDeepSeek(w http.ResponseWriter, r *http.Request) {
	InitDeepSeek()
	url := "https://api.deepseek.com/chat/completions"
	method := "POST"

	payload := strings.NewReader(`{
  "messages": [
    {
      "content": "You are a helpful assistant",
      "role": "system"
    },
    {
      "content": "Hi",
      "role": "user"
    }
  ],
  "model": "deepseek-chat", 
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {

		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("DEEPSEEK")))

	res, err := client.Do(req)
	if err != nil {

		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(body))

	json.NewEncoder(w).Encode(string(body))

}
