package main

import (
	"Order/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/placeorder", handlers.POST_placeorder_Handler).Methods("POST")

	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8090", r)
	if err != nil {
		fmt.Printf("Cannot start server: %s", err.Error())
	}
}
