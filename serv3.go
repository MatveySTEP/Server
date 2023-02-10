package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	email string `json:"email"`
}

func parseBody(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var obj Body
	json.Unmarshal(data, obj)
	fmt.Fprintf(w, fmt.Sprintf("#{obj}"))
}
