package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Config struct {
	APIKey string `json:"api_key"`
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type CurrencyResponse struct {
	Rates struct {
		RUB float64 `json:"RUB"`
	} `json:"rates"`
}

func getCurrencyRate(baseCurrency string, apiKey string) (float64, error) {
	url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s_RUB&compact=ultra&apiKey=%s", baseCurrency, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var currencyResponse CurrencyResponse
	err = json.Unmarshal(body, &currencyResponse)
	if err != nil {
		return 0, err
	}

	return currencyResponse.Rates.RUB, nil
}

func convertAmount(amount float64, baseCurrency string, apiKey string) (float64, error) {
	rate, err := getCurrencyRate(baseCurrency, apiKey)
	if err != nil {
		return 0, err
	}

	convertedAmount := amount * rate
	return convertedAmount, nil
}

type Transaction struct {
	ID                string  `json:"id"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	ConvertedAmount   float64 `json:"converted_amount,omitempty"`
	ConvertedCurrency string  `json:"converted_currency,omitempty"`
}

var transactions = map[string]Transaction{
	"1": {ID: "1", Amount: 100, Currency: "USD"},
	"2": {ID: "2", Amount: 200, Currency: "USD"},
}

func getTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	currency := r.URL.Query().Get("currency")

	transaction, ok := transactions[id]
	if !ok {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	if currency != "" {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "API key not found", http.StatusBadRequest)
			return
		}

		convertedAmount, err := convertAmount(transaction.Amount, transaction.Currency, apiKey)
		if err != nil {
			http.Error(w, "Error converting currency", http.StatusInternalServerError)
			return
		}

		transaction.ConvertedAmount = convertedAmount
		transaction.ConvertedCurrency = currency
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/transactions/{id}", getTransactionHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
