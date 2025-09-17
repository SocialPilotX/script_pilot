package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"script_pilot/internal/constants"
	"script_pilot/internal/service"
)

type RequestBody struct {
	Hint          string                `json:"hint"`
	Style         constants.ScriptStyle `json:"style"`
	Type          constants.ScriptType  `json:"type"`
	TimeInSeconds string                `json:"time_in_seconds"`
}

func GenerateScriptHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	prompt := service.BuildPrompt(body.Style, body.Type, body.Hint, body.TimeInSeconds)
	slog.Info("Prompt", "message", prompt)
	messages := []service.Content{
		{
			Role: "user",
			Parts: []service.Part{
				{Text: prompt},
			},
		},
	}
	chat, err := service.GeminiClient.SendChat(messages)
	if err != nil {
		slog.Error("error sending message", err)
		http.Error(w, `{"error":"failed to get response from LLM"}`, http.StatusInternalServerError)
		return
	}

	// Clean & extract JSON
	cleaned, err := service.CleanLLMResponse(chat)
	if err != nil {
		slog.Error("error cleaning LLM response", err)
		http.Error(w, `{"error":"failed to parse response from LLM"}`, http.StatusInternalServerError)
		return
	}

	// Return structured JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cleaned); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}
