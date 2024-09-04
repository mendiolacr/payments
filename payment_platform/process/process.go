package process

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"payment_platform/classes"
	"payment_platform/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	simulatorURL       = "http://bank_simulator:8081/simulate-payment"
	refundSimulatorURL = "http://bank_simulator:8081/simulate-refund"
)

// SimulatePayment simulates a payment request to the bank simulator
func SimulatePayment(req classes.ProcessPaymentRequest) (*classes.PaymentResponse, error) {

	// Create JSON request body
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Make the HTTP POST request to the bank simulator
	resp, err := http.Post(simulatorURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response from the bank simulator
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paymentResp classes.PaymentResponse
	err = json.Unmarshal(body, &paymentResp)
	if err != nil {
		return nil, err
	}

	return &paymentResp, nil
}

// ProcessPayment handles the payment processing fmtic
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req classes.ProcessPaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the bank simulator to check if the payment can be processed
	paymentResp, err := SimulatePayment(req)
	if err != nil {
		http.Error(w, "Failed to simulate payment", http.StatusInternalServerError)
		return
	}

	// Set the payment status based on the response from the bank simulator
	status := "failed"
	if paymentResp.Success {
		status = "approved"
	}

	// Begin transaction
	tx, err := utils.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Insert payment record
	paymentQuery := `INSERT INTO Payments (customer_id, merchant_id, amount, currency, status, created_at)
                     OUTPUT INSERTED.payment_id
                     VALUES (@customer_id, @merchant_id, @amount, @currency, @status, @created_at)`

	result := tx.QueryRow(paymentQuery,
		sql.Named("customer_id", req.CustomerID),
		sql.Named("merchant_id", req.MerchantID),
		sql.Named("amount", req.Amount),
		sql.Named("currency", req.Currency),
		sql.Named("status", status),
		sql.Named("created_at", time.Now()),
	)

	var paymentID int
	err = result.Scan(&paymentID)
	if err != nil {
		http.Error(w, "Failed to insert payment", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Insert transaction record
	transactionQuery := `INSERT INTO Transactions (payment_id, transaction_type, amount, status, created_at)
                         VALUES (@payment_id, 'Payment', @amount, @status, @created_at)`

	_, err = tx.Exec(transactionQuery,
		sql.Named("payment_id", paymentID),
		sql.Named("amount", req.Amount),
		sql.Named("status", status),
		sql.Named("created_at", time.Now()),
	)
	if err != nil {
		http.Error(w, "Failed to insert transaction", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"payment_id": paymentID,
		"status":     status,
	}
	json.NewEncoder(w).Encode(response)
}

// GetPaymentsByMerchant retrieves payment details for a specific merchant
func GetPaymentsByMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	if merchantID == "" {
		http.Error(w, "merchant_id is required", http.StatusBadRequest)
		return
	}

	query := `SELECT payment_id, customer_id, merchant_id, amount, currency, status, created_at
              FROM Payments
              WHERE merchant_id = @merchant_id`

	rows, err := utils.DB.Query(query, sql.Named("merchant_id", merchantID))
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}
	defer rows.Close()

	var payments []classes.PaymentDetails
	for rows.Next() {
		var payment classes.PaymentDetails
		err := rows.Scan(&payment.PaymentID, &payment.CustomerID, &payment.MerchantID, &payment.Amount, &payment.Currency, &payment.Status, &payment.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to read query result", http.StatusInternalServerError)
			fmt.Printf("Error: %v", err)
			return
		}
		payments = append(payments, payment)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error occurred during row iteration", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Respond to the client with the list of payments
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func ProcessRefund(w http.ResponseWriter, r *http.Request) {
	var req classes.RefundRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Begin transaction
	tx, err := utils.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Check if the payment exists and retrieve its details
	paymentQuery := `SELECT amount, status FROM Payments WHERE payment_id = @payment_id`
	row := tx.QueryRow(paymentQuery, sql.Named("payment_id", req.PaymentID))

	var paymentAmount float64
	var paymentStatus string
	err = row.Scan(&paymentAmount, &paymentStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Payment not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve payment details", http.StatusInternalServerError)
			fmt.Printf("Error: %v", err)
		}
		return
	}

	// Check if the payment has already been refunded or if it failed
	if paymentStatus == "refunded" {
		http.Error(w, "Payment has already been refunded", http.StatusBadRequest)
		return
	}

	if paymentStatus != "approved" {
		http.Error(w, "Refund can only be processed for approved payments", http.StatusBadRequest)
		return
	}

	// Call the bank simulator to approve the refund
	response, err := http.Post(refundSimulatorURL, "application/json", bytes.NewBuffer([]byte(`{"payment_id":`+strconv.Itoa(req.PaymentID)+`}`)))
	if err != nil {
		http.Error(w, "Failed to communicate with bank simulator", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}
	defer response.Body.Close()

	var bankResponse classes.BankSimulatorResponse
	err = json.NewDecoder(response.Body).Decode(&bankResponse)
	if err != nil {
		http.Error(w, "Failed to decode bank simulator response", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	if !bankResponse.Approved {
		// Handle failed refund scenario
		http.Error(w, "Refund not approved by bank simulator", http.StatusPaymentRequired)
		return
	}

	// Update payment status to 'refunded'
	updatePaymentQuery := `UPDATE Payments SET status = 'refunded' WHERE payment_id = @payment_id`
	_, err = tx.Exec(updatePaymentQuery, sql.Named("payment_id", req.PaymentID))
	if err != nil {
		http.Error(w, "Failed to update payment status", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Insert refund record
	refundQuery := `INSERT INTO Refunds (payment_id, amount, status, created_at)
                    VALUES (@payment_id, @amount, 'completed', @created_at)`
	_, err = tx.Exec(refundQuery,
		sql.Named("payment_id", req.PaymentID),
		sql.Named("amount", paymentAmount),
		sql.Named("created_at", time.Now()),
	)
	if err != nil {
		http.Error(w, "Failed to insert refund record", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		fmt.Printf("Error: %v", err)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	resp := classes.RefundResponse{
		Message: "Refund processed successfully",
	}
	json.NewEncoder(w).Encode(resp)
}

// SetupRouter sets up the routes for the API
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/process-payment", ProcessPayment).Methods("POST")
	r.HandleFunc("/get-payments", GetPaymentsByMerchant).Methods("GET")
	r.HandleFunc("/process-refund", ProcessRefund).Methods("POST")
	return r
}
