package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"example_output_module/prisma/db"
)

type Users struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
type Users_list []*Users

func (users *Users_list) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(users)
}

func GET_Users_Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	var users Users_list
	res, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, object := range res {
		ele := &Users{
			Email:    object.Email,
			Username: object.Username,
		}
		users = append(users, ele)
	}
	users.ToJSON(w)
}
