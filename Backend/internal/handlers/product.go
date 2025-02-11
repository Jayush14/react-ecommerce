package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	 query := r.URL.Query()
	filters := make(map[string][]string)
	for key, value := range query {
		if key != "_sort" && key != "_order" && key != "_limit" && key != "_page" {
			filters[key] = value
		}
	}

	sort := query.Get("_sort")
	order := query.Get("_order")

	limit, _ := strconv.Atoi(query.Get("_limit"))
	page,_:= strconv.Atoi(query.Get("_page"))

	products, totalItems := services.GetFilteredProducts(filters, sort, order, page, limit)


	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Total-Count", strconv.Itoa(totalItems))
    json.NewEncoder(w).Encode(products)

}