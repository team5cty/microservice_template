package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Product/prisma/db"
)

type AddProduct struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Productid int     `json:"productid"`
}

func (addproduct *AddProduct) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(addproduct)
}

func POST_AddProduct_Handler(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()
	ctx := context.Background()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Printf("Error connecting database: %s", err.Error())
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			fmt.Printf("Error Disconnecting database: %s", err.Error())
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	var requestData AddProduct
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err := client.Products.CreateOne(
		db.Products.Name.Set(requestData.Name),
		db.Products.Price.Set(requestData.Price),
		db.Products.Productid.Set(requestData.Productid),
	).Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
