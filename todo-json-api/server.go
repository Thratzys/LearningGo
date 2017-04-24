package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func  Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
}
