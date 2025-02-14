package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Jayush14/ecommerce-backend/internal/models"
)

// Compute SHA256 hash of a file
func computeFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("❌ Error opening file: %v", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("❌ Error hashing file: %v", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// Fetch last stored hash from the database
func getStoredHash(db *pgxpool.Pool) (string, error) {
	var storedHash string
	err := db.QueryRow(context.Background(), `SELECT value FROM metadata WHERE key = 'json_hash'`).Scan(&storedHash)
	if err != nil {
		// Return empty string if no hash exists
		return "", nil
	}
	return storedHash, nil
}

// Store the new hash in the database
func storeNewHash(db *pgxpool.Pool, newHash string) error {
	_, err := db.Exec(context.Background(),
		`INSERT INTO metadata (key, value) VALUES ('json_hash', $1) 
		 ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value`,
		newHash)
	return err
}

// Load JSON data only if there is a change
func LoadAllDataonDB(db *pgxpool.Pool, jsonFilePath string) error {
	// Compute hash of current JSON file
	newHash, err := computeFileHash(jsonFilePath)
	if err != nil {
		return err
	}

	// Get stored hash from DB
	storedHash, err := getStoredHash(db)
	if err != nil {
		return err
	}

	// Check if data has changed
	if newHash == storedHash {
		fmt.Println("✅ No changes detected, skipping data load.")
		return nil
	}

	fmt.Println("🔄 Changes detected, updating database...")

	// Open JSON file
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return fmt.Errorf("❌ Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Decode JSON into Data struct
	var data models.Data
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("❌ Error decoding JSON: %v", err)
	}

	// Insert categories
	// for _, category := range data.Categories {
	// 	_, err := db.Exec(context.Background(),
	// 		`INSERT INTO categories (value, label, checked) VALUES ($1, $2, $3) 
	// 		 ON CONFLICT (value) DO NOTHING`,
	// 		category.Value, category.Label, category.Checked)
	// 	if err != nil {
	// 		log.Printf("❌ Error inserting category: %v", err)
	// 	}
	// }

	// Insert brands
	// for _, brand := range data.Brands {
	// 	_, err := db.Exec(context.Background(),
	// 		`INSERT INTO brands (value, label, checked) VALUES ($1, $2, $3) 
	// 		 ON CONFLICT (value) DO NOTHING`,
	// 		brand.Value, brand.Label, brand.Checked)
	// 	if err != nil {
	// 		log.Printf("❌ Error inserting brand: %v", err)
	// 	}
	// }

	// Map product IDs to insert reviews later
	// productIDMap := make(map[string]int)

	// Insert products
	for _, product := range data.Products {
		var productID int
		err := db.QueryRow(context.Background(),
			`INSERT INTO products (title, description, category, price, discount_percentage, rating, stock, 
				tags, brand, sku, weight, width, height, depth, warranty_information, shipping_information, 
				availability_status, return_policy, minimum_order_quantity, barcode, qr_code, images, thumbnail) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
			RETURNING id`,
			product.Title, product.Description, product.Category, product.Price, product.DiscountPercentage,
			product.Rating, product.Stock, product.Tags, product.Brand, product.SKU, product.Weight,
			product.Width, product.Height, product.Depth,
			product.WarrantyInformation, product.ShippingInformation, product.AvailabilityStatus,
			product.ReturnPolicy, product.MinimumOrderQuantity, product.Barcode, product.QRCode,
			product.Images, product.Thumbnail).Scan(&productID)
		if err != nil {
			log.Printf("❌ Error inserting product: %v", err)
			continue
		}

		// Store the product ID for review insertion
		// productIDMap[product.Title] = productID
	}

	// Insert reviews
	// for _, product := range data.Products {
	// 	productID, exists := productIDMap[product.Title]
	// 	if !exists {
	// 		log.Printf("⚠️ Skipping reviews for product: %s (not found)", product.Title)
	// 		continue
	// 	}

	// 	for _, review := range product.Reviews {
	// 		_, err := db.Exec(context.Background(),
	// 			`INSERT INTO reviews (product_id, rating, comment, date, reviewer_name, reviewer_email) 
	// 			 VALUES ($1, $2, $3, $4, $5, $6)`,
	// 			productID, review.Rating, review.Comment, review.Date, review.ReviewerName, review.ReviewerEmail)
	// 		if err != nil {
	// 			log.Printf("❌ Error inserting review for product %d: %v", productID, err)
	// 		}
	// 	}
	// }

	// Store new hash in the database
	err = storeNewHash(db, newHash)
	if err != nil {
		log.Printf("❌ Error saving new JSON hash: %v", err)
	}

	fmt.Println("✅ Data successfully updated!")
	return nil
}
