package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"output/database"
)

type Users struct {
	Email string ` + "json:\"email\"" + `
	Username string ` + "json:\"username\"" + `
}


func (users *Users) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(users)
}



func GET_Users_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var users Users
	// Implement logic for GET /users
	users.ToJSON(w)	
}
