package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"example_output_module/database"
)

type Users struct {
	Email string   `json:"email"`
	Id int   `json:"id"`
	Password string   `json:"password"`
	Username string   `json:"username"`
}


type Users_list []*Users


func (users *Users_list) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(users)
}

func GET_Users_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
		return
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var users Users_list
	// Implement logic for GET /
	users.ToJSON(w)	
}
