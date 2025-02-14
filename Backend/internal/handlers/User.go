package handlers

import ("net/http"

	"github.com/Jayush14/ecommerce-backend/internal/models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Validate email and password
	if user.Email == "" || user.Password== "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	// Trim whitespace (optional but recommended)
	email := strings.TrimSpace(user.Email)
	password := strings.TrimSpace(user.Password)

	// Check if email already exists
	var exists bool
	err = dbpool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error while checking email", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Insert new user
	_, err = dbpool.Exec(context.Background(), `
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3)`, user.Name, email, string(hashedPassword))

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Validate email and password
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	// Trim whitespace (optional but recommended)
	email := strings.TrimSpace(user.Email)
	password := strings.TrimSpace(user.Password)

	// Get user details from database
	var dbUser models.User
	err = dbpool.QueryRow(context.Background(), "SELECT id, name, email, password FROM users WHERE email=$1", email).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Email, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}
