package models

type Ticket struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Priority       string `json:"priority"`
	Severity       string `json:"severity"`
	Environment    string `json:"environment"`
	Steps          string `json:"steps"`
	ExpectedResult string `json:"expected_result"`
	ActualResult   string `json:"actual_result"`
	Attachments    string `json:"attachments"`
	CreatedAt      string `json:"created_at"`
}

type TicketsResponse struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Tickets     []Ticket `json:"tickets"`
}

type TicketCreateRequest struct {
	Title          string `json:"title"`
	Priority       string `json:"priority"`
	Severity       string `json:"severity"`
	Environment    string `json:"	environment"`
	Steps          string `json:"steps"`
	ExpectedResult string `json:"expected_result"`
	ActualResult   string `json:"actual_result"`
	Attachments    string `json:"attachments"`
	CreatedAt      string `json:"created_at"`
}
