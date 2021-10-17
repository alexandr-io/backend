package data

type CreditCard struct {
	ID       string `json:"id,omitempty"`
	Number   string `json:"number"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	CVC      string `json:"cvc"`
}
