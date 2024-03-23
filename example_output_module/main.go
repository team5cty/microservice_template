package main

import (
	"fmt"
	"net/http"
	"example_output_module/handlers"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/user/{id}", handlers.GET_Users_Handler).Methods("GET")
	
	r.HandleFunc("/user/", handlers.GET_User_Handler).Methods("GET")
	
	r.HandleFunc("/users/", handlers.GET_Userss_Handler).Methods("GET")
	
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8080", r)
	if err!=nil{
		fmt.Printf("Cannot start server: %s",err.Error())
	}
}