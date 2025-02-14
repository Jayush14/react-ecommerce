package services

import (
	"context"
	"fmt"

	"github.com/Jayush14/ecommerce-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// GetFilteredProductsFromDB fetches filtered products from the database
func GetFilteredProductsFromDB(dbpool *pgxpool.Pool, filters map[string][]string, sortField, sortOrder string, limit, offset int) ([]models.Product, int, error) {
	// Base query
	fmt.Println("Filters:", filters)
	fmt.Println("SortField:", sortField)
	fmt.Println("SortOrder:", sortOrder)
	fmt.Println("Limit:", limit)
	fmt.Println("Offset:", offset)

	query := `
    SELECT id, title, description, category, price, discount_percentage, rating, stock, tags, 
           brand, sku, weight, width, height, depth, warranty_information, shipping_information, 
           availability_status, return_policy, minimum_order_quantity, created_at, updated_at, 
           barcode, qr_code, images, thumbnail 
    FROM products
    WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	// Apply filters
	for key, values := range filters {
		if len(values) > 0 {
			query += fmt.Sprintf(" AND %s ILIKE $%d", key, argIndex)
			args = append(args, "%"+values[0]+"%") // Use ILIKE for case-insensitive matching
			argIndex++
		}
	}

	// Sorting
	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc" // Default to ascending
		}
		query += fmt.Sprintf(" ORDER BY %s %s", sortField, sortOrder)
	}

	// Pagination
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	// Execute query
	rows, err := dbpool.Query(context.Background(), query, args...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, 0, err
	}
	defer rows.Close()

	// Fetch results
	var products []models.Product
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
			fmt.Println("Error scanning row:", err)
			return nil, 0, err
		}
		products = append(products, product)
	}

	// Count total items (for pagination)
	totalQuery := "SELECT COUNT(*) FROM products WHERE 1=1"
	totalArgs := []interface{}{}
	totalIndex := 1
	for key, values := range filters {
		if len(values) > 0 {
			totalQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, totalIndex)
			totalArgs = append(totalArgs, "%"+values[0]+"%")
			totalIndex++
		}
	}

	var totalItems int
	err = dbpool.QueryRow(context.Background(), totalQuery, totalArgs...).Scan(&totalItems)
	if err != nil {
		fmt.Println("Error executing count query:", err)
		return nil, 0, err
	}

	fmt.Println("Total items:", totalItems)
	fmt.Println("Products retrieved:", len(products))

	return products, totalItems, nil
}
