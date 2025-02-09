package services

import (
	"fmt"
	"github.com/Jayush14/ecommerce-backend/internal/models"
	"github.com/Jayush14/ecommerce-backend/internal/utils"
)

var DataStore models.Data

func LoadAllData(filepath string) error {
	 err := utils.LoadJSON(filepath, &DataStore)
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}
	fmt.Println("Data loaded successfully")
	return nil
}

func GetAllProducts() []models.Product {
	return DataStore.Products
}

func GetAllCategories() []models.Category {
	return DataStore.Categories
}

func GetAllBrands() []models.Brand {
	return DataStore.Brands
}