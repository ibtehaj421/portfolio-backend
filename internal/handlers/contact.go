package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/models"
)

func HandleContactSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Basic validation
	if contact.Name == "" || contact.Email == "" || contact.Message == "" {
		http.Error(w, "Name, email, and message are required", http.StatusBadRequest)
		return
	}

	// Insert into PostgreSQL
	query := `
		INSERT INTO contacts (name, email, project_type, message) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`
	
	err := database.DB.QueryRow(
		query, 
		contact.Name, 
		contact.Email, 
		contact.ProjectType, 
		contact.Message,
	).Scan(&contact.ID, &contact.CreatedAt)

	if err != nil {
		log.Printf("Failed to insert contact: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"message": "Message received successfully",
	})
}