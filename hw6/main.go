package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Transaction struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"` // income, expense, transfer
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

var (
	transactions []*Transaction
	nextID       = 1
	mu           sync.Mutex
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/transactions", handleTransactions).Methods("GET", "POST")
	router.HandleFunc("/transactions/{id}", handleTransactionByID).Methods("GET", "PUT", "DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTransactions(w, r)
	case http.MethodPost:
		createTransaction(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newTransaction Transaction
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTransaction.ID = strconv.Itoa(nextID)
	nextID++
	newTransaction.Date = time.Now()
	transactions = append(transactions, &newTransaction)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaction)
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	switch r.Method {
	case http.MethodGet:
		getTransactionByID(w, r, id)
	case http.MethodPut:
		updateTransactionByID(w, r, id)
	case http.MethodDelete:
		deleteTransactionByID(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTransactionByID(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	defer mu.Unlock()

	for _, t := range transactions {
		if t.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func updateTransactionByID(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	defer mu.Unlock()

	var updatedTransaction Transaction
	err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, t := range transactions {
		if t.ID == id {
			updatedTransaction.ID = id
			transactions[i] = &updatedTransaction
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTransaction)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func deleteTransactionByID(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	defer mu.Unlock()

	for i, t := range transactions {
		if t.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
