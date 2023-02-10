package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Body struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	email string `json:"email"`
}

func parseBody(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.Body.Close()
	var obj Body
	json.Unmarshal(data, obj)
	fmt.Fprintf(w, fmt.Sprintf("#{obj}"))
}
