package main

import (
	"fmt"
	"net/http"
)

func parseURLQuery2(w http.ResponseWriter, r *http.Request) {
	orderId2 := r.URL.Query()["id"]
	for _, id := range orderId2 {
		status := orderStatus(id)
		fmt.Fprintln(w, status)
	}

}
