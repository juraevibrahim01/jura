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
