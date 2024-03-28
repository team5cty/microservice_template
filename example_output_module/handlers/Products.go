package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example_output_module/prisma/db"
)

type Products struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Products_list []*Products

func (products *Products_list) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(products)
}

func GET_Products_Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	var products Products_list
	res, err := client.Products.FindMany().Exec(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, object := range res {
		ele := &Products{
			Name:  object.Name,
			Price: object.Price,
		}
		products = append(products, ele)
	}
	products.ToJSON(w)
}
