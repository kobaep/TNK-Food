package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/", http.FileServer(http.Dir("./webapp")))
	http.ListenAndServe(":8080", nil)
}
