package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/order/status", parseURLQuery)
	http.HandleFunc("/order/statusv2", parseURLQuery2)

	_ = http.ListenAndServe(":8080", nil)
}
