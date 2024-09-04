package classes

// ProcessPaymentRequest represents the JSON structure for processing a payment
type ProcessPaymentRequest struct {
	CustomerID int     `json:"customer_id"`
	MerchantID int     `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
}

// PaymentResponse represents the JSON structure for the response from the bank simulator
type PaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// PaymentDetails represents the JSON structure for payment details
type PaymentDetails struct {
	PaymentID  int     `json:"payment_id"`
	CustomerID int     `json:"customer_id"`
	MerchantID int     `json:"merchant_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
}

// RefundRequest represents the JSON structure for requesting a refund
type RefundRequest struct {
	PaymentID int `json:"payment_id"`
}

// RefundResponse represents the JSON structure for the refund response
type RefundResponse struct {
	Message string `json:"message"`
}

// BankSimulatorResponse represents the response structure from the bank simulator
type BankSimulatorResponse struct {
	Approved bool `json:"approved"`
}
