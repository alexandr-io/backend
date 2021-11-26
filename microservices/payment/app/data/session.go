package data

// Session of a user to pay
type Session struct {
	SuccessURL string `json:"success_url"`
	CancelURL  string `json:"cancel_url"`
	PriceID    string `json:"price_id"`
	CustomerID string `json:"customer_id"`
}
