package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"example_output_module/prisma/db"
)

type AddUser struct {
	Dob      string `json:"dob"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (adduser *AddUser) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(adduser)
}

func POST_AddUser_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := db.NewClient() // Initialize Prisma client
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}() // is object
	// is object

	obj := &db.UserModel{}
	obj.FromJSON(r.Body)

}
