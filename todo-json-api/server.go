package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
