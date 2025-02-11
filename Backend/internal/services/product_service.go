package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

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

// func GetFilteredProducts(filters map[string][]string, sortField, sortOrder string, page, limit int) ([]models.Product, int){
// 	filteredProducts := AllProducts

// 	for key, values := range filters {
// 		filteredProducts = filterProducts(filteredProducts, key, values)
// 	}

// 	if sortField != ""{
// 		sortProducts(filteredProducts, sortField, sortOrder)
// 	}

// 	totalItems := len(filteredProducts)
// 	start := (page - 1) * limit
// 	end := start + limit
// 	if end > totalItems {
// 		end = totalItems
// 	}
// 	if start > totalItems {
// 		start = totalItems
// 	}

// 	paginatedProducts := filteredProducts[start:end]

// 	return paginatedProducts, totalItems
// }

// func filterProducts(products []models.Product, key string, values []string) []models.Product {
//     var filtered []models.Product
//     for _, product := range products {
//         for _, value := range values {
//             if strings.Contains(strings.ToLower(getProductField(product, key)), strings.ToLower(value)) {
//                 filtered = append(filtered, product)
//                 break
//             }
//         }
//     }
//     return filtered
// }

// func getProductField(product models.Product, key string) string {
//     switch key {
//     case "category":
//         return product.Category
//     case "brand":
//         return product.Brand
//     // Add more cases as needed
//     default:
//         return ""
//     }
// }

// func sortProducts(products []models.Product, sortField, sortOrder string) {
//     sort.Slice(products, func(i, j int) bool {
//         switch sortField {
//         case "price":
//             if sortOrder == "desc" {
//                 return products[i].Price > products[j].Price
//             }
//             return products[i].Price < products[j].Price
//         case "rating":
//             if sortOrder == "desc" {
//                 return products[i].Rating > products[j].Rating
//             }
//             return products[i].Rating < products[j].Rating
//         // Add more cases as needed
//         default:
//             return false
//         }
//     })
// }


func GetFilteredProducts(filters map[string][]string, sortField, sortOrder string, page, limit int) ([]models.Product, int) {
    filteredProducts := GetAllProducts()
	fmt.Println("Allproducts",AllProducts)
	fmt.Println(filteredProducts)
    // Apply filters
    for key, values := range filters {
        filteredProducts = filterProducts(filteredProducts, key, values)
    }

    // Apply sorting
    if sortField != "" {
        sortProducts(filteredProducts, sortField, sortOrder)
    }

    // Apply pagination
    totalItems := len(filteredProducts)
    start := (page - 1) * limit
    end := start + limit
    if start > totalItems {
        start = totalItems
    }
    if end > totalItems {
        end = totalItems
    }
    paginatedProducts := filteredProducts[start:end]

    return paginatedProducts, totalItems
}

func filterProducts(products []models.Product, key string, values []string) []models.Product {
    var filtered []models.Product
    for _, product := range products {
        for _, value := range values {
            if strings.Contains(strings.ToLower(getProductField(product, key)), strings.ToLower(value)) {
                filtered = append(filtered, product)
                break
            }
        }
    }
    return filtered
}

func getProductField(product models.Product, key string) string {
    switch key {
    case "category":
        return product.Category
    case "brand":
        return product.Brand
    // Add more cases as needed
    default:
        return ""
    }
}

func sortProducts(products []models.Product, sortField, sortOrder string) {
    sort.Slice(products, func(i, j int) bool {
        switch sortField {
        case "price":
            if sortOrder == "desc" {
                return products[i].Price > products[j].Price
            }
            return products[i].Price < products[j].Price
        case "rating":
            if sortOrder == "desc" {
                return products[i].Rating > products[j].Rating
            }
            return products[i].Rating < products[j].Rating
        // Add more cases as needed
        default:
            return false
        }
    })
}