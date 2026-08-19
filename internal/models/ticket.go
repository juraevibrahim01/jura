package models

type Ticket struct {
	ID              int    `json:"id"`
	Data            string `json:"data"`
	Name            string `json:"name"`
	Module          string `json:"module"`
	Precondition    string `json:"precondition"`
	Steps           string `json:"Steps"`
	Expectation_res string `json:"expectation_res"`
	Actual_res      string `json:"actual_res"`
	Comment         string `json:"comment"`
}

type TicketsResponse struct {
	Status      string   `json:"status"`
	Description string   `json:"description"`
	Tickets     []Ticket `json:"tickets"`
}

type TicketCreateRequest struct {
	Data            string `json:"data"`
	Name            string `json:"name"`
	Module          string `json:"module"`
	Precondition    string `json:"precondition"`
	Steps           string `json:"Steps"`
	Expectation_res string `json:"expectation_res"`
	Actual_res      string `json:"actual_res"`
	Comment         string `json:"comment"`
}
