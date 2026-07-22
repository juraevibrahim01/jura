package models

type TestKey struct {
	ID             int    `json:"id"`
	Date           string `json:"date"`
	Name           string `json:"name"`
	Module         string `json:"module"`
	Precondition   string `json:"precondition"`
	Steps          string `json:"steps"`
	ExpectationRes string `json:"expectation_res"`
	ActualRes      string `json:"actual_res"`
	Comment        string `json:"comment"`
}

type TestKeysResponse struct {
	Status      string    `json:"status"`
	Description string    `json:"description"`
	TestKeys    []TestKey `json:"test_keys"`
}

type TestKeyResponse struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	TestKey     *TestKey `json:"test_key,omitempty"`
}

type TestKeyCreateRequest struct {
	Date           string `json:"date"`
	Name           string `json:"name"`
	Module         string `json:"module"`
	Precondition   string `json:"precondition"`
	Steps          string `json:"steps"`
	ExpectationRes string `json:"expectation_res"`
	ActualRes      string `json:"actual_res"`
	Comment        string `json:"comment"`
}
