package services

import (
	// "fmt"
	"context"
	"fmt"

	"github.com/Jayush14/ecommerce-backend/internal/models"
	// "github.com/Jayush14/ecommerce-backend/internal/utils"
	

	"github.com/jackc/pgx/v5/pgxpool"
)

var DataStore models.Data

// func LoadAllData(filepath string) error {
// 	 err := utils.LoadJSON(filepath, &DataStore)
// 	if err != nil {
// 		return fmt.Errorf("error loading data: %v", err)
// 	}
// 	fmt.Println("Data loaded successfully")
// 	return nil
// }
   
func LoadAllData(dbpool *pgxpool.Pool) ([]models.Product, error) {
    fmt.Println("Loading data")
    rows, err := dbpool.Query(context.Background(), `SELECT * FROM products`)
    if err != nil {
        
        return nil, err
    }
    defer rows.Close()
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
           
            return nil, err
        }
        DataStore.Products = append(DataStore.Products, product)
    }
    
    fmt.Println("Data loaded successfully")
    fmt.Println(DataStore.Products)
	return DataStore.Products, nil
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