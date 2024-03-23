package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"context"
	"example_output_module/prisma"
	"example_output_module/prisma/db"
)

type User struct {
}


type User_list []*User


func (user *User_list) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(user)
}

func GET_User_Handler (w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
	var user User_list
		
				
					res, err := client.User.FindMany()
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					for _, object := range res {
						ele := &User{
						}
						user = append(user, ele)
					}
					user.ToJSON(w)
				
		
	
}