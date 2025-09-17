package service

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// ScriptResponse defines the expected LLM JSON structure
type ScriptResponse struct {
	Topic              string   `json:"topic"`
	Keypoints          []string `json:"keypoints"`
	MimicWritingStyle  string   `json:"mimic_writing_style"`
	MoreRequirements   []string `json:"more_requirements"`
	YoutubeTitle       string   `json:"youtube_title"`
	YoutubeDescription string   `json:"youtube_description"`
	InstagramDesc      string   `json:"instagram_description"`
}

// CleanLLMResponse extracts and unmarshals JSON from messy LLM output
func CleanLLMResponse(raw string) (*ScriptResponse, error) {
	// Sometimes model wraps JSON in ```json ... ```
	re := regexp.MustCompile("(?s)```json(.*?)```")
	matches := re.FindStringSubmatch(raw)
	if len(matches) > 1 {
		raw = matches[1]
	}

	// Trim spaces, newlines
	raw = strings.TrimSpace(raw)

	var resp ScriptResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		return nil, errors.New("invalid JSON structure in LLM response")
	}
	return &resp, nil
}
