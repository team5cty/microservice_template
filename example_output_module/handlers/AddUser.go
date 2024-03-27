package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"example_output_module/prisma/db"
)

type AddUser struct {
	Dob      string `json:"dob"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// FromJSON decodes JSON data from an io.Reader into an AddUser struct
func (adduser *AddUser) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(adduser)
}

// POST_AddUser_Handler handles HTTP POST requests to add a new user
func POST_AddUser_Handler(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	ctx := context.Background()

	// Connect to the database
	if err := client.Prisma.Connect(); err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer func() {
		// Disconnect from the database
		if err := client.Prisma.Disconnect(); err != nil {
			http.Error(w, "Failed to disconnect from the database", http.StatusInternalServerError)
		}
	}()

	// Set response content type
	w.Header().Set("Content-Type", "application/json")

	// Decode request body into AddUser struct
	var requestData AddUser
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Create new user in the database
	_, err := client.User.CreateOne(
		db.User.Dob.Set(requestData.Dob),
		db.User.Email.Set(requestData.Email),
		db.User.ID.Set(requestData.Id),
		db.User.Username.Set(requestData.Username),
	).Exec(ctx)

	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Write success response
	w.WriteHeader(http.StatusCreated)
}
