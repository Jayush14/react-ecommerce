package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayush14/ecommerce-backend/internal/services"
)

func GetBrandsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Brands := services.GetAllBrands()

	if len(Brands) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No Brands found"}`))
		return
	}

	json.NewEncoder(w).Encode(Brands)
}
