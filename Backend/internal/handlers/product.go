package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayush14/ecommerce-backend/internal/services"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := services.GetAllProducts()

	if len(products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No products found"}`))
		return
	}

	json.NewEncoder(w).Encode(products)
}
