package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayush14/ecommerce-backend/internal/services"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Categories := services.GetAllCategories()

	if len(Categories) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No Categories found"}`))
		return
	}

	json.NewEncoder(w).Encode(Categories)
}
