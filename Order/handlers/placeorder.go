package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"Order/kafka"
	"Order/prisma/db"
)

type placeorder struct {
	Productid int `json:"productid"`
}

func (placeorder *placeorder) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(placeorder)
}

func POST_placeorder_Handler(w http.ResponseWriter, r *http.Request) {

	produce := kafka.Producer("orderid", 0)

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

	var requestData placeorder
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err := client.Order.CreateOne(
		db.Order.Productid.Set(requestData.Productid),
	).Exec(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	produce(strconv.Itoa(requestData.Productid))
}
