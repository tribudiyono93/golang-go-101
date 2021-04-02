package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	os.Exit(0)
}

func shutdown(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprint(writer, "The home page")
}


