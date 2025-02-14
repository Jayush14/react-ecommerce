package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayush14/ecommerce-backend/internal/models"
	// "github.com/Jayush14/ecommerce-backend/internal/services"
	 "github.com/jackc/pgx/v5/pgxpool"
)

func GetBrandsHandler(w http.ResponseWriter, r *http.Request , dbpool *pgxpool.Pool) {
	row , err := dbpool.Query(r.Context(), `SELECT DISTINCT brand FROM products`)

	if err != nil {
		http.Error(w, "Failed to load Branddata", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var brands = []models.Brand{}
	for row.Next() {
		var brand models.Brand
		err := row.Scan(&brand.Brand)
		if err != nil {
			http.Error(w, "Failed to load Branddata during scan", http.StatusInternalServerError)
			return
		}
		brands = append(brands, brand)
	}

	if len(brands) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No brands found"}`))
		return
	}

	json.NewEncoder(w).Encode(brands)
}
