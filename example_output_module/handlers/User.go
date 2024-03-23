package handlers

import (
	"context"
	"encoding/json"
	"example_output_module/prisma"

	"io"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type User_list []*User

func (user *User_list) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(user)
}

func GET_User_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
	var user User_list

	res, err := client.User.FindMany(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, object := range res {
		ele := &User{
			Email:    object.email,
			Id:       object.id,
			Username: object.username,
		}
		user = append(user, ele)
	}
	user.ToJSON(w)

}
