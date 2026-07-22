package models

type ResponseAI struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
}

type RequestAI struct {
	Question string `json:"question"`
}

type GeminiConfig struct {
	APIKey string
	APIURL string
}

// --- Нативные структуры Gemini API ---

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiContent struct {
	Role  string       `json:"role,omitempty"` // "user" или "model"
	Parts []GeminiPart `json:"parts"`
}

// Системная инструкция передается отдельным объектом
type GeminiSystemInstruction struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiNativeRequest struct {
	SystemInstruction *GeminiSystemInstruction `json:"system_instruction,omitempty"`
	Contents          []GeminiContent          `json:"contents"`
}

// Ответ от Gemini API
type GeminiNativeResponse struct {
	Candidates []struct {
		Content GeminiContent `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}
