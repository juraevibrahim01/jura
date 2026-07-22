package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/juraevibrahim01/jura/internal/models"
)

type AI_service struct {
	api_key     string
	api_url     string
	http_client *http.Client
}

func NewAI_service(config models.GeminiConfig) *AI_service {
	url := config.APIURL
	if url == "" {
		// Базовый URL нативного REST API
		url = "https://generativelanguage.googleapis.com"
	}

	return &AI_service{
		api_key: config.APIKey,
		api_url: url,
		http_client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Построение системного промпта с использованием XML-тегов
func (s *AI_service) buildSystemPrompt(userID *int) string {
	userIDStr := "Не указан"
	if userID != nil && *userID != 0 {
		userIDStr = fmt.Sprintf("%d", *userID)
	}

	// Использование XML-тегов позволяет Gemini четко разграничивать блоки правил
	return fmt.Sprintf(`
<system_instructions>
  <role>
    Ты — умный и вежливый AI-ассистент сервиса Jura.
  </role>

  <context>
    <user_id>%s</user_id>
  </context>

  <rules>
    <rule>Отвечай структурированно, информативно и по делу.</rule>
    <rule>Используй доброжелательный тон общения.</rule>
    <rule>Если не знаешь ответа на вопрос — честно признайся в этом.</rule>
  </rules>
</system_instructions>
`, userIDStr)
}

func (s *AI_service) SendMessage(userID *int, message string) (string, error) {
	systemPrompt := s.buildSystemPrompt(userID)

	// Формируем нативный запрос Gemini
	geminiReq := models.GeminiNativeRequest{
		SystemInstruction: &models.GeminiSystemInstruction{
			Parts: []models.GeminiPart{
				{Text: systemPrompt},
			},
		},
		Contents: []models.GeminiContent{
			{
				Role: "user",
				Parts: []models.GeminiPart{
					{Text: message},
				},
			},
		},
	}

	// Используем стабильную модель gemini-1.5-flash (или gemini-2.0-flash)
	modelName := "gemini-3.6-flash"
	response, err := s.callGeminiAPI(modelName, geminiReq)
	if err != nil {
		return "", fmt.Errorf("failed to call Gemini API: %w", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Gemini API error [%d]: %s", response.Error.Code, response.Error.Message)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response generated from Gemini")
	}

	// Возвращаем сгенерированный текст
	return response.Candidates[0].Content.Parts[0].Text, nil
}
func (s *AI_service) callGeminiAPI(model string, req models.GeminiNativeRequest) (*models.GeminiNativeResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Нормализуем базовый URL (убираем возможный слэш на конце)
	baseUrl := strings.TrimRight(s.api_url, "/")

	// Четкая сборка эндпоинта
	endpoint := fmt.Sprintf("%s/v1beta/models/%s:generateContent?key=%s", baseUrl, model, s.api_key)

	httpReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.http_client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var geminiResp models.GeminiNativeResponse
	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &geminiResp, nil
}

func (s *AI_service) HealthCheck() error {
	_, err := s.SendMessage(nil, "ping")
	return err
}
