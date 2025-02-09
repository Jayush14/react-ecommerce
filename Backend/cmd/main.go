package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Jayush14/ecommerce-backend/internal/handlers"
	"github.com/Jayush14/ecommerce-backend/internal/services"
)

func main() {
	err := services.LoadAllData("data/data.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/products", handlers.GetProductsHandler).Methods("GET")
	router.HandleFunc("/categories", handlers.GetCategoriesHandler).Methods("GET")
	router.HandleFunc("/brands", handlers.GetBrandsHandler).Methods("GET")
	port := ":8000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
