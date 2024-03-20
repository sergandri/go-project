package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CommissionRequest struct {
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TransactionType string  `json:"type"`
}

type CommissionResponse struct {
	Commission float64 `json:"commission"`
}

func calculateCommission(req CommissionRequest) CommissionResponse {
	var commission float64

	switch req.TransactionType {
	case "transfer":
		if req.Currency == "USD" {
			commission = req.Amount * 0.02
		} else if req.Currency == "RUB" {
			commission = req.Amount * 0.05
		}
	case "purchase", "deposit":
		commission = 0
	default:
		commission = 0
	}

	return CommissionResponse{
		Commission: commission,
	}
}

func commissionsCalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req CommissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := calculateCommission(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/commissions/calculate", commissionsCalculateHandler)
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
