package main

import (
	"fmt"
	"net/http"
	"Product/handlers"
	"Product/kafka"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/addproduct/", handlers.POST_AddProduct_Handler).Methods("POST")
	
	go kafka.Consume("orderid", 0, func(s string) {})
	
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8000", r)
	if err!=nil{
		fmt.Printf("Cannot start server: %s",err.Error())
	}
}