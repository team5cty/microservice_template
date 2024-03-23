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

type Userss struct {
	Name string   `json:"name"`
}


type Userss_list []*Userss


func (userss *Userss_list) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(userss)
}

func GET_Userss_Handler (w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
	var userss Userss_list
		
				
					//Implement logic for GET /users/
				
		
	
}