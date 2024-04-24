package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"Order/kafka"
)

type placeorder struct {
	Productid int `json:"productid"`
}

func (placeorder *placeorder) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(placeorder)
}

func POST_placeorder_Handler(w http.ResponseWriter, r *http.Request) {

	orderid_producer := kafka.Producer("orderid", 0)

	var requestData placeorder
	if err := requestData.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	s := strconv.Itoa(requestData.Productid)

	orderid_producer(s)
}
