package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
  "io"
)

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type LLMRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type LLMResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GenerateAnswerWithLLM(context, question string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"

  CurrentProfile, err := LoadAIProfile("profile.yaml")

  prompt := fmt.Sprintf(`
  Kamu adalah %s, AI yang menjawab dengan gaya %s dan nada %s.
  Jawabanmu harus dalam bahasa %s, dan gunakan gaya %s sesuai kepribadian saya.
  
  Jawablah berdasarkan pengetahuan mu. Jika tidak relevan, katakan tidak tahu.\n\n
  Berikut adalah pengetahuan yang kamu miliki:\n%s\n\n
  Pertanyaan user:\n%s
  `, 
    CurrentProfile.Name,
    CurrentProfile.Style,
    CurrentProfile.Tone,
    CurrentProfile.Language,
    CurrentProfile.Style,
    context, 
    question,
  )

	data := LLMRequest{
		Model: "meta-llama/llama-4-maverick:free",
    Messages: []Message{
        {
            Role: "system",
            Content: []Content{{Type: "text", Text: prompt}},
        },
        {
            Role: "user",
            Content: []Content{
                { Type: "text", Text: question },
            },
        },
    },
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LLM_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("LLM API error: %s", respBody)
	}

	var res LLMResponse
	if err := json.Unmarshal(respBody, &res); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(res.Choices) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

  fmt.Println("[RELEVANT KNOWLEDGE]:", context)
	return res.Choices[0].Message.Content, nil
}
