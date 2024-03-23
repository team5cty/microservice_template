package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"context"
	"github.com/gorilla/mux"
	"example_output_module/prisma"
	"example_output_module/prisma/db"
)

type Users struct {
}




func (users *Users) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(users)
}

func GET_Users_Handler (w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := prisma.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()
		var users Users
		
			
				var params map[string]string = mux.Vars(r) //access dynamic variables from this map.
				id, ok := params["id"]
				if !ok {
					http.Error(w, "ID parameter not found in the path", http.StatusBadRequest)
					return
				}
				res, err := client.User.FindUnique(db.User.id.Equals(id)).Exec(ctx)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				users = {
				}
				users.ToJSON(w)
	        
		
	
}