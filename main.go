package main

import (
	"log"
	http2 "net/http"
	"neo/internal/http"
)

func main() {
	handler := http.NewRequestHandler()
	println("noreyaga")
	err := http2.ListenAndServe(":8080", handler)
	log.Fatalln(err)
}
