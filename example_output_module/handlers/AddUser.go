package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"context"
	
	"example_output_module/prisma/db"
)

type AddUser struct {
	Dob string   `json:"dob"`
	Email string   `json:"email"`
	Id int   `json:"id"`
	Username string   `json:"username"`
}


func (adduser *AddUser) FromJSON(r io.Reader) error {
	d:= json.NewDecoder(r)
	return d.Decode(adduser)
}

func POST_AddUser_Handler (w http.ResponseWriter, r *http.Request) {
	client := db.NewClient() 
	ctx := context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	
	var requestData AddUser
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	_, err := client.User.CreateOne(
		db.User.Dob.Set(requestData.Dob),
		db.User.Email.Set(requestData.Email),
		db.User.ID.Set(requestData.Id),
		db.User.Username.Set(requestData.Username),
	).Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}