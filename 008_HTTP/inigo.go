package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello my name is inigo montoya")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
