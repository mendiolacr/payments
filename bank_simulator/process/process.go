package process

import (
	"bank_simulator/classes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SimulatePayment simulates the payment process and returns a response
func SimulatePayment(w http.ResponseWriter, r *http.Request) {
	var req classes.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Simulate a 60% chance of success
	success := rand.Intn(100) < 60

	response := classes.PaymentResponse{
		Success: success,
		Message: "Payment " + map[bool]string{true: "approved", false: "rejected"}[success],
	}
	fmt.Println("Payment " + map[bool]string{true: "approved", false: "rejected"}[success])
	w.Header().Set("Content-Type", "application/json")
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusPaymentRequired)
	}
	json.NewEncoder(w).Encode(response)
}

// SimulateRefund handles the refund simulation logic
func SimulateRefund(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming request
	var req map[string]int
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the payment_id from the request
	paymentID, ok := req["payment_id"]
	if !ok {
		http.Error(w, "Missing payment_id", http.StatusBadRequest)
		return
	}

	// Log the payment ID being processed
	log.Printf("Simulating refund for payment_id: %d", paymentID)

	// Simulate refund approval logic
	// In a real scenario, here you would include your logic to determine if the refund is approved
	approved := true // Simulate approval; change to false to simulate rejection

	// Create the response based on the simulated result
	response := classes.BankSimulatorResponse{
		Approved: approved,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetupRouter sets up the routes for the API
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/simulate-payment", SimulatePayment).Methods("POST")
	r.HandleFunc("/simulate-refund", SimulateRefund)
	return r
}
