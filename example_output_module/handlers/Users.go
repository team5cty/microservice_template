package handlers

import (
	"context"
	"encoding/json"
	"example_output_module/prisma"
	"io"
	"net/http"
)

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Users_list []*Users

func (users *Users_list) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(users)
}

func GET_Users_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
	var users Users_list

	res, err := client.User.FindMany(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, object := range res {
		ele := &Users{
			Id:       object.id,
			Username: object.username,
		}
		users = append(users, ele)
	}
	users.ToJSON(w)

}
