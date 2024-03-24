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

func (adduser *AddUser) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(adduser)
}

func POST_AddUser_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := db.NewClient() // Initialize Prisma client
	ctx := context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}() // is object
	// is object

	// Define a struct to hold the request body data
	var requestData AddUser
	// Decode the JSON request body into the requestData struct
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	// Insert data into the table

	_, err := client.User.CreateOne(
		db.User.Email.Set(requestData.Email),
		db.User.ID.Set(requestData.Id),
		db.User.Username.Set(requestData.Username),
		db.User.Dob.Set(requestData.Dob),
	).Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
