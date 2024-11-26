package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"receipt-processor/internal/logic"
	"receipt-processor/internal/models"
	"receipt-processor/internal/storage"
)

var store = storage.NewMemoryStore()

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/receipts/process", processReceipt)
	mux.HandleFunc("/receipts/", getPoints)

	return mux
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt and calculate its points
	receipt.ID = uuid.New().String()
	receipt.Points = logic.CalculatePoints(&receipt)

	// Save the receipt to the in-memory store
	store.SaveReceipt(receipt)

	// Respond with the generated ID
	response := map[string]string{"id": receipt.ID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	// Enforce exact path match for /points
	if !strings.HasSuffix(r.URL.Path, "/points") {
		http.Error(w, "Invalid endpoint", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/points"), "/receipts/")
	if id == "" {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}

	receipt, found := store.GetReceipt(id)
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": receipt.Points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}