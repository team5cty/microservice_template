package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"output/database"
)

type User struct {
	Email string ` + "json:\"email\"" + `
	Id int ` + "json:\"id\"" + `
	Password string ` + "json:\"password\"" + `
	Username string ` + "json:\"username\"" + `
}


func (user *User) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(user)
}



func GET_User_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var user User
	// Implement logic for GET /user/{id}
	user.ToJSON(w)	
}
