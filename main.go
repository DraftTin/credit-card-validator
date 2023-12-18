package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DraftTin/credit-card-validator/validator"
)

type CreditCard struct {
	CardNumber string `json:"cardNumber"`
}

func creditCardValidator(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var card CreditCard
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	isValid := validator.Luhn(card.CardNumber)
	response := map[string]bool{"isValid": isValid}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/", creditCardValidator)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
