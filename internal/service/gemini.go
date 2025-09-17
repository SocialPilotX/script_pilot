package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://generativelanguage.googleapis.com/v1beta"
)

type GeminiClientConfig struct {
	APIKey string
	Model  string
}

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Role  string `json:"role,omitempty"`
	Parts []Part `json:"parts"`
}

type GenerateContentRequest struct {
	Contents []Content `json:"contents"`
}

type PartResponse struct {
	Text string `json:"text"`
}

type ContentResponse struct {
	Parts []PartResponse `json:"parts"`
}

type GenerateContentResponse struct {
	Candidates []struct {
		Content ContentResponse `json:"content"`
	} `json:"candidates"`
}

var GeminiClient *GeminiClientConfig

// NewGeminiClient creates a new Gemini client
func NewGeminiClient(apiKey, model string) {
	GeminiClient = &GeminiClientConfig{
		APIKey: apiKey,
		Model:  model,
	}
}

// SendChat sends a list of messages to Gemini and returns the AI response
func (g *GeminiClientConfig) SendChat(messages []Content) (string, error) {
	url := fmt.Sprintf("%s/models/%s:generateContent?key=%s", baseURL, g.Model, g.APIKey)

	reqBody := GenerateContentRequest{
		Contents: messages,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		raw, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error: %s", string(raw))
	}

	var result GenerateContentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no response from Gemini")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
