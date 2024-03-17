package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"output/database"
)

type random struct {
}


func (random *random) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(random)
}



func GET_random_Handler(w http.ResponseWriter, r *http.Request) {
	db , err := database.Conn()
	if err!=nil{
		fmt.Printf("Cannot connect to database: %s",err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var random random
	// Implement logic for GET /getrandom
	random.ToJSON(w)	
}
