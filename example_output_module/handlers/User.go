package handlers

import (
	"context"
	"encoding/json"
	"example_output_module/prisma/db"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (user *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(user)
}

func GET_User_Handler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	client := db.NewClient() // Initialize Prisma client
	ctx := context.Background()
	defer client.Disconnect()

	var params map[string]string = mux.Vars(r) //access dynamic variables from this map.
	id, ok := params["id"]
	if !ok {
		http.Error(w, "ID parameter not found in the path", http.StatusBadRequest)
		return
	}
	idint, _ := strconv.Atoi(id)
	res, err := client.User.FindUnique(db.User.ID.Equals(idint)).Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ele := &User{
		Id:       res.Id,
		Username: res.Username,
	}
	ele.ToJSON(w)

}
