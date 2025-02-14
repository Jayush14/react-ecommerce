package handlers

import (
	"encoding/json"
	"net/http"
	 "strconv"
	"fmt"
	"context"
	"github.com/Jayush14/ecommerce-backend/internal/models"
	 "github.com/Jayush14/ecommerce-backend/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"

)

// func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	products := services.GetAllProducts()

// 	if len(products) == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"error": "No products found"}`))
// 		return
// 	}

// 	json.NewEncoder(w).Encode(products)
// }
 var DataStore models.Data

 func GetProductsHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
    fmt.Println("Loading data")

    rows, err := dbpool.Query(context.Background(), `
        SELECT id, title, description, category, price, discount_percentage, rating, stock, tags, 
               brand, sku, weight, width, height, depth, warranty_information, shipping_information, 
               availability_status, return_policy, minimum_order_quantity, created_at, updated_at, 
               barcode, qr_code, images, thumbnail 
        FROM products
    `)
    if err != nil {
        http.Error(w, "Failed to load Productdata", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    DataStore.Products = []models.Product{}

    for rows.Next() {
        var product models.Product
        err := rows.Scan(
            &product.ID, &product.Title, &product.Description, &product.Category, &product.Price,
            &product.DiscountPercentage, &product.Rating, &product.Stock, &product.Tags, &product.Brand,
            &product.SKU, &product.Weight, &product.Width, &product.Height, &product.Depth,
            &product.WarrantyInformation, &product.ShippingInformation, &product.AvailabilityStatus,
            &product.ReturnPolicy, &product.MinimumOrderQuantity, &product.CreatedAt, &product.UpdatedAt,
            &product.Barcode, &product.QRCode, &product.Images, &product.Thumbnail,
        )
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to load data during Scan: %v", err), http.StatusInternalServerError)
            return
        }
        DataStore.Products = append(DataStore.Products, product)
    }

    fmt.Println("Data loaded successfully")
    json.NewEncoder(w).Encode(DataStore.Products)
}
func GetFilteredProductsHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	w.Header().Set("Content-Type", "application/json")

	queryParams := r.URL.Query()
	filters := make(map[string][]string)

	// Extract filtering parameters
	for key, value := range queryParams {
		if key != "_sort" && key != "_order" && key != "_limit" && key != "_page" {
			filters[key] = value
		}
	}

	// Sorting parameters
	sortField := queryParams.Get("_sort")
	sortOrder := queryParams.Get("_order")

	// Pagination parameters
	limit, err := strconv.Atoi(queryParams.Get("_limit"))
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}
	page, err := strconv.Atoi(queryParams.Get("_page"))
	if err != nil || page <= 0 {
		page = 1 // Default page
	}
	offset := (page - 1) * limit

	// Get filtered data from database
	products, totalItems, err := services.GetFilteredProductsFromDB(dbpool, filters, sortField, sortOrder, limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch filtered products", http.StatusInternalServerError)
		return
	}

	// Set total count in response header
	w.Header().Set("X-Total-Count", strconv.Itoa(totalItems))

	// Send JSON response
	json.NewEncoder(w).Encode(products)
}
