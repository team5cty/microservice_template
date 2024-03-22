package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"example_output_module/database"
)

type AddUser struct {
}




func (adduser *AddUser) FromJSON(r io.Reader) error {
	d:= json.NewDecoder(r)
	return d.Decode(adduser)
}

func POST_AddUser_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
		return
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var params map[string]string = mux.Vars(r) //access dynamic variables from this map.
	var adduser AddUser
	adduser.FromJSON(r.Body)
	// Implement logic for POST /adduser/{id}	
}
