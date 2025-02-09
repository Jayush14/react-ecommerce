package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Jayush14/ecommerce-backend/internal/models"

)

var AllProducts []models.Product

func LoadProducts(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var data models.Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	AllProducts = data.Products
	fmt.Printf("Loaded %d products\n", len(AllProducts))
	return nil
}
