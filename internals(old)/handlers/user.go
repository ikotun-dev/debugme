package handlers

import (
	"encoding/json"
	"net/http"
)

var UserTypes = map[string]string{
	"user":                   "user",
	"bot":                    "bot",
	"customer_service_agent": "customer_service_agent",
}

type Message struct {
	Message  string `json:"message"`
	UserType string `json:"user_type"`
}

func ProcessMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(string(msg.Message)) < 1 {
		http.Error(w, "Message is empty", http.StatusBadRequest)
	}

	if _, ok := UserTypes[msg.UserType]; !ok {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
	}
}
