package handlers

import (
	"context"
	"encoding/json"
	"example_output_module/prisma"
	"io"
	"net/http"
)

type Userss struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Userss_list []*Userss

func (userss *Userss_list) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(userss)
}

func GET_Userss_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
	var userss Userss_list

	res, err := client.User.FindMany(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, object := range res {
		ele := &Userss{
			Email:    object.email,
			Id:       object.id,
			Username: object.username,
		}
		userss = append(userss, ele)
	}
	userss.ToJSON(w)

}
