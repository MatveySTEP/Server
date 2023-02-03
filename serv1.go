package main

import (
	"fmt"
	"net/http"
)

func orderStatus(id string) string {
	switch id {
	case "1":
		return "active"
	case "2":
		return "passive"
	default:
		return " no name id"
	}
}

func parseURLQuery(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("id")
	status := orderStatus(orderId)
	fmt.Fprintf(w, status)
}
