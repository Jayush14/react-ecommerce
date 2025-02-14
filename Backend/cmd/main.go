package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Jayush14/ecommerce-backend/internal/handlers"
	"github.com/Jayush14/ecommerce-backend/internal/services"
	"github.com/gorilla/mux"
)

var dbpool *pgxpool.Pool
var DATABASE_URL string
func init() {
	err := godotenv.Load(".env") 
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

}

func main() {
	var err error
    dbpool, err = pgxpool.New(context.Background(), DATABASE_URL)

	if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}
	defer dbpool.Close()

	fmt.Println("✅ Connected to database")

	
	//   err = createTables()
	//  if err != nil {
	// 	log.Fatalf("❌ Failed to create tables: %v", err)
	// }
	

	err = services.LoadAllDataonDB(dbpool, "data/data.json")
	if err != nil {
		log.Fatalf("❌ Error loading products on DB: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProductsHandler(w, r, dbpool)
	}).Methods("GET")

	router.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCategoriesHandler(w, r, dbpool)
	}).Methods("GET")

	router.HandleFunc("/brands", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBrandsHandler(w, r, dbpool)
	}).Methods("GET")

	router.HandleFunc("/products/filter", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetFilteredProductsHandler(w, r, dbpool)
	}).Methods("GET")

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, r, dbpool)
	}).Methods("POST")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginUserHandler(w, r, dbpool)
	}).Methods("POST")

	port := ":8000"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func deleteTables() error {
	queries := []string{
		"DROP TABLE IF EXISTS products CASCADE;",
		"DROP TABLE IF EXISTS categories CASCADE;",
		"DROP TABLE IF EXISTS brands CASCADE;",
		"DROP TABLE IF EXISTS reviews CASCADE;",
	}

	for _, query := range queries {
		_, err := dbpool.Exec(context.Background(), query)
		if err != nil {
			return fmt.Errorf("❌ Error deleting table: %v", err)
		}
	}
	fmt.Println("✅ All tables deleted successfully")
	return nil
}

func createTables() error {
	_, err := dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			category TEXT NOT NULL,
			price NUMERIC(10,2) NOT NULL,
			discount_percentage NUMERIC(5,2),
			rating NUMERIC(3,2),
			stock INT NOT NULL,
			tags TEXT[],
			brand TEXT NOT NULL,
			sku TEXT UNIQUE NOT NULL,
			weight INT,
			width NUMERIC(5,2),
			height NUMERIC(5,2),
			depth NUMERIC(5,2),
			warranty_information TEXT,
			shipping_information TEXT,
			availability_status TEXT,
			return_policy TEXT,
			minimum_order_quantity INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			barcode TEXT,
			qr_code TEXT,
			images TEXT[],
			thumbnail TEXT
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ Error creating products table: %v", err)
	}

	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			value TEXT UNIQUE NOT NULL,
			label TEXT NOT NULL,
			checked BOOLEAN DEFAULT FALSE
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ Error creating categories table: %v", err)
	}

	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS brands (
			id SERIAL PRIMARY KEY,
			value TEXT UNIQUE NOT NULL,
			label TEXT NOT NULL,
			checked BOOLEAN DEFAULT FALSE
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ Error creating brands table: %v", err)
	}

	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS reviews (
			id SERIAL PRIMARY KEY,
			product_id INT REFERENCES products(id) ON DELETE CASCADE,
			rating INT CHECK (rating BETWEEN 1 AND 5),
			comment TEXT,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			reviewer_name TEXT,
			reviewer_email TEXT
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ Error creating reviews table: %v", err)
	}

	_, err = dbpool.Exec(context.Background(), `
			CREATE TABLE IF NOT EXISTS metadata (
			key TEXT PRIMARY KEY,
			value TEXT
		);
	`)
	if err != nil {
		return fmt.Errorf("❌ Error creating metadata table: %v", err)
	}

	_, err  = dbpool.Exec(context.Background(),`
	  CREATE TABLE IF NOT EXISTS Users (
	   id SERIAL PRIMARY KEY,
	   name TEXT NOT NULL,
	   email TEXT NOT NULL UNIQUE,
	   password TEXT NOT NULL
	  ); 
	  `)
	if err != nil {
		return fmt.Errorf("❌ Error creating User table: %v", err)
	}
    
	fmt.Printf("✅ All tables created successfully\n")

	return nil
}