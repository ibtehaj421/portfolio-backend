package models

import "time"

type Contact struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	ProjectType string    `json:"projectType"` // Matches Next.js state
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}