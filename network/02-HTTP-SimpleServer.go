package main

import (
	"net/http"
	"io"
)

func test(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Server Test</title>
			</head>
			<body>
				<h1>Server Test</h1>
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":3000", nil)
}
