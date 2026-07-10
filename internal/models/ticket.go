package models

type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TicketsResponse struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Tickets     []Ticket `json:"tickets"`
}

type TicketCreateRequest struct {
	Title            string   `json:"title"`
	Priority         string   `json:"priority"`
	Severity         string   `json:"severity"`
	Environment      string   `json:"environment"`
	StepsToReproduce string   `json:"steps_to_reproduce"`
	ExpectedResult   string   `json:"expected_result"`
	ActualResult     string   `json:"actual_result"`
	Attachments      []string `json:"attachments"`
}
