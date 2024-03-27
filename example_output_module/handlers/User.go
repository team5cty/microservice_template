package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"context"
	
	"strconv"
	
	"github.com/gorilla/mux"
	"example_output_module/prisma/db"
)

type User struct {
	Dob string   `json:"dob"`
	Email string   `json:"email"`
	Username string   `json:"username"`
}


func (user *User) ToJSON(w io.Writer) error {
	e:= json.NewEncoder(w)
	return e.Encode(user)
}

func GET_User_Handler (w http.ResponseWriter, r *http.Request) {
	client := db.NewClient() 
	ctx := context.Background()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	m:=mux.Vars(r)
	var val string
	for _,v := range m{
		val=v
	}
	value, _ := strconv.Atoi(val)

	res, err := client.User.FindUnique(db.User.ID.Equals(value)).Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ele := &User{
		Dob:res.Dob,
		Email:res.Email,
		Username:res.Username,
	}
	ele.ToJSON(w)
}