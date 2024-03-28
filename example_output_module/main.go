package main

import (
	"fmt"
	"net/http"
	"example_output_module/handlers"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products/", handlers.GET_Products_Handler).Methods("GET")
	r.HandleFunc("/addproduct/", handlers.POST_AddProduct_Handler).Methods("POST")
	
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8080", r)
	if err!=nil{
		fmt.Printf("Cannot start server: %s",err.Error())
	}
}