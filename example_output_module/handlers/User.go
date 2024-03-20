package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"example_output_module/database"
	"github.com/gorilla/mux"
)

type User struct {
	Email string   `json:"email"`
	Username string   `json:"username"`
}




func (user *User) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(user)
}

func GET_User_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
		return
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var params map[string]string = mux.Vars(r) //access dynamic variables from this map.
	var user User
	// Implement logic for GET /user/{id}
	user.ToJSON(w)	
}
