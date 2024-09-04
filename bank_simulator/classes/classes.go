package classes

// PaymentRequest represents the JSON structure for processing a payment
type PaymentRequest struct {
	CustomerID int     `json:"customer_id"`
	MerchantID int     `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
}

// PaymentResponse represents the JSON structure for the response from the bank simulator
type PaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BankSimulatorResponse represents the response structure from the bank simulator
type BankSimulatorResponse struct {
	Approved bool `json:"approved"`
}
