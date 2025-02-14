package handlers

import (
	"encoding/json"
	"net/http"

	// "github.com/Jayush14/ecommerce-backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Jayush14/ecommerce-backend/internal/models"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request , dbpool *pgxpool.Pool) {
	row , err := dbpool.Query(r.Context(), `SELECT DISTINCT category FROM products`)

	if err != nil {
		http.Error(w, "Failed to load Categorydata", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var categories = []models.Category{}
	for row.Next() {
		var category models.Category
		err := row.Scan(&category.Category)
		if err != nil {
			http.Error(w, "Failed to load Categorydata during scan", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No categories found"}`))
		return
	}

	json.NewEncoder(w).Encode(categories)

}
