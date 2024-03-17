package main

import (
	"fmt"
	"net/http"
	"output/handlers"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{id}", handlers.GET_User_Handler).Methods("GET")
	
	r.HandleFunc("/users", handlers.GET_Users_Handler).Methods("GET")
	
	r.HandleFunc("/getrandom", handlers.GET_random_Handler).Methods("GET")
	
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", r)
}